package userservice

import (
	"context"
	"errors"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
	svc "github.com/zhora-ip/libraries-management-system/intenal/models/service"
)

func (s *UserService) Add(ctx context.Context, req *svc.AddUserRequest) (*svc.AddUserResponse, error) {

	user := &models.User{
		Login:       req.Login,
		Password:    req.Password,
		FullName:    req.FullName,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
	}

	if err := user.Validate(); err != nil {
		return nil, errors.Join(models.ErrValidationFailed, err)
	}

	if err := user.BeforeCreate(); err != nil {
		return nil, errors.Join(models.ErrEncryptionFailed, err)
	}

	user.Sanitize()

	ID, err := s.usersRepo.Add(ctx, user)
	if err != nil {
		return nil, err
	}
	return &svc.AddUserResponse{ID: ID}, nil
}
