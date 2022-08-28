# Benchmark results

## Attributes

|Test Attributes|Value|
|---------------|:-----|
|cache L1|32768|
|cache L2|1048576|
|cache L3|37486592|
|concurrency|2|
|cpu|Intel(R) Xeon(R) Platinum 8171M CPU @ 2.60GHz|
|goarch|amd64|
|goos|linux|
|pkg|github.com/DaanV2/High-Performance-Cache/benchmarks|

## Benchmark: duplicate items write test mapcache 

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10 Items per test, testing: Writing-2|615016|1961|168|4|
|100 Items per test, testing: Writing-2|127594|9385|168|4|
|1000 Items per test, testing: Writing-2|13936|84284|168|4|
|10000 Items per test, testing: Writing-2|1381|876581|168|4|

## Benchmark: duplicate items read test mapcache 

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10 Items per test, testing: Reading-2|728786|1702|176|4|
|100 Items per test, testing: Reading-2|165951|7298|176|4|
|1000 Items per test, testing: Reading-2|10000|101800|176|4|
|10000 Items per test, testing: Reading-2|986|1209058|176|4|

## Benchmark: unique items write test mapcache 

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10 Items per test, testing: Writing-2|587079|2003|168|4|
|100 Items per test, testing: Writing-2|116918|10060|168|4|
|1000 Items per test, testing: Writing-2|10000|101470|168|4|
|10000 Items per test, testing: Writing-2|828|1471901|168|4|

## Benchmark: unique items read test mapcache 

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10 Items per test, testing: Reading-2|729410|1710|176|4|
|100 Items per test, testing: Reading-2|155730|7523|176|4|
|1000 Items per test, testing: Reading-2|10000|107070|176|4|
|10000 Items per test, testing: Reading-2|1048|1136662|176|4|

