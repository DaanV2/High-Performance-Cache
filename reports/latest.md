# Benchmark results

## Benchmark: default write test mapcache, size per 

|Test Attributes|Value|
|---------------|:-----|
|pkg|github.com/DaanV2/High-Performance-Cache/benchmarks|
|cpu|Intel(R) Xeon(R) Platinum 8171M CPU @ 2.60GHz|
|cache L1|32768|
|cache L2|32768|
|cache L3|32768|
|concurrency|2|
|goos|linux|
|goarch|amd64|

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10, Test: Writing-2|749916|1828|168|4|
|100, Test: Writing-2|117913|9642|168|4|
|1000, Test: Writing-2|12924|96334|168|4|
|10000, Test: Writing-2|798|1356625|168|4|

## Benchmark: default read test mapcache, size per 

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10, Test: Writing-2|759469|1691|176|4|
|100, Test: Writing-2|156938|8004|176|4|
|1000, Test: Writing-2|12105|112873|176|4|
|10000, Test: Writing-2|1044|1133093|176|4|

