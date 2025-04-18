package txmanager

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type keyTx struct{}

type TxManager struct {
	cluster *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *TxManager {
	return &TxManager{
		cluster: pool,
	}
}
