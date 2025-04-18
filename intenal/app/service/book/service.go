package bookservice

import (
	"context"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
	svc "github.com/zhora-ip/libraries-management-system/intenal/models/service"
)

type booksRepo interface {
	Add(context.Context, *models.Book) (int64, error)
	FindAll(context.Context, *svc.FindAllRequest) (*svc.FindAllResponse, error)
}

type txManager interface {
	RunSerializable(context.Context, func(context.Context) error) error
}

type BookService struct {
	booksRepo booksRepo
	txManager txManager
}

func New(repo booksRepo, tm txManager) *BookService {
	return &BookService{
		booksRepo: repo,
		txManager: tm,
	}
}
