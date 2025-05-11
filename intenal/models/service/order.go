package service

import (
	"time"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
)

type AddOrderRequest struct {
	UserID     int64 `json:"user_id"`
	PhysBookID int64 `json:"phys_book_id"`
}

type AddOrderResponse struct {
	ID int64 `json:"id"`
}

type IssueOrderRequest struct {
	ID     int64 `json:"id"`
	UserID int64 `json:"user_id"`
}

type IssueOrderResponse struct {
}

type ReturnOrderRequest struct {
	ID     int64 `json:"id"`
	UserID int64 `json:"user_id"`
}

type ReturnOrderResponse struct {
}

type AcceptOrderRequest struct {
	ID     int64 `json:"id"`
	UserID int64 `json:"user_id"`
}

type AcceptOrderResponse struct {
}

type FindAllOrdersRequest struct {
	Cursor       time.Time `json:"cursor"`
	Limit        int64     `json:"limit"`
	Backward     bool      `json:"backward"`
	ID           *int64    `json:"id"`
	OccurrenceID *int64    `json:"occurrence_id"`
	UserID       *int64    `json:"user_id"`
	LibraryID    *int64    `json:"library_id"`
	BookID       *int64    `json:"phys_book_id"`
}

type FindAllOrdersResponse struct {
	Data []*models.Order `json:"data"`
}
