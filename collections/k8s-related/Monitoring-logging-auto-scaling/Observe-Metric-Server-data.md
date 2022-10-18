## Flow of data using `kubectl top` [#](https://www.educative.io/module/lesson/kubernetes-monitoring-logging-auto-scaling/m72Y9G9PZmG#Flow-of-data-using-kubectl-top-)

When we request metrics through `kubectl top`, the flow of data is almost the same as when the scheduler makes requests. A request is sent to the **API Server (Master Metrics API)**, which gets data from the `Metrics Server` which, in turn, was collecting information from Kubeletes running on the nodes of the cluster.



![image-20220317004749896](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/Monitoring-logging-auto-scaling/image-20220317004749896.png)



## `Metrics Server` for machines [#](https://www.educative.io/module/lesson/kubernetes-monitoring-logging-auto-scaling/m72Y9G9PZmG#Metrics-Server-for-machines-)

There are two important things to note. First of all, it provides current (or short-term) memory and CPU utilization of the containers running inside a cluster. The second and more important note is that we will not use it directly. `Metrics Server` was not designed for humans but for machines. We’ll get there later. For now, remember that there is a thing called `Metrics Server` and that you should not use it directly (once you adopt a tool that will scrape its metrics).

------

Now that we explored `Metrics Server`, we’ll try to put it to good use and learn how to auto-scale our Pods based on resource utilization, in the next lesson.





# Get Started with Auto-Scaling Pods



In this lesson, we will first deploy an app and see how to change the number of replicas based on memory, CPU, or other metrics through `HorizontalPodAutoScaler` resource.



Our goal is to deploy an application that will be automatically scaled (or de-scaled) depending on its use of resources. We’ll start by deploying an app first, and discuss how to accomplish auto-scaling later.



> I already warned you that I assume that you are familiar with Kubernetes and that in this course we’ll explore a particular topic of monitoring, alerting, scaling, and a few other things. We will not discuss Pods, StatefulSets, Deployments, Services, Ingress, and other “basic” Kubernetes resources.



# Deploy an application [#](https://www.educative.io/module/lesson/kubernetes-monitoring-logging-auto-scaling/m7mxQ0L0BgA#Deploy-an-application-)

Let’s take a look at a definition of the application we’ll use in our examples.

```shell
cat scaling/go-demo-5-no-sidecar-mem.yml
```



If you are familiar with Kubernetes, the YAML definition should be self-explanatory. We’ll comment on only the parts that are relevant for auto-scaling.

The **output**, limited to the relevant parts, is as follows.

```yaml
...
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: db
  namespace: go-demo-5
spec:
  ...
  template:
    ...
    spec:
      ...
      containers:
      - name: db
        ...
        resources:
          limits:
            memory: "150Mi"
            cpu: 0.2
          requests:
            memory: "100Mi"
            cpu: 0.1
        ...
      - name: db-sidecar
        ...

apiVersion: apps/v1
kind: Deployment
metadata:
  name: api
  namespace: go-demo-5
spec:
  ...
  template:
    ...
    spec:
      containers:
      - name: api
        ...
        resources:
          limits:
            memory: 15Mi
            cpu: 0.1
          requests:
            memory: 10Mi
            cpu: 0.01
...
```

We have two Pods that form an application. The `api` Deployment is a backend API that uses `db` StatefulSet for its state.

The essential parts of the definition are `resources`. Both the `api` and the `db` have `requests` and `limits` defined for memory and CPU. The database uses a sidecar container that will join MongoDB replicas into a replica set. Please note that, unlike other containers, the sidecar does not have `resources`. The importance behind that will be revealed later. For now, just remember that two containers have the `requests` and the `limits` defined and that one doesn’t.



``` shell
[root@10-23-75-240 k8s-specs]# kubectl apply \
>     -f scaling/go-demo-5-no-sidecar-mem.yml \
>
namespace/go-demo-5 unchanged
ingress.networking.k8s.io/api created
serviceaccount/db unchanged
role.rbac.authorization.k8s.io/db unchanged
rolebinding.rbac.authorization.k8s.io/db unchanged
statefulset.apps/db created
service/db created
deployment.apps/api created
service/api created

```



``` shell
│ Tolerations:                 node.kubernetes.io/not-ready:NoExecute op=Exists for 300s
│                              node.kubernetes.io/unreachable:NoExecute op=Exists for 300s
│ Events:
│   Type     Reason            Age   From               Message
│   ----     ------            ----  ----               -------
│   Warning  FailedScheduling  37s   default-scheduler  0/3 nodes are available: 3 pod has unbound immediate PersistentVolumeClaims.
```



