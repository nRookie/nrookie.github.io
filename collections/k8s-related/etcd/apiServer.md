![image-20221007124600687](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221007124600687.png)

<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221007124927445.png" alt="image-20221007124927445" style="zoom:50%;" />

<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221007125001698.png" alt="image-20221007125001698" style="zoom:50%;" />



<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221007125527650.png" alt="image-20221007125527650" style="zoom:50%;" />

<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221007130126109.png" alt="image-20221007130126109" style="zoom:50%;" />



<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221007130333868.png" alt="image-20221007130333868" style="zoom:50%;" />



![image-20221007135506648](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221007135506648.png)



CA.



apiServer 本身就是一个ca， 都放在 /etc/kubernetes/pki/ 下面

![image-20221007140609940](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221007140609940.png)

ca.crt 是root ca。

https 有两个作用， 

1. 加密
2. 一个网站如果运行在https模式下面，他就需要证书，需要一个可信的网站办法证书。





userAccount 和 serviceAccount ？

``` shell
k get sa
k create sa demo
k get sa demo -oyaml
```

没有secret了

``` shell
https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.24.md#urgent-upgrade-notes
```



如何添加新的用户？ (鉴权的时候会讲。先放着)

```
apiVersion: certificates.k8s.io/v1
kind: CertificateSigningRequest
metadata:
  name: myuser
spec:
  request: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURSBSRVFVRVNULS0tLS0KTUlJQ21UQ0NBWUVDQVFBd1ZERUxNQWtHQTFVRUJoTUNRMDR4RXpBUkJnTlZCQWdNQ2xOdmJXVXRVM1JoZEdVeApJVEFmQmdOVkJBb01HRWx1ZEdWeWJtVjBJRmRwWkdkcGRITWdVSFI1SUV4MFpERU5NQXNHQTFVRUF3d0VjV2x1Clp6Q0NBU0l3RFFZSktvWklodmNOQVFFQkJRQURnZ0VQQURDQ0FRb0NnZ0VCQUpHZkhVSmVjU1kvZ2hYU25kcmEKZWEweGVsc3JkZGM0TFJWOWYvd0RjR0lWdmRlSk10cGhYVDFqVnNMeW1iQ0ZkS3JHaTFqWFpZSVdwclpYdXRNSQo1L3hxSWFpK2hsazB1SVFIbncyalVuSHY1VEFwZnc2V0YzSU9aWFh6QmpST3dDSkFVR3VvbFVoMjJNSFlqSEw5CmhVT3Bnbk5YS1dqdGplaTZDY3RxYm1ibUx0M3hhUGVaRWozbld2c1BxZjVBYUEwRjRySUd0blpCK3pDUmh4UzgKODJIT0EyZE1XanpjZWRvM3ZUVy9sU29KWHQza3VtZVpJYkV5S1A2TnJpMURiL0dRWkFtd2tpMFFKRzdtUVRJNAplUnpTbTJtWWZ0VUs5cWhQNis3dFE2NWltNDl0MjM0SjJkempsUkxXRy96VVZ2QUs0RVZXU2tpZHdsZE05ZHZEClZkOENBd0VBQWFBQU1BMEdDU3FHU0liM0RRRUJDd1VBQTRJQkFRQWZPQ2gyRWJVN2J3TGhqRDArRzZseVhWR3EKMlNEbnN4OEhRMXJmenVmYlhKaUR1QzF2QXVKOVpmSzA3NlQ4ek1PUTdrY3NaS29YeG15Z0hJVnZDVUlTVmtrcApaVlVqNjJycTVLcVU0dTRhNGxjdXdMTnRINVkwMy9xWnBBNDl5anR4SHFCa0d0RWFLVlVaNVFuMWdHcDVCdlhhCjFmcjZsUTgrSjZ3dTRXLzlZTEhmRmpQZ0dhOTRhVGtvOXMyMkUreDlZQ2VJZTZwWUVXd0JuN3BMek10STlkMGsKeDdxTkZJdUkzZ3hHMjRyQnhqOGxNelJFRzYxNVZoTEdhUjRiK2NOOE8vV0loNGc3OWVqdFRYZUkvYmw2Y1JWOQpHSVo3eGMvOERVc3ExL3dGZktNZVp5SytHSDVxWlJ1M1pYMXJpNkp6SDZNdGp5WHd4c015eGhDVHhIeFEKLS0tLS1FTkQgQ0VSVElGSUNBVEUgUkVRVUVTVC0tLS0tCg==
  signerName: kubernetes.io/kube-apiserver-client
  expirationSeconds: 86400  # one day
  usages:
  - client auth
```



