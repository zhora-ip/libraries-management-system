package libcards

import (
	"context"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
)

func (r *LibCardsRepo) DeleteByUserID(ctx context.Context, ID int64) error {

	tag, err := r.db.Exec(ctx, `DELETE FROM lib_cards WHERE user_id = $1`, ID)
	if tag.RowsAffected() == 0 {
		return models.ErrObjectNotFound
	}

	return err
}
