# Benchmark results

## Attributes

|Test Attributes|Value|
|---------------|:-----|
|cache L1|32768|
|cache L2|262144|
|cache L3|52428800|
|concurrency|2|
|cpu|Intel(R) Xeon(R) CPU E5-2673 v4 @ 2.30GHz|
|goarch|amd64|
|goos|linux|
|pkg|github.com/DaanV2/High-Performance-Cache/benchmarks|

## Benchmark: duplicate items write test mapcache testing: 

|Name|Test Amount|Nanoseconds per operation|Items amount|Items per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|---:|---:|
|Writing-2|73740|78002|1000|0.01356|168|4|
|Writing#01-2|8344|784874|10000|1.198|168|4|

## Benchmark: duplicate items read test mapcache testing: 

|Name|Test Amount|Nanoseconds per operation|Items amount|Items per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|---:|---:|
|Reading-2|62907|93475|1000|0.01590|176|4|
|Reading#01-2|5996|1021681|10000|1.668|176|4|

## Benchmark: unique items write test mapcache testing: 

|Name|Test Amount|Nanoseconds per operation|Items amount|Items per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|---:|---:|
|Writing-2|57237|100574|1000|0.01747|168|4|
|Writing#01-2|5096|1214998|10000|1.962|168|4|

## Benchmark: unique items read test mapcache testing: 

|Name|Test Amount|Nanoseconds per operation|Items amount|Items per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|---:|---:|
|Reading-2|55276|104143|1000|0.01809|176|4|
|Reading#01-2|4873|1053619|10000|2.052|176|4|

