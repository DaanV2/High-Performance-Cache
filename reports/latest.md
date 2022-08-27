# Benchmark results

|Test Attributes|Value|
|---------------|:-----|
|cache L1|32768||cache L2|1048576||cache L3|37486592||concurrency|2||cpu|Intel(R) Xeon(R) Platinum 8272CL CPU @ 2.60GHz||goarch|amd64||goos|linux||pkg|github.com/DaanV2/High-Performance-Cache/benchmarks|
## Benchmark: duplicate items write test mapcache 

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10 Items per test, testing: Writing-2|820842|1526|168|4|
|100 Items per test, testing: Writing-2|164395|7590|168|4|
|1000 Items per test, testing: Writing-2|17922|66880|168|4|
|10000 Items per test, testing: Writing-2|1647|701248|168|4|

## Benchmark: duplicate items read test mapcache 

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10 Items per test, testing: Reading-2|917160|1349|176|4|
|100 Items per test, testing: Reading-2|243684|5173|176|4|
|1000 Items per test, testing: Reading-2|13560|88254|176|4|
|10000 Items per test, testing: Reading-2|1230|918552|176|4|

## Benchmark: unique items write test mapcache 

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10 Items per test, testing: Writing-2|805815|1546|168|4|
|100 Items per test, testing: Writing-2|124360|9020|168|4|
|1000 Items per test, testing: Writing-2|14306|82465|168|4|
|10000 Items per test, testing: Writing-2|1048|1102254|168|4|

## Benchmark: unique items read test mapcache 

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|
|10 Items per test, testing: Reading-2|908820|1347|176|4|
|100 Items per test, testing: Reading-2|231255|5311|176|4|
|1000 Items per test, testing: Reading-2|12858|94924|176|4|
|10000 Items per test, testing: Reading-2|1311|926432|176|4|

