package cache

import (
	"testing"

	"github.com/DaanV2/High-Performance-Cache/cache_items"
	"gotest.tools/assert"
)

func Test_FixedCache_Creation(t *testing.T) {
	settings := DefaultFixedCacheSettings[*cache_items.KeyValuePair[string]](10)

	t.Run("Default Settings", func(t *testing.T) {
		cache := NewFixedCache[*cache_items.KeyValuePair[string]](settings)

		if cache == nil {
			t.Error("Cache is nil")
		}
	})
}

func Test_FixedCache_Set(t *testing.T) {
	settings := DefaultFixedCacheSettings[*cache_items.KeyValuePair[string]](10)
	kvp := cache_items.NewKeyValuePair("key", "value")

	t.Run("Set item", func(t *testing.T) {
		cache := NewFixedCache[*cache_items.KeyValuePair[string]](settings)

		assert.Equal(t, cache.Set(kvp), true)
	})

	t.Run("Set item with existing key", func(t *testing.T) {
		cache := NewFixedCache[*cache_items.KeyValuePair[string]](settings)

		cache.Set(kvp)
		assert.Equal(t, cache.Set(kvp), true)
	})
}

func Test_FixedCache_Get(t *testing.T) {
	settings := DefaultFixedCacheSettings[*cache_items.KeyValuePair[string]](10)
	kvp := cache_items.NewKeyValuePair("key", "value")

	t.Run("Get item", func(t *testing.T) {
		cache := NewFixedCache[*cache_items.KeyValuePair[string]](settings)

		assert.Equal(t, cache.Set(kvp), true)
		item, ok := cache.Get("key")

		assert.Equal(t, ok, true)
		assert.Equal(t, item.GetValue(), kvp)
	})
}

func Test_FixedCache_ForEach(t *testing.T) {
	settings := DefaultFixedCacheSettings[*cache_items.KeyValuePair[string]](10)
	kvp1 := cache_items.NewKeyValuePair("key1", "value1")
	kvp2 := cache_items.NewKeyValuePair("key2", "value2")
	kvp3 := cache_items.NewKeyValuePair("key3", "value3")

	t.Run("Get item", func(t *testing.T) {
		cache := NewFixedCache[*cache_items.KeyValuePair[string]](settings)

		assert.Equal(t, cache.Set(kvp1), true)
		assert.Equal(t, cache.Set(kvp2), true)
		assert.Equal(t, cache.Set(kvp3), true)

		count := 0

		cache.ForEach(func(value CacheItem[*cache_items.KeyValuePair[string]]) error {
			if value.HasValue() {
				count++
			}

			return nil
		})

		assert.Equal(t, count, 3)
	})
}
