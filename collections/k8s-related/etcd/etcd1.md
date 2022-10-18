<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221006084615607.png" alt="image-20221006084615607" style="zoom:50%;" />

![image-20221006084846053](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221006084846053.png)



## 多版本机制

``` shell
etcd
```





``` shell
etcdctl --endpoints=http://localhost:2379 get b
etcdctl --endpoints=http://localhost:2379 put b v1
etcdctl --endpoints=http://localhost:2379 get b -wjson
{"header":{"cluster_id":14841639068965178418,"member_id":10276657743932975437,"revision":2,"raft_term":2},"kvs":[{"key":"Yg==","create_revision":2,"mod_revision":2,"version":1,"value":"djE="}],"count":1}
etcdctl --endpoints=http://localhost:2379 put b v2

etcdctl --endpoints=http://localhost:2379 get b -wjson

{"header":{"cluster_id":14841639068965178418,"member_id":10276657743932975437,"revision":3,"raft_term":2},"kvs":[{"key":"Yg==","create_revision":2,"mod_revision":3,"version":2,"value":"djI="}],"count":1}

etcdctl --endpoints=http://localhost:2379 get b --rev=2

```



![image-20221006094541672](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221006094541672.png) 





设计的时候，值的值不要太大。

如果包太大性能会降的比较厉害。



5. MVCC的目的。尽量在不加锁的情况下增加多版本的值(多版本并发控制)

apply 是操作mvcc模块。

先操作 treeIndex



key 和value.

Key 就是对象的key

value 是版本信息。

modified<4,0>  （4 相当于事务， 0相当于这个事务中的第几次操作）

etcd 里面 revision的值会跟着跳的。记录了全局的一次写入的动作。

revision分两个部分，一个main，一个sub。



generation 是历史的变更信息。





接下来会去写boltdb （B + 数）



会把整个信息记录在boltdb里面。



etcd 是一个多读少写的数据库。



需要确保半数以上的人提交了这个日志，我们才可以写入MVCC模块。

在这之前需要把操作写到日志里面。







``` shell
ks get po kube-controller-manager-instance-4 -oyaml
```



resourceVersion : 是一个乐观锁， 是控制器修改一个对象的时候不能版本冲突。

他其实读取的就是对象在etcd中的modified的version。 （如果变动很频繁的话，这个值会变的很大）。







etcd 的一致性怎么保证？

<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221006101117843.png" alt="image-20221006101117843" style="zoom:50%;" />



### watch 机制



![image-20221006101413629](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221006101413629.png)



<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221006101518350.png" alt="image-20221006101518350" style="zoom:50%;" />



``` shell
etcdctl --endpoints=http://localhost:2379 get --prefix a
etcdctl --endpoints=http://localhost:2379 get --prefix a --rev=6
```





![image-20221006101711054](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221006101711054.png)

![image-20221006101731983](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221006101731983.png)





主和follower都会往MVCC里面写数据





treeindex是一个inmemory 数据库。

boltdb是用来做持久化的。



Apply 成功才会返回成功的响应吧？



如果是强一致， 是apply成功。

如果是若一致，是commit成功。



![image-20221006103713465](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221006103713465.png)



watch 机制中 为啥会存在未同步完毕的情况？



1000个版本，整个的cache里面肯定是没有历史信息的，他要从db里面拉出来。

这是需要花费时间的，当内存和数据库的信息一样的情况下才会返回出来。



![image-20221006104218128](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221006104218128.png)



``` shell
etcdctl --endpoints=http://localhost:2379 member list --write-out=table

```



![image-20221006104725076](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221006104725076.png)



<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221006104814998.png" alt="image-20221006104814998" style="zoom:50%;" />





为什么不建议超过1.5M 需要messageappend把数据带过去。信息的阻塞，非常低效。



<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221006105118347.png" alt="image-20221006105118347" style="zoom:50%;" />





``` shell
etcd --quota-backend-bytes=$((16*1024*1024))

 while [1]; do dd if=/dev/urandom bs=1024 count=1024 | ETCDCTL_API=3 etcdctl put key || break; done
 
 
 Error: etcdserver: mvcc: database space exceeded
 
 etcdctl --endpoints=http://localhost:2379 --write-out=table endpoint status
 
 ETCDCTL_API=3 etcdctl alarm list
 
 ETCDCTL_API=3 etcdctl defrag
 
 ETCDCTL_API=3 etcdctl alarm disarm
```

 <img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221006125607325.png" alt="image-20221006125607325" style="zoom:50%;" />

<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221006125720359.png" alt="image-20221006125720359" style="zoom:50%;" />