package models

import "time"

type Book struct {
	ID          int64      `json:"id" db:"id"`
	Title       string     `json:"title" db:"title"`
	Author      string     `json:"author" db:"author"`
	Genre       string     `json:"genre" db:"genre"`
	Description string     `json:"description" db:"description"`
	AgeLimit    int32      `json:"age_limit" db:"age_limit"`
	Created_at  *time.Time `json:"created_at" db:"created_at"`
	Updated_at  *time.Time `json:"updated_at" db:"updated_at"`
}
