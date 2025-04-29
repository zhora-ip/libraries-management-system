package users

import (
	"context"
	"strings"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
)

func (r *UsersRepo) FindByID(ctx context.Context, ID int64) (*models.User, error) {
	var (
		user = &models.User{}
	)
	err := r.db.Get(ctx, user, `SELECT * FROM users WHERE id = $1`, ID)

	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return nil, models.ErrObjectNotFound
		}
		return nil, err
	}
	return user, nil
}
