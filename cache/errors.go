package cache

import "errors"

func NotFoundError(Key string) error {
	return errors.New("Item not found in cache: " + Key)
}
