# Benchmark results

## Benchmark: DefaultWriteTest MapCache, Per Size: 

|Test Attributes|Value|
|---------------:|:-----|
|`concurrency`|2|
|`goos`|linux|
|`goarch`|amd64|
|`pkg`|github.com/DaanV2/High-Performance-Cache/benchmarks|
|`cpu`|Intel(R) Xeon(R) CPU E5-2673 v4 @ 2.30GHz|

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----|---:|---:|---:|---:|
|10, Test: Writing-2|564366|2127|168|4|
|100, Test: Writing-2|97178|12642|168|4|
|1000, Test: Writing-2|10000|116492|168|4|
|10000, Test: Writing-2|895|1394251|174|4|

## Benchmark: DefaultReadTest MapCache, Per Size 

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----|---|---|---|---|
|10, Test: Reading-2|682707|1824|176|4|
|100, Test: Reading-2|167664|7575|176|4|
|1000, Test: Reading-2|10000|102640|176|4|
|10000, Test: Reading-2|1023|1079809|176|4|

