package users

import "github.com/zhora-ip/libraries-management-system/intenal/storage/sql_storage/db"

type UsersRepo struct {
	db db.DB
}

func NewUsers(database db.DB) *UsersRepo {
	return &UsersRepo{db: database}
}
