package cache

import "time"

// CacheItem is a single cache item.
// Treat this item as readonly.
type CacheItem[T CachableItem] struct {
	//READONLY
	hashCode int32
	//READONLY If time.Now is after this time, the item is expired.
	expiresAfter time.Time
	//READONLY
	value *T
}

func NewCacheItem[T CachableItem](hashcode int32, expiresAfter time.Time, value *T) CacheItem[T] {
	return CacheItem[T]{
		hashCode:     hashcode,
		expiresAfter: expiresAfter,
		value:        value,
	}
}

func (c CacheItem[T]) GetHashCode() int32 {
	return c.hashCode
}

func (c CacheItem[T]) GetValue() *T {
	return c.value
}

func (c CacheItem[T]) HasValue() bool {
	return c.value != nil
}

func (c CacheItem[T]) IsExpired(time time.Time) bool {
	return time.After(c.expiresAfter)
}

func (c CacheItem[T]) IsMatch(key KeyLookup) bool {
	if c.hashCode != key.HashCode {
		return false
	}
	if c.value == nil {
		return false
	}
	if (*c.value).GetKey() != key.Key {
		return false
	}

	return true
}

func (c CacheItem[T]) IsMatch2(other CacheItem[T]) bool {
	if c.hashCode != other.hashCode {
		return false
	}
	if c.value == nil {
		return false
	}
	if (*c.value).GetKey() != (*other.value).GetKey() {
		return false
	}

	return true
}
