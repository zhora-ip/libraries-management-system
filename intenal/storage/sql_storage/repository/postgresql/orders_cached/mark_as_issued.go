package orders_cached

import (
	"context"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
)

func (r *OrderRepoCached) MarkAsIssued(ctx context.Context, o *models.Order) error {
	err := r.repo.MarkAsIssued(ctx, o)
	if err != nil {
		return err
	}

	newOrder := *o
	newOrder.Status = models.StatusIssued
	r.cache.Set(&newOrder)
	return nil
}
