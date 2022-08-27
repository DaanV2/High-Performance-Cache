package benchmarks

import (
	"fmt"
	"testing"

	"github.com/DaanV2/High-Performance-Cache/cache"
	"github.com/DaanV2/High-Performance-Cache/util"
)

func BenchmarkDefaultWriteTest(b *testing.B) {
	data := DataSet

	for _, size := range Sizes {
		testdata := data[:size]
		cache := newMapCache[*BenchMarkData](size)

		b.Run(Name("Cache using: map, Size: %v, Test: Writing", size), func(b *testing.B) {
			WriteTest(b, cache, testdata)
		})
	}
}

func BenchmarkDefaultReadTest(b *testing.B) {
	data := DataSet

	for _, size := range Sizes {
		testdata := data[:size]
		cache := newMapCache[*BenchMarkData](size)

		b.Run(Name("Cache using: map, Size: %v, Test: Reading", size), func(b *testing.B) {
			ReadTest(b, cache, testdata)
		})
	}
}

func WriteTest(b *testing.B, cache cache.Cache[*BenchMarkData], testdata []*BenchMarkData) {
	for i := 0; i < b.N; i++ {
		util.Parralel(testdata, func(item *BenchMarkData, index int, current []*BenchMarkData) {
			cache.Set(item)
		})
	}
}

func ReadTest(b *testing.B, cache cache.Cache[*BenchMarkData], testdata []*BenchMarkData) {
	//Fill test
	util.Parralel(testdata, func(item *BenchMarkData, index int, current []*BenchMarkData) {
		cache.Set(item)
	})

	for i := 0; i < b.N; i++ {
		util.Parralel(testdata, func(item *BenchMarkData, index int, current []*BenchMarkData) {
			if _, err := cache.Get(item.GetKey()); err != nil {
				b.Fatal("Could not get item from cache")
			}
		})
	}
}

func Name(name string, a ...any) string {
	return fmt.Sprintf(name, a...)
}
