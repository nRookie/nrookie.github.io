<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008180455437.png" alt="image-20221008180455437" style="zoom:50%;" />





监听pod的属性，NodeName



``` shell
naqing19950908@instance-4:~/101/module6/serviceaccount$ kubectl get pods downward-api-pod   -oyaml | grep -i nodeName
  nodeName: instance-5

```



<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008181217106.png" alt="image-20221008181217106" style="zoom:50%;" />

![image-20221008181319910](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008181319910.png)



插件



<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008181941840.png" alt="image-20221008181941840" style="zoom:50%;" />



<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008182112901.png" alt="image-20221008182112901" style="zoom:50%;" />



![image-20221010100232777](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221010100232777.png)



request 本身的意思，至少需要多少memory。

limits我这个应用最多用多少资源。

cpu: 1m (千分之一的cpu)



``` shell
k get node -oyaml
```



<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221010100807934.png" alt="image-20221010100807934" style="zoom:50%;" />



allocable: 在原来的capacity的基础之上

Capacity: // 代表的是这个节点的能力，总的资源。



``` shell
crictl ps | grep nginx
crictl inspect 39e4826018353

```



``` shell
        {
          "destination": "/sys/fs/cgroup",
          "type": "cgroup",
          "source": "cgroup",
--
          ],
          "memory": {
            "limit": 1073741824
          },
          "cpu": {
            "shares": 102,
            "quota": 100000,
            "period": 100000
          }
        },
        "cgroupsPath": "kubepods-burstable-pod7a9770ba_439e_4aa2_a8af_093f9f670bbb.slice:cri-containerd:39e48260183537779b39bcf896d5b9252fa3c9c1c952393cd806df104bd34d76",
```

``` shell
        "cgroupsPath": "kubepods-burstable-pod7a9770ba_439e_4aa2_a8af_093f9f670bbb.slice:cri-containerd:39e48260183537779b39bcf896d5b9252fa3c9c1c952393cd806df104bd34d76",
naqing19950908@instance-5:~$ cd /sys/fs/cgroup/cpu
cpu/         cpu,cpuacct/ cpuacct/     cpuset/      
naqing19950908@instance-5:~$ cd /sys/fs/cgroup/cpu/kubepods.slice/kubepods-burstable.slice/
naqing19950908@instance-5:/sys/fs/cgroup/cpu/kubepods.slice/kubepods-burstable.slice$ cd kubepods-burstable-pod7a9770ba_439e_4aa2_a8af_093f9f670bbb.slice/
naqing19950908@instance-5:/sys/fs/cgroup/cpu/kubepods.slice/kubepods-burstable.slice/kubepods-burstable-pod7a9770ba_439e_4aa2_a8af_093f9f670bbb.slice$ 
naqing19950908@instance-5:/sys/fs/cgroup/cpu/kubepods.slice/kubepods-burstable.slice/kubepods-burstable-pod7a9770ba_439e_4aa2_a8af_093f9f670bbb.slice$ cat cpu.cfs_period_us // 100000 微秒的时间片里面
100000
naqing19950908@instance-5:/sys/fs/cgroup/cpu/kubepods.slice/kubepods-burstable.slice/kubepods-burstable-pod7a9770ba_439e_4aa2_a8af_093f9f670bbb.slice$ cat cpu.cfs_quota_us  // 100000 可以用10万微秒
100000
naqing19950908@instance-5:/sys/fs/cgroup/cpu/kubepods.slice/kubepods-burstable.slice/kubepods-burstable-pod7a9770ba_439e_4aa2_a8af_093f9f670bbb.slice$ cat cpu.shares 
102 // 100m, 1个cpu是1024， 0.1 个cpu是102
```

<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221010120205279.png" alt="image-20221010120205279" style="zoom:50%;" />

cpu.shares 的作用 当两个作业发生cpu竞争的时候，要按照这两个cpu-cgroup的cpushare的比例来分配cpu时间片。

可以用来做超售卖。



不定义资源：best efforts

![image-20221010120658055](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221010120658055.png)







limit-range 可以用进行默认配置。

``` shell
apiVersion: v1
kind: LimitRange
metadata:
  name: mem-limit-range
spec:
  limits:
    - default: // 如果pod没有填写过任何内存的资源的请求， 如果这个namespace里面有limitRange对象，且有这些值，
        memory: 512Mi 
      defaultRequest:
        memory: 256Mi
      type: Container
```

不过limit-range不常用, 因为也会去限制InitContainer





1. pods 本身是一堆容器的组合
2. 要在主应用跑起来之前，去跑一些初始化操作怎么？ InitContainers. InitContainers 本身也是一个数组，InitContainers不跑完，Containers不会起来的。



<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221010121634056.png" alt="image-20221010121634056" style="zoom:50%;" />

``` shell
```

<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221010125038544.png" alt="image-20221010125038544" style="zoom:50%;" />



<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221010125644179.png" alt="image-20221010125644179" style="zoom:50%;" />





Cpu.shares 是个相对值， // 没有超过上限就用shares。

一个cpu是512， 一个cpu是1024.

那么这两个进程会用1:2的比率分享cpu。



cpu.cfs_period_us 是一个绝对的值， 10万个时间片里面用多少。

cpu.cfs_quota_us 用多少。 限制了上线。

系统默认，是按10万微秒去跳的。



![image-20221010130017550](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221010130017550.png)







hpa, vpa. 原生的方法，很难解决。



<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221010130218923.png" alt="image-20221010130218923" style="zoom:50%;" />

cpuset 



