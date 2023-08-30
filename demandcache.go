package demandcache

import (
	"time"

	"github.com/num30/go-cache"
)

// DemandCache is a cache which automatically upgrades
// expiration time after getting value to the maximum life time value
// so, your most demanded value is will not be expired
type DemandCache[T any] interface {
	// Get returns item by the key
	Get(key string) (T, bool)
	// SetDefault sets value with default expiration time
	SetDefault(key string, v T)
	// Set sets value with specified expiration time
	Set(key string, v T, expiration time.Duration)
}

type demandCache[T any] struct {
	c *cache.Cache[T]
}

func New[T any](defaultExpiration, cleanupInterval time.Duration) DemandCache[T] {
	return &demandCache[T]{
		c: cache.New[T](defaultExpiration, cleanupInterval),
	}
}

func (d *demandCache[T]) Get(key string) (T, bool) {
	v, ok := d.c.Get(key)
	if !ok {
		return *new(T), false
	}
	// Upgrade expiration time to the maximum life time value
	d.SetDefault(key, v)
	return v, true
}

func (d *demandCache[T]) SetDefault(key string, value T) {
	d.c.SetDefault(key, value)
}

func (d *demandCache[T]) Set(key string, value T, expiration time.Duration) {
	d.c.Set(key, value, expiration)
}
