# Benchmark results

## Benchmark: DEFAULT WRITE TEST MAPCACHE, PER SIZE 

|Test Attributes|Value|
|---------------|:-----|
|concurrency|2|
|goos|linux|
|goarch|amd64|
|pkg|github.com/DaanV2/High-Performance-Cache/benchmarks|
|cpu|Intel(R) Xeon(R) Platinum 8171M CPU @ 2.60GHz|

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----|---:|---:|---:|---:|
|10, Test: Writing-2|709314|1548|168|4|
|100, Test: Writing-2|177890|7122|168|4|
|1000, Test: Writing-2|15303|71227|168|4|
|10000, Test: Writing-2|1261|1016107|173|4|

## Benchmark: DEFAULT READ TEST MAPCACHE, PER SIZE 

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----|---:|---:|---:|---:|
|10, Test: Reading-2|1000000|1301|176|4|
|100, Test: Reading-2|237285|5578|176|4|
|1000, Test: Reading-2|14686|78661|176|4|
|10000, Test: Reading-2|1430|892900|176|4|

