package cache

//AddCollection adds a collection to the cache
//Returns the amount of items that were added
func AddCollection[T KeyedObject](cache CacheSet[T], items []T) int {
	result := 0

	for _, item := range items {
		if cache.Set(item) {
			result++
		}
	}

	return result
}
