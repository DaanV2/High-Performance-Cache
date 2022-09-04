package benchmarks

import (
	"sync"
	"time"

	"github.com/DaanV2/High-Performance-Cache/cache"
)

type MapCache[T cache.KeyedObject] struct {
	mux          sync.RWMutex
	cache        map[string]cache.CacheItem[T]
	defaultCache time.Duration
}

func newMapCache[T cache.KeyedObject](size int) cache.Cache[T] {
	return &MapCache[T]{
		cache:        make(map[string]cache.CacheItem[T], size),
		defaultCache: time.Hour * 10,
	}
}

// Get returns the item from the cache.
func (c *MapCache[T]) Get(key string) (cache.CacheItem[T], bool) {
	c.mux.RLock()
	defer c.mux.RUnlock()

	var result cache.CacheItem[T]
	if result, ok := c.cache[key]; ok {
		return result, true
	}

	return result, false
}

// Get returns the item from the cache.
func (c *MapCache[T]) GetOrSet(key string, createFn func(key string) (T, error)) (cache.CacheItem[T], error) {
	c.mux.RLock()
	defer c.mux.RUnlock()

	var result cache.CacheItem[T]
	var ok bool
	if result, ok = c.cache[key]; ok {
		return result, nil
	}

	var value T
	var err error
	if value, err = createFn(key); err != nil {
		return result, err
	}

	result = cache.NewCacheItem(time.Now().Add(c.defaultCache), value)

	c.cache[key] = result
	return result, nil
}

// Set sets the item in the cache.
func (c *MapCache[T]) Set(value T) bool {
	c.mux.Lock()
	defer c.mux.Unlock()

	key := value.GetKey()
	time := time.Now().Add(c.defaultCache)
	item := cache.NewCacheItem(time, value)
	c.cache[key] = item
	return true
}

// Delete deletes the item from the cache.
func (c *MapCache[T]) Delete(key string) bool {
	c.mux.Lock()
	defer c.mux.Unlock()

	delete(c.cache, key)
	return true
}

// Clear clears the cache.
func (c *MapCache[T]) Clear() error {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.cache = make(map[string]cache.CacheItem[T], len(c.cache))
	return nil
}

// Close closes the cache.
func (c *MapCache[T]) ForEach(callback func(value cache.CacheItem[T]) error) error {
	c.mux.RLock()
	defer c.mux.RUnlock()

	for _, value := range c.cache {
		if err := callback(value); err != nil {
			return err
		}
	}

	return nil
}

func (c *MapCache[T]) Dispose() {
	c.cache = nil
}
