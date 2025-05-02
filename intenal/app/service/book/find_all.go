package bookservice

import (
	"context"

	svc "github.com/zhora-ip/libraries-management-system/intenal/models/service"
)

func (s *BookService) FindAll(ctx context.Context, req *svc.FindAllBooksRequest) (*svc.FindAllBooksResponse, error) {
	books, err := s.booksRepo.FindAll(ctx, req)
	if err != nil {
		return nil, err
	}

	return &svc.FindAllBooksResponse{Data: books}, nil
}
