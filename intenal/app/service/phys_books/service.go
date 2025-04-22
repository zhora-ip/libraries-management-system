package physbookservice

import (
	"context"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
)

type physBooksRepo interface {
	Add(context.Context, *models.PhysBook) (int64, error)
	FindByBookID(context.Context, int64) ([]*models.PhysBook, error)
}

type libraryRepo interface {
	FindByID(context.Context, int64) (*models.Library, error)
}

type txManager interface {
	RunSerializable(context.Context, func(context.Context) error) error
}

type PhysBookService struct {
	pbRepo    physBooksRepo
	lRepo     libraryRepo
	txManager txManager
}

func New(pbRepo physBooksRepo, lRepo libraryRepo, tm txManager) *PhysBookService {
	return &PhysBookService{
		pbRepo:    pbRepo,
		lRepo:     lRepo,
		txManager: tm,
	}
}
