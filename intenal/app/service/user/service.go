package userservice

import (
	"context"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
)

type usersRepo interface {
	Add(context.Context, *models.User) (int64, error)
}

type txManager interface {
	RunSerializable(context.Context, func(context.Context) error) error
}

type UserService struct {
	usersRepo usersRepo
	txManager txManager
}

func New(repo usersRepo, tm txManager) *UserService {
	return &UserService{
		usersRepo: repo,
		txManager: tm,
	}
}
