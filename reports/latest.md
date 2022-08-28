# Benchmark results

## Attributes

|Test Attributes|Value|
|---------------|:-----|
|cache L1|32768|
|cache L2|262144|
|cache L3|31457280|
|concurrency|2|
|cpu|Intel(R) Xeon(R) CPU E5-2673 v3 @ 2.40GHz|
|goarch|amd64|
|goos|linux|
|pkg|github.com/DaanV2/High-Performance-Cache/benchmarks|

## Benchmark: duplicate items write test mapcache 

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10 Items per test, testing: Writing-2|734497|1786|168|4|
|100 Items per test, testing: Writing-2|146502|9219|168|4|
|1000 Items per test, testing: Writing-2|15296|78222|168|4|
|10000 Items per test, testing: Writing-2|1543|794907|168|4|

## Benchmark: duplicate items read test mapcache 

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10 Items per test, testing: Reading-2|840171|1609|176|4|
|100 Items per test, testing: Reading-2|155718|7594|176|4|
|1000 Items per test, testing: Reading-2|12516|94182|176|4|
|10000 Items per test, testing: Reading-2|1272|974078|176|4|

## Benchmark: unique items write test mapcache 

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10 Items per test, testing: Writing-2|621537|1897|168|4|
|100 Items per test, testing: Writing-2|114837|9875|168|4|
|1000 Items per test, testing: Writing-2|12249|95741|168|4|
|10000 Items per test, testing: Writing-2|943|1240980|168|4|

## Benchmark: unique items read test mapcache 

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10 Items per test, testing: Reading-2|868062|1585|176|4|
|100 Items per test, testing: Reading-2|181400|7585|176|4|
|1000 Items per test, testing: Reading-2|12465|97858|176|4|
|10000 Items per test, testing: Reading-2|1202|990345|176|4|

