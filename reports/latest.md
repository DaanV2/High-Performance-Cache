# Benchmark results

## Benchmark: Default Write Test MapCache, Per Size 

|Test Attributes|Value|
|---------------|:-----|
|goos|linux|
|goarch|amd64|
|pkg|github.com/DaanV2/High-Performance-Cache/benchmarks|
|cpu|Intel(R) Xeon(R) Platinum 8171M CPU @ 2.60GHz|
|concurrency|2|

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----|---:|---:|---:|---:|
|10, Test: Writing-2|637848|1926|168|4|
|100, Test: Writing-2|117382|9647|168|4|
|1000, Test: Writing-2|12832|97995|168|4|
|10000, Test: Writing-2|770|1440182|168|4|

## Benchmark: Default Read Test MapCache, Per Size 

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----|---:|---:|---:|---:|
|10, Test: Reading-2|676890|1659|176|4|
|100, Test: Reading-2|189848|6993|176|4|
|1000, Test: Reading-2|10000|109543|176|4|
|10000, Test: Reading-2|1015|1149007|176|4|

