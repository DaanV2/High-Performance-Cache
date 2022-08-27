# Benchmark results

## Benchmark: duplicate items write test mapcache 

|Test Attributes|Value|
|---------------|:-----|
|cache L1|32768|
|cache L2|262144|
|cache L3|8388608|
|concurrency|8|
|goos|windows|
|goarch|amd64|
|pkg|github.com/DaanV2/High-Performance-Cache/benchmarks|
|cpu|Intel(R) Core(TM) i7-7700K CPU @ 4.20GHz|

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10 Items per test, testing: Writing-8|218666|5342|680|12|
|100 Items per test, testing: Writing-8|76384|15045|616|11|
|1000 Items per test, testing: Writing-8|10000|104495|552|10|
|10000 Items per test, testing: Writing-8|1054|1176858|554|10|

## Benchmark: duplicate items read test mapcache 

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10 Items per test, testing: Reading-8|222810|5399|688|12|
|100 Items per test, testing: Reading-8|124026|9768|624|11|
|1000 Items per test, testing: Reading-8|20616|53995|560|10|
|10000 Items per test, testing: Reading-8|2934|403654|560|10|

## Benchmark: unique items write test mapcache 

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10 Items per test, testing: Writing-8|214858|5351|680|12|
|100 Items per test, testing: Writing-8|76382|16704|616|11|
|1000 Items per test, testing: Writing-8|10000|127778|552|10|
|10000 Items per test, testing: Writing-8|769|1374347|552|10|

## Benchmark: unique items read test mapcache 

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10 Items per test, testing: Reading-8|200473|5219|688|12|
|100 Items per test, testing: Reading-8|126634|10480|624|11|
|1000 Items per test, testing: Reading-8|21892|54910|560|10|
|10000 Items per test, testing: Reading-8|2733|427837|560|10|

