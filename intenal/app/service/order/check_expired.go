package orderservice

import (
	"context"
	"errors"
	"strconv"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
	ntfs "github.com/zhora-ip/notification-manager/pkg/pb"
)

func (s *OrderService) CheckExpired(ctx context.Context) error {
	orders, err := s.oRepo.FindExpired(ctx)
	if err != nil && !errors.Is(err, models.ErrObjectNotFound) {
		return err
	}

	for _, order := range orders {
		err := s.processExpiredOrder(ctx, order)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *OrderService) processExpiredOrder(ctx context.Context, order *models.Order) error {
	if err := s.txManager.RunSerializable(ctx, func(ctxTx context.Context) error {
		err := s.oRepo.MarkAsExpired(ctxTx, order.ID)
		if err != nil {
			return err
		}

		user, err := s.uRepo.FindByID(ctxTx, order.UserID)
		if err != nil {
			return err
		}

		err = s.nManager.Notify(ctxTx, &ntfs.NotifyRequest{
			Email:   user.Email,
			Name:    user.FullName,
			OrderId: order.ID,
			Type:    ntfs.NotificationType_EXPIRED,
		})
		if err != nil {
			return err
		}

		errCh := make(chan error, 1)
		s.Submit(&models.AuditStatusChange{
			ID:  strconv.Itoa(int(order.ID)),
			New: models.StatusExpired.String(),
			Old: models.StatusIssued.String(),
		}, errCh)

		return <-errCh

	}); err != nil && err != models.ErrNoRows {
		return err
	}
	return nil
}
