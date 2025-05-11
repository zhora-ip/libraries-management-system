package service

import "github.com/zhora-ip/libraries-management-system/intenal/models"

type FindLibCardRequest struct {
	UserID int64 `json:"user_id"`
}

type FindLibCardResponse struct {
	Card *models.LibCard `json:"card"`
}

type ExtendLibCardRequest struct {
	UserID int64 `json:"user_id"`
}
