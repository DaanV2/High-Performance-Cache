package cache

import "time"

// KeyedObject is an interface that returns its own key
type KeyedObject interface {
	//GetKey returns the key of the item
	GetKey() string
}

// HashableObject is an object that can returns its own hash
type HashableObject interface {
	//GetHashcode returns the hashcode of the item
	GetHashcode() uint64
}

// HashKeyedObject is a KeyedObject that can be hashed and returns a key
type HashKeyedObject interface {
	KeyedObject
	HashableObject
}

// Disposable is an interface that determines if an object can be disposed
type Disposable interface {
	//Dispose disposes the object
	Dispose()
}

// CacheSet is an interface that represents a cache that can set items
type CacheSet[T KeyedObject] interface {
	//Set sets the item in the cache.
	Set(item T) bool
}

// CacheGet is an interface that represents a cache that can get items
type CacheGet[T KeyedObject] interface {
	//Get returns the item from the cache.
	Get(key string) (CacheItem[T], bool)
}

// Cache is an interface that defines the basic functionality of a cache
type Cache[T KeyedObject] interface {
	CacheSet[T]
	CacheGet[T]
	Disposable

	//Gets the item from the cache or sets it if it does not exist.
	GetOrSet(key string, createFn func(key string) (T, error)) (CacheItem[T], error)
	//Delete deletes the item from the cache.
	Delete(key string) bool
	//Clear clears the cache.
	Clear() error
	//Close closes the cache.
	ForEach(callback func(value CacheItem[T]) error) error
	//CountCapacity returns the capacity and item count of the cache
	CountCapacity() (count, capacity uint64)
}

// Cleanable is an interface that defines the basic functionality of a cache that can be cleaned
type Cleanable interface {
	//Clean cleans the cache.
	Clean(expiringDate time.Time) int
	//CleanParallel cleans the cache in parallel.
	CleanParallel(expiringDate time.Time) int
}

// CacheCleanable is an interface that defines the basic functionality of a cache that can be cleaned
type CacheCleanable[T KeyedObject] interface {
	Cache[T]
	Cleanable
}
