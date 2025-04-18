package db

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	txmanager "github.com/zhora-ip/libraries-management-system/intenal/storage/sql_storage/tx_manager"
)

type Config struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"ssl_mode"`
}

// NewDb initializes a new Database instance with a connection pool based on the provided database URL.
func NewDb(ctx context.Context, databaseURL string) (*Database, error) {
	pool, err := pgxpool.Connect(ctx, databaseURL)
	if err != nil {
		return nil, err
	}

	return newDatabase(txmanager.New(pool)), nil
}
