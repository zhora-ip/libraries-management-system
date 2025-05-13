package userservice

import (
	"context"
	"log"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
	svc "github.com/zhora-ip/libraries-management-system/intenal/models/service"
	ntfs "github.com/zhora-ip/notification-manager/pkg/pb"
)

func (s *UserService) Update(ctx context.Context, req *svc.UpdateUserRequest) (*svc.UpdateUserResponse, error) {
	if err := s.txManager.RunSerializable(ctx, func(ctxTx context.Context) error {
		user, err := s.uRepo.FindByID(ctxTx, req.ID)
		if err != nil {
			return err
		}

		if err := formUpdateUserReq(req, user); err != nil {
			return err
		}

		if err := s.uRepo.Update(ctxTx, user); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	if req.Email != nil {

		gResp, err := s.nManager.VerifyEmail(ctx, &ntfs.VerifyEmailRequest{
			Email: *req.Email,
		})
		if !gResp.Success {
			log.Print(gResp.Message, err)
		}

	}

	return nil, nil
}

func formUpdateUserReq(req *svc.UpdateUserRequest, user *models.User) error {
	if req.Login != nil {
		user.Login = *req.Login
	}

	if req.FullName != nil {
		user.FullName = *req.FullName
	}

	if req.PhoneNumber != nil {
		user.PhoneNumber = *req.PhoneNumber
	}

	if req.Email != nil {
		user.Email = *req.Email
	}

	if req.Password != nil {
		if user.ComparePassword(*req.Password) {
			return models.ErrRepeatedPassword
		}
		user.Password = *req.Password
	}

	if err := user.Validate(); err != nil {
		return err
	}

	if err := user.BeforeCreate(); err != nil {
		return err
	}

	user.Sanitize()
	return nil
}
