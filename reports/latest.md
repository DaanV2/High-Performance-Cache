# Benchmark results

## Attributes

|Test Attributes|Value|
|---------------|:-----|
|cache L1|32768|
|cache L2|262144|
|cache L3|31457280|
|concurrency|2|
|cpu|Intel(R) Xeon(R) CPU E5-2673 v3 @ 2.40GHz|
|goarch|amd64|
|goos|linux|
|pkg|github.com/DaanV2/High-Performance-Cache/benchmarks|

## Benchmark: duplicate items write test mapcache testing: 

|Name|Test Amount|Nanoseconds per operation|Items amount|Items per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|---:|---:|
|Writing-2|27535|206386|1000|0.03632|168|4|
|Writing#01-2|14817|414379|2000|0.1350|168|4|
|Writing#02-2|5995|1016398|5000|0.8340|168|4|
|Writing#03-2|2973|2087162|10000|3.364|168|4|
|Writing#04-2|1494|4063965|20000|13.39|168|4|
|Writing#05-2|595|10263034|50000|84.03|168|4|
|Writing#06-2|266|21684084|100000|375.9|168|4|

## Benchmark: duplicate items read test mapcache testing: 

|Name|Test Amount|Nanoseconds per operation|Items amount|Items per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|---:|---:|
|Reading-2|53886|112535|1000|0.01856|176|4|
|Reading#01-2|27954|213179|2000|0.07155|176|4|
|Reading#02-2|10000|553616|5000|0.5000|176|4|
|Reading#03-2|5968|1062449|10000|1.676|176|4|
|Reading#04-2|2803|2040634|20000|7.135|176|4|
|Reading#05-2|1114|5655503|50000|44.88|176|4|
|Reading#06-2|594|10176889|100000|168.4|176|4|

## Benchmark: unique items write test mapcache testing: 

|Name|Test Amount|Nanoseconds per operation|Items amount|Items per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|---:|---:|
|Writing-2|26163|229852|1000|0.03822|168|4|
|Writing#01-2|12591|479883|2000|0.1588|168|4|
|Writing#02-2|4479|1268788|5000|1.116|168|4|
|Writing#03-2|2337|2495573|10000|4.279|168|4|
|Writing#04-2|1201|5213299|20000|16.65|168|4|
|Writing#05-2|447|13381507|50000|111.9|168|4|
|Writing#06-2|186|31661663|100000|537.6|168|4|

## Benchmark: unique items read test mapcache testing: 

|Name|Test Amount|Nanoseconds per operation|Items amount|Items per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|---:|---:|
|Reading-2|56776|109610|1000|0.01761|176|4|
|Reading#01-2|29719|224633|2000|0.06730|176|4|
|Reading#02-2|9666|567533|5000|0.5173|176|4|
|Reading#03-2|5024|1168007|10000|1.990|176|4|
|Reading#04-2|2498|2201623|20000|8.006|176|4|
|Reading#05-2|968|5825733|50000|51.65|176|4|
|Reading#06-2|516|11334902|100000|193.8|176|4|

