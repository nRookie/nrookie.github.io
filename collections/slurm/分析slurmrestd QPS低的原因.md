

# 环境



``` shell
[root@primary ~]# sinfo -Nl
Sun Nov 14 16:58:02 2021
NODELIST   NODES PARTITION       STATE CPUS    S:C:T MEMORY TMP_DISK WEIGHT AVAIL_FE REASON
backup         1   control        idle 2       1:1:2   3804        0      1   (null) none
node1          1  compute*        idle 2       1:1:2  15772        0      1   (null) none
node2          1  compute*        idle 2       1:1:2  15772        0      1   (null) none
node3          1  compute*        idle 2       1:1:2  15772        0      1   (null) none
primary        1   control        idle 2       1:1:2   3804        0      1   (null) none
```







## 本地写一个简单的并发请求测试



### 1. 发送100个请求

``` golang
var wg sync.WaitGroup

const (
	GetJobURLTest       = "http://{ip}:6820/slurmdb/v0.0.37/job/39"
)
func GetSlurmJob() error {

	resBody, err := utils.HttpGet(context.Background(), GetJobURLTest)
	if err != nil {
		log.Fatal(err)
	}

	res := new(map[string]interface{})
	err = json.Unmarshal(resBody, res)
	if err != nil {
		return err
	}
	// fmt.Println(res)
	wg.Done()
	return err
}

func ProcessSlurm() {

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go GetSlurmJob()
	}

	wg.Wait()
}
```



``` shell
FVFF87EFQ6LR :: ~/developer/tools ‹3-refactor-tools-refactor-the-tools*› » time go run cmd/tools/main.go
finished
go run cmd/tools/main.go  0.59s user 0.44s system 15% cpu 6.509 total
```



#### Vmstat



``` shell
[root@primary ~]# vmstat -t 1 10
procs -----------memory---------- ---swap-- -----io---- -system-- ------cpu----- -----timestamp-----
 r  b   swpd   free   buff  cache   si   so    bi    bo   in   cs us sy id wa st                 CST
 0  0      0 1208048   3104 2109604    0    0     0     3   14   18  0  0 100  0  0 2021-11-14 17:31:52
 0  0      0 1208040   3104 2109588    0    0     0     0  189  327  0  0 100  0  0 2021-11-14 17:31:53
 0  0      0 1208040   3104 2109588    0    0     0    12  220  380  1  0 100  0  0 2021-11-14 17:31:54
 0  0      0 1208040   3104 2109580    0    0     0     0  200  344  0  0 100  0  0 2021-11-14 17:31:55
 0  0      0 1205796   3104 2109520    0    0     0     0 3396 18174 12 12 76  0  0 2021-11-14 17:31:56
 0  0      0 1206692   3104 2109544    0    0     0     0 2254 5987  7  5 88  0  0 2021-11-14 17:31:57
 0  0      0 1206756   3104 2109588    0    0     0     0  895 1385  0  2 98  0  0 2021-11-14 17:31:58
 0  0      0 1206756   3104 2109588    0    0     0     0  774 1642  0  1 99  0  0 2021-11-14 17:31:59
 0  0      0 1206788   3104 2109584    0    0     0     0  563 1049  1  1 99  0  0 2021-11-14 17:32:00
 0  0      0 1207040   3104 2109580    0    0     0     0  392  805  1  1 99  0  0 2021-11-14 17:32:01
```



#### iostat



