package userservice

import (
	"context"
	"time"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
	svc "github.com/zhora-ip/libraries-management-system/intenal/models/service"
)

const (
	extensionTime = 365 * 24 * time.Hour
)

func (s *UserService) ExtendLibCard(ctx context.Context, req *svc.ExtendLibCardRequest) error {
	if err := s.txManager.RunSerializable(ctx, func(ctxTx context.Context) error {
		card, err := s.lcRepo.FindByUserID(ctxTx, req.UserID)
		if err != nil {
			return err
		}

		if card.ExpiresAt.After(time.Now()) {
			return models.ErrCardNotExpired
		}

		err = s.lcRepo.Extend(ctxTx, card.ID, time.Now().Add(extensionTime))
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}
