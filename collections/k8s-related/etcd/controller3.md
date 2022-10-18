![image-20221017040003066](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221017040003066.png)



<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221017040906255.png" alt="image-20221017040906255" style="zoom:33%;" />



<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221017041525004.png" alt="image-20221017041525004" style="zoom:33%;" />





![image-20221017041658387](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221017041658387.png)



![image-20221017050123794](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221017050123794.png)



statefulset 里面会有service-name， servicename的作用，

statefulsetcontroller 创建pod的时候会写两个额外的值。

（hostname, subdomain).



Core-dns 就会去做一件事情，会在coredns里面记录一条记录。

<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221017050243514.png" alt="image-20221017050243514" style="zoom:50%;" />

这个dns就会指向ss-0这个pod的IP



<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221017050651828.png" alt="image-20221017050651828" style="zoom:50%;" />

![image-20221017051115990](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221017051115990.png)





``` shell
k api-resoureces
```

