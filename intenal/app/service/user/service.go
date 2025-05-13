package userservice

import (
	"context"
	"time"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
	ntfs "github.com/zhora-ip/notification-manager/pkg/pb"
)

type usersRepo interface {
	Add(context.Context, *models.User) (int64, error)
	FindByLogin(context.Context, string) (*models.User, error)
	FindByID(context.Context, int64) (*models.User, error)
	Delete(context.Context, int64) error
	Update(context.Context, *models.User) error
	MarkAsVerified(context.Context, string) error
}

type libCardsRepo interface {
	Add(context.Context, *models.LibCard) (int64, error)
	FindByUserID(context.Context, int64) (*models.LibCard, error)
	DeleteByUserID(context.Context, int64) error
	Extend(context.Context, int64, time.Time) error
}

type ordersRepo interface {
	FindBlockedByUserID(context.Context, int64) ([]*models.Order, error)
}

type txManager interface {
	RunSerializable(context.Context, func(context.Context) error) error
}

type tkManager interface {
	NewJWT(int64, int32, time.Duration) (string, error)
}

type nManager interface {
	VerifyEmail(context.Context, *ntfs.VerifyEmailRequest) (*ntfs.VerifyEmailResponse, error)
	ConfirmEmail(ctx context.Context, req *ntfs.ConfirmationRequest) (*ntfs.ConfirmationResponse, error)
}

type UserService struct {
	uRepo     usersRepo
	lcRepo    libCardsRepo
	oRepo     ordersRepo
	txManager txManager
	tkManager tkManager
	nManager  nManager
}

func New(uRepo usersRepo, lcRepo libCardsRepo, oRepo ordersRepo,
	tm txManager, tkm tkManager, nm nManager) *UserService {
	return &UserService{
		uRepo:     uRepo,
		lcRepo:    lcRepo,
		oRepo:     oRepo,
		txManager: tm,
		tkManager: tkm,
		nManager:  nm,
	}
}
