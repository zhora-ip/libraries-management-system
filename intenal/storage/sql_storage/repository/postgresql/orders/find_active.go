package orders

import (
	"context"
	"errors"

	"github.com/jackc/pgx"
	"github.com/zhora-ip/libraries-management-system/intenal/models"
)

func (r *OrdersRepo) FindActive(ctx context.Context) ([]*models.Order, error) {
	var (
		orders = []*models.Order{}
		query  = `SELECT * FROM orders WHERE status in ($1, $2, $3);`
	)
	err := r.db.Select(
		ctx,
		&orders,
		query,
		models.StatusAvailable,
		models.StatusIssued,
		models.StatusReturned,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, models.ErrObjectNotFound
		}
		return nil, err
	}

	if len(orders) == 0 {
		return nil, models.ErrObjectNotFound
	}

	return orders, nil
}
