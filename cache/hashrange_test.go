package cache

import (
	"fmt"
	"math"
	"testing"

	"gotest.tools/assert"
)

func Test_IsInRange(t *testing.T) {

	t.Run("Returns false when new", func(t *testing.T) {
		hr := NewHashRange()

		max := uint64(math.MaxUint64)
		step := max / 100
		max -= step * 2

		for i := uint64(0); i < max; i += step {
			t.Run(fmt.Sprintf("Returns false when new with value: %d", i), func(t *testing.T) {
				assert.Equal(t, hr.IsInRange(uint64(i)), false)
			})
		}
	})

	t.Run("Test egdes", func(t *testing.T) {
		hr := NewHashRangeFrom(10, 20)

		assert.Equal(t, hr.IsInRange(10), true)
		assert.Equal(t, hr.IsInRange(20), true)

		assert.Equal(t, hr.IsInRange(11), true)
		assert.Equal(t, hr.IsInRange(19), true)

		assert.Equal(t, hr.IsInRange(9), false)
		assert.Equal(t, hr.IsInRange(21), false)
	})
}

func Test_Update(t *testing.T) {	


	t.Run("Test update with egdes", func(t *testing.T) {
		hr := NewHashRangeFrom(10, 20)

		assert.Equal(t, hr.IsInRange(10), true)
		assert.Equal(t, hr.IsInRange(20), true)

		assert.Equal(t, hr.IsInRange(11), true)
		assert.Equal(t, hr.IsInRange(19), true)

		assert.Equal(t, hr.IsInRange(9), false)
		assert.Equal(t, hr.IsInRange(21), false)

		// Update
		hr.UpdateRange(30)

		assert.Equal(t, hr.IsInRange(10), true)
		assert.Equal(t, hr.IsInRange(30), true)

		assert.Equal(t, hr.IsInRange(11), true)
		assert.Equal(t, hr.IsInRange(29), true)

		assert.Equal(t, hr.IsInRange(9), false)
		assert.Equal(t, hr.IsInRange(31), false)

	})
}