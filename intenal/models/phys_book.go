package models

import "time"

type PhysBook struct {
	ID          int64      `db:"id"`
	LibraryID   int64      `db:"library_id"`
	BookID      int64      `db:"book_id"`
	IsAvailable bool       `db:"is_available"`
	CreatedAt   *time.Time `db:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at"`
}
