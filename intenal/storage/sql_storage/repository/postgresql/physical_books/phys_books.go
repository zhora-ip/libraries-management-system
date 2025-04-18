package physbooks

import "github.com/zhora-ip/libraries-management-system/intenal/storage/sql_storage/db"

type PhysBooksRepo struct {
	db db.DB
}

func NewPhysBooks(database db.DB) *PhysBooksRepo {
	return &PhysBooksRepo{db: database}
}
