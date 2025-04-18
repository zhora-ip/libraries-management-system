package txmanager

import (
	"context"

	"github.com/jackc/pgx/v4"
)

func (tm *TxManager) RunSerializable(ctx context.Context, action func(context.Context) error) error {

	ops := pgx.TxOptions{
		IsoLevel:   pgx.Serializable,
		AccessMode: pgx.ReadWrite,
	}

	return tm.beginTx(ctx, ops, action)
}

func (tm *TxManager) beginTx(ctx context.Context, ops pgx.TxOptions, action func(context.Context) error) error {

	tx, err := tm.cluster.BeginTx(ctx, ops)
	if err != nil {
		return err
	}

	ctx = context.WithValue(ctx, keyTx{}, tx)

	defer tx.Rollback(ctx)

	if err := action(ctx); err != nil {
		return err
	}

	return tx.Commit(ctx)
}
