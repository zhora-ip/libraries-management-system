package in_memory

func (c *OrdersCache) Del(ID int64) {
	c.active.mx.Lock()
	defer c.active.mx.Unlock()
	delete(c.active.data, ID)
}
