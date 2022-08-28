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
|10 Items per test, testing: Writing-2|636854|1940|168|4|
|100 Items per test, testing: Writing-2|126295|9212|168|4|
|1000 Items per test, testing: Writing-2|13914|85008|168|4|
|10000 Items per test, testing: Writing-2|1290|881996|168|4|

## Benchmark: duplicate items read test mapcache 

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10 Items per test, testing: Reading-2|764448|1669|176|4|
|100 Items per test, testing: Reading-2|169272|7227|176|4|
|1000 Items per test, testing: Reading-2|10000|112541|176|4|
|10000 Items per test, testing: Reading-2|1078|1045952|176|4|

## Benchmark: unique items write test mapcache 

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10 Items per test, testing: Writing-2|588066|2086|168|4|
|100 Items per test, testing: Writing-2|108153|10916|168|4|
|1000 Items per test, testing: Writing-2|10000|105403|168|4|
|10000 Items per test, testing: Writing-2|879|1483078|168|4|

## Benchmark: unique items read test mapcache 

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10 Items per test, testing: Reading-2|707022|1673|176|4|
|100 Items per test, testing: Reading-2|147648|7154|176|4|
|1000 Items per test, testing: Reading-2|10000|104402|176|4|
|10000 Items per test, testing: Reading-2|1111|1008909|176|4|

