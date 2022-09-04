package cache

import (
	"time"

	util "github.com/DaanV2/High-Performance-Cache/util"
)

// CacheItem is a single cache item.
// Treat this item as readonly.
type CacheItem[T KeyedObject] struct {
	//READONLY
	hashCode uint64
	//READONLY If time.Now is after this time, the item is expired.
	expiresAfter time.Time
	//READONLY
	value T
}

// NewCacheItem creates a new CacheItem.
func NewCacheItem[T KeyedObject](expiresAfter time.Time, value T) CacheItem[T] {
	return CacheItem[T]{
		hashCode:     util.GetHashcode(value.GetKey()),
		expiresAfter: expiresAfter,
		value:        value,
	}
}

// EmptyCacheItem returns an empty CacheItem.
func EmptyCacheItem[T KeyedObject]() CacheItem[T] {
	return CacheItem[T]{}
}

// GetHashcode returns the hashcode of the key.
func (c CacheItem[T]) GetHashcode() uint64 {
	return c.hashCode
}

// GetValue returns the value of the item.
func (c CacheItem[T]) GetValue() T {
	return c.value
}

// GetValue returns the value of the item.
func (c CacheItem[T]) GetKey() string {
	return c.value.GetKey()
}

// HasValue returns true if the item has a value.
func (c CacheItem[T]) HasValue() bool {
	return c.hashCode != 0
}

// IsExpired returns true if the item is expired.
func (c CacheItem[T]) IsExpired(time time.Time) bool {
	return time.After(c.expiresAfter)
}

// IsMatch returns true if the item is a match for the given key.
func (c CacheItem[T]) IsMatch(key HashKeyedObject) bool {
	if c.hashCode != key.GetHashcode() {
		return false
	}
	if c.GetKey() != key.GetKey() {
		return false
	}

	return true
}

// CanPlaceHere returns true if the item can be placed in the space the other item is occuping
// expiringDate is the date the item(s) will expire.
// other is the item that is already in the space.
func (c CacheItem[T]) CanPlaceHere(expiringDate time.Time, other HashKeyedObject) bool {
	if !c.HasValue() {
		return true
	}
	if c.IsExpired(expiringDate) {
		return true
	}
	if c.hashCode != other.GetHashcode() {
		return false
	}
	if c.GetKey() != other.GetKey() {
		return false
	}

	return true
}
