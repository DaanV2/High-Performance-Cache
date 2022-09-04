package cache

import "errors"

// NotFoundError is returned when an item is not found in the cache.
func NotFoundError(Key string) error {
	return errors.New("Item not found in cache: " + Key)
}


// NotFoundError is returned when an item is not found in the cache.
func CouldntSetOrGetError(Key string) error {
	return errors.New("Couldn't set or get the item: " + Key)
}