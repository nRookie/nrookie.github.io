

# 用户目录隔离问题



## 问题描述



假设目前home目录结构如下，其中user1代表用户1的数据目录， user2代表用户2的数据目录。

本例描述用户1是一个恶意用户，其尝试获取用户2的敏感数据。

``` shell
[root@primary home]# tree
.
├── user1
│   ├── job1.sh
│   └── job.269.out
└── user2
    └── secret
```

由用户隔离，user1禁止访问user2的数据，反之亦然。但是目前所有的sbatch命令都由root用户执行，因此不能使用原生的linux用户隔离。



如上用户1的job1.sh 的脚本如下所示

``` shell
[root@primary user1]# cat job1.sh 
#!/bin/bash
#SBATCH --partition=compute
#SBATCH --qos=low
#SBATCH -J chroot
#SBATCH --nodes=1
#SBATCH --ntasks-per-node=1

# 可以考虑执行多次获取其他用户的目录
cd ../user2
pwd
ls
```

用户1通过控制台提交作业, 可以看到可以自由的登录到用户2的user2文件夹里面。

``` shell
[root@primary user1]# sbatch job1.sh 
Submitted batch job 270
[root@primary user1]# cat slurm-270.out 
/nfs1/home/user2
secret
[root@primary user1]# 
```



## 解决方法1: 使用chroot将sbatch的根目录锁定在提交用户的目录下（没走通）



### chroot jail 介绍

https://www.geeksforgeeks.org/linux-virtualization-using-chroot-jail/



### 实操

1. 查看sbatch 所需要的依赖文件

2. ```shell
   [root@primary user1]# ldd $(which sbatch)
   	linux-vdso.so.1 =>  (0x00007ffc6bb95000)
   	libslurmfull.so => /usr/lib64/slurm/libslurmfull.so (0x00007f3059d7b000)
   	libdl.so.2 => /lib64/libdl.so.2 (0x00007f3059b77000)
   	libm.so.6 => /lib64/libm.so.6 (0x00007f3059875000)
   	libresolv.so.2 => /lib64/libresolv.so.2 (0x00007f305965b000)
   	libpthread.so.0 => /lib64/libpthread.so.0 (0x00007f305943f000)
   	libc.so.6 => /lib64/libc.so.6 (0x00007f3059071000)
   	libgcc_s.so.1 => /lib64/libgcc_s.so.1 (0x00007f3058e5b000)
   	/lib64/ld-linux-x86-64.so.2 (0x00007f305a15e000)
   ```

3. 添加相应的依赖到user1文件夹下面

4. ``` shell
   [root@primary user1]# ldd $(which sbatch) | awk ' $(NF - 1) != "=>" {print $(NF - 1)}' | sed -e 's/\(\/\)\(.*\)\(\/.*\)/mkdir -p \2 \&\&  cp   \0 \2\3/' > command.sh
   [root@primary user1]# chmod 700 command.sh
   
   [root@primary user1]# ./command.sh 
   [root@primary user1]# ls
   command.sh  job1.sh  lib64  usr
   [root@primary user1]# tree
   .
   ├── command.sh
   ├── job1.sh
   ├── lib64
   │   ├── ld-linux-x86-64.so.2
   │   ├── libc.so.6
   │   ├── libdl.so.2
   │   ├── libgcc_s.so.1
   │   ├── libm.so.6
   │   ├── libpthread.so.0
   │   └── libresolv.so.2
   └── usr
       └── lib64
           └── slurm
               └── libslurmfull.so
   ```

5. 通过chroot 运行 sbatch job1.sh

6. ``` shell
   [root@primary user1]# cd ..
   [root@primary home]# chroot user1 sbatch user1/job1.sh 
   chroot: failed to run command ‘sbatch’: No such file or directory
   ```

7. 拷贝sbatch二进制文件并执行

8. ```shell
   [root@primary user1]# cp $(which sbatch) bin/sbatch
   cp: 无法创建普通文件"bin/sbatch": 没有那个文件或目录
   [root@primary user1]# mkdir bin
   [root@primary user1]# cp $(which sbatch) bin/sbatch
   ```

9. 再运行一次

10. ``` shell
    [root@primary home]# chroot user1 /bin/sbatch user1/job1.sh 
    sbatch: error: resolve_ctls_from_dns_srv: res_nsearch error: Host name lookup failure
    sbatch: error: fetch_config: DNS SRV lookup failed
    sbatch: error: _establish_config_source: failed to fetch config
    sbatch: fatal: Could not establish a configuration source
    ```

11. 拷贝 /etc/文件夹再运行

12. ``` shell
    [root@primary user1]# cp -r /etc/ etc
    [root@primary user1]# cd ..
    [root@primary home]# chroot user1 /bin/sbatch user1/job1.sh 
    sbatch: error: Invalid user for SlurmUser slurm, ignored
    sbatch: fatal: Unable to process configuration file
    ```

13. 待续



