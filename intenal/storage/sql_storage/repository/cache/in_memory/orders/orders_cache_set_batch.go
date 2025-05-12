package in_memory

import "github.com/zhora-ip/libraries-management-system/intenal/models"

func (c *OrdersCache) SetBatch(data []*models.Order) {
	c.history.mx.Lock()
	defer c.history.mx.Unlock()
	c.history.data = data
}
