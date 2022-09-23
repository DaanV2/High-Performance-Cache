package benchmarks

import "github.com/DaanV2/High-Performance-Cache/cache"

var CachesToTest = []*TestSettings{
	{
		Name:        "MapCache",
		CreateCache: newMapCache[*BenchmarkData],
	},
	{
		Name: "FixedCache",
		CreateCache: func(size int) cache.Cache[*BenchmarkData] {
			settings := cache.DefaultFixedCacheSettings[*BenchmarkData](size)

			return cache.NewFixedCache[*BenchmarkData](settings)
		},
	},
}
