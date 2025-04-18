package libcards

import (
	"context"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
)

func (r *LibCardsRepo) Add(ctx context.Context, libCard *models.LibCard) (int64, error) {
	var ID int64
	err := r.db.ExecQueryRow(
		ctx,
		`INSERT INTO
			lib_cards(
				code,
				user_id
			)
			VALUES($1,$2)
			RETURNING id;`,
		libCard.Code,
		libCard.UserID,
	).Scan(&ID)

	return ID, err
}
