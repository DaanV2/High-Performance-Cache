# Benchmark results

## Attributes

|Test Attributes|Value|
|---------------|:-----|
|cache L1|32768|
|cache L2|262144|
|cache L3|8388608|
|concurrency|8|
|cpu|Intel(R) Core(TM) i7-7700K CPU @ 4.20GHz|
|goarch|amd64|
|goos|windows|
|pkg|github.com/DaanV2/High-Performance-Cache/benchmarks|

## Benchmark: duplicate items write test mapcache 

|Name|Test Amount|Nanoseconds per operation|Items amount|Items per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|---:|---:|
|10 Items per test, testing: Writing-8|220911|5409|10.00|0.0000453|680|12|
|100 Items per test, testing: Writing-8|94731|12963|100.0|0.001056|616|11|
|1000 Items per test, testing: Writing-8|10000|112775|1000|0.1000|552|10|
|10000 Items per test, testing: Writing-8|1129|1059373|10000|8.857|552|10|

## Benchmark: duplicate items read test mapcache 

|Name|Test Amount|Nanoseconds per operation|Items amount|Items per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|---:|---:|
|10 Items per test, testing: Reading-8|229153|5237|10.00|0.0000436|688|12|
|100 Items per test, testing: Reading-8|143175|8703|100.0|0.0006984|624|11|
|1000 Items per test, testing: Reading-8|22694|53145|1000|0.04406|560|10|
|10000 Items per test, testing: Reading-8|2614|453535|10000|3.826|560|10|

## Benchmark: unique items write test mapcache 

|Name|Test Amount|Nanoseconds per operation|Items amount|Items per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|---:|---:|
|10 Items per test, testing: Writing-8|214804|5227|10.00|0.0000466|680|12|
|100 Items per test, testing: Writing-8|94254|12822|100.0|0.001061|616|11|
|1000 Items per test, testing: Writing-8|10000|119506|1000|0.1000|552|10|
|10000 Items per test, testing: Writing-8|1014|1183974|10000|9.862|552|10|

## Benchmark: unique items read test mapcache 

|Name|Test Amount|Nanoseconds per operation|Items amount|Items per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|---:|---:|
|10 Items per test, testing: Reading-8|233600|5035|10.00|0.0000428|688|12|
|100 Items per test, testing: Reading-8|139885|8622|100.0|0.0007149|624|11|
|1000 Items per test, testing: Reading-8|22806|52974|1000|0.04385|560|10|
|10000 Items per test, testing: Reading-8|2558|465783|10000|3.909|560|10|

