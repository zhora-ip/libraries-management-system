package libraries

import "github.com/zhora-ip/libraries-management-system/intenal/storage/sql_storage/db"

type LibrariesRepo struct {
	db db.DB
}

func NewLibraries(database db.DB) *LibrariesRepo {
	return &LibrariesRepo{db: database}
}
