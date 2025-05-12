package in_memory

import "github.com/zhora-ip/libraries-management-system/intenal/models"

func (c *OrdersCache) Set(o *models.Order) {
	c.active.mx.Lock()
	defer c.active.mx.Unlock()
	c.active.data[o.ID] = o
}
