package orders_cached

import (
	"context"
	"log"
	"time"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
	svc "github.com/zhora-ip/libraries-management-system/intenal/models/service"
)

const (
	cronTime   = 1
	historyLen = 5
)

type repo interface {
	Add(context.Context, *models.Order) (int64, error)
	FindActive(context.Context) ([]*models.Order, error)
	FindAll(context.Context, *svc.FindAllOrdersRequest) ([]*models.Order, error)
	FindBlockedByUserID(context.Context, int64) ([]*models.Order, error)
	FindByID(context.Context, int64) (*models.Order, error)
	FindCanceled(context.Context) ([]*models.Order, error)
	MarkAsAccepted(context.Context, int64) error
	MarkAsCanceled(context.Context, int64) error
	MarkAsIssued(context.Context, *models.Order) error
	MarkAsReturned(context.Context, int64) error
}

type cache interface {
	Set(*models.Order)
	Get(int64) (*models.Order, error)
	Del(int64)
	SetBatch(orders []*models.Order)
	GetBatch(req *svc.FindAllOrdersRequest) ([]*models.Order, error)
	PrintCache()
}

type OrderRepoCached struct {
	repo  repo
	cache cache
}

func New(repo repo, cache cache) *OrderRepoCached {
	orderRepoCached := &OrderRepoCached{
		repo:  repo,
		cache: cache,
	}

	go func() {
		ticker := time.NewTicker(time.Second * cronTime)
		defer ticker.Stop()

		for {
			<-ticker.C
			orders, err := repo.FindAll(context.Background(), &svc.FindAllOrdersRequest{
				Cursor:   time.Now(),
				Limit:    historyLen,
				Backward: false,
			})

			if err != nil {
				log.Print(err)
				continue
			}

			cache.SetBatch(orders)
		}
	}()

	return orderRepoCached
}
