package orderservice

import (
	"context"
	"fmt"
	"strconv"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
	svc "github.com/zhora-ip/libraries-management-system/intenal/models/service"
)

func (s *OrderService) Accept(ctx context.Context, req *svc.AcceptOrderRequest) (*svc.AcceptOrderResponse, error) {
	var (
		resp *svc.AcceptOrderResponse
	)
	if err := s.txManager.RunSerializable(ctx, func(ctxTx context.Context) error {
		order, err := s.oRepo.FindByID(ctxTx, req.ID)

		switch {
		case err != nil:
			return fmt.Errorf("s.oRepo.FindByID: %w", err)
		case order.Status != models.StatusReturned:
			return models.ErrIncorrectOrderStatus
		}

		err = s.oRepo.MarkAsAccepted(ctxTx, req.ID)
		if err != nil {
			return fmt.Errorf("s.oRepo.MarkAsAccepted: %w", err)
		}

		audit := &models.AuditStatusChange{
			ID:  strconv.Itoa(int(order.ID)),
			Old: order.Status.String(),
			New: models.StatusAccepted.String(),
		}

		errCh := make(chan error, 1)
		s.Submit(audit, errCh)

		return <-errCh

	}); err != nil {
		return nil, fmt.Errorf("s.txManager.RunSerializable: %w", err)
	}

	return resp, nil
}