``` shell
[root@primary ~]# iostat -t 1 10
Linux 4.19.0-9.el7.xcloud.x86_64 (primary) 	2021年11月14日 	_x86_64_	(2 CPU)

2021年11月14日 17时30分38秒
avg-cpu:  %user   %nice %system %iowait  %steal   %idle
           0.20    0.00    0.22    0.00    0.00   99.58

Device:            tps    kB_read/s    kB_wrtn/s    kB_read    kB_wrtn
vda               0.53         0.82         5.15     642114    4057329
vdb               0.00         0.01         0.00       6567       2068
loop0             0.07         0.08         0.00      65346          0

2021年11月14日 17时30分39秒
avg-cpu:  %user   %nice %system %iowait  %steal   %idle
           0.50    0.00    0.00    0.00    0.00   99.50

Device:            tps    kB_read/s    kB_wrtn/s    kB_read    kB_wrtn
vda               0.00         0.00         0.00          0          0
vdb               0.00         0.00         0.00          0          0
loop0             0.00         0.00         0.00          0          0

2021年11月14日 17时30分40秒
avg-cpu:  %user   %nice %system %iowait  %steal   %idle
           0.00    0.00    0.50    0.00    0.00   99.50

Device:            tps    kB_read/s    kB_wrtn/s    kB_read    kB_wrtn
vda               0.00         0.00         0.00          0          0
vdb               0.00         0.00         0.00          0          0
loop0             0.00         0.00         0.00          0          0

2021年11月14日 17时30分41秒
avg-cpu:  %user   %nice %system %iowait  %steal   %idle
           0.50    0.00    0.00    0.00    0.00   99.50

Device:            tps    kB_read/s    kB_wrtn/s    kB_read    kB_wrtn
vda               5.00         0.00        72.00          0         72
vdb               0.00         0.00         0.00          0          0
loop0             0.00         0.00         0.00          0          0

2021年11月14日 17时30分42秒
avg-cpu:  %user   %nice %system %iowait  %steal   %idle
           5.70    0.00    5.70    0.00    0.00   88.60

Device:            tps    kB_read/s    kB_wrtn/s    kB_read    kB_wrtn
vda               0.00         0.00         0.00          0          0
vdb               0.00         0.00         0.00          0          0
loop0             0.00         0.00         0.00          0          0

2021年11月14日 17时30分43秒
avg-cpu:  %user   %nice %system %iowait  %steal   %idle
           7.45    0.00    7.45    0.00    0.00   85.11

Device:            tps    kB_read/s    kB_wrtn/s    kB_read    kB_wrtn
vda              16.00         0.00       164.50          0        164
vdb               0.00         0.00         0.00          0          0
loop0             0.00         0.00         0.00          0          0

2021年11月14日 17时30分44秒
avg-cpu:  %user   %nice %system %iowait  %steal   %idle
           0.51    0.00    1.02    0.00    0.00   98.48

Device:            tps    kB_read/s    kB_wrtn/s    kB_read    kB_wrtn
vda               0.00         0.00         0.00          0          0
vdb               0.00         0.00         0.00          0          0
loop0             0.00         0.00         0.00          0          0

2021年11月14日 17时30分45秒
avg-cpu:  %user   %nice %system %iowait  %steal   %idle
           0.50    0.00    1.00    0.00    0.00   98.50

Device:            tps    kB_read/s    kB_wrtn/s    kB_read    kB_wrtn
vda               0.00         0.00         0.00          0          0
vdb               0.00         0.00         0.00          0          0
loop0             0.00         0.00         0.00          0          0

2021年11月14日 17时30分46秒
avg-cpu:  %user   %nice %system %iowait  %steal   %idle
           0.00    0.00    0.50    0.00    0.00   99.50

Device:            tps    kB_read/s    kB_wrtn/s    kB_read    kB_wrtn
vda               0.00         0.00         0.00          0          0
vdb               0.00         0.00         0.00          0          0
loop0             0.00         0.00         0.00          0          0

2021年11月14日 17时30分47秒
avg-cpu:  %user   %nice %system %iowait  %steal   %idle
           0.50    0.00    0.50    0.00    0.00   99.00

Device:            tps    kB_read/s    kB_wrtn/s    kB_read    kB_wrtn
vda               0.00         0.00         0.00          0          0
vdb               0.00         0.00         0.00          0          0
loop0             0.00         0.00         0.00          0          0
```





### 2. 发送1000个请求

``` shell
FVFF87EFQ6LR :: ~/developer/tools ‹3-refactor-tools-refactor-the-tools*› » time go run cmd/tools/main.go                                        130 ↵
2021/11/14 17:09:43 context deadline exceeded (Client.Timeout or context cancellation while reading body)
exit status 1
go run cmd/tools/main.go  0.99s user 0.71s system 5% cpu 31.516 total
```



slurmrestd 没有响应了。





## 使用ApacheBench



### 本机测试



#### 1000个请求，并发100

``` shell
[root@primary ~]# ab -n 1000 -c 100 -H "Content-type: application/json" -H "x-slurm-user-name: slurmrestd" -H "x-slurm-user-token:xxxxxxxxx" http://localhost:6820/slurmdb/v0.0.37/job/39
This is ApacheBench, Version 2.3 <$Revision: 1430300 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:
Server Hostname:        localhost
Server Port:            6820

Document Path:          /slurmdb/v0.0.37/job/39
Document Length:        5955 bytes

Concurrency Level:      100
Time taken for tests:   17.125 seconds
Complete requests:      1000
Failed requests:        0
Write errors:           0
Total transferred:      6028000 bytes
HTML transferred:       5955000 bytes
Requests per second:    58.39 [#/sec] (mean)
Time per request:       1712.519 [ms] (mean)
Time per request:       17.125 [ms] (mean, across all concurrent requests)
Transfer rate:          343.75 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.4      0       2
Processing:   338 1610 248.0   1668    2019
Waiting:      336 1577 249.6   1661    2011
Total:        338 1610 247.7   1668    2020

Percentage of the requests served within a certain time (ms)
  50%   1668
  66%   1674
  75%   1678
  80%   1681
  90%   1691
  95%   1704
  98%   1926
  99%   1959
 100%   2020 (longest request)

```



