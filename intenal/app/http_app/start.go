package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/zhora-ip/libraries-management-system/intenal/app/http_app/server"
	bookservice "github.com/zhora-ip/libraries-management-system/intenal/app/service/book"
	userservice "github.com/zhora-ip/libraries-management-system/intenal/app/service/user"
	sqldb "github.com/zhora-ip/libraries-management-system/intenal/storage/sql_storage/db"
	"github.com/zhora-ip/libraries-management-system/intenal/storage/sql_storage/repository/postgresql/books"
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
	bService := bookservice.New(bRepo, db.GetTM())

	uRepo := users.NewUsers(db)
	uService := userservice.New(uRepo, db.GetTM())

	srv := server.New(bService, uService)
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
