package main

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	mu sync.RWMutex
	data map[string]string
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]string),
	}
}

func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, exists := c.data[key]
	return value, exists
}

func (c *Cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

func main() {
	cache:= NewCache()

	for i := 0; i < 5; i++ {
		go func(id int){
			for{
				val,found:=cache.Get("КлЮч")
				if found{
					fmt.Printf("Читатель %d получил: %s \n", id, val)
				} else{
					fmt.Printf("Читатель %d: Ключ не найден", id)
				}
				time.Sleep(100*time.Millisecond)
			}
		}(i)
	}

	go func(){
		for i := 0; ;i++{
			value := fmt.Sprintf("value_%d", i)
			cache.Set("КлЮч", value)
			fmt.Printf("Писатель установил: %s\n", value)
			time.Sleep(500*time.Millisecond)
		}
	}()


	time.Sleep(3*time.Second)
	
}