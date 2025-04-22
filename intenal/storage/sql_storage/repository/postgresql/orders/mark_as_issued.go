package orders

import (
	"context"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
)

func (r *OrdersRepo) MarkAsIssued(ctx context.Context, order *models.Order) error {
	var (
		query = `UPDATE
					orders
				SET
					status = $1,
					updated_at = now(),
					expires_at = $2
				WHERE
					id = $3;
		`
	)
	tag, err := r.db.Exec(ctx, query, models.StatusIssued, order.ExpiresAt, order.ID)

	if tag.RowsAffected() == 0 {
		return models.ErrNoRows
	}
	return err
}
