# Benchmark results

## Attributes

| Test Attributes | Value                                               |
| --------------- | :-------------------------------------------------- |
| cache L1        | 32768                                               |
| cache L2        | 262144                                              |
| cache L3        | 8388608                                             |
| concurrency     | 8                                                   |
| cpu             | Intel(R) Core(TM) i7-7700K CPU @ 4.20GHz            |
| goarch          | amd64                                               |
| goos            | windows                                             |
| pkg             | github.com/DaanV2/High-Performance-Cache/benchmarks |

## Benchmark: duplicate items write test 

|                                 Name | Test Amount | Nanoseconds per operation |        | Bytes per operation | Allocations per operation |
| -----------------------------------: | ----------: | ------------------------: | -----: | ------------------: | ------------------------: |
|      type:MapCache,testing:Writing-8 |       51604 |                     23335 |  100.0 |                 616 |                        11 |
|    type:FixedCache,testing:Writing-8 |       49168 |                     25264 |  100.0 |                5416 |                       111 |
|   type:MapCache,testing:Writing#01-8 |       28983 |                     42485 |  200.0 |                 552 |                        10 |
| type:FixedCache,testing:Writing#01-8 |       27428 |                     42944 |  200.0 |               10152 |                       210 |
|   type:MapCache,testing:Writing#02-8 |       12196 |                     98904 |  500.0 |                 616 |                        11 |
| type:FixedCache,testing:Writing#02-8 |       15430 |                     75679 |  500.0 |               24617 |                       511 |
|   type:MapCache,testing:Writing#03-8 |        6666 |                    186703 |   1000 |                 553 |                        10 |
| type:FixedCache,testing:Writing#03-8 |        9999 |                    111091 |   1000 |               48553 |                      1010 |
|   type:MapCache,testing:Writing#04-8 |        3332 |                    347819 |   2000 |                 553 |                        10 |
| type:FixedCache,testing:Writing#04-8 |        5334 |                    221737 |   2000 |               96556 |                      2010 |
|   type:MapCache,testing:Writing#05-8 |        1428 |                    844433 |   5000 |                 557 |                        10 |
| type:FixedCache,testing:Writing#05-8 |        2607 |                    383665 |   5000 |              240553 |                      5010 |
|   type:MapCache,testing:Writing#06-8 |         705 |                   1774032 |  10000 |                 559 |                        10 |
| type:FixedCache,testing:Writing#06-8 |        1339 |                    781101 |  10000 |              480552 |                     10010 |
|   type:MapCache,testing:Writing#07-8 |         363 |                   3325056 |  20000 |                 556 |                        10 |
| type:FixedCache,testing:Writing#07-8 |         847 |                   1312538 |  20000 |              960555 |                     20010 |
|   type:MapCache,testing:Writing#08-8 |         134 |                   8647696 |  50000 |                 578 |                        10 |
| type:FixedCache,testing:Writing#08-8 |         400 |                   3167664 |  50000 |             2400560 |                     50010 |
|   type:MapCache,testing:Writing#09-8 |          66 |                  17660024 | 100000 |                 552 |                        10 |
| type:FixedCache,testing:Writing#09-8 |         189 |                   8898749 | 100000 |             4800552 |                    100010 |

## Benchmark: duplicate items read test 

|                             Name | Test Amount | Nanoseconds per operation |        | Bytes per operation | Allocations per operation |
| -------------------------------: | ----------: | ------------------------: | -----: | ------------------: | ------------------------: |
|      MapCache testing: Reading-8 |      112147 |                     12111 |  100.0 |                 624 |                        11 |
|    FixedCache testing: Reading-8 |       40641 |                     27310 |  100.0 |                3024 |                       111 |
|   MapCache testing: Reading#01-8 |       61366 |                     19496 |  200.0 |                 560 |                        10 |
| FixedCache testing: Reading#01-8 |       25470 |                     44164 |  200.0 |                5360 |                       210 |
|   MapCache testing: Reading#02-8 |       32161 |                     38014 |  500.0 |                 624 |                        11 |
| FixedCache testing: Reading#02-8 |       16002 |                     70027 |  500.0 |               12625 |                       511 |
|   MapCache testing: Reading#03-8 |       19939 |                     58973 |   1000 |                 560 |                        10 |
| FixedCache testing: Reading#03-8 |       10000 |                    107483 |   1000 |               24561 |                      1010 |
|   MapCache testing: Reading#04-8 |       10000 |                    106351 |   2000 |                 560 |                        10 |
| FixedCache testing: Reading#04-8 |        7058 |                    181104 |   2000 |               48563 |                      2010 |
|   MapCache testing: Reading#05-8 |        4999 |                    237032 |   5000 |                 560 |                        10 |
| FixedCache testing: Reading#05-8 |        3636 |                    303400 |   5000 |              120562 |                      5010 |
|   MapCache testing: Reading#06-8 |        2284 |                    516157 |  10000 |                 560 |                        10 |
| FixedCache testing: Reading#06-8 |        1654 |                    639763 |  10000 |              240562 |                     10010 |
|   MapCache testing: Reading#07-8 |        1146 |                   1079183 |  20000 |                 560 |                        10 |
| FixedCache testing: Reading#07-8 |        1080 |                    992038 |  20000 |              480561 |                     20010 |
|   MapCache testing: Reading#08-8 |         470 |                   2533397 |  50000 |                 560 |                        10 |
| FixedCache testing: Reading#08-8 |         446 |                   2349312 |  50000 |             1200560 |                     50010 |
|   MapCache testing: Reading#09-8 |         268 |                   4710333 | 100000 |                 560 |                        10 |
| FixedCache testing: Reading#09-8 |         169 |                   6777500 | 100000 |             2400561 |                    100010 |

