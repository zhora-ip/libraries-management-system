package bookservice

import (
	"context"

	svc "github.com/zhora-ip/libraries-management-system/intenal/models/service"
)

func (s *BookService) FindAll(ctx context.Context, req *svc.FindAllRequest) (*svc.FindAllResponse, error) {
	return s.booksRepo.FindAll(ctx, req)
}
