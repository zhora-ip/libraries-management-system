package users

import (
	"context"
	"errors"

	"github.com/jackc/pgconn"
	"github.com/zhora-ip/libraries-management-system/intenal/models"
)

func (r *UsersRepo) Update(ctx context.Context, user *models.User) error {
	var (
		query = `UPDATE
					users
				SET
					login = $1,
					encrypted_password = $2,
					full_name = $3,
					phone_number = $4,
					email = $5
				WHERE
					id = $6
				`
	)
	tag, err := r.db.Exec(
		ctx,
		query,
		user.Login,
		user.EncryptedPassword,
		user.FullName,
		user.PhoneNumber,
		user.Email,
		user.ID,
	)

	if tag.RowsAffected() == 0 {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			switch pgErr.ConstraintName {
			case "users_login_key":
				return models.ErrLoginAlreadyExists
			case "users_phone_number_key":
				return models.ErrPhoneNumberAlreadyExists
			case "users_email_key":
				return models.ErrEmailAlreadyExists
			}
		}
		return models.ErrObjectNotFound
	}
	return err
}
