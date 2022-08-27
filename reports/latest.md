# Benchmark results

## Benchmark: default write test mapcache, size per 

|Test Attributes|Value|
|---------------|:-----|
|goarch|amd64|
|pkg|github.com/DaanV2/High-Performance-Cache/benchmarks|
|cpu|Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz|
|cache L1|32768|
|cache L2|1310720|
|cache L3|50331648|
|concurrency|2|
|goos|linux|

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10, Test: Writing-2|782769|1524|168|4|
|100, Test: Writing-2|160269|6815|168|4|
|1000, Test: Writing-2|16856|71285|168|4|
|10000, Test: Writing-2|1417|948119|168|4|

## Benchmark: default read test mapcache, size per 

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10, Test: Writing-2|901734|1344|176|4|
|100, Test: Writing-2|243831|5234|176|4|
|1000, Test: Writing-2|12957|84955|176|4|
|10000, Test: Writing-2|1257|952746|176|4|

