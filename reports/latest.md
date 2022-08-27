# Benchmark results

## Benchmark 0

|Test Attributes|Value|
|---------------|-----|
|concurrency|8|
|goos|windows|
|goarch|amd64|
|pkg|github.com/DaanV2/High-Performance-Cache/benchmarks|
|cpu|Intel(R) Core(TM) i7-7700K CPU @ 4.20GHz|

|Name|Test Amount|Nanoseconds per operation|Bytes per operation|Allocations per operation|
|----|---|---|---|---|
|DefaultWriteTest Cache using: map, Size: 10, Test: Writing-8|212872|5317|680|12|
|DefaultWriteTest Cache using: map, Size: 100, Test: Writing-8|85623|15367|616|11|
|DefaultWriteTest Cache using: map, Size: 1000, Test: Writing-8|10000|107338|554|10|
|DefaultWriteTest Cache using: map, Size: 10000, Test: Writing-8|993|1141165|553|10|
|DefaultReadTest Cache using: map, Size: 10, Test: Reading-8|231214|5097|688|12|
|DefaultReadTest Cache using: map, Size: 100, Test: Reading-8|132202|10552|624|11|
|DefaultReadTest Cache using: map, Size: 1000, Test: Reading-8|22171|55566|560|10|
|DefaultReadTest Cache using: map, Size: 10000, Test: Reading-8|2796|461322|560|10|