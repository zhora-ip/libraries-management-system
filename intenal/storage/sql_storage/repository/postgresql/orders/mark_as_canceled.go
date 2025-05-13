package orders

import (
	"context"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
)

func (r *OrdersRepo) MarkAsCanceled(ctx context.Context, ID int64) error {
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
	tag, err := r.db.Exec(ctx, query, models.StatusCanceled, ID)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return models.ErrNoRows
	}
	return err
}
