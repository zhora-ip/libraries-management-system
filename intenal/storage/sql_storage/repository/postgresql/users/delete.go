package users

import (
	"context"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
)

func (r *UsersRepo) Delete(ctx context.Context, ID int64) error {
	tag, err := r.db.Exec(ctx, `DELETE FROM users WHERE id = $1`, ID)

	if tag.RowsAffected() == 0 {
		return models.ErrObjectNotFound
	}

	return err
}
