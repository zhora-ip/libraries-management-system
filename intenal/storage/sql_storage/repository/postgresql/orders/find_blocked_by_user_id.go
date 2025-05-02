package orders

import (
	"context"
	"errors"

	"github.com/jackc/pgx"
	"github.com/zhora-ip/libraries-management-system/intenal/models"
)

func (r *OrdersRepo) FindBlockedByUserID(ctx context.Context, ID int64) ([]*models.Order, error) {
	var (
		orders = []*models.Order{}
		query  = `SELECT * FROM orders WHERE user_id = $1 AND status in ($2, $3)`
	)

	err := r.db.Select(ctx, &orders, query, ID, models.StatusIssued, models.StatusExpired)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	if len(orders) == 0 {
		return nil, nil
	}

	return orders, nil
}
