package cache

import (
	"time"
)

// Cache is the interface that defines methods for a key-value cache.
type Cache interface {
	Set(key string, value interface{}, duration time.Duration)
	Get(key string) (interface{}, error)
	Delete(key string) error
}
