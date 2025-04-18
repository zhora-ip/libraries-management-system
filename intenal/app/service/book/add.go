package bookservice

import (
	"context"
	"errors"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
	svc "github.com/zhora-ip/libraries-management-system/intenal/models/service"
)

func (s *BookService) Add(ctx context.Context, req *svc.AddBookRequest) (*svc.AddBookResponse, error) {

	book := &models.Book{
		Title:       req.Title,
		Author:      req.Author,
		Genre:       req.Genre,
		Description: req.Description,
		AgeLimit:    req.AgeLimit,
	}

	if err := book.Validate(); err != nil {
		return nil, errors.Join(models.ErrValidationFailed, err)
	}

	ID, err := s.booksRepo.Add(ctx, book)
	if err != nil {
		return nil, err
	}
	return &svc.AddBookResponse{ID: ID}, nil
}
