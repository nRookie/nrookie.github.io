<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221014155945121.png" alt="image-20221014155945121" style="zoom:33%;" />



kubelet 定时获取容器的日志和容器可写层的磁盘使用情况， 如果超过限制， 则会对pod 进行驱逐。



### Init Container 的资源需求



![image-20221014160105368](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221014160105368.png)

<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221014160332475.png" alt="image-20221014160332475" style="zoom:33%;" />



<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221014160651726.png" alt="image-20221014160651726" style="zoom:33%;" />





``` shell
k label node instance-5 disktype=ssd
```





<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221014161357812.png" alt="image-20221014161357812" style="zoom:33%;" />



``` shell
k label node instance-5 disktype-
```





<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221014162936743.png" alt="image-20221014162936743" style="zoom:33%;" />



可以用来把pod 逼到不同的节点上去。



![image-20221014163648795](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221014163648795.png)  



// 代表同一个节点。





<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221014163722430.png" alt="image-20221014163722430" style="zoom:33%;" />



// 污点 taint

使用场景。 



1, key, 2, value. 3. 产生的效果

``` shell
 kubectl taint nodes instance-4 for-special-user=qing:NoSchedule
```











``` shell
ks get pods coredns-565d847f94-lwkcn -o yaml
```

``` shell
  tolerations:
  - key: CriticalAddonsOnly
    operator: Exists
  - effect: NoSchedule
    key: node-role.kubernetes.io/master
  - effect: NoSchedule
    key: node-role.kubernetes.io/control-plane
  - effect: NoExecute
    key: node.kubernetes.io/not-ready // cni 坏了
    operator: Exists
    tolerationSeconds: 300
  - effect: NoExecute
    key: node.kubernetes.io/unreachable // ping 不通
    operator: Exists
    tolerationSeconds: 300
```

let's why Kubernetes can failover



1. kubernetes 实现故障转移的机制。

2. 每个租户打上污点，来进行隔离。



<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221014181454206.png" alt="image-20221014181454206" style="zoom:33%;" />





<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221014182002578.png" alt="image-20221014182002578" style="zoom:33%;" />



kubectl describe, 去看他的event。,看pending， 是不是tolerant的问题。





<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221014182057597.png" alt="image-20221014182057597" style="zoom:33%;" />



<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221014182220766.png" alt="image-20221014182220766" style="zoom:33%;" />



<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221014182400188.png" alt="image-20221014182400188" style="zoom:33%;" />

Eviction is **the process of proactively terminating one or more Pods on resource-starved Nodes**





有状态的应用跑起来会很麻烦，需要驱逐掉。 

传统数据库上kubernetes要考虑，（如果接受1天的数据丢失，单节点可以每天进行备份）



kubernetes 推崇的是无状态。



kubernetes 跑 有状态的复杂度要比无状态服务很复杂。

<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221014183738454.png" alt="image-20221014183738454" style="zoom:33%;" />

keepalive，绑定一个externalIP。





ingress 是整个集群的一个网关。

![image-20221014185307549](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221014185307549.png)