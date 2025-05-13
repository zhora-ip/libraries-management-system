package userservice

import (
	"context"

	svc "github.com/zhora-ip/libraries-management-system/intenal/models/service"
	ntfs "github.com/zhora-ip/notification-manager/pkg/pb"
)

func (s *UserService) ConfirmEmail(ctx context.Context, req *svc.ConfirmEmailRequest) (*svc.ConfirmEmailResponse, error) {
	gResp, err := s.nManager.ConfirmEmail(ctx, &ntfs.ConfirmationRequest{Token: req.Token})
	if err != nil {
		return &svc.ConfirmEmailResponse{
			Message: gResp.Message,
		}, err
	}
	err = s.uRepo.MarkAsVerified(ctx, gResp.Email)
	if err != nil {
		return nil, err
	}
	return &svc.ConfirmEmailResponse{}, nil
}
