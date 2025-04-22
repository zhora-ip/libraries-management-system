package service

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
