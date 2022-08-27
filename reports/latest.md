# Benchmark results

## Benchmark: default write test mapcache, size per 

|Test Attributes|Value|
|---------------|:-----|
|goos|linux|
|goarch|amd64|
|pkg|github.com/DaanV2/High-Performance-Cache/benchmarks|
|cpu|Intel(R) Xeon(R) CPU E5-2673 v3 @ 2.40GHz|
|cache L1|32768|
|cache L2|262144|
|cache L3|31457280|
|concurrency|2|

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10, Test: Writing-2|580708|1963|168|4|
|100, Test: Writing-2|106957|10749|168|4|
|1000, Test: Writing-2|10000|103723|168|4|
|10000, Test: Writing-2|914|1290680|168|4|

## Benchmark: default read test mapcache, size per 

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10, Test: Writing-2|721532|1662|176|4|
|100, Test: Writing-2|177375|6884|176|4|
|1000, Test: Writing-2|14666|103368|176|4|
|10000, Test: Writing-2|1221|991004|176|4|

