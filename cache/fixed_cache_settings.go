package cache

import (
	"math"
	"time"

	"github.com/DaanV2/High-Performance-Cache/util"
)

// FixedCacheSettings is a struct that contains the settings of a FixedCache
type FixedCacheSettings struct {
	// Capacity is the capacity of the cache
	Capacity int
	// The amount of buckets to create for the given amount of items
	BucketAmount func(items int) int
	// The settings of the buckets
	Buckets CacheBucketSliceSettings
	// The amount of locks to create for the given amount of items
	LockAmount int
	// The duration of an item in the cache
	CacheDuration time.Duration
}

// DefaultFixedCacheSettings creates a new default FixedCacheSettings
func DefaultFixedCacheSettings[T KeyedObject](capacity int) FixedCacheSettings {
	return FixedCacheSettings{
		Capacity: capacity,
		BucketAmount: func(items int) int {
			return int(math.Sqrt(float64(items)))
		},
		Buckets:       DefaultBucketSettings[T](util.CacheL1),
		LockAmount:    1024,
		CacheDuration: time.Minute * 3,
	}
}
