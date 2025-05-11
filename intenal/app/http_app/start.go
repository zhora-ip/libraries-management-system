package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	kafkaConsumer "github.com/zhora-ip/libraries-management-system/infrastructure/kafka/consumer"
	kafkaProducer "github.com/zhora-ip/libraries-management-system/infrastructure/kafka/producer"
	auth "github.com/zhora-ip/libraries-management-system/pkg/token_manager"

	"github.com/zhora-ip/libraries-management-system/intenal/app/audit"
	"github.com/zhora-ip/libraries-management-system/intenal/app/http_app/server"
	bookservice "github.com/zhora-ip/libraries-management-system/intenal/app/service/book"
	orderservice "github.com/zhora-ip/libraries-management-system/intenal/app/service/order"
	physbookservice "github.com/zhora-ip/libraries-management-system/intenal/app/service/phys_books"
	userservice "github.com/zhora-ip/libraries-management-system/intenal/app/service/user"
	"github.com/zhora-ip/libraries-management-system/intenal/storage/redis"
	sqldb "github.com/zhora-ip/libraries-management-system/intenal/storage/sql_storage/db"
	auditlogs "github.com/zhora-ip/libraries-management-system/intenal/storage/sql_storage/repository/postgresql/audit_logs"
	"github.com/zhora-ip/libraries-management-system/intenal/storage/sql_storage/repository/postgresql/books"
	libcards "github.com/zhora-ip/libraries-management-system/intenal/storage/sql_storage/repository/postgresql/lib_cards"
	"github.com/zhora-ip/libraries-management-system/intenal/storage/sql_storage/repository/postgresql/libraries"
	"github.com/zhora-ip/libraries-management-system/intenal/storage/sql_storage/repository/postgresql/orders"
	physbooks "github.com/zhora-ip/libraries-management-system/intenal/storage/sql_storage/repository/postgresql/physical_books"
	"github.com/zhora-ip/libraries-management-system/intenal/storage/sql_storage/repository/postgresql/tasks"
	"github.com/zhora-ip/libraries-management-system/intenal/storage/sql_storage/repository/postgresql/users"
)

const (
	timeoutCheckExpired = 1
	insecurePort        = "8000"
	securePort          = "8001"
)

type shutdowner interface {
	ShutDown()
}

func Start(cfg *Config) error {
	var (
		shutdowners []shutdowner
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err := sqldb.NewDb(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.GetPool().Close()

	tkManager, err := auth.NewManager(os.Getenv("JWT_SECRET"))
	if err != nil {
		log.Fatal(err)
	}

	aRepo := auditlogs.NewAuditLogs(db)
	tRepo := tasks.NewTasks(db)
	bRepo := books.NewBooks(db)
	lRepo := libraries.NewLibraries(db)
	lcRepo := libcards.NewLibCards(db)
	uRepo := users.NewUsers(db)
	pbRepo := physbooks.NewPhysBooks(db)
	oRepo := orders.NewOrders(db)

	wPool := audit.NewWP(aRepo, false)
	wPool.SetNext(audit.NewWP(nil, true))
	wPool.Run()
	shutdowners = append(shutdowners, wPool)

	wPool2 := audit.NewWP(tRepo, false)
	wPool2.Run()
	shutdowners = append(shutdowners, wPool2)

	bService := bookservice.New(bRepo, db.GetTM())
	uService := userservice.New(uRepo, lcRepo, oRepo, db.GetTM(), tkManager)
	pbService := physbookservice.New(pbRepo, lRepo, db.GetTM())
	oService := orderservice.New(pbRepo, oRepo, lcRepo, db.GetTM(), wPool2)

	p, err := kafkaProducer.New(cfg.KafkaCfg)
	if err != nil {
		log.Print(err)
	}
	aProducer := audit.NewAuditProducer(tRepo, p)
	shutdowners = append(shutdowners, p)

	go aProducer.Produce(ctx)

	for i := 1; i <= 3; i++ {
		c, err := kafkaConsumer.New(cfg.KafkaCfg)
		if err != nil {
			log.Print(err)
			continue
		}
		aConsumer := audit.NewAuditConsumer(c, wPool, true)
		shutdowners = append(shutdowners, c)
		go aConsumer.Consume(ctx)
	}

	cache := redis.New(cfg.RedisCfg)

	srv := server.New(bService, uService, pbService, oService, tkManager, cache)

	go runCanceledOrdersCron(ctx, oService)
	err = runServer(srv, shutdowners)
	return err
}

func runServer(srv *server.Server, shutdowners []shutdowner) error {

	errCh := make(chan error, 2)
	go func() {
		log.Printf("http-server is up and running on %s", insecurePort)
		errCh <- http.ListenAndServe(":"+insecurePort, srv)
	}()

	go func() {
		log.Printf("https-server is up and running on %s", securePort)
		errCh <- http.ListenAndServeTLS(":"+securePort, "./ssl/server.crt", "./ssl/server.key", srv)
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	select {
	case <-sigCh:
		log.Print("Received terminate, graceful shutdown!")
	case err := <-errCh:
		if err != nil && err != http.ErrServerClosed {
			log.Printf("HTTP server error, %v", err)
			return err
		}
	}

	for _, sh := range shutdowners {
		if sh != nil {
			sh.ShutDown()
		}
	}

	log.Print("Server exited gracefully")

	return nil
}

func runCanceledOrdersCron(ctx context.Context, oService *orderservice.OrderService) {
	ticker := time.NewTicker(time.Second * timeoutCheckExpired)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			err := oService.CheckCanceled(ctx)
			if err != nil {
				log.Print(err)
			}
		}
	}
}
