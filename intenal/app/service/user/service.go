package userservice

import (
	"context"
	"time"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
)

type usersRepo interface {
	Add(context.Context, *models.User) (int64, error)
	FindByLogin(context.Context, string) (*models.User, error)
}

type libCardsRepo interface {
	Add(context.Context, *models.LibCard) (int64, error)
}

type txManager interface {
	RunSerializable(context.Context, func(context.Context) error) error
}

type tkManager interface {
	NewJWT(int64, int32, time.Duration) (string, error)
}

type UserService struct {
	uRepo     usersRepo
	lcRepo    libCardsRepo
	txManager txManager
	tkManager tkManager
}

func New(uRepo usersRepo, lcRepo libCardsRepo, tm txManager, tkm tkManager) *UserService {
	return &UserService{
		uRepo:     uRepo,
		lcRepo:    lcRepo,
		txManager: tm,
		tkManager: tkm,
	}
}
