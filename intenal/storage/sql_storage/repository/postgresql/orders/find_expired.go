package orders

import (
	"context"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
)

func (r *OrdersRepo) FindExpired(ctx context.Context) ([]*models.Order, error) {
	var (
		orders []*models.Order
		query  = `SELECT
					 * 
				FROM
					orders
				WHERE 
					expires_at < now() AND
					status = $1;
		`
	)

	err := r.db.Select(ctx, &orders, query, models.StatusIssued)
	if len(orders) == 0 {
		return nil, models.ErrObjectNotFound
	}
	if err != nil {
		return nil, err
	}

	return orders, nil
}
