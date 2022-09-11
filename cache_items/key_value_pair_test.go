package cache_items

import (
	"testing"

	"gotest.tools/assert"
)

func Test_Items_Implement_Interface(t *testing.T) {
	value := "test"
	item := CacheItemString(value)
	key := item.GetKey()

	assert.Equal(t, key, value)

	kvp := NewKeyValuePair(key, value)
	assert.Equal(t, kvp.GetKey(), key)
}
