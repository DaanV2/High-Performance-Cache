package benchmarks

import (
	"testing"

	"github.com/DaanV2/High-Performance-Cache/cache_items"
)

func Test_MapCache(t *testing.T) {
	kvp := cache_items.NewKeyValuePair("key", "value")

	t.Run("Can make cache", func(t *testing.T) {
		cache := newMapCache[*cache_items.KeyValuePair[string]](10)

		if cache == nil {
			t.Error("Cache is nil")
		}
	})

	t.Run("Can Set", func(t *testing.T) {
		cache := newMapCache[*cache_items.KeyValuePair[string]](10)

		if !cache.Set(kvp) {
			t.Error("Cache.Set returned false")
		}
	})
}
