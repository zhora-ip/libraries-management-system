package service

import (
	"time"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
)

type AddBookRequest struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Genre       string `json:"genre"`
	Description string `json:"description"`
	AgeLimit    int32  `json:"age_limit"`
}

type AddBookResponse struct {
	ID int64 `json:"id"`
}

type FindAllBooksRequest struct {
	Cursor   time.Time `json:"cursor"`
	Limit    int64     `json:"limit"`
	Backward bool      `json:"backward"`
	ID       *int64    `json:"id"`
	Title    *string   `json:"title"`
	Author   *string   `json:"author"`
	Genre    *string   `json:"genre"`
	AgeLimit *int32    `json:"age_limit"`
}

type FindAllBooksResponse struct {
	Data []*models.Book `json:"data"`
}
