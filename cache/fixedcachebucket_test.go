package cache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_FixedCacheBucket_Creation(t *testing.T) {
	settings := DefaultFixedCacheSettings[*KeyValuePair[string]](10)
	settings.Cleaning.AutoClean = false

	t.Run("Default Settings", func(t *testing.T) {
		cache := NewFixedCacheBucket[*KeyValuePair[string]](10, settings.Buckets)

		if cache == nil {
			t.Error("Cache is nil")
		}
	})
}

func Test_FixedCacheBucket_Set(t *testing.T) {
	settings := DefaultFixedCacheSettings[*KeyValuePair[string]](10)
	settings.Cleaning.AutoClean = false
	kvp := NewKeyValuePair("key", "value")
	item := NewCacheItem(time.Now(), kvp)

	t.Run("Set item", func(t *testing.T) {
		cache := NewFixedCacheBucket[*KeyValuePair[string]](10, settings.Buckets)

		assert.Equal(t, cache.Set(item), true)
	})

	t.Run("Set item with existing key", func(t *testing.T) {
		cache := NewFixedCacheBucket[*KeyValuePair[string]](10, settings.Buckets)

		cache.Set(item)
		assert.Equal(t, cache.Set(item), true)
	})
}

func Test_FixedCacheBucket_Get(t *testing.T) {
	settings := DefaultFixedCacheSettings[*KeyValuePair[string]](10)
	settings.Cleaning.AutoClean = false
	kvp := NewKeyValuePair("key", "value")
	item := NewCacheItem(time.Now(), kvp)

	t.Run("Get item", func(t *testing.T) {
		cache := NewFixedCacheBucket[*KeyValuePair[string]](10, settings.Buckets)

		assert.Equal(t, cache.Set(item), true)
		item, ok := cache.Get(NewKeyLookup("key"))

		assert.Equal(t, ok, true)
		assert.Equal(t, item.GetValue(), kvp)
	})
}

func Test_FixedCacheBucket_ForEach(t *testing.T) {
	settings := DefaultFixedCacheSettings[*KeyValuePair[string]](10)
	settings.Cleaning.AutoClean = false
	kvp1 := NewCacheItem(time.Now(), NewKeyValuePair("key1", "value1"))
	kvp2 := NewCacheItem(time.Now(), NewKeyValuePair("key2", "value2"))
	kvp3 := NewCacheItem(time.Now(), NewKeyValuePair("key3", "value3"))

	t.Run("Get item", func(t *testing.T) {
		cache := NewFixedCacheBucket[*KeyValuePair[string]](10, settings.Buckets)

		assert.Equal(t, cache.Set(kvp1), true)
		assert.Equal(t, cache.Set(kvp2), true)
		assert.Equal(t, cache.Set(kvp3), true)

		count := 0

		cache.ForEach(func(value CacheItem[*KeyValuePair[string]]) error {
			if value.HasValue() {
				count++
			}

			return nil
		})

		assert.Equal(t, count, 3)
	})
}
