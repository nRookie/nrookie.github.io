``` shell
kubectl top -h
Display Resource (CPU/Memory) usage.

 The top command allows you to see the resource consumption for nodes or pods.

 This command requires Metrics Server to be correctly configured and working on the server.

Available Commands:
  node        Display resource (CPU/memory) usage of nodes
  pod         Display resource (CPU/memory) usage of pods

Usage:
  kubectl top [flags] [options]

Use "kubectl <command> --help" for more information about a given command.
Use "kubectl options" for a list of global command-line options (applies to all commands).
```





but when we type

``` shell
[root@10-23-75-240 k8s-specs]# kubectl top pod
error: Metrics API not available
```



## Install metrics server 



https://github.com/kubernetes-sigs/metrics-server



``` shell
[root@10-23-75-240 k8s-specs]# kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml
Unable to connect to the server: proxyconnect tcp: EOF
```

https://docs.mattermost.com/configure/using-outbound-proxy.html



didn't figure out a better solution, unset the https_proxy to install it.

``` shell
 unset https_proxy 
```



``` shell
kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml
serviceaccount/metrics-server created
clusterrole.rbac.authorization.k8s.io/system:aggregated-metrics-reader created
clusterrole.rbac.authorization.k8s.io/system:metrics-server created
rolebinding.rbac.authorization.k8s.io/metrics-server-auth-reader created
clusterrolebinding.rbac.authorization.k8s.io/metrics-server:system:auth-delegator created
clusterrolebinding.rbac.authorization.k8s.io/system:metrics-server created
service/metrics-server created
deployment.apps/metrics-server created
apiservice.apiregistration.k8s.io/v1beta1.metrics.k8s.io created

```





``` shell
[root@10-23-75-240 k8s]# docker pull nginx
Using default tag: latest
Error response from daemon: Get "https://registry-1.docker.io/v2/": proxyconnect tcp: EOF

Error response from daemon: Get "https://registry-1.docker.io/v2/": proxyconnect tcp: EOF
```

https://stackoverflow.com/questions/64137423/docker-error-response-from-daemon-get-https-registry-1-docker-io-v2-proxyc

#### should use http in docker https://proxy

``` shell
vim /etc/systemd/system/docker.service.d/http-proxy.conf
systemctl daemon-reload 
systemctl restart docker

[root@10-23-75-240 k8s]# docker pull nginx
Using default tag: latest
latest: Pulling from library/nginx
Digest: sha256:1c13bc6de5dfca749c377974146ac05256791ca2fe1979fc8e8278bf0121d285
Status: Image is up to date for nginx:latest
docker.io/library/nginx:latest
```





#### not running



``` shell
[root@10-23-75-240 k8s]# kubectl -n kube-system get pods
NAME                                   READY   STATUS    RESTARTS         AGE
antrea-agent-4b7bf                     2/2     Running   4 (3m37s ago)    5h31m
antrea-agent-gpfxb                     2/2     Running   12 (2m52s ago)   5h34m
antrea-agent-x8pwg                     2/2     Running   4 (3m32s ago)    5h31m
antrea-controller-7974576775-8cc48     1/1     Running   9 (2m36s ago)    5h34m
coredns-64897985d-9z46t                1/1     Running   4 (3m11s ago)    5h34m
coredns-64897985d-pjvpp                1/1     Running   4 (3m11s ago)    5h34m
etcd-10-23-75-240                      1/1     Running   12 (3m5s ago)    5h34m
kube-apiserver-10-23-75-240            1/1     Running   7 (2m43s ago)    5h34m
kube-controller-manager-10-23-75-240   1/1     Running   5 (3m16s ago)    5h34m
kube-proxy-6gk5l                       1/1     Running   2 (3m37s ago)    5h31m
kube-proxy-kjzpv                       1/1     Running   5 (3m16s ago)    5h34m
kube-proxy-ls4g6                       1/1     Running   2 (3m33s ago)    5h31m
kube-scheduler-10-23-75-240            1/1     Running   7 (3m16s ago)    5h34m
metrics-server-847dcc659d-8zbnt        0/1     Running   0                12m
```



