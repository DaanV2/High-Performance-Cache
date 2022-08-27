# Benchmark results

|Test Attributes|Value|
|---------------|:-----|
cache L1: 32768cache L2: 1048576cache L3: 37486592concurrency: 2cpu: Intel(R) Xeon(R) Platinum 8171M CPU @ 2.60GHzgoarch: amd64goos: linuxpkg: github.com/DaanV2/High-Performance-Cache/benchmarks
## Benchmark: duplicate items write test mapcache 

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10 Items per test, testing: Writing-2|684678|1899|168|4|
|100 Items per test, testing: Writing-2|142282|8706|168|4|
|1000 Items per test, testing: Writing-2|14394|83365|168|4|
|10000 Items per test, testing: Writing-2|1315|879603|168|4|

## Benchmark: duplicate items read test mapcache 

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10 Items per test, testing: Reading-2|758938|1636|176|4|
|100 Items per test, testing: Reading-2|188647|6632|176|4|
|1000 Items per test, testing: Reading-2|10000|116933|176|4|
|10000 Items per test, testing: Reading-2|997|1219882|176|4|

## Benchmark: unique items write test mapcache 

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10 Items per test, testing: Writing-2|662408|1881|168|4|
|100 Items per test, testing: Writing-2|113260|9907|168|4|
|1000 Items per test, testing: Writing-2|10000|106207|168|4|
|10000 Items per test, testing: Writing-2|734|1524920|168|4|

## Benchmark: unique items read test mapcache 

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10 Items per test, testing: Reading-2|804410|1612|176|4|
|100 Items per test, testing: Reading-2|190862|6742|176|4|
|1000 Items per test, testing: Reading-2|10000|109863|176|4|
|10000 Items per test, testing: Reading-2|939|1242503|176|4|

