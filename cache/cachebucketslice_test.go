package cache

import (
	"testing"
)

func Test_CacheBucketSlice_Implements_Interface(t *testing.T) {
	settings := CacheBucketSliceSettings{
		MaxSize: 10,
	}

	c := NewCacheBucketSlice[*CacheItemString](settings)

	var _ Cache[*CacheItemString] = c
}