#### 1000个请求并发200

``` shell
[root@primary ~]# ab -n 1000 -c 200 -H "Content-type: application/json" -H "x-slurm-user-name: slurmrestd" -H "x-slurm-user-token:xxxxxxxxx" http://localhost:6820/slurmdb/v0.0.37/job/39
This is ApacheBench, Version 2.3 <$Revision: 1430300 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:
Server Hostname:        localhost
Server Port:            6820

Document Path:          /slurmdb/v0.0.37/job/39
Document Length:        5955 bytes

Concurrency Level:      200
Time taken for tests:   17.097 seconds
Complete requests:      1000
Failed requests:        0
Write errors:           0
Total transferred:      6028000 bytes
HTML transferred:       5955000 bytes
Requests per second:    58.49 [#/sec] (mean)
Time per request:       3419.348 [ms] (mean)
Time per request:       17.097 [ms] (mean, across all concurrent requests)
Transfer rate:          344.32 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   1.1      0       4
Processing:   340 3048 695.8   3328    3691
Waiting:      336 3017 692.4   3317    3673
Total:        340 3049 694.8   3328    3693

Percentage of the requests served within a certain time (ms)
  50%   3328
  66%   3334
  75%   3337
  80%   3339
  90%   3345
  95%   3357
  98%   3570
  99%   3589
 100%   3693 (longest request)
```



#### 1000个请求并发500

``` shell
[root@primary ~]# ab -n 1000 -c 500 -H "Content-type: application/json" -H "x-slurm-user-name: slurmrestd" -H "x-slurm-user-token:xxxxxxxxx" http://localhost:6820/slurmdb/v0.0.37/job/39
This is ApacheBench, Version 2.3 <$Revision: 1430300 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:
Server Hostname:        localhost
Server Port:            6820

Document Path:          /slurmdb/v0.0.37/job/39
Document Length:        5955 bytes

Concurrency Level:      500
Time taken for tests:   17.161 seconds
Complete requests:      1000
Failed requests:        0
Write errors:           0
Total transferred:      6028000 bytes
HTML transferred:       5955000 bytes
Requests per second:    58.27 [#/sec] (mean)
Time per request:       8580.349 [ms] (mean)
Time per request:       17.161 [ms] (mean, across all concurrent requests)
Transfer rate:          343.03 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    6   5.8      8      14
Processing:   344 6390 2376.5   7682    8682
Waiting:      335 6360 2372.0   7663    8679
Total:        344 6396 2372.6   7694    8693

Percentage of the requests served within a certain time (ms)
  50%   7694
  66%   8330
  75%   8358
  80%   8370
  90%   8397
  95%   8412
  98%   8441
  99%   8595
 100%   8693 (longest request)
[root@primary ~]#

```



#### 1000个请求并发10

``` shell
[root@primary ~]# ab -n 1000 -c 10 -H "Content-type: application/json" -H "x-slurm-user-name: slurmrestd" -H "x-slurm-user-token:xxxxxxxxx" http://localhost:6820/slurmdb/v0.0.37/job/39
This is ApacheBench, Version 2.3 <$Revision: 1430300 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:
Server Hostname:        localhost
Server Port:            6820

Document Path:          /slurmdb/v0.0.37/job/39
Document Length:        5955 bytes

Concurrency Level:      10
Time taken for tests:   33.887 seconds
Complete requests:      1000
Failed requests:        0
Write errors:           0
Total transferred:      6028000 bytes
HTML transferred:       5955000 bytes
Requests per second:    29.51 [#/sec] (mean)
Time per request:       338.872 [ms] (mean)
Time per request:       33.887 [ms] (mean, across all concurrent requests)
Transfer rate:          173.71 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       0
Processing:   330  335   4.4    334     356
Waiting:      330  335   4.4    333     356
Total:        330  335   4.5    334     357

Percentage of the requests served within a certain time (ms)
  50%    334
  66%    334
  75%    336
  80%    337
  90%    341
  95%    346
  98%    351
  99%    354
 100%    357 (longest request)
```



## 通过外网测试



#### 1000个请求并发20

