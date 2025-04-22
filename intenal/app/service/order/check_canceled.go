package orderservice

import (
	"context"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
)

func (s *OrderService) CheckCanceled(ctx context.Context) error {

	orders, err := s.oRepo.FindCanceled(ctx)
	if err != nil && err != models.ErrObjectNotFound {
		return err
	}

	for _, order := range orders {
		if err := s.txManager.RunSerializable(ctx, func(ctxTx context.Context) error {
			err := s.oRepo.MarkAsCanceled(ctxTx, order.ID)
			if err != nil {
				return err
			}

			return s.pbRepo.MarkAsAvailable(ctxTx, order.PhysicalBookID)

		}); err != nil && err != models.ErrNoRows {
			return err
		}
	}
	return nil
}
