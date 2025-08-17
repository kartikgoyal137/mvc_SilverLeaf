# Performance Benchmark
## 1000 concurrent users ; 100,000 requests

These results were generated using Apache Bench after implementing the proactive, single-query caching strategy.

```
This is ApacheBench, Version 2.3 <$Revision: 1923142 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, [http://www.zeustech.net/](http://www.zeustech.net/)
Licensed to The Apache Software Foundation, [http://www.apache.org/](http://www.apache.org/)

Benchmarking localhost (be patient)
Completed 10000 requests
Completed 20000 requests
Completed 30000 requests
Completed 40000 requests
Completed 50000 requests
Completed 60000 requests
Completed 70000 requests
Completed 80000 requests
Completed 90000 requests
Completed 100000 requests
Finished 100000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /api/v1/menu/cat/1
Document Length:        3363 bytes

Concurrency Level:      1000
Time taken for tests:   4.945 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      346500000 bytes
HTML transferred:       336300000 bytes
Requests per second:    20223.50 [#/sec] (mean)
Time per request:       49.447 [ms] (mean)
Time per request:       0.049 [ms] (mean, across all concurrent requests)
Transfer rate:          68432.05 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   24   3.8     23      42
Processing:     6   25   4.5     25      63
Waiting:        1   19   3.9     19      41
Total:         33   49   5.4     48      78

Percentage of the requests served within a certain time (ms)
  50%     48
  66%     48
  75%     49
  80%     50
  90%     52
  95%     61
  98%     72
  99%     73
 100%     78 (longest request)
```