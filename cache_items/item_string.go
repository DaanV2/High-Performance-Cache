package cache_items

// CacheItemString an item that stores a string
type CacheItemString string

// NewCacheItemString creates a new CacheItemString.
func NewCacheItemString(key string) *CacheItemString {
	item := CacheItemString(key)
	return &item
}

// GetKey returns the key of the KeyValuePair.
func (c *CacheItemString) GetKey() string {
	return string((*c))
}

func (c *CacheItemString) String() string {
	return string(*c)
}
