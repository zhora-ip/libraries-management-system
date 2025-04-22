package service

import "github.com/zhora-ip/libraries-management-system/intenal/models"

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
}

type FindAllBooksResponse struct {
	Data []*models.Book `json:"data"`
}
