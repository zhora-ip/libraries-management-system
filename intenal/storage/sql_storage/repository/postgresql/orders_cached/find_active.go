package orders_cached

import (
	"context"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
)

func (r *OrderRepoCached) FindActive(ctx context.Context) ([]*models.Order, error) {
	orders, err := r.repo.FindActive(ctx)
	if err != nil {
		return nil, err
	}

	for _, o := range orders {
		r.cache.Set(o)
	}

	return orders, nil
}
