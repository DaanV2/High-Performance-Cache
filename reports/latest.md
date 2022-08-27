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
|10, Test: Writing-2|645619|1921|168|4|
|100, Test: Writing-2|93130|11031|168|4|
|1000, Test: Writing-2|10000|102073|168|4|
|10000, Test: Writing-2|1000|1206480|168|4|

## Benchmark: DEFAULT READ TEST MAPCACHE, PER SIZE 

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10, Test: Reading-2|736604|1640|176|4|
|100, Test: Reading-2|171710|6835|176|4|
|1000, Test: Reading-2|12313|98459|176|4|
|10000, Test: Reading-2|1132|1024624|176|4|

