package orders_cached

import (
	"context"
	"time"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
)

func (r *OrderRepoCached) Add(ctx context.Context, o *models.Order) (int64, error) {
	var (
		now = time.Now()
	)
	o.CreatedAt = &now
	o.UpdatedAt = &now
	ID, err := r.repo.Add(ctx, o)
	if err != nil {
		return 0, err
	}

	o.ID = ID
	r.cache.Set(o)
	return ID, nil
}
