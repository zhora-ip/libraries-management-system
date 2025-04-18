package libcards

import "github.com/zhora-ip/libraries-management-system/intenal/storage/sql_storage/db"

type LibCardsRepo struct {
	db db.DB
}

func NewLibCards(database db.DB) *LibCardsRepo {
	return &LibCardsRepo{db: database}
}
