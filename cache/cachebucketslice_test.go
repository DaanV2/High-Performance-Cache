package cache

import (
	"fmt"
	"strconv"
	"testing"
	"time"
	"unsafe"

	"github.com/DaanV2/High-Performance-Cache/util"
	"github.com/stretchr/testify/assert"
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

func GetFilled() (*CacheBucketSlice[*CacheItemString], []*CacheItemString) {
	cache := NewCache()
	data := GenerateTestData()

	for _, item := range data {
		citem := NewCacheItem(time.Now(), item)
		cache.Set(citem)
	}

	return cache, data
}

func Test_CacheBucketSlice(t *testing.T) {
	t.Run("Validation Check", func(t *testing.T) {
		cache := NewCache()

		assert.NotNil(t, cache, "Cache is nil")
		assert.NotNil(t, cache.items, "Cache items is nil")
		assert.Equal(t, cache.itemCount, 0, "Cache item count is not 0")
	})

	t.Run("Can set all items", func(t *testing.T) {
		cache := NewCache()
		data := GenerateTestData()

		for index, item := range data {
			citem := NewCacheItem(time.Now(), item)
			result := cache.Set(citem)

			assert.Equal(t, result, true, "Set failed")
			assert.Equal(t, cache.Count(), index+1, "Cache item count is not correct")
		}
	})

	t.Run("Can get all items", func(t *testing.T) {
		cache, data := GetFilled()
		t.Logf("Cache: %v", cache)
		t.Logf("Data: %v", data)
		t.Logf("Cache Items: %v", cache.items)

		for index, item := range data {
			lookup := NewKeyLookup(item.GetKey())
			result, ok := cache.Get(lookup)

			startindex := cache.GetStartIndex(lookup.Hashcode)
			t.Logf("StartIndex: %d for item: %s", startindex, item.GetKey())

			assert.Equal(t, ok, true, fmt.Sprintf("Item not found: %d", index))
			assert.Equal(t, result.GetValue(), item, "Value not equal")
		}
	})

	t.Run("StartIndex cannot be larger then internal size", func(t *testing.T) {
		cache, _ := GetFilled()
		max := len(cache.items)

		for i := 0; i < max*3; i++ {
			startIndex := cache.GetStartIndex(uint64(i))

			assert.Less(t, startIndex, max, "StartIndex is larger then internal size")
		}
	})

	t.Run("ForEach works", func(t *testing.T) {
		cache, _ := GetFilled()

		count := 0
		cache.ForEach(func(value CacheItem[*CacheItemString]) error {
			if value.HasValue() {
				count++
			}

			return nil
		})

		assert.Equal(t, count, testSize, "ForEach did hit not all items")
	})

	t.Run("Is Full", func(t *testing.T) {
		cache, _ := GetFilled()

		assert.True(t, cache.IsFull())
	})

	t.Run("Clean", func(t *testing.T) {
		cache, _ := GetFilled()

		amount := cache.Clean(time.Now().Add(time.Hour * 3))
		assert.Equal(t, cache.Capacity(), amount)
		assert.Equal(t, cache.Count(), 0)
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
			assert.LessOrEqual(t, space, cacheSize)
		})
	}
}
