package orders_cached

import (
	"context"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
)

func (r *OrderRepoCached) MarkAsAccepted(ctx context.Context, ID int64) error {
	var (
		o   *models.Order
		err error
	)
	err = r.repo.MarkAsAccepted(ctx, ID)
	if err != nil {
		return err
	}

	o, err = r.cache.Get(ID)
	if err != nil {
		return nil
	}

	newOrder := *o
	newOrder.Status = models.StatusAccepted
	r.cache.Set(&newOrder)
	return nil
}
