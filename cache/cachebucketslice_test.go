package cache

import (
	"strconv"
	"testing"
	"time"
	"unsafe"

	"github.com/DaanV2/High-Performance-Cache/util"
	"gotest.tools/assert"
)

const testSize = 10

func NewCache() *CacheBucketSlice[*CacheItemString] {
	return NewCacheBucketSlice[*CacheItemString](CacheBucketSliceSettings{
		MaxSize: testSize,
	})
}

func GenerateTestData() []*CacheItemString {
	result := make([]*CacheItemString, 0, testSize)

	for i := 0; i < testSize; i++ {
		key := "key" + strconv.FormatInt(int64(i), 10)
		item := CacheItemString(key)
		result = append(result, &item)
	}

	return result
}

func Test_CacheBucketSlice(t *testing.T) {
	cache := NewCache()
	data := GenerateTestData()

	t.Run("Can set all items", func(t *testing.T) {
		for _, item := range data {
			citem := NewCacheItem(time.Now(), item)
			result := cache.Set(citem)

			assert.Equal(t, result, true)
		}
	})
}

func Test_CacheBucketSliceSettings(t *testing.T) {
	cachesTarget := []util.CacheKind{
		util.CacheL1,
		util.CacheL2,
		util.CacheL3,
	}

	kvp := NewKeyValuePair("key", "value")

	for _, cacheTarget := range cachesTarget {
		t.Run("Slice size is fit for cache: "+cacheTarget.String(), func(t *testing.T) {
			settings := DefaultBucketSettings[*KeyValuePair[string]](cacheTarget)
			size := int64(unsafe.Sizeof(kvp))

			space := size * int64(settings.MaxSize)
			cacheSize := int64(cacheTarget.GetCacheSize())

			t.Logf("Space: %d, CacheSize: %d", space, cacheSize)
			assert.Assert(t, space <= cacheSize)
		})
	}
}
