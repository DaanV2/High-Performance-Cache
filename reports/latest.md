# Benchmark results

## Attributes

|Test Attributes|Value|
|---------------|:-----|
|cache L1|32768|
|cache L2|1310720|
|cache L3|50331648|
|concurrency|2|
|cpu|Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz|
|goarch|amd64|
|goos|linux|
|pkg|github.com/DaanV2/High-Performance-Cache/benchmarks|

## Benchmark: duplicate items write test mapcache testing: 

|Name|Test Amount|Nanoseconds per operation|Items amount|Items per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|---:|---:|
|Writing-2|791052|1567|10.00|0.0000126|168|4|
|Writing#01-2|170348|6988|100.0|0.0005870|168|4|
|Writing#02-2|19562|61558|1000|0.05112|168|4|
|Writing#03-2|2035|611546|10000|4.914|168|4|

## Benchmark: duplicate items read test mapcache testing: 

|Name|Test Amount|Nanoseconds per operation|Items amount|Items per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|---:|---:|
|Reading-2|871070|1398|10.00|0.0000115|176|4|
|Reading#01-2|258882|4743|100.0|0.0003863|176|4|
|Reading#02-2|12931|87963|1000|0.07733|176|4|
|Reading#03-2|1381|880949|10000|7.241|176|4|

## Benchmark: unique items write test mapcache testing: 

|Name|Test Amount|Nanoseconds per operation|Items amount|Items per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|---:|---:|
|Writing-2|817585|1567|10.00|0.0000122|168|4|
|Writing#01-2|174432|7016|100.0|0.0005733|168|4|
|Writing#02-2|17690|70176|1000|0.05653|168|4|
|Writing#03-2|1318|843226|10000|7.587|168|4|

## Benchmark: unique items read test mapcache testing: 

|Name|Test Amount|Nanoseconds per operation|Items amount|Items per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|---:|---:|
|Reading-2|897058|1397|10.00|0.0000111|176|4|
|Reading#01-2|242205|4619|100.0|0.0004129|176|4|
|Reading#02-2|13005|92258|1000|0.07689|176|4|
|Reading#03-2|1171|1020742|10000|8.540|176|4|

