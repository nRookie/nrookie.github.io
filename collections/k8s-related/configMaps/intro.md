## The ConfigMap



ConfigMap allows us to “inject” configuration into containers. The source of the configs can be files, directories, or literal values. The destination can be files or environment variables.

> **ConfigMap** takes a configuration from a source and mounts it into running containers as a *volume*.

That’s all the theory you’ll get up-front. Instead of a lengthy explanation, we’ll run some examples, and comment on the features we experience. We’ll be learning by doing, instead of learning by memorizing theory.

Let’s prepare the cluster and see ConfigMaps in action.



## Creating a ConfigMap[#](https://www.educative.io/module/lesson/a-practical-guide-to-kubernetes/JYRnmn7xom2#Creating-a-ConfigMap)

In its purest, and probably the most common form, a ConfigMap takes a single file. For example, we can create one from the `cm/prometheus-conf.yml` file.



``` shell
kubectl create cm my-config \
    --from-file=cm/prometheus-conf.yml
```





``` yaml
[root@10-23-75-240 k8s-specs]# kubectl describe cm my-config
Name:         my-config
Namespace:    default
Labels:       <none>
Annotations:  <none>

Data
====
prometheus-conf.yml:
----
global:
  scrape_interval:     15s

scrape_configs:
  - job_name: prometheus
    metrics_path: /prometheus/metrics
    static_configs:
      - targets:
        - localhost:9090


BinaryData
====

Events:  <none>
```

The important part is located below `Data`. We can see the key which, in this case, is the name of the file (`prometheus-conf.yml`). Further down you can see the content of the file. If you execute `cat cm/prometheus-conf.yml`, you’ll see that it is the same as what we saw from the ConfigMap’s description.



## Mounting the ConfigMap[#](https://www.educative.io/module/lesson/a-practical-guide-to-kubernetes/JYRnmn7xom2#Mounting-the-ConfigMap)

ConfigMap is useless by itself. It is yet another Volume which, like all the others, needs a mount



### Pod with Mounted ConfigMap[#](https://www.educative.io/module/lesson/a-practical-guide-to-kubernetes/JYRnmn7xom2#Pod-with-Mounted-ConfigMap)



``` shell
[root@10-23-75-240 k8s-specs]# cat cm/alpine.yml
apiVersion: v1
kind: Pod
metadata:
  name: alpine
spec:
  containers:
  - name: alpine
    image: alpine
    command: ["sleep"]
    args: ["100000"]
    volumeMounts:
    - name: config-vol
      mountPath: /etc/config
  volumes:
  - name: config-vol
    configMap:
      name: my-config
```



The essential sections are `volumeMounts` and `volumes`. Since `volumeMounts` are the same no matter the type of the Volume, there’s nothing special about it. We defined that it should be based on the volume called `config-vol` and that it should mount the path `/etc/config`. The `volumes` section uses `configMap` as the type and, in this case, has a single item `name`, that coincides with the name of the ConfigMap we created earlier.

### Creating the Pod[#](https://www.educative.io/module/lesson/a-practical-guide-to-kubernetes/JYRnmn7xom2#Creating-the-Pod)

Let’s create the Pod and see what happens.

``` shell
kubectl create -f cm/alpine.yml
kubectl get pods
```





Please confirm that the Pod is indeed running before moving on.

``` shell
kubectl exec -it alpine -- \
    ls /etc/config
    
[root@10-23-75-240 k8s-specs]# kubectl exec -it alpine -- \
>     ls /etc/config
prometheus-conf.yml
```



``` shell
[root@10-23-75-240 k8s-specs]# kubectl exec -it alpine --     ls -l /etc/config
total 0
lrwxrwxrwx    1 root     root            26 Mar 16 02:51 prometheus-conf.yml -> ..data/prometheus-conf.ym
```



You’ll see that `prometheus-conf.yml` is a link to `..data/prometheus-conf.yml`.

If you dig deeper, you’ll see that `..data` is also a link to the directory named from a timestamp. And so on, and so forth. For now, the exact logic behind all the links and the actual files is not of great importance. From the functional point of view, there is `prometheus-conf.yml`, and our application can do whatever it needs to do with it.

Let’s confirm that the content of the file inside the container is indeed the same as the source file we used to create the ConfigMap.

``` shell
[root@10-23-75-240 k8s-specs]# kubectl exec -it alpine -- \
>     cat /etc/config/prometheus-conf.yml
global:
  scrape_interval:     15s

scrape_configs:
  - job_name: prometheus
    metrics_path: /prometheus/metrics
    static_configs:
      - targets:
        - localhost:9090
```



The **output** should be the same as the contents of the `cm/prometheus-conf.yml` file.

We saw one combination of ConfigMap. Let’s see what else we can do with it.



## Deleting the Objects[#](https://www.educative.io/module/lesson/a-practical-guide-to-kubernetes/JYRnmn7xom2#Deleting-the-Objects)



``` shell
[root@10-23-75-240 k8s-specs]# kubectl delete -f cm/alpine.yml
pod "alpine" deleted
kubectl delete cm my-config
[root@10-23-75-240 k8s-specs]# kubectl delete cm my-config
configmap "my-config" deleted
```

