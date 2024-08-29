package inmemorycache

import (
	"cache-service/cache"
	"errors"
	"sync"
	"time"
)

type cacheItem struct {
	value      interface{}
	expiration int64
}

// memoryCache is a thread-safe in-memory cache.
type memoryCache struct {
	items map[string]*cacheItem
	mu    sync.RWMutex
}

func NewCache() cache.Cache {
	cache := &memoryCache{
		items: make(map[string]*cacheItem),
	}

	return cache
}

// Set adds a key-value pair to the cache with an expiration time.
func (c *memoryCache) Set(key string, value interface{}, duration time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	expiration := time.Now().Add(duration).UnixNano()
	c.items[key] = &cacheItem{
		value:      value,
		expiration: expiration,
	}
}

func (c *memoryCache) Get(key string) (interface{}, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, exists := c.items[key]
	if !exists {
		return nil, errors.New("cache error: key not found")
	}

	if time.Now().UnixNano() > item.expiration {
		return nil, errors.New("cache error: data is expired")
		// may delete the key value from chache
	}

	return item.value, nil
}

func (c *memoryCache) Delete(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, exists := c.items[key]; !exists {
		return errors.New("cache error: key not found")
	}

	delete(c.items, key)
	return nil
}

// startEviction periodically checks and removes expired items.

// func (c *memoryCache) startEviction() {
// 	ticker := time.NewTicker(time.Minute)
// 	defer ticker.Stop()

// 	for range ticker.C {
// 		c.mu.Lock()
// 		for key, item := range c.items {
// 			if time.Now().UnixNano() > item.expiration {
// 				delete(c.items, key)
// 			}
// 		}
// 		c.mu.Unlock()
// 	}
// }
