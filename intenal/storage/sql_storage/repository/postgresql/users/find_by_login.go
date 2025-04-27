package users

import (
	"context"
	"strings"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
)

func (r *UsersRepo) FindByLogin(ctx context.Context, login string) (*models.User, error) {
	var (
		user = &models.User{}
	)

	err := r.db.Get(ctx, user, `SELECT * FROM users WHERE login = $1`, login)

	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return nil, models.ErrObjectNotFound
		}
		return nil, err
	}

	return user, nil
}
