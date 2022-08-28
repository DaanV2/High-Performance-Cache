# Benchmark results

## Attributes

|Test Attributes|Value|
|---------------|:-----|
|cache L1|32768|
|cache L2|1048576|
|cache L3|37486592|
|concurrency|2|
|cpu|Intel(R) Xeon(R) Platinum 8171M CPU @ 2.60GHz|
|goarch|amd64|
|goos|linux|
|pkg|github.com/DaanV2/High-Performance-Cache/benchmarks|

## Benchmark: duplicate items write test mapcache testing: 

|Name|Test Amount|Nanoseconds per operation|Items amount|Items per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|---:|---:|
|Writing-2|70599|87481|1000|0.01416|168|4|
|Writing#01-2|35038|172401|2000|0.05708|168|4|
|Writing#02-2|13726|442836|5000|0.3643|168|4|
|Writing#03-2|6622|889082|10000|1.510|168|4|
|Writing#04-2|3344|1810290|20000|5.981|168|4|
|Writing#05-2|1369|4570376|50000|36.52|168|4|
|Writing#06-2|660|8997335|100000|151.5|168|4|

## Benchmark: duplicate items read test mapcache testing: 

|Name|Test Amount|Nanoseconds per operation|Items amount|Items per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|---:|---:|
|Reading-2|57146|110060|1000|0.01750|176|4|
|Reading#01-2|27457|217503|2000|0.07284|176|4|
|Reading#02-2|10000|549597|5000|0.5000|176|4|
|Reading#03-2|5064|1191316|10000|1.975|176|4|
|Reading#04-2|2775|2177278|20000|7.207|176|4|
|Reading#05-2|1092|5478643|50000|45.79|176|4|
|Reading#06-2|536|11170697|100000|186.6|176|4|

## Benchmark: unique items write test mapcache testing: 

|Name|Test Amount|Nanoseconds per operation|Items amount|Items per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|---:|---:|
|Writing-2|59030|103609|1000|0.01694|168|4|
|Writing#01-2|26804|220453|2000|0.07462|168|4|
|Writing#02-2|8157|678948|5000|0.6130|168|4|
|Writing#03-2|4165|1538039|10000|2.401|168|4|
|Writing#04-2|1948|3172072|20000|10.27|168|4|
|Writing#05-2|709|8382467|50000|70.52|168|4|
|Writing#06-2|324|18853312|100000|308.6|184|4|

## Benchmark: unique items read test mapcache testing: 

|Name|Test Amount|Nanoseconds per operation|Items amount|Items per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|---:|---:|
|Reading-2|56390|107482|1000|0.01773|176|4|
|Reading#01-2|26659|209836|2000|0.07502|176|4|
|Reading#02-2|10000|589539|5000|0.5000|176|4|
|Reading#03-2|6022|1038083|10000|1.661|176|4|
|Reading#04-2|2486|2408381|20000|8.045|176|4|
|Reading#05-2|1086|5808344|50000|46.04|176|4|
|Reading#06-2|489|12208566|100000|204.5|176|4|

