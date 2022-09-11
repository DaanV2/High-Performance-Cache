package cache

import (
	"testing"

	util "github.com/DaanV2/High-Performance-Cache/util"
	"gotest.tools/assert"
)

func Test_KeyLookup(t *testing.T) {
	keyStr := "0123456789"
	key := NewKeyLookup(keyStr)

	t.Run("Test key", func(t *testing.T) {
		assert.Equal(t, key.GetKey(), keyStr)
	})

	t.Run("Test hashcode", func(t *testing.T) {
		hashcode := util.GetHashcode(keyStr)
		assert.Equal(t, key.GetHashcode(), hashcode)
	})
}