``` shell
│ Events:                                                                                                                                                                                          │
│   Type     Reason     Age                   From               Message                                                                                                                           │
│   ----     ------     ----                  ----               -------                                                                                                                           │
│   Normal   Scheduled  15m                   default-scheduler  Successfully assigned kube-system/metrics-server-847dcc659d-8zbnt to 10-23-184-141                                                │
│   Normal   Pulling    14m (x4 over 15m)     kubelet            Pulling image "k8s.gcr.io/metrics-server/metrics-server:v0.6.1"                                                                   │
│   Warning  Failed     14m (x4 over 15m)     kubelet            Failed to pull image "k8s.gcr.io/metrics-server/metrics-server:v0.6.1": rpc error: code = Unknown desc = Error response from daem │
│ on: Get "https://k8s.gcr.io/v2/": proxyconnect tcp: EOF                                                                                                                                          │
│   Warning  Failed     14m (x4 over 15m)     kubelet            Error: ErrImagePull                                                                                                               │
│   Warning  Failed     14m (x6 over 15m)     kubelet            Error: ImagePullBackOff                                                                                                           │
│   Normal   BackOff    5m48s (x44 over 15m)  kubelet            Back-off pulling image "k8s.gcr.io/metrics-server/metrics-server:v0.6.1"                                                          │
│   Warning  Unhealthy  58s (x25 over 4m30s)  kubelet            Readiness probe failed: HTTP probe failed with statuscode: 500                                                                    │
│                                                                                                                                
```



The reason is the pod cannot communicate with the API server



``` shell
kubectl  -n kube-system edit deploy metrics-server

containers:
      - args:
        - --cert-dir=/tmp
        - --secure-port=4443
        - --kubelet-preferred-address-types=InternalIP,ExternalIP,Hostname
        - --kubelet-use-node-status-port
        - --metric-resolution=15s
        command:
        - /metrics-server
        - --kubelet-insecure-tls
        - --kubelet-preferred-address-types=InternalIP
```



``` shell
[root@10-23-75-240 k8s]# kubectl -n kube-system  get pods
NAME                                   READY   STATUS    RESTARTS       AGE
antrea-agent-4b7bf                     2/2     Running   4 (12m ago)    5h40m
antrea-agent-gpfxb                     2/2     Running   12 (11m ago)   5h43m
antrea-agent-x8pwg                     2/2     Running   4 (12m ago)    5h40m
antrea-controller-7974576775-8cc48     1/1     Running   9 (11m ago)    5h43m
coredns-64897985d-9z46t                1/1     Running   4 (12m ago)    5h43m
coredns-64897985d-pjvpp                1/1     Running   4 (12m ago)    5h43m
etcd-10-23-75-240                      1/1     Running   12 (12m ago)   5h43m
kube-apiserver-10-23-75-240            1/1     Running   7 (11m ago)    5h43m
kube-controller-manager-10-23-75-240   1/1     Running   5 (12m ago)    5h43m
kube-proxy-6gk5l                       1/1     Running   2 (12m ago)    5h40m
kube-proxy-kjzpv                       1/1     Running   5 (12m ago)    5h43m
kube-proxy-ls4g6                       1/1     Running   2 (12m ago)    5h40m
kube-scheduler-10-23-75-240            1/1     Running   7 (12m ago)    5h43m
metrics-server-77b7f4f884-qjltk        1/1     Running   0              69s
```





``` shell
[root@10-23-75-240 k8s]# kubectl top node
Error from server (ServiceUnavailable): the server is currently unable to handle the request (get nodes.metrics.k8s.io)
```



https://github.com/kubernetes-sigs/metrics-server/issues/448

https://kubernetes.io/docs/tasks/extend-kubernetes/configure-aggregation-layer/



``` shell
E0312 13:54:21.518287       1 available_controller.go:524] v1beta2.controlplane.antrea.tanzu.vmware.com failed with: failing or missing response from https://10.96.212.119:443/apis/controlplane.antrea.tanzu.vmware.com/v1beta2: Get "https://10.96.212.119:443/apis/controlplane.antrea.tanzu.vmware.com/v1beta2": EOF
```



``` shell
https://github.com/kubernetes-sigs/metrics-server/issues/744
```

tried to solve by removing the proxy.

did not worked.



## other solutions

https://www.linuxsysadmins.com/service-unavailable-kubernetes-metrics/





does not found a solution



