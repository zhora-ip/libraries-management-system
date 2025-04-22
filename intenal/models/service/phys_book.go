package service

import "github.com/zhora-ip/libraries-management-system/intenal/models"

type BookAvailability struct {
	Library     *models.Library `json:"library"`
	PhysBookIDs []int64         `json:"phys_book_ids"`
	Amount      int64           `json:"amount"`
}

type FindPBookByBookIDResponse struct {
	Data []*BookAvailability `json:"data"`
}
