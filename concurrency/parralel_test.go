package concurrency

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

type testData struct {
	Value int
}

func Test_Parralel_CanSet(t *testing.T) {
	for size := 17; size < 1000; size *= 2 {
		t.Run(fmt.Sprintf("Parralel Can Set, %v size", size), func(t *testing.T) {
			values := make([]*testData, size)

			for i := 0; i < len(values); i++ {
				values[i] = &testData{Value: 0}
			}

			Parralel(values, func(item *testData, index int, current []*testData) {
				item.Value = 1
			})

			count := 0
			for _, item := range values {
				if item.Value == 1 {
					count++
				}
			}

			assert.Equal(t, count, len(values))
		})
	}
}
