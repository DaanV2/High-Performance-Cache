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
|10 Items per test, testing: Writing-2|662586|1805|168|4|
|100 Items per test, testing: Writing-2|136026|8700|168|4|
|1000 Items per test, testing: Writing-2|14012|83813|168|4|
|10000 Items per test, testing: Writing-2|1393|866536|171|4|

## Benchmark: duplicate items read test mapcache 

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10 Items per test, testing: Reading-2|834451|1578|176|4|
|100 Items per test, testing: Reading-2|196477|7092|176|4|
|1000 Items per test, testing: Reading-2|10000|110225|176|4|
|10000 Items per test, testing: Reading-2|990|1146188|176|4|

## Benchmark: unique items write test mapcache 

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10 Items per test, testing: Writing-2|675837|1797|168|4|
|100 Items per test, testing: Writing-2|136905|8958|168|4|
|1000 Items per test, testing: Writing-2|10000|100988|168|4|
|10000 Items per test, testing: Writing-2|810|1402179|168|4|

## Benchmark: unique items read test mapcache 

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10 Items per test, testing: Reading-2|751969|1565|176|4|
|100 Items per test, testing: Reading-2|185106|7346|176|4|
|1000 Items per test, testing: Reading-2|10000|107023|176|4|
|10000 Items per test, testing: Reading-2|1044|1051777|176|4|

