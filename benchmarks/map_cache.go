package benchmarks

import (
	"sync"

	"github.com/DaanV2/High-Performance-Cache/cache"
)

type MapCache[T cache.CachableItem] struct {
	mux   sync.RWMutex
	cache map[string]T
}

func newMapCache[T cache.CachableItem](size int) cache.Cache[T] {
	return &MapCache[T]{
		cache: make(map[string]T, size),
	}
}

//Get returns the item from the cache.
func (c *MapCache[T]) Get(key string) (T, error) {
	c.mux.RLock()
	defer c.mux.RUnlock()

	if item, ok := c.cache[key]; ok {
		return item, nil
	}

	var result T
	return result, cache.NotFoundError(key)
}

//Set sets the item in the cache.
func (c *MapCache[T]) Set(value T) error {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.cache[value.GetKey()] = value
	return nil
}

//Delete deletes the item from the cache.
func (c *MapCache[T]) Delete(key string) error {
	c.mux.Lock()
	defer c.mux.Unlock()
	
	delete(c.cache, key)
	return nil
}

//Clear clears the cache.
func (c *MapCache[T]) Clear() error {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.cache = make(map[string]T, len(c.cache))
	return nil
}

//Close closes the cache.
func (c *MapCache[T]) ForEach(callback func(value T) error) error {
	c.mux.RLock()
	defer c.mux.RUnlock()

	for _, value := range c.cache {
		if err := callback(value); err != nil {
			return err
		}
	}

	return nil
}
