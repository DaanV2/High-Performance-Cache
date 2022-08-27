package benchmarks

import (
	"fmt"
	"testing"

	"github.com/DaanV2/High-Performance-Cache/cache"
	"github.com/DaanV2/High-Performance-Cache/concurrency"
)


type TestSettings struct {
	Name string
	CreateCache func (size int) cache.Cache[*BenchmarkData]
}

func RunBenchMarks(b *testing.B, testdata []*BenchmarkData, settings []*TestSettings, benchmark func(b *testing.B, settings *TestSettings, testdata []*BenchmarkData)) {
	sizes := GenerateSizes(testdata)

	for _, size := range sizes {
		for _, settings := range settings {
			benchmark(b, settings, testdata[:size])
		}
	}

	fmt.Println("Done")
}

func WriteTest(b *testing.B, settings *TestSettings, testdata []*BenchmarkData) {	
	cache := settings.CreateCache(len(testdata))

	b.Run(Name("%s %v Items per test, testing: Writing", settings.Name, len(testdata)), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			concurrency.Parralel(testdata, func(item *BenchmarkData, index int, current []*BenchmarkData) {
				cache.Set(item)
			})
		}
	})
}

func ReadTest(b *testing.B, settings *TestSettings, testdata []*BenchmarkData) {
	cache := settings.CreateCache(len(testdata))

	//Fill test
	concurrency.Parralel(testdata, func(item *BenchmarkData, index int, current []*BenchmarkData) {
		cache.Set(item)
	})


	b.Run(Name("%s %v Items per test, testing: Reading", settings.Name, len(testdata)), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			concurrency.Parralel(testdata, func(item *BenchmarkData, index int, current []*BenchmarkData) {
				if _, err := cache.Get(item.GetKey()); err != nil {
					b.Fatal("Could not get item from cache")
				}
			})
		}
	})
}


func Name(name string, a ...any) string {
	return fmt.Sprintf(name, a...)
}