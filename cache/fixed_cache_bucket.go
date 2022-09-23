package cache

import "time"

// FixedCacheBucket is a cache bucket that has a fixed size
type FixedCacheBucket[T KeyedObject] struct {
	//The hashrange is used to find quickly the correct bucket for a given key.
	hashRange HashRange

	//The items in the bucket
	slices []*CacheBucketSlice[T]
}

// NewFixedCacheBucket creates a new FixedCacheBucket with the spaces for the given amount of items
func NewFixedCacheBucket[T KeyedObject](itemCount int, settings CacheBucketSliceSettings) *FixedCacheBucket[T] {
	result := &FixedCacheBucket[T]{
		slices:    []*CacheBucketSlice[T]{},
		hashRange: NewHashRange(),
	}

	for i := 0; i < itemCount; {
		set := NewCacheBucketSlice[T](settings)
		result.slices = append(result.slices, set)
		i += set.Capacity()
	}

	return result
}

// Get returns the item from the cache.
func (fcb *FixedCacheBucket[T]) Get(key KeyLookup) (CacheItem[T], bool) {
	if !fcb.hashRange.Contains(key.Hashcode) {
		return CacheItem[T]{}, false
	}

	for _, slice := range fcb.slices {
		if item, ok := slice.Get(key); ok {
			return item, true
		}
	}

	return CacheItem[T]{}, false
}

// Set sets the item in the cache.
func (fcb *FixedCacheBucket[T]) Set(item CacheItem[T]) bool {
	return fcb.SetWithExpire(item, time.Now())
}

// SetWithExpire sets the item in the cache with the given expire time.
func (fcb *FixedCacheBucket[T]) SetWithExpire(value CacheItem[T], expiringTime time.Time) bool {
	for _, slice := range fcb.slices {
		if slice.SetWithExpire(value, expiringTime) {
			fcb.hashRange.UpdateRange(value.hashcode)
			return true
		}
	}

	return false
}

// Delete deletes the item from the cache.
func (fcb *FixedCacheBucket[T]) Delete(key KeyLookup) bool {
	result := false

	for _, slice := range fcb.slices {
		result = result || slice.Delete(key)
	}

	return result
}

// Clear clears the cache.
func (fcb *FixedCacheBucket[T]) Clear() error {
	for _, slice := range fcb.slices {
		if err := slice.Clear(); err != nil {
			return err
		}
	}

	fcb.hashRange = NewHashRange()

	return nil
}

// Clean cleans the cache of any expired items.
func (fcb *FixedCacheBucket[T]) Clean(expiringDate time.Time) int {
	amount := 0

	hr := NewHashRange()

	for _, slice := range fcb.slices {
		amount += slice.Clean(expiringDate)
		hr.Update(slice.hashRange)
	}

	fcb.hashRange = hr

	return amount
}

// ForEach calls the given function for each item in the cache.
func (fcb *FixedCacheBucket[T]) ForEach(callback func(value CacheItem[T]) error) error {
	for _, slice := range fcb.slices {
		if err := slice.ForEach(callback); err != nil {
			return err
		}
	}

	return nil
}

// CountCapacity returns the total item count and capacity of the cache.
func (fcb *FixedCacheBucket[T]) CountCapacity() (count, capacity uint64) {
	count = 0
	capacity = 0

	for _, slice := range fcb.slices {
		count += uint64(slice.Count())
		capacity += uint64(slice.Capacity())
	}

	return count, capacity
}

func (fbc *FixedCacheBucket[T]) Dispose() {
	slices := fbc.slices
	fbc.hashRange = NewHashRange()
	fbc.slices = nil

	for _, slice := range slices {
		slice.Dispose()
	}
}