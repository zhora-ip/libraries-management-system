package userservice

import (
	"context"
	"fmt"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
	svc "github.com/zhora-ip/libraries-management-system/intenal/models/service"
)

func (s *UserService) FindByID(ctx context.Context, req *svc.FindUserByIDRequest) (*svc.FindUserByIDResponse, error) {
	var (
		user *models.User
		card *models.LibCard
		err  error
	)

	if err := s.txManager.RunSerializable(ctx, func(ctxTx context.Context) error {
		user, err = s.uRepo.FindByID(ctxTx, req.ID)
		if err != nil {
			return fmt.Errorf("s.uRepo.FindByID: %w", err)
		}

		card, err = s.lcRepo.FindByUserID(ctxTx, req.ID)
		if err != nil {
			return fmt.Errorf("s.lcRepo.FindByUserID: %w", err)
		}

		return nil
	}); err != nil {
		return nil, fmt.Errorf("s.txManager.RunSerializable: %w", err)
	}

	return &svc.FindUserByIDResponse{
		Login:       user.Login,
		FullName:    user.FullName,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Role:        user.Role.String(),
		Code:        card.Code,
		ExpiresAt:   card.ExpiresAt,
	}, nil

}
