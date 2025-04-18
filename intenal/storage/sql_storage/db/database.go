package db

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	txmanager "github.com/zhora-ip/libraries-management-system/intenal/storage/sql_storage/tx_manager"

	"github.com/jackc/pgconn"
)

// DB defines an interface for database operations.
type DB interface {
	// Get retrieves a single database record and maps it to the provided destination.
	Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error

	// Select retrieves multiple database records and maps them to the provided destination.
	Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error

	// Exec executes an SQL statement and returns the result.
	Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error)

	// ExecQueryRow executes an SQL query that is expected to return a single row.
	ExecQueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row
}

// Database implements the DB interface, providing access to a database cluster.
type Database struct {
	txManager *txmanager.TxManager
}

func newDatabase(tm *txmanager.TxManager) *Database {
	return &Database{txManager: tm}
}

// GetPool retrieves the current database connection pool.
func (db Database) GetPool() *pgxpool.Pool {
	return db.txManager.GetQuerier(context.Background()).(*pgxpool.Pool)
}

// GetTM retrieves the transaction manager.
func (db Database) GetTM() *txmanager.TxManager {
	return db.txManager
}

// Get retrieves a single database record and maps it to the provided destination.
func (db Database) Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return pgxscan.Get(ctx, db.txManager.GetQuerier(ctx), dest, query, args...)
}

// Select retrieves multiple database records and maps them to the provided destination
func (db Database) Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return pgxscan.Select(ctx, db.txManager.GetQuerier(ctx), dest, query, args...)
}

// Exec executes an SQL statement and returns the result.
func (db Database) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	return db.txManager.GetQuerier(ctx).Exec(ctx, query, args...)
}

// ExecQueryRow executes an SQL query that is expected to return a single row.
func (db Database) ExecQueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return db.txManager.GetQuerier(ctx).QueryRow(ctx, query, args...)
}
