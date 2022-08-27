package cache

type FixedCacheBucket[T CachableItem] struct {
	slices []*CacheBucketSlice[T]
}