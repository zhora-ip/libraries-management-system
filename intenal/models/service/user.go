package service

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type AddUserRequest struct {
	Login       string `json:"login"`
	Password    string `json:"password"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Role        int32  `json:"role"`
}

type AddUserResponse struct {
	ID int64 `json:"id"`
}

type ConfirmEmailRequest struct {
	Token string
}

type ConfirmEmailResponse struct {
	Message string
}

type FindUserByIDRequest struct {
	ID int64 `json:"id"`
}

type FindUserByIDResponse struct {
	Login       string     `json:"login"`
	FullName    string     `json:"full_name"`
	PhoneNumber string     `json:"phone_number"`
	Email       string     `json:"email"`
	Role        string     `json:"role"`
	Code        string     `json:"lib_card"`
	ExpiresAt   *time.Time `json:"expires_at"`
}

type DeleteUserRequest struct {
	ID int64 `json:"id"`
}

type DeleteUserResponse struct {
}

type UpdateUserRequest struct {
	ID                int64   `json:"id"`
	Login             *string `json:"login"`
	Password          *string `json:"password"`
	EncryptedPassword *string `json:"-"`
	FullName          *string `json:"full_name"`
	PhoneNumber       *string `json:"phone_number"`
	Email             *string `json:"email"`
}

type UpdateUserResponse struct {
}

type GenerateTokenRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type GenerateTokenResponse struct {
	Token string `json:"token"`
}

func (g *GenerateTokenRequest) Validate() error {
	return validation.ValidateStruct(g,
		validation.Field(
			&g.Login,
			validation.Required.Error("login is required"),
			validation.Length(5, 100),
		),

		validation.Field(
			&g.Password,
			validation.Required.Error("password is required"),
			validation.Length(6, 100),
		))
}