``` shell
[root@10-13-175-37 ~]# ab -n 1000 -c 20 -H "Content-type: application/json" -H "x-slurm-user-name: slurmrestd" -H "x-slurm-user-token:xxxxxxxxx" http://{ip}:6820/slurmdb/v0.0.37/job/39
This is ApacheBench, Version 2.3 <$Revision: 1843412 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking {ip} (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:
Server Hostname:        {ip}
Server Port:            6820

Document Path:          /slurmdb/v0.0.37/job/39
Document Length:        5955 bytes

Concurrency Level:      20
Time taken for tests:   61.108 seconds
Complete requests:      1000
Failed requests:        0
Total transferred:      6028000 bytes
HTML transferred:       5955000 bytes
Requests per second:    16.36 [#/sec] (mean)
Time per request:       1222.155 [ms] (mean)
Time per request:       61.108 [ms] (mean, across all concurrent requests)
Transfer rate:          96.33 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:       27   33  45.3     31    1049
Processing:   359 1071 1873.0    696   31456
Waiting:      359  565 1193.9    404   30688
Total:        387 1103 1873.0    727   31487

Percentage of the requests served within a certain time (ms)
  50%    727
  66%    789
  75%   1180
  80%   1238
  90%   1756
  95%   2683
  98%   4720
  99%   8004
 100%  31487 (longest request)

```

#### 1000个请求并发100

``` shell
[root@10-13-175-37 ~]# ab -n 1000 -c 100 -H "Content-type: application/json" -H "x-slurm-user-name: slurmrestd" -H "x-slurm-user-token:xxxxxxxxx" http://{ip}:6820/slurmdb/v0.0.37/job/39
This is ApacheBench, Version 2.3 <$Revision: 1843412 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking {ip} (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
apr_pollset_poll: The timeout specified has expired (70007)
Total of 969 requests completed

[root@10-13-175-37 ~]# ab -n 1000 -c 100 -H "Content-type: application/json" -H "x-slurm-user-name: slurmrestd" -H "x-slurm-user-token:xxxxxxxxx" http://{ip}:6820/slurmdb/v0.0.37/job/39
This is ApacheBench, Version 2.3 <$Revision: 1843412 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking {ip} (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests

apr_pollset_poll: The timeout specified has expired (70007)
Total of 984 requests completed
```

有丢包的情况





## Slurmrest load变多



### 通过go 写的并发脚本执行.

![image-20211115103001540](/Users/kestrel/developer/nrookie.github.io/collections/slurm/image-20211115103001540.png)



### 通过apachebench 测试 （广东区域的虚拟机发送）

![image-20211115105042636](/Users/kestrel/developer/nrookie.github.io/collections/slurm/image-20211115105042636.png)

### 通过apachebench 测试 (办公网)

![image-20211115104748708](/Users/kestrel/developer/nrookie.github.io/collections/slurm/image-20211115104748708.png)



###



经过测试发现是办公网发送和广东区域虚机发送导致的负载差异。



再使用手机5G网络热点确认一下

![image-20211115115106141](/Users/kestrel/developer/nrookie.github.io/collections/slurm/image-20211115115106141.png)



cpu负载很低。



办公网络和其他网络有什么差异？









## 分析有什么不同点



### Sar



#### 办公网

![image-20211115112144030](/Users/kestrel/developer/nrookie.github.io/collections/slurm/image-20211115112144030.png)



#### 广东2网络

![image-20211115112321693](/Users/kestrel/developer/nrookie.github.io/collections/slurm/image-20211115112321693.png)



为什么办公网system的usage 这么高?

![image-20211115134756443](/Users/kestrel/developer/nrookie.github.io/collections/slurm/image-20211115134756443.png)



大部分时间花在了 context switch 上面

#### 办公网

<img src="/Users/kestrel/developer/nrookie.github.io/collections/slurm/image-20211115140859244.png" alt="image-20211115140859244" style="zoom:80%;" />



#### 广东2网访问

<img src="/Users/kestrel/developer/nrookie.github.io/collections/slurm/image-20211115141315768.png" alt="image-20211115141315768" style="zoom:80%;" />





更改配置，缩小线程数

``` shell
ExecStart=/usr/sbin/slurmrestd -f /etc/slurm/slurm.conf -t 4 -u slurmrestd -a rest_auth/jwt -s openapi/v0.0.37,dbv0.0.37 -vvvvv 0.0.0.0:6820
```



``` shell
% time   seconds usecs/call   calls  errors syscall

------ ----------- ----------- --------- --------- ----------------

 48.03  1.813128     17  108235   20969 futex
```



























watch interrupts

``` shell
watch -tdn1 cat /proc/interrupts
```



#### Interrups explained

Local timer interrupts





Context switch



```shell
pidstat -wt 1
```







发和不发请求好像没有什么区别, interrupt 没什么区别



![image-20211115171748294](/Users/kestrel/developer/nrookie.github.io/collections/slurm/image-20211115171748294.png)



![image-20211115172012605](/Users/kestrel/developer/nrookie.github.io/collections/slurm/image-20211115172012605.png)







https://www.ibm.com/docs/en/db2/9.7?topic=consumption-diagnosing-high-context-switch-rate-problem
