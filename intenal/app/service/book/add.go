package bookservice

import (
	"context"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
)

func (s *BookService) Add(ctx context.Context, req *AddBookRequest) (*AddBookResponse, error) {
	book := &models.Book{
		Title:       req.Title,
		Author:      req.Author,
		Genre:       req.Genre,
		Description: req.Description,
		AgeLimit:    req.AgeLimit,
	}
	ID, err := s.booksRepo.Add(ctx, book)
	if err != nil {
		return nil, err
	}
	return &AddBookResponse{ID: ID}, nil
}
