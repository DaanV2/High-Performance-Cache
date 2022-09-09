package cache

import (
	"fmt"
	"time"

	util "github.com/DaanV2/High-Performance-Cache/util"
)

// CacheItem is a single cache item.
// Treat this item as readonly.
type CacheItem[T KeyedObject] struct {
	//READONLY
	hashcode uint64
	//READONLY If time.Now is after this time, the item is expired.
	expiresAfter time.Time
	//READONLY
	value T
}

// NewCacheItem creates a new CacheItem.
func NewCacheItem[T KeyedObject](expiresAfter time.Time, value T) CacheItem[T] {
	key := value.GetKey()

	return CacheItem[T]{
		hashcode:     util.GetHashcode(key),
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
	return c.hashcode
}

// GetValue returns the value of the item.
func (c CacheItem[T]) GetValue() T {
	return c.value
}

const emptyKey = ""

// GetValue returns the value of the item.
func (c CacheItem[T]) GetKey() string {
	if c.HasValue() {
		return c.value.GetKey()
	}

	return emptyKey
}

// HasValue returns true if the item has a value.
func (c CacheItem[T]) HasValue() bool {
	return c.hashcode != 0
}

// IsExpired returns true if the item is expired.
func (c CacheItem[T]) IsExpired(time time.Time) bool {
	return time.After(c.expiresAfter)
}

// IsMatch returns true if the item is a match for the given key.
func (c CacheItem[T]) IsMatch(key HashKeyedObject) bool {
	if c.hashcode != key.GetHashcode() {
		return false
	}

	keyThis := c.GetKey()
	keyOther := key.GetKey()
	if keyThis != keyOther {
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
	if c.hashcode != other.GetHashcode() {
		return false
	}
	if c.GetKey() != other.GetKey() {
		return false
	}

	return true
}

func (c CacheItem[T]) String() string {
	return fmt.Sprintf("key: %s, hashcode: %d, expiresAfter: %s", c.GetKey(), c.hashcode, c.expiresAfter)
}
