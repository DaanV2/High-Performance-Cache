package cache

import (
	"testing"

	"gotest.tools/assert"
)

func Test_Items_Implement_Interface(t *testing.T) {
	value := "test"
	item := CacheItemString(value)
	key := item.GetKey()
	
	assert.Equal(t, key, value)
}