package orders

import (
	"context"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
)

func (r *OrdersRepo) Add(ctx context.Context, order *models.Order) (int64, error) {
	var ID int64
	err := r.db.ExecQueryRow(
		ctx,
		`INSERT INTO
			orders(
				library_id,
				physical_book_id,
				user_id,
				status,
				expires_at
			)
			VALUES($1,$2,$3,$4,$5)
			RETURNING id;`,
		order.LibraryID,
		order.PhysicalBookID,
		order.UserID,
		order.Status,
		order.ExpiresAt,
	).Scan(&ID)

	return ID, err
}
