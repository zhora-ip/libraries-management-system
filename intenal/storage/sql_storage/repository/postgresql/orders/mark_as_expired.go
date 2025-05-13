package orders

import (
	"context"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
)

func (r *OrdersRepo) MarkAsExpired(ctx context.Context, ID int64) error {
	var (
		query = `UPDATE
					orders
				SET
					status = $1,
					updated_at = now()
				WHERE
					id = $2;
		`
	)
	tag, err := r.db.Exec(ctx, query, models.StatusExpired, ID)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return models.ErrNoRows
	}
	return err
}
