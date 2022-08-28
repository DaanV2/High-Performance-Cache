package cache

import "errors"

// NotFoundError is returned when an item is not found in the cache.
func NotFoundError(Key string) error {
	return errors.New("Item not found in cache: " + Key)
}
