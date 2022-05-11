``` shell
kubectl apply \
    -f scaling/go-demo-5-no-sidecar-mem.yml \
    --record
    
    
unable to recognize "scaling/go-demo-5-no-sidecar-mem.yml": no matches for kind "Role" in version "rbac.authorization.k8s.io/v1beta1"
unable to recognize "scaling/go-demo-5-no-sidecar-mem.yml": no matches for kind "RoleBinding" in version "rbac.authorization.k8s.io/v1beta1"
```



## Enabling or disabling API groups 

Certain resources and API groups are enabled by default. You can enable or disable them by setting `--runtime-config` on the API server. The `--runtime-config` flag accepts comma separated `<key>[=<value>]` pairs describing the runtime configuration of the API server. If the `=<value>` part is omitted, it is treated as if `=true` is specified. For example:

- to disable `batch/v1`, set `--runtime-config=batch/v1=false`
- to enable `batch/v2alpha1`, set `--runtime-config=batch/v2alpha1`



``` shell
Modify /etc/kubernetes/manifests/kube-apiserver.manifest

And then restart kubelet: systemctl restart kubelet
```



### in here we change the rbac version from v1beta1 to v1

 vi scaling/go-demo-5-no-sidecar-mem.yml 

``` yaml
[root@10-23-75-240 k8s-specs]# kubectl apply     -f scaling/go-demo-5-no-sidecar-mem.yml     --record
Flag --record has been deprecated, --record will be removed in the future
namespace/go-demo-5 unchanged
ingress.networking.k8s.io/api unchanged
serviceaccount/db unchanged
role.rbac.authorization.k8s.io/db unchanged
rolebinding.rbac.authorization.k8s.io/db unchanged
statefulset.apps/db configured
service/db unchanged
deployment.apps/api configured
service/api unchanged
```



### Prepare



``` yaml
[root@10-23-75-240 k8s-specs]# cat pv/mongo-nfs-pv.yml 
apiVersion: v1
kind: PersistentVolume
metadata:
  name: mongo-data
  namespace: go-demo-5
  labels:
    type: local
    author: qing.na
spec:
  storageClassName: nfs-client
  capacity:
    storage: 40Gi
  accessModes:
    - ReadWriteMany
  volumeMode: Filesystem
  mountOptions:
    - hard
    - nfsvers=4.1
 # persistentVolumeReclaimPolicy: Retain
  nfs:
    path: /data/
    server: 172.16.0.51
    readOnly: false
```



``` shell
[root@10-23-75-240 k8s-specs]# cat pv/mongo-nfs-pvc.yml 
kind: PersistentVolumeClaim
apiVersion: v1
metadata:
  name: mongo-data
  namespace: go-demo-5
spec:
  storageClassName: nfs-client
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 10Gi
```



``` yaml
[root@10-23-75-240 k8s-specs]# cat scaling/go-demo-5-no-sidecar-mem.yml 
apiVersion: v1
kind: Namespace
metadata:
  name: go-demo-5

---

apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: api
  namespace: go-demo-5
  annotations:
    kubernetes.io/ingress.class: "nginx"
    ingress.kubernetes.io/ssl-redirect: "false"
    nginx.ingress.kubernetes.io/ssl-redirect: "false"
spec:
  rules:
  - http:
      paths:
      - path: /demo
        pathType: ImplementationSpecific
        backend:
          service:
            name: api
            port:
              number: 8080

---

apiVersion: v1
kind: ServiceAccount
metadata:
  name: db
  namespace: go-demo-5

---

kind: Role
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: db
  namespace: go-demo-5
rules:
- apiGroups: [""]
  resources: ["pods"]
  verbs: ["list"]

---

apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: db
  namespace: go-demo-5
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: db
subjects:
- kind: ServiceAccount
  name: db

---

apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: db
  namespace: go-demo-5
spec:
  serviceName: db
  selector:
    matchLabels:
      app: db
  template:
    metadata:
      labels:
        app: db
    spec:
      serviceAccountName: db
      terminationGracePeriodSeconds: 10
      containers:
      - name: db
        image: mongo:3.3
        command:
          - mongod
          - "--replSet"
          - rs0
          - "--smallfiles"
          - "--noprealloc"
        ports:
          - containerPort: 27017
        resources:
          limits:
            memory: "150Mi"
            cpu: 0.2
          requests:
            memory: "100Mi"
            cpu: 0.1
        volumeMounts:
        - name: mongo-data
          mountPath: /data/db
      - name: db-sidecar
        image: cvallance/mongo-k8s-sidecar
        env:
        - name: MONGO_SIDECAR_POD_LABELS
          value: "app=db"
        - name: KUBE_NAMESPACE
          value: go-demo-5
        - name: KUBERNETES_MONGO_SERVICE_NAME
          value: db
      volumes:
        - name: mongo-data
          persistentVolumeClaim:
            claimName: mongo-data
---

apiVersion: v1
kind: Service
metadata:
  name: db
  namespace: go-demo-5
spec:
  ports:
  - port: 27017
  clusterIP: None
  selector:
    app: db

---

apiVersion: apps/v1
kind: Deployment
metadata:
  name: api
  namespace: go-demo-5
spec:
  selector:
    matchLabels:
      app: api
  template:
    metadata:
      labels:
        app: api
    spec:
      containers:
      - name: api
        image: vfarcic/go-demo-5
        env:
        - name: DB
          value: db
        readinessProbe:
          httpGet:
            path: /demo/hello
            port: 8080
          periodSeconds: 1
        livenessProbe:
          httpGet:
            path: /demo/hello
            port: 8080
        resources:
          limits:
            memory: 15Mi
            cpu: 0.1
          requests:
            memory: 10Mi
            cpu: 0.01

---

apiVersion: v1
kind: Service
metadata:
  name: api
  namespace: go-demo-5
spec:
  ports:
  - port: 8080
  selector:
    app: api
[root@10-23-75-240 k8s-specs]# 

```





