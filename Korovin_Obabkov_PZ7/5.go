package main

import (
	"fmt"
	"time"
)

func main() {
	cache := NewCache()

	cache.Set("key1", "value1", 4*time.Second)
	cache.Set("key2", "value2", 10*time.Second)

	if val, ok := cache.Get("key1"); ok {
		fmt.Printf("key1: %v\n", val)
	}

	time.Sleep(3 * time.Second)

	if val, ok := cache.Get("key1"); ok {
		fmt.Printf("key1: %v\n", val)
	} else {
		fmt.Println("key1 не найден или истёк")
	}

	if val, ok := cache.Get("key2"); ok {
		fmt.Printf("key2: %v\n", val)
	}

	cache.Delete("key2")

	if val, ok := cache.Get("key2"); ok {
		fmt.Printf("key2: %v\n", val)
	} else {
		fmt.Println("key2 не найден или истек")
	}
}

type cacheItem struct {
	value      interface{}
	expireTime time.Time
}

type Cache struct {
	items map[string]cacheItem
}

func NewCache() *Cache {
	return &Cache{
		items: make(map[string]cacheItem),
	}
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	expireTime := time.Now().Add(ttl)
	c.items[key] = cacheItem{
		value:      value,
		expireTime: expireTime,
	}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	item, found := c.items[key]
	if !found {
		return nil, false
	}

	if time.Now().After(item.expireTime) {
		delete(c.items, key)
		return nil, false
	}

	return item.value, true
}

func (c *Cache) Delete(key string) {
	delete(c.items, key)
	fmt.Printf("%v удалён\n", key)
}
