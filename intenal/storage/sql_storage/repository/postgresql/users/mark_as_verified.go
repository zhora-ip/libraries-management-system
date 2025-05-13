package users

import (
	"context"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
)

func (r *UsersRepo) MarkAsVerified(ctx context.Context, email string) error {
	tag, err := r.db.Exec(ctx, "UPDATE users SET verified = true WHERE email = $1;", email)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return models.ErrObjectNotFound
	}

	return nil
}
