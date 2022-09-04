package cache

import (
	"strconv"
	"testing"
	"time"

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
