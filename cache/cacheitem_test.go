package cache

import (
	"testing"
	"time"

	util "github.com/DaanV2/High-Performance-Cache/util"
	"gotest.tools/assert"
)

func Test_Empty(t *testing.T) {
	item := EmptyCacheItem[*CacheItemString]()

	t.Run("HasValue is false", func(t *testing.T) {
		assert.Equal(t, item.HasValue(), false)
	})

	t.Run("GetHashCode is 0", func(t *testing.T) {
		assert.Equal(t, item.GetHashCode(), uint64(0))
	})

	t.Run("IsExipred is true, because time 0", func(t *testing.T) {
		assert.Equal(t, item.IsExpired(time.Now()), true)
		assert.Equal(t, item.IsExpired(time.UnixMicro(0)), true)
	})

	t.Run("GetValue is nil", func(t *testing.T) {
		var empty *CacheItemString
		v := item.GetValue()
		assert.Equal(t, v, empty)
	})
}

func Test_Value(t *testing.T) {
	key := "0123456789"
	hashcode := util.GetHashcode(key)
	item := NewCacheItem(time.Now(), NewCacheItemString(key))

	t.Run("HasValue is true", func(t *testing.T) {
		assert.Equal(t, item.HasValue(), true)
	})

	t.Run("GetHashCode is 0", func(t *testing.T) {
		assert.Equal(t, item.GetHashCode(), hashcode)
	})

	t.Run("GetValue is still its same item", func(t *testing.T) {
		v1 := string(*item.GetValue())

		assert.Equal(t, v1, key)
	})

	t.Run("IsExpired works as intented", func(t *testing.T) {
		now := time.Now()
		expires := now.Add(time.Second * 5)

		item := NewCacheItem(expires, NewCacheItemString(key))

		assert.Equal(t, item.IsExpired(now), false)

		now = now.Add(time.Second * 10)

		assert.Equal(t, item.IsExpired(now), true)
	})
}