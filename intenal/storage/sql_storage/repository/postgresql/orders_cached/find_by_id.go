package orders_cached

import (
	"context"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
)

func (r *OrderRepoCached) FindByID(ctx context.Context, ID int64) (*models.Order, error) {
	o, err := r.cache.Get(ID)
	if err == nil {
		return o, nil
	}

	o, err = r.repo.FindByID(ctx, ID)
	if err != nil {
		return nil, err
	}

	r.cache.Set(o)
	return o, nil
}