userAccount 和 serviceAccount的区别

来自kubernetes外部？

serviceAccount存在kubernetes里面。

![image-20221007142641660](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221007142641660.png)



![image-20221007165535530](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221007165535530.png)



![image-20221007165547871](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221007165547871.png)





![image-20221007165635596](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221007165635596.png)

![image-20221007165658894](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221007165658894.png)

![image-20221007165720903](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221007165720903.png)







![image-20221007184523058](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221007184523058.png)





<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221007185141406.png" alt="image-20221007185141406" style="zoom:50%;" />

<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221007185322337.png" alt="image-20221007185322337" style="zoom:50%;" />

<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221007185351783.png" alt="image-20221007185351783" style="zoom:50%;" />



<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221007185558139.png" alt="image-20221007185558139" style="zoom:33%;" />



![image-20221007185916531](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221007185916531.png)





<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221007190144878.png" alt="image-20221007190144878" style="zoom:33%;" />

![image-20221007190315211](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221007190315211.png)



<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221007190357934.png" alt="image-20221007190357934" style="zoom:50%;" />



<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221007190510920.png" alt="image-20221007190510920" style="zoom:33%;" />



<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221007192935867.png" alt="image-20221007192935867" style="zoom:33%;" />

![image-20221008004112451](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008004112451.png)





![image-20221008004641735](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008004641735.png)



![image-20221008004934142](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008004934142.png)



<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008004958978.png" alt="image-20221008004958978" style="zoom:50%;" />



<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008005747732.png" alt="image-20221008005747732" style="zoom:33%;" />

<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008005801574.png" alt="image-20221008005801574" style="zoom:33%;" />





<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008010610055.png" alt="image-20221008010610055" style="zoom:33%;" />



https://kubernetes.io/docs/reference/access-authn-authz/webhook/



![image-20221008011212381](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008011212381.png)





https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/



![image-20221008014406560](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008014406560.png)



![image-20221008014606891](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008014606891.png)



![image-20221008014756288](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008014756288.png)





![image-20221008015015067](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008015015067.png)



![image-20221008015324368](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008015324368.png)



![image-20221008015801752](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008015801752.png)





![image-20221008020026421](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008020026421.png)



![image-20221008020440736](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008020440736.png)

![image-20221008111536712](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008111536712.png)

![image-20221008111558328](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008111558328.png)



![image-20221008111741857](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008111741857.png)



![image-20221008111756086](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008111756086.png)

<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008112107084.png" alt="image-20221008112107084" style="zoom:50%;" />



<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008112348357.png" alt="image-20221008112348357" style="zoom:50%;" />



![image-20221008113040447](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008113040447.png)





``` shell
k get PriorityLevelConfiguration exempt -oyaml

k get flowschema exempt -oyaml
 
k get flowschema 

k get PriorityLevelConfiguration

kubectl get --raw /debug/api_priority_and_fairness/dump_priority_levels
```



flowschema 和 prioritylevel的关系

<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008113644953.png" alt="image-20221008113644953" style="zoom:50%;" />

需要对request进行分类，哪一类的优先级需要给高优先级，哪一类需要给低优先级。

prioritylevel 是在通过 flow归类以后，在同一个flow内的优先级

``` shell
 k get flowschema service-accounts -oyaml
```





<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008114005633.png" alt="image-20221008114005633" style="zoom:50%;" />

<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008114051616.png" alt="image-20221008114051616" style="zoom:50%;" />

不同的serviceaccount 就是不同的flow。

