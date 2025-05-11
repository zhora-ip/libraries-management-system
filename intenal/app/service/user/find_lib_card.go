package userservice

import (
	"context"

	svc "github.com/zhora-ip/libraries-management-system/intenal/models/service"
)

func (s *UserService) FindLibCard(ctx context.Context, req *svc.FindLibCardRequest) (*svc.FindLibCardResponse, error) {
	card, err := s.lcRepo.FindByUserID(ctx, req.UserID)
	if err != nil {
		return nil, err
	}
	return &svc.FindLibCardResponse{Card: card}, nil
}
