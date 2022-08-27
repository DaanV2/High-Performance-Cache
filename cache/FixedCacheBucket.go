package cache

type FixedCacheBucket[T CachableItem] struct {
	slices []*CacheBucketSlice[T]
}

// NewFixedCacheBucket creates a new FixedCacheBucket with the spaces for the given amount of items
func NewFixedCacheBucket[T CachableItem](itemCount int, settings CacheBucketSliceSettings) *FixedCacheBucket[T] {
	result := &FixedCacheBucket[T]{
		slices: []*CacheBucketSlice[T]{},
	}

	for i := 0; i < itemCount; {
		set := NewCacheBucketSlice[T](settings)
		result.slices = append(result.slices, set)
		i += set.Capacity()
	}

	return result
}