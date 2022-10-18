``` shell
 I0312 07:22:10.834263       1 log_file.go:99] Set log file max size to 104857600                                                                                                                 
│ I0312 07:22:10.836055       1 agent.go:85] Starting Antrea agent (version v1.5.0)                                                                                                                
│ I0312 07:22:10.836082       1 client.go:96] No kubeconfig file was specified. Falling back to in-cluster config                                                                                  
│ I0312 07:22:10.837795       1 client.go:96] No kubeconfig file was specified. Falling back to in-cluster config                                                                                  
│ I0312 07:22:10.838535       1 prometheus.go:171] Initializing prometheus metrics                                                                                                                 
│ I0312 07:22:10.840392       1 ovs_client.go:67] Connecting to OVSDB at address /var/run/openvswitch/db.sock                                                                                      
│ I0312 07:22:10.843274       1 agent.go:338] Setting up node network                                                                                                                              
│ I0312 07:22:10.853277       1 agent.go:873] Setting Node MTU=1402                                                                                                                                
│ E0312 07:22:10.853290       1 agent.go:907] Spec.PodCIDR is empty for Node 10-23-184-141. Please make sure --allocate-node-cidrs is enabled for kube-controller-manager and --cluster-cidr spec  
│ F0312 07:22:10.853819       1 main.go:58] Error running agent: error initializing agent: CIDR string is empty for node 10-23-184-141 
```



looks like the node is not getting an ip address  

``` shell
error initializing agent: CIDR string is empty for`node 10-23-184-141
```

Lets look at the kube-controller logs as well

``` shell
│ E0312 07:16:53.993207       1 controller_utils.go:260] Error while processing Node Add/Delete: failed to allocate cidr from cluster cidr at idx:0: CIDR allocation failed; there are no remainin     
│ g CIDRs left to allocate in the accepted range                                                                                                                                                       
│ I0312 07:16:53.993408       1 event.go:294] "Event occurred" object="10-23-184-141" kind="Node" apiVersion="v1" type="Normal" reason="CIDRNotAvailable" message="Node 10-23-184-141 status is no     
│ w: CIDRNotAvailable"                                                                                                                                                                                 
│ E0312 07:17:03.578491       1 controller_utils.go:260] Error while processing Node Add/Delete: failed to allocate cidr from cluster cidr at idx:0: CIDR allocation failed; there are no remainin     
│ g CIDRs left to allocate in the accepted range     
```

 

So why is the kube-controller complaining that there are no CIDRs remaining?



192.168.10.1

0000 0000  0000 0000  0000 0000   0000 0000







I restart it with the 

``` shell
kubeadm init --pod-network-cidr=10.23.0.0/10
```



 and it worked!





**Unable to connect to the server: x509: certificate signed by unknown authority (possibly because of "crypto/rsa: verification error" while trying to verify candidate authority certificate "kubernetes")**

``` shell'
[root@10-23-75-240 ~]# kubectl get pods

Unable to connect to the server: x509: certificate signed by unknown authority (possibly because of "crypto/rsa: verification error" while trying to verify candidate authority certificate "kubernetes")

[root@10-23-75-240 ~]# ls

Desktop Downloads install-release.sh k8s k9s thinclient_drives

[root@10-23-75-240 ~]# export KUBECONFIG=/etc/kubernetes/admin.conf

[root@10-23-75-240 ~]# kubectl get pods

NAME            READY  STATUS  RESTARTS  AGE

my-nginx-59d54b77c8-lshfr  1/1   Running  0     16m

