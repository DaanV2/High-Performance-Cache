# Benchmark results

## Attributes

|Test Attributes|Value|
|---------------|:-----|
|cache L1|32768|
|cache L2|262144|
|cache L3|52428800|
|concurrency|2|
|cpu|Intel(R) Xeon(R) CPU E5-2673 v4 @ 2.30GHz|
|goarch|amd64|
|goos|linux|
|pkg|github.com/DaanV2/High-Performance-Cache/benchmarks|

## Benchmark: duplicate items write test mapcache testing: 

|Name|Test Amount|Nanoseconds per operation|Items amount|Items per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|---:|---:|
|Writing-2|31275|196914|1000|0.03197|168|4|
|Writing#01-2|15272|399910|2000|0.1310|168|4|
|Writing#02-2|6112|952512|5000|0.8181|168|4|
|Writing#03-2|3200|1887331|10000|3.125|168|4|
|Writing#04-2|1622|3914655|20000|12.33|168|4|
|Writing#05-2|626|9591758|50000|79.87|168|4|
|Writing#06-2|315|19182418|100000|317.5|168|4|

## Benchmark: duplicate items read test mapcache testing: 

|Name|Test Amount|Nanoseconds per operation|Items amount|Items per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|---:|---:|
|Reading-2|58173|108497|1000|0.01719|176|4|
|Reading#01-2|27990|212067|2000|0.07145|176|4|
|Reading#02-2|10000|535178|5000|0.5000|176|4|
|Reading#03-2|5545|1065933|10000|1.803|176|4|
|Reading#04-2|2780|2137247|20000|7.194|176|4|
|Reading#05-2|1072|5470637|50000|46.64|176|4|
|Reading#06-2|540|10802103|100000|185.2|176|4|

## Benchmark: unique items write test mapcache testing: 

|Name|Test Amount|Nanoseconds per operation|Items amount|Items per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|---:|---:|
|Writing-2|27007|218196|1000|0.03703|168|4|
|Writing#01-2|12896|440244|2000|0.1551|168|4|
|Writing#02-2|5389|1170805|5000|0.9278|168|4|
|Writing#03-2|2512|2424200|10000|3.981|168|4|
|Writing#04-2|1222|4996053|20000|16.37|168|4|
|Writing#05-2|459|13335312|50000|108.9|168|4|
|Writing#06-2|190|30727341|100000|526.3|168|4|

## Benchmark: unique items read test mapcache testing: 

|Name|Test Amount|Nanoseconds per operation|Items amount|Items per operation|Bytes per operation|Allocations per operation|
|----:|---:|---:|---:|---:|---:|---:|
|Reading-2|53456|105245|1000|0.01871|176|4|
|Reading#01-2|26847|223770|2000|0.07450|176|4|
|Reading#02-2|10000|548102|5000|0.5000|176|4|
|Reading#03-2|5804|987849|10000|1.723|176|4|
|Reading#04-2|2548|2199730|20000|7.849|176|4|
|Reading#05-2|1066|5571150|50000|46.90|176|4|
|Reading#06-2|528|11625387|100000|189.4|176|4|