## Benchmark: unique items write test 

|                                 Name | Test Amount | Nanoseconds per operation |        | Bytes per operation | Allocations per operation |
| -----------------------------------: | ----------: | ------------------------: | -----: | ------------------: | ------------------------: |
|      type:MapCache,testing:Writing-8 |       47967 |                     25663 |  100.0 |                 616 |                        11 |
|    type:FixedCache,testing:Writing-8 |       46408 |                     26620 |  100.0 |                5464 |                       112 |
|   type:MapCache,testing:Writing#01-8 |       25770 |                     45295 |  200.0 |                 552 |                        10 |
| type:FixedCache,testing:Writing#01-8 |       27700 |                     42955 |  200.0 |               10681 |                       221 |
|   type:MapCache,testing:Writing#02-8 |       10000 |                    100496 |  500.0 |                 617 |                        11 |
| type:FixedCache,testing:Writing#02-8 |       16038 |                     75992 |  500.0 |               25097 |                       521 |
|   type:MapCache,testing:Writing#03-8 |        5454 |                    227918 |   1000 |                 552 |                        10 |
| type:FixedCache,testing:Writing#03-8 |       10000 |                    147202 |   1000 |               49177 |                      1023 |
|   type:MapCache,testing:Writing#04-8 |        3157 |                    378500 |   2000 |                 554 |                        10 |
| type:FixedCache,testing:Writing#04-8 |        4772 |                    248344 |   2000 |              103514 |                      2155 |
|   type:MapCache,testing:Writing#05-8 |        1269 |                    963140 |   5000 |                 553 |                        10 |
| type:FixedCache,testing:Writing#05-8 |        1874 |                    623176 |   5000 |              682346 |                     14214 |
|   type:MapCache,testing:Writing#06-8 |         608 |                   1979951 |  10000 |                 554 |                        10 |
| type:FixedCache,testing:Writing#06-8 |        1210 |                    956220 |  10000 |              836857 |                     17433 |
|   type:MapCache,testing:Writing#07-8 |         297 |                   3960658 |  20000 |                 562 |                        10 |
| type:FixedCache,testing:Writing#07-8 |         885 |                   1445215 |  20000 |             1110264 |                     23129 |
|   type:MapCache,testing:Writing#08-8 |         100 |                  10699713 |  50000 |                 592 |                        10 |
| type:FixedCache,testing:Writing#08-8 |         374 |                   3267256 |  50000 |             2634552 |                     54885 |
|   type:MapCache,testing:Writing#09-8 |          51 |                  22931410 | 100000 |                 552 |                        10 |
| type:FixedCache,testing:Writing#09-8 |          56 |                  20484571 | 100000 |            31949844 |                    665620 |

## Attributes

| Test Attributes | Value                                 |
| --------------- | :------------------------------------ |
| fatal error     | all goroutines are asleep - deadlock! |

## Benchmark: unique items read test 

|                             Name | Test Amount | Nanoseconds per operation |        | Bytes per operation | Allocations per operation |
| -------------------------------: | ----------: | ------------------------: | -----: | ------------------: | ------------------------: |
|      MapCache testing: Reading-8 |      121777 |                     10337 |  100.0 |                 624 |                        11 |
|    FixedCache testing: Reading-8 |       61214 |                     20378 |  100.0 |                3048 |                       112 |
|   MapCache testing: Reading#01-8 |       76350 |                     16187 |  200.0 |                 560 |                        10 |
| FixedCache testing: Reading#01-8 |       36356 |                     33924 |  200.0 |                5624 |                       221 |
|   MapCache testing: Reading#02-8 |       32511 |                     37153 |  500.0 |                 624 |                        11 |
| FixedCache testing: Reading#02-8 |       19893 |                     62379 |  500.0 |               12866 |                       521 |
|   MapCache testing: Reading#03-8 |       17514 |                     68880 |   1000 |                 560 |                        10 |
| FixedCache testing: Reading#03-8 |       13563 |                     87890 |   1000 |               24872 |                      1023 |
|   MapCache testing: Reading#04-8 |       10000 |                    108034 |   2000 |                 560 |                        10 |
| FixedCache testing: Reading#04-8 |        7059 |                    215162 |   2000 |               52042 |                      2155 |
|   MapCache testing: Reading#05-8 |        4137 |                    246527 |   5000 |                 560 |                        10 |
| FixedCache testing: Reading#05-8 |        2262 |                    890768 |   5000 |              341460 |                     14214 |
|   MapCache testing: Reading#06-8 |        2552 |                    464049 |  10000 |                 560 |                        10 |
| FixedCache testing: Reading#06-8 |        1347 |                    772770 |  10000 |              418714 |                     17433 |
|   MapCache testing: Reading#07-8 |        1310 |                    949088 |  20000 |                 560 |                        10 |
| FixedCache testing: Reading#07-8 |         892 |                   1159190 |  20000 |              555416 |                     23129 |
|   MapCache testing: Reading#08-8 |         480 |                   2638558 |  50000 |                 560 |                        10 |
| FixedCache testing: Reading#08-8 |         436 |                   2733909 |  50000 |             1317560 |                     54885 |
|   MapCache testing: Reading#09-8 |         240 |                   5137793 | 100000 |                 560 |                        10 |