my-nginx-59d54b77c8-pgt5b  1/1   Running  0     16m
```



 



## cannot  communicate with pod through service in master node

``` shell
[root@10-23-75-240 k8s]# kubectl get svc
NAME            TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)           AGE
go-demo-2-svc   NodePort    10.103.225.77    <none>        28017:30286/TCP   4m25s
kubernetes      ClusterIP   10.96.0.1        <none>        443/TCP           30m
my-nginx        ClusterIP   10.108.193.250   <none>        8080/TCP          64s
[root@10-23-75-240 k8s]# curl 10.108.193.250 -v
* About to connect() to proxy 127.0.0.1 port 8001 (#0)
*   Trying 127.0.0.1...
* Connected to 127.0.0.1 (127.0.0.1) port 8001 (#0)
> GET HTTP://10.108.193.250/ HTTP/1.1
> User-Agent: curl/7.29.0
> Host: 10.108.193.250
> Accept: */*
> Proxy-Connection: Keep-Alive
> 
< HTTP/1.1 503 Service Unavailable
< Connection: close
* HTTP/1.1 proxy connection set close!
< Proxy-Connection: close
< Content-Length: 0
< 
* Closing connection 0


[root@10-23-75-240 k8s]# curl 10.108.193.250 -v
* About to connect() to 10.108.193.250 port 80 (#0)
*   Trying 10.108.193.250...
```



change the yaml



``` shell
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-nginx
spec:
  selector:
    matchLabels:
      run: my-nginx
  replicas: 2
  template:
    metadata:
      labels:
        run: my-nginx
    spec:
      containers:
      - name: my-nginx
        image: nginx:latest
        ports:
        - containerPort: 80
          hostPort: 80
        readinessProbe:
          failureThreshold: 3
          httpGet:
            port: 80
            scheme: HTTP
          periodSeconds: 10
          successThreshold: 1
          timeoutSeconds: 1

---
apiVersion: v1
kind: Service
metadata:
  name: my-nginx
  labels:
    run: my-nginx
spec:
  ports:
  - port: 8080
    protocol: TCP
  selector:
    run: my-nginx
```



working now

``` shell
/docker-entrypoint.sh: Looking for shell scripts in /docker-entrypoint.d/                                                                                                                         
│ /docker-entrypoint.sh: Launching /docker-entrypoint.d/10-listen-on-ipv6-by-default.sh                                                                                                             
│ 10-listen-on-ipv6-by-default.sh: info: Getting the checksum of /etc/nginx/conf.d/default.conf                                                                                                     
│ 10-listen-on-ipv6-by-default.sh: info: Enabled listen on IPv6 in /etc/nginx/conf.d/default.conf                                                                                                   
│ /docker-entrypoint.sh: Launching /docker-entrypoint.d/20-envsubst-on-templates.sh                                                                                                                 
│ /docker-entrypoint.sh: Launching /docker-entrypoint.d/30-tune-worker-processes.sh                                                                                                                 
│ /docker-entrypoint.sh: Configuration complete; ready for start up                                                                                                                                 
│ 2022/03/13 03:31:23 [notice] 1#1: using the "epoll" event method                                                                                                                                  
│ 2022/03/13 03:31:23 [notice] 1#1: nginx/1.21.6                                                                                                                                                    
│ 2022/03/13 03:31:23 [notice] 1#1: built by gcc 10.2.1 20210110 (Debian 10.2.1-6)                                                                                                                  
│ 2022/03/13 03:31:23 [notice] 1#1: OS: Linux 4.19.188-10.el7.ucloud.x86_64                                                                                                                         
│ 2022/03/13 03:31:23 [notice] 1#1: getrlimit(RLIMIT_NOFILE): 1048576:1048576                                                                                                                       
│ 2022/03/13 03:31:23 [notice] 1#1: start worker processes                                                                                                                                          
│ 2022/03/13 03:31:23 [notice] 1#1: start worker process 30                                                                                                                                         
│ 2022/03/13 03:31:23 [notice] 1#1: start worker process 31                                                                                                                                         
│ 10.16.1.1 - - [13/Mar/2022:03:31:23 +0000] "GET / HTTP/1.1" 200 615 "-" "kube-probe/1.23" "-"                                                                                                     
│ 10.16.1.1 - - [13/Mar/2022:03:31:27 +0000] "GET / HTTP/1.1" 200 615 "-" "kube-probe/1.23" "-"                                                                                                     
│ 10.16.1.1 - - [13/Mar/2022:03:31:37 +0000] "GET / HTTP/1.1" 200 615 "-" "kube-probe/1.23" "-"                                                                                                     
│                                                                                                   
```



**main reason , http_proxy**





``` shell
[root@10-23-75-240 k8s]# curl 10.16.1.8:8080

curl: (7) Failed connect to 10.16.1.8:8080; 拒绝连接
```

change service to node port

``` shell
---
apiVersion: v1
kind: Service
metadata:
  name: my-nginx
  labels:
    run: my-nginx
spec:
  type: NodePort
  ports:
  - port: 8080
    protocol: TCP
  selector:
    run: my-nginx
```

 

``` shell
[root@10-23-75-240 k8s]# kubectl get svc
NAME            TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)           AGE
go-demo-2-svc   NodePort    10.103.225.77    <none>        28017:30286/TCP   16m
kubernetes      ClusterIP   10.96.0.1        <none>        443/TCP           41m
my-nginx        NodePort    10.108.193.250   <none>        8080:31744/TCP    12m
[root@10-23-75-240 k8s]# curl 10.108.193.250:8080 -v
* About to connect() to 10.108.193.250 port 8080 (#0)
*   Trying 10.108.193.250...
* 拒绝连接
* Failed connect to 10.108.193.250:8080; 拒绝连接
* Closing connection 0
curl: (7) Failed connect to 10.108.193.250:8080; 拒绝连接
```



**set to the correct port here**

``` yaml
[root@10-23-75-240 k8s]# vi run-my-nginx.yaml 

kind: Deployment
metadata:
  name: my-nginx
spec:
  selector:
    matchLabels:
      run: my-nginx
  replicas: 2
  template:
    metadata:
      labels:
        run: my-nginx
    spec:
      containers:
      - name: my-nginx
        image: nginx:latest
        ports:
        - containerPort: 80
          hostPort: 80
        readinessProbe:
          failureThreshold: 3
          httpGet:
            port: 80
            scheme: HTTP
          periodSeconds: 10
          successThreshold: 1
          timeoutSeconds: 1

---
apiVersion: v1
kind: Service
metadata:
  name: my-nginx
  labels:
    run: my-nginx
spec:
  type: NodePort
  ports:
  - port: 80
    nodePort: 30080
    protocol: TCP
  selector:
    run: my-nginx
```





Fixed here



```shell
[root@10-23-75-240 k8s]# PORT=$(kubectl get svc my-nginx \
    -o jsonpath="{.spec.ports[0].nodePort}")
[root@10-23-75-240 k8s]# curl 10.23.184.141:$PORT
<!DOCTYPE html>
<html>
<head>
<title>Welcome to nginx!</title>
<style>
html { color-scheme: light dark; }
body { width: 35em; margin: 0 auto;
font-family: Tahoma, Verdana, Arial, sans-serif; }
</style>
</head>
<body>
<h1>Welcome to nginx!</h1>
<p>If you see this page, the nginx web server is successfully installed and
working. Further configuration is required.</p>

<p>For online documentation and support please refer to
<a href="http://nginx.org/">nginx.org</a>.<br/>
Commercial support is available at
<a href="http://nginx.com/">nginx.com</a>.</p>

<p><em>Thank you for using nginx.</em></p>
</body>
</html>
```









## antrea node is not running

After a while, redeploying using v.1.23.5 version



1. kubelet cannot find chi.sock

``` shell
4月 01 22:10:24 10-23-74-207 kubelet[10084]: E0401 22:10:24.140537   10084 pod_workers.go:949] "Error syncing pod, skipping" err="failed to \"CreatePodSandbox\" for \"coredns-64897985d-68h96_kube-system(3805cfe0-c98a-4ce3-bccd-1f5aba24f400)\" with CreatePodSandboxError: \"Failed to create sandbox for pod \\\"coredns-64897985d-68h96_kube-system(3805cfe0-c98a-4ce3-bccd-1f5aba24f400)\\\": rpc error: code = Unknown desc = failed to set up sandbox container \\\"b9703240ec37437c0cf7b1c5151be80b0a3728671fc04c85f3249d331e370999\\\" network for pod \\\"coredns-64897985d-68h96\\\": networkPlugin cni failed to set up pod \\\"coredns-64897985d-68h96_kube-system\\\" network: rpc error: code = Unavailable desc = connection error: desc = \\\"transport: Error while dialing dial unix /var/run/antrea/cni.sock: connect: no such file or directory\\\"\"" pod="kube-system/coredns-64897985d-68h96" podUID=3805cfe0-c98a-4ce3-bccd-1f5aba24f400
[root@10-23-74-207 ~]# kubectl get pods -A

```



Try different version of antrea.

origin version is v1.5.1

``` shell

kubectl apply -f https://github.com/antrea-io/antrea/releases/download/v1.5.2/antrea.yml
```







### Kube init failed unknown service runtime.v1alpha2.RuntimeService

https://github.com/kubernetes-sigs/cri-tools/issues/710

