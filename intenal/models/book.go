package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Book struct {
	ID          int64      `json:"id" db:"id"`
	Title       string     `json:"title" db:"title"`
	Author      string     `json:"author" db:"author"`
	Genre       string     `json:"genre" db:"genre"`
	Description string     `json:"description" db:"description"`
	AgeLimit    int32      `json:"age_limit" db:"age_limit"`
	CreatedAt   *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at" db:"updated_at"`
}

func (b *Book) Validate() error {
	return validation.ValidateStruct(b,
		validation.Field(
			&b.Title,
			validation.Required.Error("title is required"),
		),

		validation.Field(
			&b.Author,
			validation.Required.Error("author is required"),
		),

		validation.Field(
			&b.Genre,
			validation.Required.Error("genre is required"),
		),

		validation.Field(
			&b.Description,
			validation.Required.Error("description required"),
		),
	)
}
