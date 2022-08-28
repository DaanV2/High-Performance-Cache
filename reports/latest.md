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
|10 Items per test, testing: Writing-2|688651|1842|168|4|
|100 Items per test, testing: Writing-2|121998|8978|168|4|
|1000 Items per test, testing: Writing-2|14230|82998|168|4|
|10000 Items per test, testing: Writing-2|1400|870630|169|4|

## Benchmark: duplicate items read test mapcache 

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10 Items per test, testing: Reading-2|713815|1659|176|4|
|100 Items per test, testing: Reading-2|180578|6714|176|4|
|1000 Items per test, testing: Reading-2|10000|103778|176|4|
|10000 Items per test, testing: Reading-2|1087|1083887|176|4|

## Benchmark: unique items write test mapcache 

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10 Items per test, testing: Writing-2|629646|1856|168|4|
|100 Items per test, testing: Writing-2|132177|9560|168|4|
|1000 Items per test, testing: Writing-2|12094|99544|168|4|
|10000 Items per test, testing: Writing-2|825|1469417|168|4|

## Benchmark: unique items read test mapcache 

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10 Items per test, testing: Reading-2|744636|1728|176|4|
|100 Items per test, testing: Reading-2|181286|6741|176|4|
|1000 Items per test, testing: Reading-2|10000|109925|176|4|
|10000 Items per test, testing: Reading-2|1018|1173706|176|4|

