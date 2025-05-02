package userservice

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
	svc "github.com/zhora-ip/libraries-management-system/intenal/models/service"
)

const (
	tokenTTL = 30 * 24 * time.Hour
)

func (s *UserService) GenerateToken(ctx context.Context, req *svc.GenerateTokenRequest) (*svc.GenerateTokenResponse, error) {

	var (
		user *models.User
	)

	if err := req.Validate(); err != nil {
		return nil, errors.Join(models.ErrValidationFailed, err)
	}

	user, err := s.uRepo.FindByLogin(ctx, req.Login)
	if err != nil {
		return nil, fmt.Errorf("s.uRepo.FindByLogin: %w", err)
	}

	if !user.ComparePassword(req.Password) {
		return nil, models.ErrForbidden
	}

	token, err := s.tkManager.NewJWT(user.ID, int32(user.Role), tokenTTL)
	if err != nil {
		return nil, err
	}

	return &svc.GenerateTokenResponse{
		Token: token,
	}, nil

}
