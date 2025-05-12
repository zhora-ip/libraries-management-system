package orders_cached

import (
	"context"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
	svc "github.com/zhora-ip/libraries-management-system/intenal/models/service"
)

func (r *OrderRepoCached) FindAll(ctx context.Context, req *svc.FindAllOrdersRequest) ([]*models.Order, error) {
	var (
		cond = req.ID != nil || req.OccurrenceID != nil ||
			req.UserID != nil || req.LibraryID != nil || req.BookID != nil
	)

	if cond {
		return r.repo.FindAll(ctx, req)
	}

	orders, err := r.cache.GetBatch(req)
	if err != nil {
		return r.repo.FindAll(ctx, req)
	}
	return orders, nil
}
