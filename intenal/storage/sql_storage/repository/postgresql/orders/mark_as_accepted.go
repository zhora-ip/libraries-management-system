package orders

import (
	"context"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
)

func (r *OrdersRepo) MarkAsAccepted(ctx context.Context, ID int64) error {
	var (
		query = `UPDATE
					orders
				SET
					updated_at = now(),
					status = $1
				WHERE
					id = $2;
		`
	)
	tag, err := r.db.Exec(ctx, query, models.StatusAccepted, ID)
	if tag.RowsAffected() == 0 {
		return models.ErrNoRows
	}
	return err
}
