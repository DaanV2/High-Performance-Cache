# Benchmark results

|Test Attributes|Value|
|---------------|:-----|
|cache L1|32768||cache L2|262144||cache L3|31457280||concurrency|2||cpu|Intel(R) Xeon(R) CPU E5-2673 v3 @ 2.40GHz||goarch|amd64||goos|linux||pkg|github.com/DaanV2/High-Performance-Cache/benchmarks|
## Benchmark: duplicate items write test mapcache 

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10 Items per test, testing: Writing-2|643772|1909|168|4|
|100 Items per test, testing: Writing-2|118773|9953|168|4|
|1000 Items per test, testing: Writing-2|14566|77816|168|4|
|10000 Items per test, testing: Writing-2|1389|836346|168|4|

## Benchmark: duplicate items read test mapcache 

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10 Items per test, testing: Reading-2|740259|1714|176|4|
|100 Items per test, testing: Reading-2|181414|6587|176|4|
|1000 Items per test, testing: Reading-2|10000|106892|176|4|
|10000 Items per test, testing: Reading-2|1213|1010347|176|4|

## Benchmark: unique items write test mapcache 

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10 Items per test, testing: Writing-2|725613|1891|168|4|
|100 Items per test, testing: Writing-2|108181|10914|168|4|
|1000 Items per test, testing: Writing-2|10000|101189|168|4|
|10000 Items per test, testing: Writing-2|964|1292951|168|4|

## Benchmark: unique items read test mapcache 

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10 Items per test, testing: Reading-2|745600|1678|176|4|
|100 Items per test, testing: Reading-2|149745|6984|176|4|
|1000 Items per test, testing: Reading-2|10000|101185|176|4|
|10000 Items per test, testing: Reading-2|1213|947807|176|4|

