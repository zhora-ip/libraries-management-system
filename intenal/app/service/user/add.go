package userservice

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
	svc "github.com/zhora-ip/libraries-management-system/intenal/models/service"
	ntfs "github.com/zhora-ip/notification-manager/pkg/pb"
)

const (
	expirationTime = 24 * 365
)

func (s *UserService) Add(ctx context.Context, req *svc.AddUserRequest) (*svc.AddUserResponse, error) {
	var (
		resp      *svc.AddUserResponse
		expiresAt = time.Now().Add(time.Hour * expirationTime)
	)

	if err := s.txManager.RunSerializable(ctx, func(ctxTx context.Context) error {

		user := &models.User{
			Login:       req.Login,
			Password:    req.Password,
			FullName:    req.FullName,
			PhoneNumber: req.PhoneNumber,
			Email:       req.Email,
			Role:        models.UserRole(req.Role),
		}

		if err := user.Validate(); err != nil {
			return errors.Join(models.ErrValidationFailed, err)
		}

		if err := user.BeforeCreate(); err != nil {
			return errors.Join(models.ErrEncryptionFailed, err)
		}

		user.Sanitize()

		ID, err := s.uRepo.Add(ctxTx, user)
		if err != nil {
			return fmt.Errorf("s.uRepo.Add: %w", err)
		}

		_, err = s.lcRepo.Add(ctxTx, &models.LibCard{
			Code:      "code" + req.Login,
			UserID:    ID,
			ExpiresAt: &expiresAt,
		})
		if err != nil {
			return fmt.Errorf("s.lcRepo.Add: %w", err)
		}

		resp = &svc.AddUserResponse{ID: ID}
		return nil
	}); err != nil {
		return nil, fmt.Errorf("s.txManager.RunSerializable: %w", err)
	}

	gResp, err := s.nManager.VerifyEmail(ctx, &ntfs.VerifyEmailRequest{
		Email: req.Email,
	})
	if !gResp.Success {
		log.Print(gResp.Message, err)
	}

	return resp, nil

}