通过什么样的方式限流， 就是通过prioritylevelconfiguration定义的。

``` shell
k get PriorityLevelConfiguration workload-low -oyaml

```

<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008114258386.png" alt="image-20221008114258386" style="zoom:50%;" />







![image-20221008115048414](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008115048414.png)





![image-20221008115059035](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008115059035.png)



API Server 是无状态的 Rest Server





![image-20221008115409410](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008115409410.png)



<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008115738924.png" alt="image-20221008115738924" style="zoom:50%;" />

<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008115841771.png" alt="image-20221008115841771" style="zoom:50%;" />



<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008120122369.png" alt="image-20221008120122369" style="zoom:50%;" />



<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008120252866.png" alt="image-20221008120252866" style="zoom:50%;" />





``` shell
 k get svc
NAME         TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)           AGE
envoy        NodePort    10.98.214.99    <none>        10000:32349/TCP   3d9h
kubernetes   ClusterIP   10.96.0.1       <none>        443/TCP           18d
nginx        NodePort    10.101.66.251   <none>        80:32310/TCP      3d23h

curl https://10.96.0.1 -k
 
```



![image-20221008120916641](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008120916641.png)



![image-20221008121524757](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008121524757.png)





![image-20221008121409497](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008121409497.png)





k8s所有的对象是如何注册进k8s里面的？ （除了apiHandler）

每个对象都有自己的一个apiService。就是把api的那些handler注册一下。

ApiServer是一个aggregator， 是一个分布式的api网关。

``` shell
k get apiservice
```







![image-20221008121818686](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008121818686.png)





group 定义对象放在哪里面？

把一类对象放在一个group里面。

同一个group在访问etcd的时候会共用一个connection。

Node 和 pod 是 共用一个group的



Kind 代表他是啥



Version 

Internal version，

 external version 是面向用户的。





![image-20221008122312709](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008122312709.png)



![image-20221008122414402](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008122414402.png)

TypeMeta 定义它是什么

<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008123051813.png" alt="image-20221008123051813" style="zoom:50%;" />





![image-20221008122830717](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008122830717.png)





external version 都带着版本信息



<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008123203023.png" alt="image-20221008123203023" style="zoom:50%;" />



<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008123524138.png" alt="image-20221008123524138" style="zoom:50%;" />



<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008123556306.png" alt="image-20221008123556306" style="zoom:50%;" />



<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008123904381.png" alt="image-20221008123904381" style="zoom:50%;" />



![image-20221008124027109](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008124027109.png)



<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008124113054.png" alt="image-20221008124113054" style="zoom:50%;" />





![image-20221008124135650](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008124135650.png)





``` shell
k get po centos -oyaml

```



Subresource

<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008124317959.png" alt="image-20221008124317959" style="zoom:50%;" />

<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008130413360.png" alt="image-20221008130413360" style="zoom:50%;" />



<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008130447980.png" alt="image-20221008130447980" style="zoom:50%;" />



![image-20221008130558290](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008130558290.png)



<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008130744313.png" alt="image-20221008130744313" style="zoom:50%;" />





不要从apiServer开始读代码



为什么要注册apiGroup



新加的apiserver 是怎么访问到。



metric-server的配置



![image-20221008132032611](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008132032611.png)



<img src="/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/etcd/image-20221008132308639.png" alt="image-20221008132308639" style="zoom:50%;" />



kubelet监听这个目录（/etc/kubernetes/manifest，如果有组件，就会读取这些组件，创建pod

apiServer这个pod是apiServer补出来的 叫做mirrorpod







## User accounts versus service accounts



Kubernetes distinguishes between the concept of a user account and a service account for a number of reasons:

- User accounts are for humans. Service accounts are for processes, which run in pods.
- User accounts are intended to be global. Names must be unique across all namespaces of a cluster. Service accounts are namespaced.
- Typically, a cluster's user accounts might be synced from a corporate database, where new user account creation requires special privileges and is tied to complex business processes. Service account creation is intended to be more lightweight, allowing cluster users to create service accounts for specific tasks by following the principle of least privilege.
- Auditing considerations for humans and service accounts may differ.
- A config bundle for a complex system may include definition of various service accounts for components of that system. Because service accounts can be created without many constraints and have namespaced names, such config is portable.



