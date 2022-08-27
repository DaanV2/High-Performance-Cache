# Benchmark results

## Benchmark: DefaultWriteTest MapCache, Per Size 

|Test Attributes|Value|
|---------------|:-----|
|cpu|Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz|
|concurrency|2|
|goos|linux|
|goarch|amd64|
|pkg|github.com/DaanV2/High-Performance-Cache/benchmarks|

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----|---:|---:|---:|---:|
|10, Test: Writing-2|828039|1504|168|4|
|100, Test: Writing-2|171572|6602|168|4|
|1000, Test: Writing-2|17292|67566|168|4|
|10000, Test: Writing-2|1351|898551|168|4|

## Benchmark: DefaultReadTest MapCache, Per Size 

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----|---:|---:|---:|---:|
|10, Test: Reading-2|927782|1342|176|4|
|100, Test: Reading-2|230473|4746|176|4|
|1000, Test: Reading-2|12076|98599|176|4|
|10000, Test: Reading-2|1089|1084074|176|4|

