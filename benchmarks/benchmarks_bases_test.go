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
	sizes := Sizes()

	for _, size := range sizes {
		for _, settings := range settings {
			benchmark(b, settings, testdata[:size])
		}
	}

	fmt.Println("Done")
}

func WriteTest(b *testing.B, settings *TestSettings, testdata []*BenchmarkData) {	
	cache := settings.CreateCache(len(testdata))

	b.Run(Name("%s testing: Writing", settings.Name), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			concurrency.Parralel(testdata, func(item *BenchmarkData, index int, current []*BenchmarkData) {
				cache.Set(item)
			})
		}

		b.ReportMetric(float64(len(testdata)), "N")
		b.ReportMetric(float64(len(testdata)) / float64(b.N), "N/op")
	})
}

func ReadTest(b *testing.B, settings *TestSettings, testdata []*BenchmarkData) {
	cache := settings.CreateCache(len(testdata))

	//Fill test
	concurrency.Parralel(testdata, func(item *BenchmarkData, index int, current []*BenchmarkData) {
		cache.Set(item)
	})


	b.Run(Name("%s testing: Reading", settings.Name), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			concurrency.Parralel(testdata, func(item *BenchmarkData, index int, current []*BenchmarkData) {
				if _, ok := cache.Get(item.GetKey()); !ok {
					b.Fatal("Could not get item from cache")
				}
			})
		}

		b.ReportMetric(float64(len(testdata)), "N")
		b.ReportMetric(float64(len(testdata)) / float64(b.N), "N/op")
	})
}


func Name(name string, a ...any) string {
	return fmt.Sprintf(name, a...)
}