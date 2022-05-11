# Auto-scale based on resource usage [#](https://www.educative.io/module/lesson/kubernetes-monitoring-logging-auto-scaling/B8WwZX39Y0N#Auto-scale-based-on-resource-usage-)

So far, the `HPA` has not yet performed auto-scaling based on resource usage. Let’s do that now. First, we’ll try to create another `HorizontalPodAutoscaler` but, this time, we’ll target the StatefulSet that runs our MongoDB. So, let’s take a look at yet another YAML definition.



## Create `HPA` [#](https://www.educative.io/module/lesson/kubernetes-monitoring-logging-auto-scaling/B8WwZX39Y0N#Create-HPA-)

```shell
cat scaling/go-demo-5-db-hpa.yml
```

The **output** is as follows.

``` yaml
apiVersion: autoscaling/v2beta1
kind: HorizontalPodAutoscaler
metadata:
  name: db
  namespace: go-demo-5
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: StatefulSet
    name: db
  minReplicas: 3
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





That definition is almost the same as the one we used before. The only difference is that this time we’re targeting `StatefulSet` called `db` and that the minimum number of replicas should be `3`.

Let’s apply it.



``` shell
kubectl apply \
    -f scaling/go-demo-5-db-hpa.yml \
    --record
```



``` shell
[root@10-23-75-240 k8s-specs]# kubectl -n go-demo-5 get hpa
NAME   REFERENCE        TARGETS                        MINPODS   MAXPODS   REPLICAS   AGE
api    Deployment/api   40%/80%, 10%/80%               2         5         2          14m
db     StatefulSet/db   <unknown>/80%, <unknown>/80%   3         5         1          24s
```



### Resource utilization not getting shown [#](https://www.educative.io/module/lesson/kubernetes-monitoring-logging-auto-scaling/B8WwZX39Y0N#Resource-utilization-not-getting-shown-)

There might be something wrong since the resource utilization continued being unknown. Let’s describe the newly created `HPA` and see whether we’ll be able to find the cause behind the issue.



``` shell
[root@10-23-75-240 k8s-specs]# kubectl -n go-demo-5 describe hpa db
Warning: autoscaling/v2beta2 HorizontalPodAutoscaler is deprecated in v1.23+, unavailable in v1.26+; use autoscaling/v2 HorizontalPodAutoscaler
Name:                                                     db
Namespace:                                                go-demo-5
Labels:                                                   <none>
Annotations:                                              kubernetes.io/change-cause: kubectl apply --filename=scaling/go-demo-5-db-hpa.yml --record=true
CreationTimestamp:                                        Mon, 21 Mar 2022 11:06:11 +0800
Reference:                                                StatefulSet/db
Metrics:                                                  ( current / target )
  resource memory on pods  (as a percentage of request):  <unknown> / 80%
  resource cpu on pods  (as a percentage of request):     <unknown> / 80%
Min replicas:                                             3
Max replicas:                                             5
StatefulSet pods:                                         3 current / 3 desired
Conditions:
  Type           Status  Reason                   Message
  ----           ------  ------                   -------
  AbleToScale    True    SucceededGetScale        the HPA controller was able to get the target's current scale
  ScalingActive  False   FailedGetResourceMetric  the HPA was unable to compute the replica count: failed to get memory utilization: missing request for memory
Events:
  Type     Reason                        Age                From                       Message
  ----     ------                        ----               ----                       -------
  Normal   SuccessfulRescale             44s                horizontal-pod-autoscaler  New size: 3; reason: Current number of replicas below Spec.MinReplicas
  Warning  FailedGetResourceMetric       14s (x2 over 29s)  horizontal-pod-autoscaler  failed to get memory utilization: missing request for memory
  Warning  FailedGetResourceMetric       14s (x2 over 29s)  horizontal-pod-autoscaler  failed to get cpu utilization: missing request for cpu
  Warning  FailedComputeMetricsReplicas  14s (x2 over 29s)  horizontal-pod-autoscaler  invalid metrics (2 invalid out of 2), first error is: failed to get memory utilization: missing request for memory
