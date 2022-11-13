<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221018170328680.png" alt="image-20221018170328680" style="zoom:33%;" />

<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221018171036894.png" alt="image-20221018171036894" style="zoom:33%;" />



![image-20221018172641348](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221018172641348.png)





<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221018172926460.png" alt="image-20221018172926460" style="zoom:33%;" />





<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221018175039133.png" alt="image-20221018175039133" style="zoom:33%;" />

![image-20221018175928975](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221018175928975.png)



<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221018180024019.png" alt="image-20221018180024019" style="zoom:33%;" />

<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221018180152560.png" alt="image-20221018180152560" style="zoom:33%;" />







<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221018180331790.png" alt="image-20221018180331790" style="zoom:33%;" />

<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221018180442386.png" alt="image-20221018180442386" style="zoom:33%;" />

<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221018180704574.png" alt="image-20221018180704574" style="zoom:33%;" />

<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221018180827256.png" alt="image-20221018180827256" style="zoom:33%;" />



<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221018181002391.png" alt="image-20221018181002391" style="zoom:25%;" />

Containers 的 pods怎么看 ?

他本身符合 cri-o命令，可以用

``` shell
 crictl pods
```

 



在containerd里面 sandbox是一个配置项。

![image-20221018181347590](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221018181347590.png)

![image-20221018182028485](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221018182028485.png)



crictl 把sandbox 和用户容器分开了，要通过crictl pods的命令看sandbox



crictl ps 去看用户容器。



crictl inspect sandboxid



crictl inspect containerid



crictl 本身不提供 push 命令。



因此docker 变成了镜像构建工具了。



unix 套接字，kubelet会去连接那个地址和容器运行时通信。



![image-20221018182050584](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221018182050584.png)

