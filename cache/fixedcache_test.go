package cache

import (
	"testing"

	"gotest.tools/assert"
)

func Test_FixedCache_Creation(t *testing.T) {
	settings := DefaultFixedCacheSettings[*KeyValuePair[string]](10)
	settings.Cleaning.AutoClean = false

	t.Run("Default Settings", func(t *testing.T) {
		cache := NewFixedCache[*KeyValuePair[string]](settings)

		if cache == nil {
			t.Error("Cache is nil")
		}
	})
}

func Test_FixedCache_Set(t *testing.T) {
	settings := DefaultFixedCacheSettings[*KeyValuePair[string]](10)
	settings.Cleaning.AutoClean = false
	kvp := NewKeyValuePair("key", "value")

	t.Run("Set item", func(t *testing.T) {
		cache := NewFixedCache[*KeyValuePair[string]](settings)

		assert.Equal(t, cache.Set(kvp), true)
	})

	t.Run("Set item with existing key", func(t *testing.T) {
		cache := NewFixedCache[*KeyValuePair[string]](settings)

		cache.Set(kvp)
		assert.Equal(t, cache.Set(kvp), true)
	})
}

func Test_FixedCache_Get(t *testing.T) {
	settings := DefaultFixedCacheSettings[*KeyValuePair[string]](10)
	settings.Cleaning.AutoClean = false
	kvp := NewKeyValuePair("key", "value")

	t.Run("Get item", func(t *testing.T) {
		cache := NewFixedCache[*KeyValuePair[string]](settings)
		
		assert.Equal(t, cache.Set(kvp), true)
		item, ok := cache.Get("key")

		assert.Equal(t, ok, true)
		assert.Equal(t, item.GetValue(), kvp)
	})
}