```





Please note that your **output** could have only one event or even none of those. If that’s the case, please wait for a few minutes and repeat the previous command.

If we focus on the first message, we can see that it started well. `HPA` detected that the current number of replicas is below the limit and increased them to three. That is the expected behavior, so let’s move to the other two messages.

`HPA` could not calculate the percentage because we did not specify how much memory we are requesting for the `db-sidecar` container. Without `requests`, `HPA` cannot calculate the percentage of the actual memory usage. In other words, we missed specifying resources for the `db-sidecar` container and `HPA` could not do its work. We’ll fix that by applying `go-demo-5-no-hpa.yml`.



## Create `HPA` with new definition [#](https://www.educative.io/module/lesson/kubernetes-monitoring-logging-auto-scaling/B8WwZX39Y0N#Create-HPA-with-new-definition-)



Let’s take a quick look at the new definition.



```shell
cat scaling/go-demo-5-no-hpa.yml
```



``` yaml
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
      - name: db-sidecar
        ...
        resources:
          limits:
            memory: "100Mi"
            cpu: 0.2
          requests:
            memory: "50Mi"
            cpu: 0.1
...
```



The only noticeable difference, when compared with the initial definition, is that this time we defined the resources for the `db-sidecar` container. Let’s apply it.



> Please change the volume template if necessary

``` shell

kubectl apply \
    -f scaling/go-demo-5-no-hpa.yml \
    --record
```

``` shell
kubectl -n go-demo-5 get hpa
```



``` shell
db-sidecar     at Socket.<anonymous> (/opt/cvallance/mongo-k8s-sidecar/node_modules/mongodb-core/lib/connection/connection.js:189:49)                             
│ db-sidecar     at Object.onceWrapper (events.js:273:13)                                                                                                           
│ db-sidecar     at Socket.emit (events.js:182:13)                                                                                                                  
│ db-sidecar     at emitErrorNT (internal/streams/destroy.js:82:8)                                                                                                  
│ db-sidecar     at emitErrorAndCloseNT (internal/streams/destroy.js:50:3)                                                                                          
│ db-sidecar   name: 'MongoError',                                                                                                                                  
│ db-sidecar   message:                                                                                                                                             
│ db-sidecar    'failed to connect to server [127.0.0.1:27017] on first connect [MongoError: connect ECONNREFUSED 127.0.0.1:27017]' }    
```

https://www.google.com/search?q=pod+in+container+127.0.0.1%3A27017&sxsrf=APq-WBtXKHMd2MSBhxpI-AiIQNkz3qnT1A%3A1647833028791&ei=xO83YvHqL6TQkPIP06y9yA0&ved=0ahUKEwjx3NGyoNb2AhUkKEQIHVNWD9kQ4dUDCA4&uact=5&oq=pod+in+container+127.0.0.1%3A27017&gs_lcp=Cgdnd3Mtd2l6EAMyBggAEBYQHjIICAAQFhAKEB4yBggAEBYQHjIJCAAQyQMQFhAeMgYIABAWEB4yBggAEBYQHjIGCAAQFhAeMgYIABAWEB4yBggAEBYQHjIGCAAQFhAeOgcIABBHELADOgQIIxAnOg4ILhCABBCxAxDHARCjAjoOCC4QgAQQsQMQxwEQ0QM6CwguEIAEELEDEIMBOggIABCABBCxAzoLCC4QsQMQgwEQ1AI6CwgAEIAEELEDEIMBOgUIABCRAjoECAAQQzoFCAAQgAQ6BwgAELEDEEM6BwgAEIAEEApKBAhBGABKBAhGGABQngZYvBVggRdoAXABeAGAAcUCiAGrIJIBBjItMTYuMZgBAKABAaABAsgBCMABAQ&sclient=gws-wiz



``` shell
i was able to resolve the issue after disable IPv6 for MongoDB POD:
you can edit the mongoDB deployment in kubeapps and add this values:

name: MONGODB_ENABLE_IPV6
value: “no”
image: docker.io/bitnami/mongodb:4.0.1-debian-9-r12
```



``` shell
    spec:
      serviceAccountName: db
      terminationGracePeriodSeconds: 10
      containers:
      - name: db
        image: mongo:3.3
        env:
        - name: MONGODB_ENABLE_IPV6
          value: "no"
```

not working.

