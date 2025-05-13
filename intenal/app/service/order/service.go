package orderservice

import (
	"context"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
	svc "github.com/zhora-ip/libraries-management-system/intenal/models/service"
	ntfs "github.com/zhora-ip/notification-manager/pkg/pb"
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
	FindExpired(context.Context) ([]*models.Order, error)
	MarkAsCanceled(context.Context, int64) error
	MarkAsReturned(context.Context, int64) error
	MarkAsExpired(context.Context, int64) error
	MarkAsAccepted(context.Context, int64) error
	FindAll(context.Context, *svc.FindAllOrdersRequest) ([]*models.Order, error)
}

type libCardsRepo interface {
	FindByUserID(context.Context, int64) (*models.LibCard, error)
}

type usersRepo interface {
	FindByID(context.Context, int64) (*models.User, error)
}

type ntfsManager interface {
	Notify(context.Context, *ntfs.NotifyRequest) error
}

type txManager interface {
	RunSerializable(context.Context, func(context.Context) error) error
}

type workerPool interface {
	Submit(any, chan<- error)
}

type OrderService struct {
	pbRepo    physBooksRepo
	oRepo     ordersRepo
	lcRepo    libCardsRepo
	uRepo     usersRepo
	txManager txManager
	nManager  ntfsManager
	audit     workerPool
}

func New(pbRepo physBooksRepo, oRepo ordersRepo, lcRepo libCardsRepo,
	uRepo usersRepo, tm txManager, nm ntfsManager, wp workerPool) *OrderService {
	return &OrderService{
		pbRepo:    pbRepo,
		oRepo:     oRepo,
		lcRepo:    lcRepo,
		uRepo:     uRepo,
		txManager: tm,
		nManager:  nm,
		audit:     wp,
	}
}

// Submit sends a task for processing
func (o *OrderService) Submit(l any, errCh chan<- error) {
	o.audit.Submit(l, errCh)
}
