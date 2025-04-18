package txmanager

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type QueryEngine interface {
	Exec(ctx context.Context, query string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row
}

func (tm *TxManager) GetQuerier(ctx context.Context) QueryEngine {
	tx, ok := ctx.Value(keyTx{}).(QueryEngine)
	if ok && tx != nil {
		return tx
	}

	return tm.cluster
}
