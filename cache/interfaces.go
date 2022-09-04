package cache

import "time"

// CachableItem
type CachableItem interface {
	//GetKey returns the key of the item
	GetKey() string
}

// Cache
type Cache[T CachableItem] interface {
	//Get returns the item from the cache.
	Get(key string) (CacheItem[T], error)
	//Set sets the item in the cache.
	Set(item T) error
	//Delete deletes the item from the cache.
	Delete(key string) error
	//Clear clears the cache.
	Clear() error
	//Close closes the cache.
	ForEach(callback func(value CacheItem[T]) error) error
}



type Cleanable interface {
	Clean(time time.Time) int
}