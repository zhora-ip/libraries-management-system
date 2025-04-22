package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

type UserRole int32

const (
	UserRoleUnknown   UserRole = iota // 0
	UserRoleAdmin                     // 1
	UserRoleLibrarian                 // 2
	UserRoleReader                    // 3
	Admin             = "ADMIN"
	Librarian         = "LIBRARIAN"
	Reader            = "READER"
)

func (u UserRole) String() string {
	return [...]string{
		Unknown,
		Admin,
		Librarian,
		Reader,
	}[u]
}

type User struct {
	ID                int64      `db:"id"`
	Login             string     `db:"login"`
	Password          string     `db:"-"`
	EncryptedPassword string     `db:"encrypted_password"`
	FullName          string     `db:"full_name"`
	PhoneNumber       string     `db:"phone_number"`
	Email             string     `db:"email"`
	Role              UserRole   `db:"role"`
	CreatedAt         *time.Time `db:"created_at"`
	UpdatedAt         *time.Time `db:"updated_at"`
}

func (u *User) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(
			&u.Login,
			validation.Required.Error("login is required"),
			validation.Length(5, 100),
		),

		validation.Field(
			&u.Password,
			validation.By(RequiredIf(u.EncryptedPassword == "")),
			validation.Length(6, 100),
		),

		validation.Field(
			&u.FullName,
			validation.Required.Error("fullname is required"),
		),

		validation.Field(
			&u.PhoneNumber,
			validation.Required.Error("phone_number required"),
			is.E164,
		),

		validation.Field(
			&u.Email,
			validation.Required.Error("email required"),
			is.Email,
		),
	)
}

func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		enc, err := encryptString(u.Password)
		if err != nil {
			return err
		}
		u.EncryptedPassword = enc
	}
	return nil
}

func (u *User) Sanitize() {
	u.Password = ""
}

func (u *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password)) == nil
}

func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)

	if err != nil {
		return "", err
	}

	return string(b), nil
}
