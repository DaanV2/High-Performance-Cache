# Benchmark results

## Attributes

|Test Attributes|Value|
|---------------|:-----|
|cache L1|32768|
|cache L2|1048576|
|cache L3|37486592|
|concurrency|2|
|cpu|Intel(R) Xeon(R) Platinum 8272CL CPU @ 2.60GHz|
|goarch|amd64|
|goos|linux|
|pkg|github.com/DaanV2/High-Performance-Cache/benchmarks|

## Benchmark: duplicate items write test mapcache 

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10 Items per test, testing: Writing-2|803432|1552|168|4|
|100 Items per test, testing: Writing-2|155336|7786|168|4|
|1000 Items per test, testing: Writing-2|18386|66060|168|4|
|10000 Items per test, testing: Writing-2|1748|677626|168|4|

## Benchmark: duplicate items read test mapcache 

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10 Items per test, testing: Reading-2|931220|1356|176|4|
|100 Items per test, testing: Reading-2|196638|5876|176|4|
|1000 Items per test, testing: Reading-2|12603|95724|176|4|
|10000 Items per test, testing: Reading-2|1209|998409|176|4|

## Benchmark: unique items write test mapcache 

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10 Items per test, testing: Writing-2|819900|1567|168|4|
|100 Items per test, testing: Writing-2|140608|8749|168|4|
|1000 Items per test, testing: Writing-2|14599|78007|168|4|
|10000 Items per test, testing: Writing-2|1062|1105110|168|4|

## Benchmark: unique items read test mapcache 

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10 Items per test, testing: Reading-2|919068|1365|176|4|
|100 Items per test, testing: Reading-2|220015|5625|176|4|
|1000 Items per test, testing: Reading-2|12895|92262|176|4|
|10000 Items per test, testing: Reading-2|1119|1073795|176|4|

