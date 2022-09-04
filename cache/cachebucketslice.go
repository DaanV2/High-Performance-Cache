package cache

import (
	"fmt"
	"time"
	"unsafe"

	util "github.com/DaanV2/High-Performance-Cache/util"
)

func init() {

}

// CacheBucketSliceSettings is the settings for creating a CacheBucketSlice.
type CacheBucketSliceSettings struct {
	//MaxSize is the maximum item count in a slice.
	MaxSize int
}

// DefaultSettings returns the default settings for a CacheBucketSlice.
func DefaultBucketSettings[T KeyedObject](targetCache util.CacheKind) CacheBucketSliceSettings {
	cacheSize := util.GetCacheSize(targetCache)

	var template CacheItem[T]
	size := int64(unsafe.Sizeof(template))

	//How many items can we fit in the cache?
	itemsCount := cacheSize / size

	//Use only about 60% of the cache for the items, so 40% is left for the code.
	//TODO make this configurable
	//TODO maybe reduce the cache size by the size of the code use for reading writing the cache (don't forget to add some space for other stuff in the cache)
	itemsCount = (itemsCount * 100) / 60
	fmt.Printf("cache bucket slice size: %v\n", itemsCount)

	return CacheBucketSliceSettings{
		MaxSize: int(itemsCount),
	}
}

// CacheBucketSlice is a cache bucket that stores items in a slice.
type CacheBucketSlice[T KeyedObject] struct {
	//The hashrange is used to find quickly the correct bucket for a given key.
	hashRange HashRange
	//The amount of items in the cache
	itemCount int
	//The items in the cache
	items []CacheItem[T]
}

// NewCacheBucketSlice creates a new CacheBucketSlice.
func NewCacheBucketSlice[T KeyedObject](settings CacheBucketSliceSettings) *CacheBucketSlice[T] {
	return &CacheBucketSlice[T]{
		hashRange: NewHashRange(),
		itemCount: 0,
		items:     make([]CacheItem[T], settings.MaxSize),
	}
}

// GetStartIndex returns the start index for a given key.
func (bucketSlice *CacheBucketSlice[T]) GetStartIndex(hashcode uint64) int {
	return int(hashcode) % len(bucketSlice.items)
}

// IsFull returns true if the bucket is full.
func (bucketSlice *CacheBucketSlice[T]) IsFull() bool {
	return bucketSlice.itemCount >= len(bucketSlice.items)
}

// Clean removes any item that is expired.
func (bucketSlice *CacheBucketSlice[T]) Clean(expiringDate time.Time) int {
	count := 0
	itemCount := 0
	items := bucketSlice.items
	//Build new hashrange after cleaning
	newRange := NewHashRange()

	for i := 0; i < len(items); i++ {
		item := items[i]

		//Is an item in the stored at i
		if item.HasValue() {
			//Is the item expired?
			if item.IsExpired(expiringDate) {
				items[i] = CacheItem[T]{}
				count++
			} else {
				//Else we update the range and count
				(&newRange).UpdateRange(item.hashCode)
				itemCount++
			}
		}
	}

	bucketSlice.itemCount = itemCount
	bucketSlice.hashRange = newRange

	return count
}

// Get returns the item for a given key. if it can be found, else returns false on the second parameter
func (bucketSlice *CacheBucketSlice[T]) Get(key KeyLookup) (CacheItem[T], bool) {
	var result CacheItem[T]

	if !bucketSlice.hashRange.IsInRange(key.Hashcode) {
		return result, false
	}

	//Get the index we expect the item to be at, based on the hashcode
	//Can still be somewhere else in the cache
	start := bucketSlice.GetStartIndex(key.Hashcode)
	items := bucketSlice.items
	max := len(items)

	//Start at the given index to lookat and loop all around
	for i := start; i < max; i++ {
		item := items[i]
		if item.IsMatch(key) {
			return item, true
		}
	}
	for i := 0; i < start; i++ {
		item := items[i]
		if item.IsMatch(key) {
			return item, true
		}
	}

	return result, false
}

// Set sets the item for a given key. return true is successfull, false is it failed
func (bucketSlice *CacheBucketSlice[T]) Set(value CacheItem[T]) bool {
	return bucketSlice.SetWithExpire(value, time.Now())
}

// SetWithExpire sets the item for a given key. return true is successfull, false is it failed
// If the item is expired, it will be replaced with the given item.
func (bucketSlice *CacheBucketSlice[T]) SetWithExpire(value CacheItem[T], expiringTime time.Time) bool {
	//Bucket is full
	if bucketSlice.IsFull() {
		return false
	}

	start := bucketSlice.GetStartIndex(value.hashCode)
	items := bucketSlice.items
	max := len(items)

	//Start at the given index to lookat and loop all around until you find an empty spot
	//Or replaceable item
	for i := start; i < max; i++ {
		item := items[i]
		//If value check is a match, we replace the item
		if item.HasValue() {
			if item.CanPlaceHere(expiringTime, value) {
				items[i] = value
				return true
			}
		} else {
			items[i] = value
			return true
		}
	}
	for i := 0; i < start; i++ {
		item := items[i]
		if item.HasValue() {
			if item.CanPlaceHere(expiringTime, value) {
				items[i] = value
				return true
			}
		} else {
			items[i] = value
			return true
		}
	}

	return false
}

// Count returns the amount of items in the cache.
func (bucketSlice *CacheBucketSlice[T]) Count() int {
	return bucketSlice.itemCount
}

// Capacity returns the maximum amount of items the cache can hold.
func (bucketSlice *CacheBucketSlice[T]) Capacity() int {
	return len(bucketSlice.items)
}

// Clear clears the cache.
func (bucketSlice *CacheBucketSlice[T]) Clear() error {
	bucketSlice.items = make([]CacheItem[T], len(bucketSlice.items))
	bucketSlice.itemCount = 0
	bucketSlice.hashRange = NewHashRange()

	return nil
}

// ForEach iterates over the items in the cache.
func (bucketSlice *CacheBucketSlice[T]) ForEach(callback func(value CacheItem[T]) error) error {
	for _, item := range bucketSlice.items {
		if item.HasValue() {
			if err := callback(item); err != nil {
				return err
			}
		}
	}
	return nil
}

// Delete removes an item from the cache.
func (bucketSlice *CacheBucketSlice[T]) Delete(key KeyLookup) bool {
	start := bucketSlice.GetStartIndex(key.Hashcode)
	items := bucketSlice.items
	max := len(items)

	//Start at the given index to lookat and loop all around until you find an empty spot
	//Or replaceable item
	for i := start; i < max; i++ {
		item := items[i]
		//If value check is a match, we replace the item
		if item.HasValue() {
			if item.IsMatch(key) {
				items[i] = CacheItem[T]{}
				return true
			}
		}
	}
	for i := 0; i < start; i++ {
		item := items[i]
		if item.HasValue() {
			if item.IsMatch(key) {
				items[i] = CacheItem[T]{}
				return true
			}
		}
	}

	return false
}
