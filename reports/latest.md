# Benchmark results

## Benchmark: DEFAULT WRITE TEST MAPCACHE, PER SIZE 

|Test Attributes|Value|
|---------------|:-----|
|goarch|amd64|
|pkg|github.com/DaanV2/High-Performance-Cache/benchmarks|
|cpu|Intel(R) Xeon(R) CPU E5-2673 v4 @ 2.30GHz|
|concurrency|2|
|goos|linux|

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10, Test: Writing-2|557539|1994|168|4|
|100, Test: Writing-2|102694|11603|168|4|
|1000, Test: Writing-2|10000|113592|168|4|
|10000, Test: Writing-2|943|1312225|169|4|

## Benchmark: DEFAULT READ TEST MAPCACHE, PER SIZE 

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10, Test: Reading-2|687027|1731|176|4|
|100, Test: Reading-2|177386|7144|176|4|
|1000, Test: Reading-2|10000|100742|176|4|
|10000, Test: Reading-2|1102|1030561|176|4|

