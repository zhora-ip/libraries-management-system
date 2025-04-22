package users

import (
	"context"
	"errors"

	"github.com/jackc/pgconn"
	"github.com/zhora-ip/libraries-management-system/intenal/models"
)

func (r *UsersRepo) Add(ctx context.Context, user *models.User) (int64, error) {
	var ID int64

	err := r.db.ExecQueryRow(
		ctx,
		`INSERT INTO
			users(
				login,
				encrypted_password,
				full_name,
				phone_number,
				email,
				role
			)
			VALUES($1,$2,$3,$4,$5,$6)
			RETURNING id;`,
		user.Login,
		user.EncryptedPassword,
		user.FullName,
		user.PhoneNumber,
		user.Email,
		user.Role,
	).Scan(&ID)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			switch pgErr.ConstraintName {
			case "users_login_key":
				return 0, models.ErrLoginAlreadyExists
			case "users_phone_number_key":
				return 0, models.ErrPhoneNumberAlreadyExists
			case "users_email_key":
				return 0, models.ErrEmailAlreadyExists
			}
		}
	}

	return ID, err
}
