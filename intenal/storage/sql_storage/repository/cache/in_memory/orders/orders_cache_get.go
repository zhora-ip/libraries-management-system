package in_memory

import (
	"github.com/zhora-ip/libraries-management-system/intenal/models"
)

func (c *OrdersCache) Get(ID int64) (*models.Order, error) {
	c.active.mx.RLock()
	defer c.active.mx.RUnlock()
	o, ok := c.active.data[ID]
	if !ok {
		return nil, ErrCacheMiss
	}
	return o, nil
}
