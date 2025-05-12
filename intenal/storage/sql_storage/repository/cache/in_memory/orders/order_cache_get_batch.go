package in_memory

import (
	"slices"
	"time"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
	svc "github.com/zhora-ip/libraries-management-system/intenal/models/service"
)

func (c *OrdersCache) GetBatch(req *svc.FindAllOrdersRequest) ([]*models.Order, error) {
	c.history.mx.RLock()
	defer c.history.mx.RUnlock()

	var (
		lenCache = len(c.history.data)
		limit    = int(req.Limit)
		cursor   = req.Cursor
		corr     = 0
	)

	pos, ok := slices.BinarySearchFunc(c.history.data, cursor, func(a *models.Order, b time.Time) int {
		return -a.UpdatedAt.Compare(b)
	})

	if ok {
		corr = 1
	}

	if !req.Backward {
		if pos+limit+corr > lenCache || pos+corr == pos+limit+corr {
			return nil, ErrCacheMiss
		}

		return c.history.data[pos+corr : pos+limit+corr], nil
	}

	if req.Backward {
		if pos >= lenCache || pos == max(0, pos-limit) {
			return nil, ErrCacheMiss
		}
		return c.history.data[max(0, pos-limit):pos], nil
	}

	return nil, ErrCacheMiss

}
