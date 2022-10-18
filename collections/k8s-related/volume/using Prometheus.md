``` yaml
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: prometheus
  annotations:
    kubernetes.io/ingress.class: "nginx"
    ingress.kubernetes.io/ssl-redirect: "false"
    nginx.ingress.kubernetes.io/ssl-redirect: "false"
spec:
  rules:
  - http:
      paths:
      - path: /prometheus
        backend:
          serviceName: prometheus
          servicePort: 9090

---

apiVersion: apps/v1
kind: Deployment
metadata:
  name: prometheus
spec:
  selector:
    matchLabels:
      type: monitor
      service: prometheus
  strategy:
    type: Recreate
  template:
    metadata:
      labels:
        type: monitor
        service: prometheus
    spec:
      containers:
      - name: prometheus
        image: prom/prometheus:v2.0.0
        command:
        - /bin/prometheus
        args:
        - "--config.file=/etc/prometheus/prometheus.yml"
        - "--storage.tsdb.path=/prometheus"
        - "--web.console.libraries=/usr/share"
        - "--web.external-url=http://192.168.99.100/prometheus"

---

apiVersion: v1
kind: Service
metadata:
  name: prometheus
spec:
  ports:
  - port: 9090
  selector:
    type: monitor
    service: prometheus
```





``` shell
cat volume/prometheus.yml | sed -e \
    "s/192.168.99.100/$IP/g" \
    | kubectl create -f - \
    --record --save-config
```



📝 Please note that, this time, the `create` command has dash (`-`) instead of the path to the file. That’s an indication that `stdin` should be used instead.



``` shell
docker add the proxy could pull the image now.
```



``` html
Configuration
global:
  scrape_interval: 15s
  scrape_timeout: 10s
  evaluation_interval: 15s
alerting:
  alertmanagers:
  - static_configs:
    - targets: []
    scheme: http
    timeout: 10s
scrape_configs:
- job_name: prometheus
  scrape_interval: 15s
  scrape_timeout: 10s
  metrics_path: /metrics
  scheme: http
  static_configs:
  - targets:
    - localhost:9090
```

The problem is with the `metrics_path` field. By default, it is set to `/metrics`. However, since we changed the base path to `/prometheus`, the field should have `/prometheus/metrics` as the value.

### Changing Prometheus Configuration



Long story short, we must change the Prometheus configuration.

We could, for example, enter the container, update the configuration file, and send the reload request to Prometheus. That would be a terrible solution since it would last only until the next time we update the application, or until the container fails, and Kubernetes decides to reschedule it.

Let’s explore alternative solutions. We could, for example, use `hostPath` Volume for this as well. If we can guarantee that the correct configuration file is inside the VM, the Pod could attach it to the `prometheus` container. Let’s try it out.



We could, for example, enter the container, update the configuration file, and send the reload request to Prometheus. That would be a terrible solution since it would last only until the next time we update the application, or until the container fails, and Kubernetes decides to reschedule it.



Let’s explore alternative solutions. We could, for example, use `hostPath` Volume for this as well. If we can guarantee that the correct configuration file is inside the VM, the Pod could attach it to the `prometheus` container. Let’s try it out.





``` shell
cat volume/prometheus-host-path.yml
[root@10-23-75-240 k8s-specs]# cat volume/prometheus-host-path.yml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: prometheus
  annotations:
    kubernetes.io/ingress.class: "nginx"
    ingress.kubernetes.io/ssl-redirect: "false"
    nginx.ingress.kubernetes.io/ssl-redirect: "false"
spec:
  rules:
  - http:
      paths:
      - path: /prometheus
        pathType: ImplementationSpecific
        backend:
          service:
            name: prometheus
            port:
              number: 9090

---

apiVersion: apps/v1
kind: Deployment
metadata:
  name: prometheus
spec:
  selector:
    matchLabels:
      type: monitor
      service: prometheus
  strategy:
    type: Recreate
  template:
    metadata:
      labels:
        type: monitor
        service: prometheus
    spec:
      containers:
      - name: prometheus
        image: prom/prometheus:v2.0.0
        command:
        - /bin/prometheus
        args:
        - "--config.file=/etc/prometheus/prometheus.yml"
        - "--storage.tsdb.path=/prometheus"
        - "--web.console.libraries=/usr/share"
        - "--web.external-url=http://192.168.99.100/prometheus"
        volumeMounts:
        - mountPath: /etc/prometheus/prometheus.yml
          name: prom-conf
      volumes:
      - name: prom-conf
        hostPath:
          path: /files/prometheus-conf.yml
          type: File

---

apiVersion: v1
kind: Service
metadata:
  name: prometheus
spec:
  ports:
  - port: 9090
  selector:
    type: monitor
    service: prometheus

```