server is created for the moment.



``` shell
[root@10-23-75-240 k8s-specs]# kubectl -n go-demo-5 get pods
NAME                   READY   STATUS    RESTARTS        AGE
api-57dd9cd978-lpr5g   1/1     Running   2 (3m43s ago)   4m12s
db-0                   2/2     Running   0               4m12s
```





![image-20220321104418628](/Users/user/playground/share/nrookie.github.io/collections/k8s-related/metrics/image-20220321104418628.png)



As you hopefully know, we should aim at having at least two replicas of each Pod, as long as they are scalable. Still, neither of the two had `replicas` defined. That is intentional. The fact that we can specify the number of replicas of a Deployment or a StatefulSet does not mean that we should. At least, not always.



# Where to set replicas? [#](https://www.educative.io/module/lesson/kubernetes-monitoring-logging-auto-scaling/m7mxQ0L0BgA#Where-to-set-replicas?-)



If the number of replicas is static and you have no intention to scale (or de-scale) your application over time, set `replicas` as part of your Deployment or StatefulSet definition. If, on the other hand, you plan to change the number of replicas based on memory, CPU, or other metrics, use `HorizontalPodAutoscaler` resource instead.



Let’s take a look at a simple example of a `HorizontalPodAutoscaler`.



``` shell
cat scaling/go-demo-5-api-hpa.yml
```

``` yaml
apiVersion: autoscaling/v2beta1
kind: HorizontalPodAutoscaler
metadata:
  name: api
  namespace: go-demo-5
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: api
  minReplicas: 2
  maxReplicas: 5
  metrics:
  - type: Resource
    resource:
      name: cpu
      targetAverageUtilization: 80
  - type: Resource
    resource:
      name: memory
      targetAverageUtilization: 80
```

The definition uses `HorizontalPodAutoscaler` targeting the `api` Deployment. Its boundaries are a minimum of two and a maximum of five replicas. Those limits are fundamental. Without them, we’d run a risk of scaling up into infinity or scaling down to zero replicas. The `minReplicas` and `maxReplicas` fields are a safety net.



## Scale or Descale [#](https://www.educative.io/module/lesson/kubernetes-monitoring-logging-auto-scaling/m7mxQ0L0BgA#Scale-or-Descale-)



The key section of the definition is `metrics`. It provides formulas Kubernetes should use to decide whether it should scale (or de-scale) a resource. In our case, we’re using the `Resource` type entries. They are targeting the average utilization of eighty percent for memory and CPU. If the actual usage of either of the two deviates, Kubernetes will scale (or de-scale) the resource.



The key section of the definition is `metrics`. It provides formulas Kubernetes should use to decide whether it should scale (or de-scale) a resource. In our case, we’re using the `Resource` type entries. They are targeting the average utilization of eighty percent for memory and CPU. If the actual usage of either of the two deviates, Kubernetes will scale (or de-scale) the resource



