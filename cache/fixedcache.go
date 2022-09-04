package cache

import (
	"math"
	"sync"
	"sync/atomic"
	"time"

	"github.com/DaanV2/High-Performance-Cache/concurrency"
	"github.com/DaanV2/High-Performance-Cache/util"
)

type FixedCacheSettings struct {
	Capacity      int
	Cleaning      CacheCleaningSettings
	BucketAmount  func(items int) int
	Buckets       CacheBucketSliceSettings
	LockAmount    int
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
		Cleaning:      DefaultCacheCleaningSettings(),
		CacheDuration: time.Minute * 3,
	}
}

// FixedCache is a cache that has a fixed size
type FixedCache[T KeyedObject] struct {
	//The buckets in the cache
	buckets []*FixedCacheBucket[T]
	//The locks for the buckets
	locks concurrency.Locks
	//The settings for the cache
	settings FixedCacheSettings
	//The cleaning handler
	Clearer *CacheCleaner
}

// Dispose disposes the cache
func (fx *FixedCache[T]) Dispose() {
	if fx.Clearer != nil {
		fx.Clearer.Dispose()
	}

	fx.buckets = nil
	fx.Clearer = nil
}

// NewFixedCache creates a new FixedCache
func NewFixedCache[T KeyedObject](settings FixedCacheSettings) *FixedCache[T] {
	bucketCount := settings.BucketAmount(settings.Capacity)
	bucketItemCount := settings.Capacity / bucketCount

	result := &FixedCache[T]{
		buckets: make([]*FixedCacheBucket[T], bucketCount),
		locks:   concurrency.NewLocks(1024),
		settings: settings,
	}

	result.Clearer = StartCacheCleaner(settings.Cleaning, result)

	for i := 0; i < bucketCount; i++ {
		result.buckets[i] = NewFixedCacheBucket[T](bucketItemCount, settings.Buckets)
	}

	return result
}

// GetBucket returns the bucket for the given key
func (fx *FixedCache[T]) GetBucket(hashcode uint64) (*FixedCacheBucket[T], *sync.Mutex) {
	bucketIndex := int(hashcode % uint64(len(fx.buckets)))
	bucket := fx.buckets[bucketIndex]
	lock := fx.GetLock(bucketIndex)

	return bucket, lock
}

// GetLock returns the lock for the bucket index
func (fx *FixedCache[T]) GetLock(bucketIndex int) *sync.Mutex {
	return fx.locks.GetLockByIndex(bucketIndex % fx.locks.Count())
}

// GetLockForHashcode returns the lock for the given hashcode
func (fx *FixedCache[T]) GetLockForHashcode(hashcode uint64) *sync.Mutex {
	bucketIndex := int(hashcode % uint64(len(fx.buckets)))
	return fx.GetLock(bucketIndex)
}

// Get returns the item from the cache.
func (fx *FixedCache[T]) Get(key string) (CacheItem[T], bool) {
	lookup := NewKeyLookup(key)
	bucket, lock := fx.GetBucket(lookup.Hashcode)

	lock.Lock()
	defer lock.Unlock()

	return bucket.Get(lookup)
}

// Set sets the item in the cache. If the item is already in the cache, it will be updated.
// Returns true if the item was added or updated. false if failed (or full)
func (fx *FixedCache[T]) Set(item T) bool {
	value := NewCacheItem[T](time.Now(), item)
	bucket, lock := fx.GetBucket(value.hashCode)

	lock.Lock()
	defer lock.Unlock()

	return bucket.Set(value)
}

// Gets the item from the cache or sets it if it does not exist.
func (fx *FixedCache[T]) GetOrSet(key string, createFn func(key string) (T, error)) (CacheItem[T], error) {
	lookup := NewKeyLookup(key)
	bucket, lock := fx.GetBucket(lookup.Hashcode)

	lock.Lock()
	defer lock.Unlock()

	if item, ok := bucket.Get(lookup); ok {
		return item, nil
	}

	var item T
	var err error
	if item, err = createFn(key); err != nil {
		return CacheItem[T]{}, err
	}

	value := NewCacheItem[T](time.Now(), item)
	if bucket.Set(value) {
		return value, nil
	}

	return value, CouldntSetOrGetError(key)
}

// Delete deletes the item from the cache.
func (fx *FixedCache[T]) Delete(key string) bool {
	lookup := NewKeyLookup(key)
	bucket, lock := fx.GetBucket(lookup.Hashcode)

	lock.Lock()
	defer lock.Unlock()

	return bucket.Delete(lookup)
}

// Clear clears the cache.
func (fx *FixedCache[T]) Clear() error {
	max := len(fx.buckets)

	for I := 0; I < max; I++ {
		if err := fx.bucketClear(I); err != nil {
			return err
		}
	}

	return nil
}

// Clear clears the cache.
func (fx *FixedCache[T]) ClearParralel() []error {
	errors := make([]error, len(fx.buckets))
	lock := sync.Mutex{}

	concurrency.Parralel(fx.buckets, func(item *FixedCacheBucket[T], index int, current []*FixedCacheBucket[T]) {
		err := fx.bucketClear(index)

		lock.Lock()
		defer lock.Unlock()

		errors = append(errors, err)
	})

	return errors
}

// Clear clears the specified bucket
func (fx *FixedCache[T]) bucketClear(index int) error {
	bucket := fx.buckets[index]
	lock := fx.GetLock(index)

	lock.Lock()
	defer lock.Unlock()

	return bucket.Clear()
}

// ForEach iterates over the cache and calls the given function for each item.
func (fx *FixedCache[T]) ForEach(callback func(value CacheItem[T]) error) error {
	max := len(fx.buckets)

	for I := 0; I < max; I++ {
		if err := fx.bucketForEach(I, callback); err != nil {
			return err
		}
	}

	return nil
}

// bucketForEach iterates over the specified bucket and calls the given function for each item.
func (fx *FixedCache[T]) bucketForEach(index int, callback func(value CacheItem[T]) error) error {
	bucket := fx.buckets[index]
	lock := fx.GetLock(index)

	lock.Lock()
	defer lock.Unlock()

	return bucket.ForEach(callback)
}

// Clean cleans the cache of any expired items.
func (fx *FixedCache[T]) Clean(expiringDate time.Time) int {
	max := len(fx.buckets)
	amount := 0

	for I := 0; I < max; I++ {
		amount += fx.bucketClean(I, expiringDate)
	}

	return amount
}

// CleanParralel cleans the cache of any expired items in parralel.
func (fx *FixedCache[T]) CleanParralel(expiringDate time.Time) int {
	amount := &atomic.Int32{}

	concurrency.Parralel(fx.buckets, func(item *FixedCacheBucket[T], index int, current []*FixedCacheBucket[T]) {
		temp := fx.bucketClean(index, expiringDate)
		amount.Add(int32(temp))
	})

	return int(amount.Load())
}

// bucketClean cleans the specified bucket of any expired items.
func (fx *FixedCache[T]) bucketClean(index int, expiringDate time.Time) int {
	bucket := fx.buckets[index]
	lock := fx.GetLock(index)

	lock.Lock()
	defer lock.Unlock()

	return bucket.Clean(expiringDate)
}
