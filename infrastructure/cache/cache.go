package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

// Cache ...
type Cache struct {
	*cache.Cache
}

// New ...
func New() *Cache {
	return &Cache{cache.New(5*time.Minute, 10*time.Minute)}
}

// DefaultExpiration ...
func (c *Cache) DefaultExpiration() time.Duration {
	return cache.DefaultExpiration
}