## Service account automation

Three separate components cooperate to implement the automation around service accounts:

- A `ServiceAccount` admission controller
- A Token controller
- A `ServiceAccount` controller



### ServiceAccount Admission Controller

The modification of pods is implemented via a plugin called an [Admission Controller](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/). It is part of the API server. It acts synchronously to modify pods as they are created or updated. When this plugin is active (and it is by default on most distributions), then it does the following when a pod is created or modified:

1. If the pod does not have a `ServiceAccount` set, it sets the `ServiceAccount` to `default`.
2. It ensures that the `ServiceAccount` referenced by the pod exists, and otherwise rejects it.
3. It adds a `volume` to the pod which contains a token for API access if neither the ServiceAccount `automountServiceAccountToken` nor the Pod's `automountServiceAccountToken` is set to `false`.
4. It adds a `volumeSource` to each container of the pod mounted at `/var/run/secrets/kubernetes.io/serviceaccount`, if the previous step has created a volume for the ServiceAccount token.
5. If the pod does not contain any `imagePullSecrets`, then `imagePullSecrets` of the `ServiceAccount` are added to the pod.



### ServiceAccount Admission Controller

The modification of pods is implemented via a plugin called an [Admission Controller](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/). It is part of the API server. It acts synchronously to modify pods as they are created or updated. When this plugin is active (and it is by default on most distributions), then it does the following when a pod is created or modified:

1. If the pod does not have a `ServiceAccount` set, it sets the `ServiceAccount` to `default`.
2. It ensures that the `ServiceAccount` referenced by the pod exists, and otherwise rejects it.
3. It adds a `volume` to the pod which contains a token for API access if neither the ServiceAccount `automountServiceAccountToken` nor the Pod's `automountServiceAccountToken` is set to `false`.
4. It adds a `volumeSource` to each container of the pod mounted at `/var/run/secrets/kubernetes.io/serviceaccount`, if the previous step has created a volume for the ServiceAccount token.
5. If the pod does not contain any `imagePullSecrets`, then `imagePullSecrets` of the `ServiceAccount` are added to the pod.

### ServiceAccount Admission Controller

The modification of pods is implemented via a plugin called an [Admission Controller](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/). It is part of the API server. It acts synchronously to modify pods as they are created or updated. When this plugin is active (and it is by default on most distributions), then it does the following when a pod is created or modified:

1. If the pod does not have a `ServiceAccount` set, it sets the `ServiceAccount` to `default`.
2. It ensures that the `ServiceAccount` referenced by the pod exists, and otherwise rejects it.
3. It adds a `volume` to the pod which contains a token for API access if neither the ServiceAccount `automountServiceAccountToken` nor the Pod's `automountServiceAccountToken` is set to `false`.
4. It adds a `volumeSource` to each container of the pod mounted at `/var/run/secrets/kubernetes.io/serviceaccount`, if the previous step has created a volume for the ServiceAccount token.
5. If the pod does not contain any `imagePullSecrets`, then `imagePullSecrets` of the `ServiceAccount` are added to the pod.



#### Bound Service Account Token Volume

**FEATURE STATE:** `Kubernetes v1.22 [stable]`

The ServiceAccount admission controller will add the following projected volume instead of a Secret-based volume for the non-expiring service account token created by the Token controller.



``` shell
- name: kube-api-access-<random-suffix>
  projected:
    defaultMode: 420 # 0644
    sources:
      - serviceAccountToken:
          expirationSeconds: 3607
          path: token
      - configMap:
          items:
            - key: ca.crt
              path: ca.crt
          name: kube-root-ca.crt
      - downwardAPI:
          items:
            - fieldRef:
                apiVersion: v1
                fieldPath: metadata.namespace
              path: namespace

```

https://kubernetes.io/docs/reference/access-authn-authz/service-accounts-admin/



