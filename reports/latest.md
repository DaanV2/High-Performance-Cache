# Benchmark results

## Benchmark: default write test mapcache, per size 

|Test Attributes|Value|
|---------------|:-----|
|pkg|github.com/DaanV2/High-Performance-Cache/benchmarks|
|cpu|Intel(R) Xeon(R) CPU E5-2673 v3 @ 2.40GHz|
|concurrency|2|
|goos|linux|
|goarch|amd64|

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10, Test: Writing-2|670378|1804|168|4|
|100, Test: Writing-2|100370|11258|168|4|
|1000, Test: Writing-2|10000|101008|168|4|
|10000, Test: Writing-2|963|1203674|168|4|

## Benchmark: default read test mapcache, per size 

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10, Test: Reading-2|746310|1598|176|4|
|100, Test: Reading-2|170182|6585|176|4|
|1000, Test: Reading-2|12728|94610|176|4|
|10000, Test: Reading-2|1080|995172|176|4|

