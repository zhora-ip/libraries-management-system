package userservice

import (
	"context"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
)

type usersRepo interface {
	Add(context.Context, *models.User) (int64, error)
}

type libCardsRepo interface {
	Add(context.Context, *models.LibCard) (int64, error)
}

type txManager interface {
	RunSerializable(context.Context, func(context.Context) error) error
}

type UserService struct {
	uRepo     usersRepo
	lcRepo    libCardsRepo
	txManager txManager
}

func New(uRepo usersRepo, lcRepo libCardsRepo, tm txManager) *UserService {
	return &UserService{
		uRepo:     uRepo,
		lcRepo:    lcRepo,
		txManager: tm,
	}
}