The only significant difference, when compared with the previous definition, is in the added `volumeMounts` and `volumes`fields. We’re using the same schema as before, except that, this time, the `type` is set to `File`. Once we `apply` this Deployment, the file `/files/prometheus-conf.yml` on the host will be available as `/etc/prometheus/prometheus.yml` inside the container.

If you recall, we copied one file to the `~/.minikube/files` directory, and Minikube copied it to the `/files` directory inside the VM.





``` shell
[root@10-23-75-240 k8s-specs]# cp volume/prometheus-conf.yml  /files

[root@10-23-75-240 k8s-specs]# scp volume/prometheus-conf.yml root@106.75.225.217:/files/
prometheus-conf.yml                                                                                                                                               100%  177    93.5KB/s   00:00
[root@10-23-75-240 k8s-specs]# scp volume/prometheus-conf.yml root@106.75.245.207:/files/
root@106.75.245.207's password:
prometheus-conf.yml
```





``` shell
cat volume/prometheus-host-path.yml \
    | sed -e \
    "s/192.168.99.100/$IP/g" \
    | kubectl apply -f -

kubectl rollout status deploy prometheus
```





```
google-chrome --no-sandbox --user-data-dir
```



![image-20220316101706937](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/volume/image-20220316101706937.png)



A `hostPath` Volume maps a directory from a host to where the Pod is running. Using it to “inject” configuration files into containers would mean that we’d have to make sure that the file is present on every node of the cluster.



Working with Minikube can be potentially misleading. The fact that we’re running a single-node cluster means that every Pod we run will be scheduled on one node. Copying a configuration file to that single node, as we did in our example, ensures that it can be mounted in any Pod. However, the moment we add more nodes to the cluster, we’d experience side effects. We’d need to make sure that each node in our cluster has the same file we wish to mount, as we would not be able to predict where individual Pods would be scheduled. This would introduce far too much unnecessary work and added complexity.

## Exploring the Solutions[#](https://www.educative.io/module/lesson/a-practical-guide-to-kubernetes/JY1mmlMA3Og#Exploring-the-Solutions)



An alternative solution would be to mount an NFS drive to all the nodes and store the file there. That would provide the guarantee that the file will be available on all the nodes, as long as we do *NOT* forget to mount NFS on each.



Another solution could be to create a custom Prometheus image. It could be based on the official image, with a single `COPY` instruction that would add the configuration. The advantage of that solution is that the image would be entirely immutable. Its state would not be polluted with unnecessary Volume mounts. Anyone could run that image and expect the same result. That is my preferred solution. However, in some cases, you might want to deploy the same application with a slightly different configuration. Should we, in those cases, fall back to mounting an NFS drive on each node and continue using `hostPath`?

Even though mounting an NFS drive would solve some of the problems, it is still not a great solution. In order to mount a file from NFS, we need to use the [nfs](https://kubernetes.io/docs/concepts/storage/volumes/#nfs) Volume type instead of `hostPath`. Even then it would be a sub-optimal solution. A much better approach would be to use `configMap`. We’ll explore it in the next chapter.





``` txt
Do use hostPath to mount host resources like /var/run/docker.sock and /dev/cgroups. Do not use it to inject configuration files or store the state of an application.
```

destroying the pod

``` shell
kubectl delete \
    -f volume/prometheus-host-path.yml
```





![image-20220321091920236](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/volume/image-20220321091920236.png)





