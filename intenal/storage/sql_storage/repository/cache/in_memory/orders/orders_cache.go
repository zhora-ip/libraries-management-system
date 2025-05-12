package in_memory

import (
	"errors"
	"sync"

	"github.com/zhora-ip/libraries-management-system/intenal/models"
)

var (
	ErrCacheMiss = errors.New("cache miss")
)

type OrdersCache struct {
	active  *activeOrders
	history *historyOrders
}

type activeOrders struct {
	mx   *sync.RWMutex
	data map[int64]*models.Order
}

type historyOrders struct {
	mx   *sync.RWMutex
	data []*models.Order
}

func New() *OrdersCache {
	return &OrdersCache{
		active: &activeOrders{
			mx:   &sync.RWMutex{},
			data: map[int64]*models.Order{},
		},
		history: &historyOrders{
			mx:   &sync.RWMutex{},
			data: []*models.Order{},
		},
	}
}
