package cache

import (
	"github.com/bluele/gcache"
)

// CacheBuilder is defined about a simple cache configuraion.
// It contains cache name and cache size.
type CacheBuilder struct {
	name string
	size int
}

var (
	caches map[string]gcache.Cache
)

func init() {
	caches = make(map[string]gcache.Cache)
}

// New return a cache builder accroding to configuration.
func New(cacheName string, size int) *CacheBuilder {
	return &CacheBuilder{
		name: cacheName,
		size: size,
	}
}

// Build is used to build a cache entity.
// It's base on LRU policy and can be used anywhere in this project.
func (cb *CacheBuilder) Build() {
	caches[cb.name] = gcache.New(cb.size).
		LRU().
		Build()
}

// Use is used to specify a cache by cache name.
func Use(cacheName string) gcache.Cache {
	return caches[cacheName]
}