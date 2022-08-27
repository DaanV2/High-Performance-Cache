package cache

import (
	"time"
	"unsafe"

	util "github.com/DaanV2/High-Performance-Cache/util"
)

func init() {

}

type CacheBucketSliceSettings struct {
	MaxSize int
}

func DefaultSettings[T CachableItem](targetCache util.CacheKind) CacheBucketSliceSettings {
	cacheSize := util.GetCacheSize(targetCache)

	var template CacheItem[T]
	size := int64(unsafe.Sizeof(template))

	//How many items can we fit in the cache?
	itemsCount := cacheSize / size

	//Use only about 60% of the cache for the items, so 40% is left for the code.
	//TODO make this configurable
	//TODO maybe reduce the cache size by the size of the code use for reading writing the cache (don't forget to add some space for other stuff in the cache)
	itemsCount = (itemsCount * 100) / 60

	return CacheBucketSliceSettings{
		MaxSize: int(itemsCount),
	}
}

type CacheBucketSlice[T CachableItem] struct {
	//The hashrange is used to find quickly the correct bucket for a given key.
	hashRange HashRange
	//The amount of items in the cache
	itemCount int
	//The items in the cache
	items []CacheItem[T]
}

func NewCacheBucketSlice[T CachableItem](settings CacheBucketSliceSettings) *CacheBucketSlice[T] {
	return &CacheBucketSlice[T]{
		hashRange: NewHashRange(),
		itemCount: 0,
		items:     make([]CacheItem[T], settings.MaxSize),
	}
}

func (bucketSlice *CacheBucketSlice[T]) GetStartIndex(hashcode uint64) int {
	return int(hashcode) % len(bucketSlice.items)
}

func (bucketSlice *CacheBucketSlice[T]) IsFull() bool {
	return bucketSlice.itemCount >= len(bucketSlice.items)
}

// Clean removes any item that is expired.
func (bucketSlice *CacheBucketSlice[T]) Clean(time time.Time) int {
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
			if item.IsExpired(time) {
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

func (bucketSlice *CacheBucketSlice[T]) Get(key KeyLookup) (T, bool) {
	var result T

	if !bucketSlice.hashRange.IsInRange(key.HashCode) {
		return result, false
	}

	//Get the index we expect the item to be at, based on the hashcode
	//Can still be somewhere else in the cache
	start := bucketSlice.GetStartIndex(key.HashCode)
	items := bucketSlice.items
	max := len(items)

	//Start at the given index to lookat and loop all around
	for i := start; i < max; i++ {
		item := items[i]
		if item.IsMatch(key) {
			return item.GetValue(), true
		}
	}
	for i := 0; i < start; i++ {
		item := items[i]
		if item.IsMatch(key) {
			return item.GetValue(), true
		}
	}

	return result, false
}

func (bucketSlice *CacheBucketSlice[T]) Set(value CacheItem[T]) bool {
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
			if item.IsMatch2(value) {
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
			if item.IsMatch2(value) {
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

//Clear clears the cache.
func (bucketSlice *CacheBucketSlice[T]) Clear() error {
	bucketSlice.items = make([]CacheItem[T], len(bucketSlice.items))
	bucketSlice.itemCount = 0
	bucketSlice.hashRange = NewHashRange()
	
	return nil
}

//ForEach iterates over the items in the cache.
func (bucketSlice *CacheBucketSlice[T]) ForEach(callback func(value T) error) error {
	for _, item := range bucketSlice.items {
		if item.HasValue() {
			if err := callback(item.GetValue()); err != nil {
				return err
			}
		}
	}
	return nil
}