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
|goos|windows|
|pkg|github.com/DaanV2/High-Performance-Cache/benchmarks|

## Benchmark: duplicate items write test mapcache 

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10 Items per test, testing: Writing-2|674824|1957|168|4|
|100 Items per test, testing: Writing-2|122132|10084|168|4|
|1000 Items per test, testing: Writing-2|15541|73555|168|4|
|10000 Items per test, testing: Writing-2|1586|789149|168|4|

## Benchmark: duplicate items read test mapcache 

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10 Items per test, testing: Reading-2|717018|1724|176|4|
|100 Items per test, testing: Reading-2|168006|7079|176|4|
|1000 Items per test, testing: Reading-2|13477|90722|176|4|
|10000 Items per test, testing: Reading-2|1131|1060812|176|4|

## Benchmark: unique items write test mapcache 

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10 Items per test, testing: Writing-2|611095|1946|168|4|
|100 Items per test, testing: Writing-2|115588|10910|168|4|
|1000 Items per test, testing: Writing-2|10000|102735|168|4|
|10000 Items per test, testing: Writing-2|920|1293559|168|4|

## Benchmark: unique items read test mapcache 

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10 Items per test, testing: Reading-2|648715|1755|176|4|
|100 Items per test, testing: Reading-2|183188|6538|176|4|
|1000 Items per test, testing: Reading-2|10000|104402|176|4|
|10000 Items per test, testing: Reading-2|1150|1003665|176|4|

