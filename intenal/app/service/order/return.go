package orderservice

import (
	"context"
	"fmt"
	"strconv"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
	svc "github.com/zhora-ip/libraries-management-system/intenal/models/service"
)

func (s *OrderService) Return(ctx context.Context, req *svc.ReturnOrderRequest) (*svc.ReturnOrderResponse, error) {
	var (
		resp *svc.ReturnOrderResponse
	)

	if err := s.txManager.RunSerializable(ctx, func(ctxTx context.Context) error {
		order, err := s.oRepo.FindByID(ctxTx, req.ID)

		switch {
		case err != nil:
			return fmt.Errorf("s.oRepo.FindByID: %w", err)
		case order.Status != models.StatusExpired && order.Status != models.StatusIssued:
			return models.ErrIncorrectOrderStatus
		case req.UserID != order.UserID:
			return models.ErrForbidden
		}

		err = s.oRepo.MarkAsReturned(ctxTx, order.ID)
		if err != nil {
			return fmt.Errorf("s.oRepo.MarkAsReturned: %w", err)
		}

		audit := &models.AuditStatusChange{
			ID:  strconv.Itoa(int(order.ID)),
			Old: order.Status.String(),
			New: models.StatusReturned.String(),
		}

		errCh := make(chan error, 100)
		s.Submit(audit, errCh)

		return <-errCh

	}); err != nil {
		return nil, fmt.Errorf("s.txManager.RunSerializable: %w", err)
	}

	return resp, nil
}
