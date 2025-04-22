package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/zhora-ip/libraries-management-system/intenal/app/http_app/server"
	bookservice "github.com/zhora-ip/libraries-management-system/intenal/app/service/book"
	orderservice "github.com/zhora-ip/libraries-management-system/intenal/app/service/order"
	physbookservice "github.com/zhora-ip/libraries-management-system/intenal/app/service/phys_books"
	userservice "github.com/zhora-ip/libraries-management-system/intenal/app/service/user"
	sqldb "github.com/zhora-ip/libraries-management-system/intenal/storage/sql_storage/db"
	"github.com/zhora-ip/libraries-management-system/intenal/storage/sql_storage/repository/postgresql/books"
	libcards "github.com/zhora-ip/libraries-management-system/intenal/storage/sql_storage/repository/postgresql/lib_cards"
	"github.com/zhora-ip/libraries-management-system/intenal/storage/sql_storage/repository/postgresql/libraries"
	"github.com/zhora-ip/libraries-management-system/intenal/storage/sql_storage/repository/postgresql/orders"
	physbooks "github.com/zhora-ip/libraries-management-system/intenal/storage/sql_storage/repository/postgresql/physical_books"
	"github.com/zhora-ip/libraries-management-system/intenal/storage/sql_storage/repository/postgresql/users"
)

const (
	timeoutCheckExpired = 1
	serverPort          = "8000"
)

func Start(cfg *Config) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err := sqldb.NewDb(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.GetPool().Close()

	bRepo := books.NewBooks(db)
	lRepo := libraries.NewLibraries(db)
	lcRepo := libcards.NewLibCards(db)
	uRepo := users.NewUsers(db)
	pbRepo := physbooks.NewPhysBooks(db)
	oRepo := orders.NewOrders(db)

	bService := bookservice.New(bRepo, db.GetTM())
	uService := userservice.New(uRepo, lcRepo, db.GetTM())
	pbService := physbookservice.New(pbRepo, lRepo, db.GetTM())
	oService := orderservice.New(pbRepo, oRepo, lcRepo, db.GetTM())

	srv := server.New(bService, uService, pbService, oService)

	go runExpiredOrdersCron(ctx, oService)
	err = runServer(srv)
	return err
}

func runServer(srv *server.Server) error {

	errCh := make(chan error, 1)
	go func() {
		log.Print("Server is up and running")
		errCh <- http.ListenAndServe(":"+serverPort, srv)
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

	log.Print("Server exited gracefully")

	return nil
}

func runExpiredOrdersCron(ctx context.Context, oService *orderservice.OrderService) {
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
