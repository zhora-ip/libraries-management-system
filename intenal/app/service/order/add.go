package orderservice

import (
	"context"
	"fmt"
	"time"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
	svc "github.com/zhora-ip/libraries-management-system/intenal/models/service"
)

const (
	expirationTime = 24
)

func (s *OrderService) Add(ctx context.Context, req *svc.AddOrderRequest) (*svc.AddOrderResponse, error) {
	var (
		resp      *svc.AddOrderResponse
		expiresAt = time.Now().Add(time.Hour * expirationTime)
	)

	if err := s.txManager.RunSerializable(ctx, func(ctxTx context.Context) error {

		ID, err := s.addHelper(ctxTx, req, expiresAt)
		if err != nil {
			return err
		}

		resp = &svc.AddOrderResponse{ID: ID}
		return nil
	}); err != nil {
		return nil, fmt.Errorf("s.txManager.RunSerializable: %w", err)
	}
	return resp, nil
}

func (s *OrderService) addHelper(ctxTx context.Context, req *svc.AddOrderRequest, expiresAt time.Time) (int64, error) {
	lCard, err := s.lcRepo.FindByUserID(ctxTx, req.UserID)
	if err != nil {
		return 0, fmt.Errorf("s.lcRepo.FindByUserID: %w", err)
	}

	if lCard.ExpiresAt.Before(time.Now()) {
		return 0, models.ErrLibCardExpired
	}

	book, err := s.pbRepo.FindByID(ctxTx, req.PhysBookID)
	if err != nil {
		return 0, fmt.Errorf("s.pbRepo.FindByID: %w", err)
	}

	err = s.pbRepo.MarkAsUnavailable(ctxTx, req.PhysBookID)
	if err != nil {
		return 0, fmt.Errorf("s.pbRepo.MarkAsUnavailable: %w", err)
	}

	order := &models.Order{
		LibraryID:      book.LibraryID,
		PhysicalBookID: book.ID,
		UserID:         req.UserID,
		Status:         models.StatusAvailable,
		ExpiresAt:      &expiresAt,
	}

	ID, err := s.oRepo.Add(ctxTx, order)
	if err != nil {
		return 0, fmt.Errorf("s.oRepo.Add: %w", err)
	}

	return ID, nil
}
