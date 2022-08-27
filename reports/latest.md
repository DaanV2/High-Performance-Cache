# Benchmark results

|Test Attributes|Value|
|---------------|:-----|
|goarch|amd64|
|pkg|github.com/DaanV2/High-Performance-Cache/benchmarks|
|cpu|Intel(R) Xeon(R) CPU E5-2673 v4 @ 2.30GHz|
|cache L1|32768|
|cache L2|262144|
|cache L3|52428800|
|concurrency|2|
|goos|linux|

## Benchmark: duplicate items write test mapcache 

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10 Items per test, testing: Writing-2|591194|2102|168|4|
|100 Items per test, testing: Writing-2|105813|10922|168|4|
|1000 Items per test, testing: Writing-2|14602|87278|168|4|
|10000 Items per test, testing: Writing-2|1362|885557|168|4|

## Benchmark: duplicate items read test mapcache 

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10 Items per test, testing: Reading-2|681489|1870|176|4|
|100 Items per test, testing: Reading-2|167390|7457|176|4|
|1000 Items per test, testing: Reading-2|10000|107659|176|4|
|10000 Items per test, testing: Reading-2|1142|1051712|176|4|

## Benchmark: unique items write test mapcache 

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10 Items per test, testing: Writing-2|590065|2021|168|4|
|100 Items per test, testing: Writing-2|100728|11547|168|4|
|1000 Items per test, testing: Writing-2|10000|112555|168|4|
|10000 Items per test, testing: Writing-2|865|1377990|168|4|

## Benchmark: unique items read test mapcache 

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10 Items per test, testing: Reading-2|691371|1782|176|4|
|100 Items per test, testing: Reading-2|161805|7205|176|4|
|1000 Items per test, testing: Reading-2|10000|109356|176|4|
|10000 Items per test, testing: Reading-2|1203|1080604|176|4|