>  Please note that we used the `v2beta1` version of the API and you might be wondering why we chose that one instead of the stable and production-ready v1. After all, `beta 1` releases are still far from being polished enough for general usage. The reason is simple. HorizontalPodAutoscaler v1 is too basic. It only allows scaling based on the CPU. Even our simple example goes beyond that by adding memory to the mix. Later on, we’ll extend it even more. So, while v1 is considered stable, it does not provide much value, and we can either wait until v2 is released or start experimenting with v2beta releases right away. We are opting for the latter option. By the time you read this, more stable releases are likely to exist and to be supported in your Kubernetes cluster. If that’s the case, feel free to change `apiVersion` before applying the definition.

Let’s apply the definition that creates the `HorizontalPodAutoscaler (HPA)`.



``` shell
kubectl apply \
    -f scaling/go-demo-5-api-hpa.yml \
    --record
```



Next, we’ll take a look at the information we’ll get by retrieving the `HPA` resources.

```shell
[root@10-23-75-240 k8s-specs]# kubectl -n go-demo-5 get hpa
NAME   REFERENCE        TARGETS            MINPODS   MAXPODS   REPLICAS   AGE
api    Deployment/api   39%/80%, 10%/80%   2         5         2          3m42s

```





### Current number of replicas below `minReplicas` 



We can see that both CPU and memory utilization are way below the expected utilization of `80%`. Still, Kubernetes increased the number of replicas from one to two because that’s the minimum we defined. We made the contract stating that the `api` Deployment should never have less than two replicas, and Kubernetes complied with that by scaling up even if the resource utilization is way below the expected average utilization. We can confirm that behavior through the events of the `HorizontalPodAutoscaler`.







### Current number of replicas below `minReplicas` [#](https://www.educative.io/module/lesson/kubernetes-monitoring-logging-auto-scaling/m7mxQ0L0BgA#Current-number-of-replicas-below-minReplicas-)

We can see that both CPU and memory utilization are way below the expected utilization of `80%`. Still, Kubernetes increased the number of replicas from one to two because that’s the minimum we defined. We made the contract stating that the `api` Deployment should never have less than two replicas, and Kubernetes complied with that by scaling up even if the resource utilization is way below the expected average utilization. We can confirm that behavior through the events of the `HorizontalPodAutoscaler`.



![image-20220321110050222](/Users/user/playground/share/nrookie.github.io/collections/k8s-related/metrics/image-20220321110050222.png)









## trouble-shooting



``` shell
│ on error: desc = "transport: Error while dialing dial unix /var/run/antrea/cni.sock: connect: no such file or directory"                                                                      
│   Warning  FailedCreatePodSandBox  5m13s                  kubelet            Failed to create pod sandbox: rpc error: code = Unknown desc = failed to set up sandbox container "fac2e49559d0  
│ 9b89e9137c12646104b155c627628351765033b8cc25bf64ef94" network for pod "db-0": networkPlugin cni failed to set up pod "db-0_go-demo-5" network: rpc error: code = Unavailable desc = connecti  
│ on error: desc = "transport: Error while dialing dial unix /var/run/antrea/cni.sock: connect: no such file or directory"                                                                      
│   Normal   SandboxChanged          5m9s (x12 over 5m20s)  kubelet            Pod sandbox changed, it will be killed and re-created.                                                           
│   Warning  FailedCreatePodSandBox  21s (x274 over 5m12s)  kubelet            (combined from similar events): Failed to create pod sandbox: rpc error: code = Unknown desc = failed to set up  
│  sandbox container "eb1887bd20ac39cd061f73383caf6076420eeed4d47931d42377bf14d61fa611" network for pod "db-0": networkPlugin cni failed to set up pod "db-0_go-demo-5" network: rpc error: co  
│ de = Unavailable desc = connection error: desc = "transport: Error while dialing dial unix /var/run/antrea/cni.sock: connect: no such file or directory"                                      
│                                                                                                                                                           
```



![image-20220403110227448](/Users/user/playground/share/nrookie.github.io/collections/k8s-related/metrics/image-20220403110227448.png)

``` shell
[root@10-23-87-58 k8s]# rm /etc/cni/net.d/10-antrea.conflist 
rm：是否删除普通文件 "/etc/cni/net.d/10-antrea.conflist"？y
```



it seems like kubelet searching cni through etc cni folder
