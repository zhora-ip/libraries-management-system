package books

import "github.com/zhora-ip/libraries-management-system/intenal/storage/sql_storage/db"

// BooksRepo represents the repository for managing books data in the database.
type BooksRepo struct {
	db db.DB
}

// NewBooks creates and initializes a new instance of BooksRepo with the given database connection.
func NewBooks(database db.DB) *BooksRepo {
	return &BooksRepo{db: database}
}
