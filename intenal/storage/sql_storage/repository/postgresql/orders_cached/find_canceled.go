package orders_cached

import (
	"context"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
)

func (r *OrderRepoCached) FindCanceled(ctx context.Context) ([]*models.Order, error) {
	return r.repo.FindCanceled(ctx)
}
