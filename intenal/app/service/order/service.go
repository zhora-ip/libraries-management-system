package orderservice

import (
	"context"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
)

type physBooksRepo interface {
	FindByID(context.Context, int64) (*models.PhysBook, error)
	MarkAsUnavailable(context.Context, int64) error
	MarkAsAvailable(context.Context, int64) error
}

type ordersRepo interface {
	Add(context.Context, *models.Order) (int64, error)
	FindByID(context.Context, int64) (*models.Order, error)
	MarkAsIssued(context.Context, *models.Order) error
	FindCanceled(context.Context) ([]*models.Order, error)
	MarkAsCanceled(context.Context, int64) error
	MarkAsReturned(context.Context, int64) error
}

type libCardsRepo interface {
	FindByUserID(context.Context, int64) (*models.LibCard, error)
}

type txManager interface {
	RunSerializable(context.Context, func(context.Context) error) error
}

type OrderService struct {
	pbRepo    physBooksRepo
	oRepo     ordersRepo
	lcRepo    libCardsRepo
	txManager txManager
}

func New(pbRepo physBooksRepo, oRepo ordersRepo, lcRepo libCardsRepo, tm txManager) *OrderService {
	return &OrderService{
		pbRepo:    pbRepo,
		oRepo:     oRepo,
		lcRepo:    lcRepo,
		txManager: tm,
	}
}
