package orderservice

import (
	"context"
	"errors"
	"strconv"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
)

func (s *OrderService) CheckCanceled(ctx context.Context) error {

	orders, err := s.oRepo.FindCanceled(ctx)
	if err != nil && !errors.Is(err, models.ErrObjectNotFound) {
		return err
	}

	for _, order := range orders {
		err := s.processCanceledOrder(ctx, order)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *OrderService) processCanceledOrder(ctx context.Context, order *models.Order) error {
	if err := s.txManager.RunSerializable(ctx, func(ctxTx context.Context) error {
		err := s.oRepo.MarkAsCanceled(ctxTx, order.ID)
		if err != nil {
			return err
		}

		err = s.pbRepo.MarkAsAvailable(ctxTx, order.PhysicalBookID)
		if err != nil {
			return err
		}

		errCh := make(chan error, 1)
		s.Submit(&models.AuditStatusChange{
			ID:  strconv.Itoa(int(order.ID)),
			New: models.StatusCanceled.String(),
			Old: models.StatusAvailable.String(),
		}, errCh)

		return <-errCh

	}); err != nil && err != models.ErrNoRows {
		return err
	}
	return nil
}
