







[root@master cpu]# cat cpu.cfs_quota_us 

-1

[root@master cpu]# cat cpu.cfs_period_us

100000



``` shell
[root@master cpudemo]# cat cgroup.procs 
[root@master cpudemo]# echo 17736 > cgroup.procs
[root@master cpudemo]# echo 10000 > cpu.cfs_quota_us 

Tasks: 192 total,   3 running, 110 sleeping,   0 stopped,   0 zombie
%Cpu(s):  3.7 us,  0.5 sy,  0.0 ni, 95.6 id,  0.0 wa,  0.2 hi,  0.1 si,  0.0 st
KiB Mem :  7959028 total,  2352248 free,  2663420 used,  2943360 buff/cache
KiB Swap:        0 total,        0 free,        0 used.  4451068 avail Mem 

  PID USER      PR  NI    VIRT    RES    SHR S  %CPU %MEM     TIME+ COMMAND                                                                            
17736 root      20   0  702416   4704    516 R  10.0  0.1   5:50.27 busyloop              



[root@master cpudemo]# echo 100000 > cpu.cfs_quota_us 
top - 08:56:07 up 6 days, 11:20,  2 users,  load average: 0.62, 0.73, 0.33
Tasks: 191 total,   2 running, 110 sleeping,   0 stopped,   0 zombie
%Cpu(s): 25.8 us,  0.2 sy,  0.0 ni, 73.6 id,  0.0 wa,  0.2 hi,  0.1 si,  0.0 st
KiB Mem :  7959028 total,  2352720 free,  2662816 used,  2943492 buff/cache
KiB Swap:        0 total,        0 free,        0 used.  4451556 avail Mem 

  PID USER      PR  NI    VIRT    RES    SHR S  %CPU %MEM     TIME+ COMMAND       
```







### memory 子系统





- memory.usage_in_bytes
  - cgroup下进程使用的内存，包含cgroup及其子cgroup下的进程使用的内存
- memory.max_usage_in_bytes
  - cgroup下进程使用内存的最大值，包含cgroup的内存使用量。
- Memory.limit_in_bytes
  - 设置cgroup下进程最多能使用的内存。如果设置为 -1，表示对该cgroup的内存使用不做限制。
- memory.oom_control
  - 设置是否在Cgroup中使用OOM(Out of Memory) Killer， 默认为使用。 当属于该cgroup的进程使用的内存超过最大的限定值时， 会立刻被OOm Killer 处理。



### memory 子系统练习







