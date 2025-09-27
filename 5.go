package main

import (
    "fmt"
    "sync"
    "time"
)

type CacheItem struct {
    value      interface{}
    timeStamp  time.Time
}

type TTL struct {
    items map[string]*CacheItem
    mutex sync.Mutex
}

func NewCache() *TTL {
    return &TTL{
        items: make(map[string]*CacheItem),
    }
}

func (c *TTL) Get(key string) (interface{}, bool) {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    item, exists := c.items[key]
    if !exists || item.timeStamp.Before(time.Now()) {
        delete(c.items, key)
        return nil, false
    }
    item.timeStamp = time.Now().Add(item.timeStamp.Sub(time.Now()))
    return item.value, true
}

func (c *TTL) Set(key string, value interface{}, ttl time.Duration) {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    c.items[key] = &CacheItem{
        value:    value,
        timeStamp: time.Now().Add(ttl),
    }
}

func main() {
    cache := NewCache()
    cache.Set("key", "cached_value", 5*time.Second)
    val, found := cache.Get("key")
    fmt.Println(val, found)
    time.Sleep(6 * time.Second)
    val, found = cache.Get("key")
    fmt.Println(val, found)
}