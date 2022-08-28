# Benchmark results

## Attributes

|Test Attributes|Value|
|---------------|:-----|
|cache L1|32768|
|cache L2|1310720|
|cache L3|50331648|
|concurrency|2|
|cpu|Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz|
|goarch|amd64|
|goos|linux|
|pkg|github.com/DaanV2/High-Performance-Cache/benchmarks|

## Benchmark: duplicate items write test mapcache testing: 

|Name|Test Amount|Nanoseconds per operation|Items amount|Items per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|---:|---:|
|Writing-2|20097|60883|1000|0.04976|168|4|
|Writing#01-2|2012|587089|10000|4.970|168|4|

## Benchmark: duplicate items read test mapcache testing: 

|Name|Test Amount|Nanoseconds per operation|Items amount|Items per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|---:|---:|
|Reading-2|12618|87455|1000|0.07925|176|4|
|Reading#01-2|1138|1056239|10000|8.787|176|4|

## Benchmark: unique items write test mapcache testing: 

|Name|Test Amount|Nanoseconds per operation|Items amount|Items per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|---:|---:|
|Writing-2|17360|69643|1000|0.05760|168|4|
|Writing#01-2|1356|871653|10000|7.375|168|4|

## Benchmark: unique items read test mapcache testing: 

|Name|Test Amount|Nanoseconds per operation|Items amount|Items per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|---:|---:|
|Reading-2|12480|96826|1000|0.08013|176|4|
|Reading#01-2|1218|987740|10000|8.210|176|4|

