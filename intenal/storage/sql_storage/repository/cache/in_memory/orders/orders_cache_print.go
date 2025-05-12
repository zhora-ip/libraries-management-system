package in_memory

import "fmt"

func (c *OrdersCache) PrintCache() {
	fmt.Println("Active: ")
	for k, v := range c.active.data {
		fmt.Printf("%d: %#v\n", k, *v)
	}

	fmt.Println("History: ")
	for k, v := range c.history.data {
		fmt.Printf("%d: %#v\n", k, *v)
	}
}
