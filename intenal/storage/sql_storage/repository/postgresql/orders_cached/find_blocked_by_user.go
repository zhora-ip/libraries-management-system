package orders_cached

import (
	"context"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
)

func (r *OrderRepoCached) FindBlockedByUserID(ctx context.Context, ID int64) ([]*models.Order, error) {
	return r.repo.FindBlockedByUserID(ctx, ID)
}
