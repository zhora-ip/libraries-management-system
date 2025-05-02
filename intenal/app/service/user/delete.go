package userservice

import (
	"context"
	"errors"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
	svc "github.com/zhora-ip/libraries-management-system/intenal/models/service"
)

func (s *UserService) Delete(ctx context.Context, req *svc.DeleteUserRequest) (*svc.DeleteUserResponse, error) {
	if err := s.txManager.RunSerializable(ctx, func(ctxTx context.Context) error {
		orders, err := s.oRepo.FindBlockedByUserID(ctxTx, req.ID)
		if err != nil {
			return err
		}

		if len(orders) != 0 {
			return models.ErrPendingOrders
		}

		err = s.lcRepo.DeleteByUserID(ctxTx, req.ID)
		if err != nil && !errors.Is(err, models.ErrObjectNotFound) {
			return err
		}

		if err := s.uRepo.Delete(ctxTx, req.ID); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}
	return nil, nil
}
