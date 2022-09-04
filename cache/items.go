package cache

// CacheItemString an item that stores a string
type CacheItemString string

// GetKey returns the key of the KeyValuePair.
func (c *CacheItemString) GetKey() string {
	return string((*c))
}

// NewCacheItemString creates a new CacheItemString.
func NewCacheItemString(key string) *CacheItemString {
	item := CacheItemString(key)
	return &item
}

// KeyValuePair is a key/value pair of a key and generic value.
// Treat this as an immutable object.
type KeyValuePair[T any] struct {
	// Key is the key of the key/value pair.
	Key string
	// Value is the value of the key/value pair.
	Value T
}

// NewKeyValuePair creates a new KeyValuePair.
func NewKeyValuePair[T any](key string, value T) *KeyValuePair[T] {
	return &KeyValuePair[T]{
		Key:   key,
		Value: value,
	}
}

// GetKey returns the key of the key/value pair.
func (kvp *KeyValuePair[T]) GetKey() string {
	return kvp.Key
}
