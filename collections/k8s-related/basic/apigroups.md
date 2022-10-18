

![image-20220313142508421](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/basic/image-20220313142508421.png)



![image-20220313142533666](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/basic/image-20220313142533666.png)



![image-20220313142632176](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/basic/image-20220313142632176.png)





![image-20220313142730571](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/basic/image-20220313142730571.png)





![image-20220313142905440](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/basic/image-20220313142905440.png)





![image-20220313142926925](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/basic/image-20220313142926925.png)





![image-20220313143041713](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/basic/image-20220313143041713.png)





![image-20220313143211100](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/basic/image-20220313143211100.png)





![image-20220313143327153](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/basic/image-20220313143327153.png)



![image-20220313143507252](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/basic/image-20220313143507252.png)





## How to enable api groups

![image-20220313143851416](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/basic/image-20220313143851416.png)

![](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/basic/image-20220313143836789.png)



/etc/kubernetes/manifests/kube-apiserver.yaml

``` shell
[root@10-23-75-240 k8s]# vi /etc/kubernetes/manifests/kube-apiserver.yaml
[root@10-23-75-240 k8s]# systemctl restart kubelet
```



``` shell
│     Command:
│       kube-apiserver
│       --advertise-address=10.23.75.240
│       --allow-privileged=true
│       --authorization-mode=Node,RBAC
│       --client-ca-file=/etc/kubernetes/pki/ca.crt
│       --enable-admission-plugins=NodeRestriction
│       --enable-bootstrap-token-auth=true
│       --etcd-cafile=/etc/kubernetes/pki/etcd/ca.crt
│       --etcd-certfile=/etc/kubernetes/pki/apiserver-etcd-client.crt
│       --etcd-keyfile=/etc/kubernetes/pki/apiserver-etcd-client.key
│       --etcd-servers=https://127.0.0.1:2379
│       --kubelet-client-certificate=/etc/kubernetes/pki/apiserver-kubelet-client.crt
│       --kubelet-client-key=/etc/kubernetes/pki/apiserver-kubelet-client.key
│       --kubelet-preferred-address-types=InternalIP,ExternalIP,Hostname
│       --proxy-client-cert-file=/etc/kubernetes/pki/front-proxy-client.crt
│       --proxy-client-key-file=/etc/kubernetes/pki/front-proxy-client.key
│       --requestheader-allowed-names=front-proxy-client
│       --requestheader-client-ca-file=/etc/kubernetes/pki/front-proxy-ca.crt
│       --requestheader-extra-headers-prefix=X-Remote-Extra-
│       --requestheader-group-headers=X-Remote-Group
│       --requestheader-username-headers=X-Remote-User
│       --runtime-config=batch/v1beta1
│       --secure-port=6443
│       --service-account-issuer=https://kubernetes.default.svc.cluster.local
│       --service-account-key-file=/etc/kubernetes/pki/sa.pub
│       --service-account-signing-key-file=/etc/kubernetes/pki/sa.key
│       --service-cluster-ip-range=10.96.0.0/12
│       --tls-cert-file=/etc/kubernetes/pki/apiserver.crt
│       --tls-private-key-file=/etc/kubernetes/pki/apiserver.key
```





