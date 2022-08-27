package cache

import "github.com/DaanV2/High-Performance-Cache/concurrency"

type FixedCache[T CachableItem] struct {
	buckets []*FixedCacheBucket[T]
	locks   concurrency.Locks
}