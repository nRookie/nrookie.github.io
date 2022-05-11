

unshare - run program with some namespaces unshared from parent







``` shell
[root@master ~]# lsns -t net
        NS TYPE NPROCS   PID USER    COMMAND
4026531992 net     183     1 root    /usr/lib/systemd/systemd --switched-root --system --deserialize 22
4026532216 net       1 14424 polkitd redis-server *:6379                               
4026532355 net       1 14431 polkitd mongod --bind_ip_all
4026532419 net       1 14647 root    /usr/local/bin/etcd --name etcd0 --data-dir /etcd-data --listen-client-urls http://0.0.0.0:2379 --advertise-client-
4026532481 net       4 14652 root    /bin/sh -c /usr/sbin/sshd && bash /usr/bin/start-zk.sh
4026532543 net       1 14723 polkitd mysqld
4026532609 net       2 13416 root    unshare -fn sleep 120
```



Nsenter enther the namespace 

``` shell
[root@master ~]# nsenter -t 13416 -n ip a 
1: lo: <LOOPBACK> mtu 65536 qdisc noop state DOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
```





check network namespace



``` shell
lsns -t net

[root@master ~]# lsns -t net
        NS TYPE NPROCS   PID USER    COMMAND
4026531992 net     183     1 root    /usr/lib/systemd/systemd --switched-root --system --deserialize 22
4026532216 net       1 14424 polkitd redis-server *:6379                               
4026532355 net       1 14431 polkitd mongod --bind_ip_all
4026532419 net       1 14647 root    /usr/local/bin/etcd --name etcd0 --data-dir /etcd-data --listen-client-urls http://0.0.0.0:2379 --advertise-client-
4026532481 net       4 14652 root    /bin/sh -c /usr/sbin/sshd && bash /usr/bin/start-zk.sh
4026532543 net       1 14723 polkitd mysqld
4026532609 net       2 13701 root    unshare -fn sleep 120
```





Linux 对 Namespace 操作方法



- clone

在创建新进程的系统调用时， 可以通过flags参数指定需要新建的Namespace 类型：



int clone(int (*fn) (void *), void *child_stack, int flags, void *arg)



- setns

  该系统调用可以让调用进程加入某个已经存在的Namespace中：

- unshare

  该系统调用可以将调用进程移动到新的Namespace 下：

  int unshare(int flags)

