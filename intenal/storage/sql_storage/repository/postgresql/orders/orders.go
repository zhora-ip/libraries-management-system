package orders

import "github.com/zhora-ip/libraries-management-system/intenal/storage/sql_storage/db"

type OrdersRepo struct {
	db db.DB
}

func NewOrders(database db.DB) *OrdersRepo {
	return &OrdersRepo{db: database}
}
