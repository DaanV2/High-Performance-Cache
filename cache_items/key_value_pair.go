package cache_items

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

// String returns the string representation of the KeyValuePair.
func (kvp *KeyValuePair[T]) String() string {
	return kvp.GetKey()
}