package orderservice

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
	svc "github.com/zhora-ip/libraries-management-system/intenal/models/service"
)

const (
	expirationTimeIssue = 24 * 7
)

func (s *OrderService) Issue(ctx context.Context, req *svc.IssueOrderRequest) (*svc.IssueOrderResponse, error) {
	var (
		resp      *svc.IssueOrderResponse
		expiresAt = time.Now().Add(expirationTimeIssue)
	)
	if err := s.txManager.RunSerializable(ctx, func(ctxTx context.Context) error {
		order, err := s.oRepo.FindByID(ctxTx, req.ID)

		switch {
		case err != nil:
			return fmt.Errorf("s.oRepo.FindByID: %w", err)
		case order.Status != models.StatusAvailable:
			return models.ErrIncorrectOrderStatus
		case req.UserID != order.UserID:
			return models.ErrForbidden
		case order.ExpiresAt.Before(time.Now()):
			return models.ErrAlreadyExpired
		}

		order.ExpiresAt = &expiresAt
		err = s.oRepo.MarkAsIssued(ctxTx, order)
		if err != nil {
			return fmt.Errorf("s.oRepo.MarkAsIssued: %w", err)
		}

		audit := &models.AuditStatusChange{
			ID:  strconv.Itoa(int(order.ID)),
			Old: order.Status.String(),
			New: models.StatusIssued.String(),
		}

		errCh := make(chan error, 1)
		s.Submit(audit, errCh)

		return <-errCh

	}); err != nil {
		return nil, fmt.Errorf("s.txManager.RunSerializable: %w", err)
	}

	return resp, nil
}
