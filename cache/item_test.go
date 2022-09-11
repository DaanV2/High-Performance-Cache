package cache

import (
	"testing"
	"time"

	util "github.com/DaanV2/High-Performance-Cache/util"
	"gotest.tools/assert"

	"github.com/DaanV2/High-Performance-Cache/cache_items"
)

func Test_Empty(t *testing.T) {
	item := EmptyCacheItem[*cache_items.CacheItemString]()

	t.Run("HasValue is false", func(t *testing.T) {
		assert.Equal(t, item.HasValue(), false)
	})

	t.Run("GetHashcode is 0", func(t *testing.T) {
		assert.Equal(t, item.GetHashcode(), uint64(0))
	})

	t.Run("IsExpired is true, because time 0", func(t *testing.T) {
		assert.Equal(t, item.IsExpired(time.Now()), true)
		assert.Equal(t, item.IsExpired(time.UnixMicro(0)), true)
	})

	t.Run("GetValue is nil", func(t *testing.T) {
		var empty *cache_items.CacheItemString
		v := item.GetValue()
		assert.Equal(t, v, empty)
	})
}

func Test_Value(t *testing.T) {
	key := "0123456789"
	hashcode := util.GetHashcode(key)
	item := NewCacheItem(time.Now(), cache_items.NewCacheItemString(key))

	t.Run("HasValue is true", func(t *testing.T) {
		assert.Equal(t, item.HasValue(), true)
	})

	t.Run("GetHashcode is 0", func(t *testing.T) {
		assert.Equal(t, item.GetHashcode(), hashcode)
	})

	t.Run("GetValue is still its same item", func(t *testing.T) {
		v1 := string(*item.GetValue())

		assert.Equal(t, v1, key)
	})

	t.Run("IsExpired works as intended", func(t *testing.T) {
		now := time.Now()
		expires := now.Add(time.Second * 5)

		item := NewCacheItem(expires, cache_items.NewCacheItemString(key))

		assert.Equal(t, item.IsExpired(now), false)

		now = now.Add(time.Second * 10)

		assert.Equal(t, item.IsExpired(now), true)
	})
}

func Test_IsExpired(t *testing.T) {
	expiresAfter := time.Now()
	item := NewCacheItem(expiresAfter, cache_items.NewCacheItemString("0123456789"))

	t.Run("IsExpired is true when time is after expiresAfter", func(t *testing.T) {
		assert.Equal(t, item.IsExpired(expiresAfter.Add(time.Second*10)), true)
	})

	t.Run("IsExpired is false when time is before expiresAfter", func(t *testing.T) {
		assert.Equal(t, item.IsExpired(expiresAfter.Add(time.Second*-10)), false)
	})
}

func Test_CanPlaceHere(t *testing.T) {
	var time0 time.Time

	key1 := cache_items.NewCacheItemString("0123456789")
	key2 := cache_items.NewCacheItemString("abcdefgh")

	filledItem := NewCacheItem(time.Now(), key1)
	//filledItem2 := NewCacheItem(time.Now(), key2)
	emptyItem := CacheItem[*cache_items.CacheItemString]{}

	t.Run("IsMatch returns true if empty", func(t *testing.T) {
		assert.Equal(t, emptyItem.CanPlaceHere(time0, filledItem), true)
	})

	t.Run("IsMatch returns true if expired", func(t *testing.T) {
		timeNow := time.Now()
		expiredItem := NewCacheItem(timeNow.Add(time.Hour*-1), key2)

		assert.Equal(t, expiredItem.CanPlaceHere(timeNow, filledItem), true)
	})

	t.Run("IsMatch returns false if not expired and not empty", func(t *testing.T) {
		expiredItem := NewCacheItem(time.Now().Add(time.Hour*1), key2)

		assert.Equal(t, expiredItem.CanPlaceHere(time0, filledItem), false)
	})

	t.Run("IsMatch returns true if a match", func(t *testing.T) {
		item := NewCacheItem(time.Now(), key1)

		assert.Equal(t, item.CanPlaceHere(time0, filledItem), true)
	})
}
