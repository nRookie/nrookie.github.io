# Env 

Centos 7.9

### Preparing the required container images

#### setup the proxy for image pull

``` shell
sudo mkdir -p /etc/systemd/system/docker.service.d

vim /etc/systemd/system/docker.service.d/http-proxy.conf

[Service]
Environment="HTTP_PROXY=http://localhost:8001"
Environment="HTTPS_PROXY=http://localhost:8001"
```



``` shell
systemctl daemon-reload 
systemctl restart docker

sudo systemctl show --property=Environment docker
Environment=HTTP_PROXY=http://localhost:8001 HTTPS_PROXY=http://localhost:8001
```



#### pull images

``` shell
[root@10-23-245-35 ~]# kubeadm config images pull
[config/images] Pulled k8s.gcr.io/kube-apiserver:v1.23.4
[config/images] Pulled k8s.gcr.io/kube-controller-manager:v1.23.4
[config/images] Pulled k8s.gcr.io/kube-scheduler:v1.23.4
[config/images] Pulled k8s.gcr.io/kube-proxy:v1.23.4
[config/images] Pulled k8s.gcr.io/pause:3.6
[config/images] Pulled k8s.gcr.io/etcd:3.5.1-0
[config/images] Pulled k8s.gcr.io/coredns/coredns:v1.8.6
```



### Initializing your control-plane node



https://kubernetes.io/docs/reference/setup-tools/kubeadm/kubeadm-init/



#### initializing

``` shell
[root@10-23-75-240 ~]# kubeadm init
[init] Using Kubernetes version: v1.23.4
[preflight] Running pre-flight checks
[preflight] Pulling images required for setting up a Kubernetes cluster
[preflight] This might take a minute or two, depending on the speed of your internet connection
[preflight] You can also perform this action in beforehand using 'kubeadm config images pull'

```



#### set the cgroup driver to systemd

https://kubernetes.io/docs/tasks/administer-cluster/kubeadm/configure-cgroup-driver/



https://www.devopsschool.com/blog/how-to-change-the-cgroup-driver-from-cgroupfs-systemd-in-docker/

``` shell
$ vi /usr/lib/systemd/system/docker.service
Modify this line
ExecStart=/usr/bin/dockerd -H fd:// --containerd=/run/containerd/containerd.sock
To
ExecStart=/usr/bin/dockerd -H fd:// --containerd=/run/containerd/containerd.sock --exec-opt native.cgroupdriver=systemd

#Restart the Docker service by running the following command:
systemctl daemon-reload
systemctl restart docker

# Verify the cgroups driver to systemd
docker info
```





#### remove the kubeadm init

```bash
kubeadm reset

rm ~/.kube/config 

The reset process does not clean CNI configuration. To do so, you must remove /etc/cni/net.d

If you wish to reset iptables, you must do so manually by using the "iptables" command
```



#### restart

``` shell
[root@10-23-75-240 ~]# kubeadm init 
[init] Using Kubernetes version: v1.23.4
[preflight] Running pre-flight checks
[preflight] Pulling images required for setting up a Kubernetes cluster
[preflight] This might take a minute or two, depending on the speed of your internet connection
[preflight] You can also perform this action in beforehand using 'kubeadm config images pull'
[certs] Using certificateDir folder "/etc/kubernetes/pki"
[certs] Generating "ca" certificate and key
[certs] Generating "apiserver" certificate and key
[certs] apiserver serving cert is signed for DNS names [10-23-75-240 kubernetes kubernetes.default kubernetes.default.svc kubernetes.default.svc.cluster.local] and IPs [10.96.0.1 10.23.75.240]
[certs] Generating "apiserver-kubelet-client" certificate and key
[certs] Generating "front-proxy-ca" certificate and key
[certs] Generating "front-proxy-client" certificate and key
[certs] Generating "etcd/ca" certificate and key
[certs] Generating "etcd/server" certificate and key
[certs] etcd/server serving cert is signed for DNS names [10-23-75-240 localhost] and IPs [10.23.75.240 127.0.0.1 ::1]
[certs] Generating "etcd/peer" certificate and key
[certs] etcd/peer serving cert is signed for DNS names [10-23-75-240 localhost] and IPs [10.23.75.240 127.0.0.1 ::1]
[certs] Generating "etcd/healthcheck-client" certificate and key
[certs] Generating "apiserver-etcd-client" certificate and key
[certs] Generating "sa" key and public key
[kubeconfig] Using kubeconfig folder "/etc/kubernetes"
[kubeconfig] Writing "admin.conf" kubeconfig file
[kubeconfig] Writing "kubelet.conf" kubeconfig file
[kubeconfig] Writing "controller-manager.conf" kubeconfig file
[kubeconfig] Writing "scheduler.conf" kubeconfig file
[kubelet-start] Writing kubelet environment file with flags to file "/var/lib/kubelet/kubeadm-flags.env"
[kubelet-start] Writing kubelet configuration to file "/var/lib/kubelet/config.yaml"
[kubelet-start] Starting the kubelet
[control-plane] Using manifest folder "/etc/kubernetes/manifests"
[control-plane] Creating static Pod manifest for "kube-apiserver"
[control-plane] Creating static Pod manifest for "kube-controller-manager"
[control-plane] Creating static Pod manifest for "kube-scheduler"
[etcd] Creating static Pod manifest for local etcd in "/etc/kubernetes/manifests"
[wait-control-plane] Waiting for the kubelet to boot up the control plane as static Pods from directory "/etc/kubernetes/manifests". This can take up to 4m0s
[apiclient] All control plane components are healthy after 6.004286 seconds
[upload-config] Storing the configuration used in ConfigMap "kubeadm-config" in the "kube-system" Namespace
[kubelet] Creating a ConfigMap "kubelet-config-1.23" in namespace kube-system with the configuration for the kubelets in the cluster
NOTE: The "kubelet-config-1.23" naming of the kubelet ConfigMap is deprecated. Once the UnversionedKubeletConfigMap feature gate graduates to Beta the default name will become just "kubelet-config". Kubeadm upgrade will handle this transition transparently.
[upload-certs] Skipping phase. Please see --upload-certs
[mark-control-plane] Marking the node 10-23-75-240 as control-plane by adding the labels: [node-role.kubernetes.io/master(deprecated) node-role.kubernetes.io/control-plane node.kubernetes.io/exclude-from-external-load-balancers]
[mark-control-plane] Marking the node 10-23-75-240 as control-plane by adding the taints [node-role.kubernetes.io/master:NoSchedule]
[bootstrap-token] Using token: tx5otm.356deacqpiioal1p
[bootstrap-token] Configuring bootstrap tokens, cluster-info ConfigMap, RBAC Roles
[bootstrap-token] configured RBAC rules to allow Node Bootstrap tokens to get nodes
[bootstrap-token] configured RBAC rules to allow Node Bootstrap tokens to post CSRs in order for nodes to get long term certificate credentials
[bootstrap-token] configured RBAC rules to allow the csrapprover controller automatically approve CSRs from a Node Bootstrap Token
[bootstrap-token] configured RBAC rules to allow certificate rotation for all node client certificates in the cluster
[bootstrap-token] Creating the "cluster-info" ConfigMap in the "kube-public" namespace
[kubelet-finalize] Updating "/etc/kubernetes/kubelet.conf" to point to a rotatable kubelet client certificate and key
[addons] Applied essential addon: CoreDNS
[addons] Applied essential addon: kube-proxy

Your Kubernetes control-plane has initialized successfully!

To start using your cluster, you need to run the following as a regular user:

  mkdir -p $HOME/.kube
  sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
  sudo chown $(id -u):$(id -g) $HOME/.kube/config

Alternatively, if you are the root user, you can run:

  export KUBECONFIG=/etc/kubernetes/admin.conf

You should now deploy a pod network to the cluster.
Run "kubectl apply -f [podnetwork].yaml" with one of the options listed at:
  https://kubernetes.io/docs/concepts/cluster-administration/addons/

Then you can join any number of worker nodes by running the following on each as root:

kubeadm join 10.23.75.240:6443 --token tx5otm.356deacqpiioal1p \
	--discovery-token-ca-cert-hash sha256:a304d318be204cf2c659d3695c414ca888ad4b8df09a82155dffea2e906e31aa 
```



To make kubectl work for your non-root user, run these commands, which are also part of the `kubeadm init` output:



``` shell
mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config
```



Alternatively, if you are the `root` user, you can run:

``` shell
export KUBECONFIG=/etc/kubernetes/admin.conf
```



``` shell
[root@10-23-75-240 ~]# kubectl get nodes
NAME           STATUS     ROLES                  AGE     VERSION
10-23-75-240   NotReady   control-plane,master   2m33s   v1.23.4
```



### Add a new node to the cluster

#### set no proxy for the netowrk

``` shell
[root@10-23-184-141 ~]# export NO_PROXY=10.23.0.0/12
```



#### join the cluster

``` shell
[root@10-23-184-141 ~]# kubeadm join 10.23.75.240:6443 --token tx5otm.356deacqpiioal1p --discovery-token-ca-cert-hash sha256:a304d318be204cf2c659d3695c414ca888ad4b8df09a82155dffea2e906e31aa 
[preflight] Running pre-flight checks
[preflight] Reading configuration from the cluster...
[preflight] FYI: You can look at this config file with 'kubectl -n kube-system get cm kubeadm-config -o yaml'
[kubelet-start] Writing kubelet configuration to file "/var/lib/kubelet/config.yaml"
[kubelet-start] Writing kubelet environment file with flags to file "/var/lib/kubelet/kubeadm-flags.env"
[kubelet-start] Starting the kubelet
[kubelet-start] Waiting for the kubelet to perform the TLS Bootstrap...

This node has joined the cluster:
* Certificate signing request was sent to apiserver and a response was received.
* The Kubelet was informed of the new secure connection details.

Run 'kubectl get nodes' on the control-plane to see this node join the cluster
```



### Installing a Pod network add-on

You can install a Pod network add-on with the following command on the control-plane node or a node that has the kubeconfig credentials:

``` shell
kubectl apply -f <add-on.yaml>
```



You can install only one Pod network per cluster.

Once a Pod network has been installed, you can confirm that it is working by checking that the CoreDNS Pod is `Running` in the output of `kubectl get pods --all-namespaces`. And once the CoreDNS Pod is up and running, you can continue by joining your nodes.



https://antrea.io/docs/v1.4.0/docs/getting-started/



``` shell
kubeadm init --pod-network-cidr=10.23.0.0/10
```



``` shell
kubectl apply -f https://github.com/antrea-io/antrea/releases/download/v1.5.1/antrea.yml
```





### Joining your nodes

``` shell
kubeadm join 10.23.75.240:6443 --token bncfcw.2aqncopypuid2cqk --discovery-token-ca-cert-hash sha256:7da3c94757189038f7c6036ab05278e502fee429ea74221594f19e6cd0b75078
```



``` shell
kubeadm token list

TOKEN           TTL     EXPIRES        USAGES          DESCRIPTION                        EXTRA GROUPS

bncfcw.2aqncopypuid2cqk  23h     2022-03-13T03:40:11Z  authentication,signing  The default bootstrap token generated by 'kubeadm init'.  system:bootstrappers:kubeadm:default-node-token

[root@10-23-75-240 ~]# 
```









##  add new node to existing cluster



``` shell
[root@10-23-130-9 ~]# kubeadm join 10.23.75.240:6443 --token 1z8da4.mk9soruz9eqkkenz --discovery-token-ca-cert-hash sha256:e41e47ba0b518ef4a0e1b1e0cd978ebf2fa330c7d14ef921f6ae64b28e217c63 --v=2
I0321 14:57:53.645606   10937 join.go:413] [preflight] found NodeName empty; using OS hostname as NodeName
I0321 14:57:53.645708   10937 initconfiguration.go:117] detected and using CRI socket: /var/run/dockershim.sock
[preflight] Running pre-flight checks
I0321 14:57:53.645770   10937 preflight.go:92] [preflight] Running general checks
I0321 14:57:53.645814   10937 checks.go:283] validating the existence of file /etc/kubernetes/kubelet.conf
I0321 14:57:53.645822   10937 checks.go:283] validating the existence of file /etc/kubernetes/bootstrap-kubelet.conf
I0321 14:57:53.645829   10937 checks.go:107] validating the container runtime
I0321 14:57:53.751839   10937 checks.go:133] validating if the "docker" service is enabled and active
I0321 14:57:53.760490   10937 checks.go:332] validating the contents of file /proc/sys/net/bridge/bridge-nf-call-iptables
I0321 14:57:53.760551   10937 checks.go:332] validating the contents of file /proc/sys/net/ipv4/ip_forward
I0321 14:57:53.760571   10937 checks.go:654] validating whether swap is enabled or not
I0321 14:57:53.760595   10937 checks.go:373] validating the presence of executable conntrack
I0321 14:57:53.760611   10937 checks.go:373] validating the presence of executable ip
I0321 14:57:53.760626   10937 checks.go:373] validating the presence of executable iptables
I0321 14:57:53.760636   10937 checks.go:373] validating the presence of executable mount
I0321 14:57:53.760654   10937 checks.go:373] validating the presence of executable nsenter
I0321 14:57:53.760667   10937 checks.go:373] validating the presence of executable ebtables
I0321 14:57:53.760678   10937 checks.go:373] validating the presence of executable ethtool
I0321 14:57:53.760686   10937 checks.go:373] validating the presence of executable socat
I0321 14:57:53.760698   10937 checks.go:373] validating the presence of executable tc
I0321 14:57:53.760708   10937 checks.go:373] validating the presence of executable touch
I0321 14:57:53.760727   10937 checks.go:521] running all checks
I0321 14:57:53.859066   10937 checks.go:404] checking whether the given node name is valid and reachable using net.LookupHost
I0321 14:57:53.859264   10937 checks.go:620] validating kubelet version
I0321 14:57:53.908911   10937 checks.go:133] validating if the "kubelet" service is enabled and active
I0321 14:57:53.915269   10937 checks.go:206] validating availability of port 10250
I0321 14:57:53.915449   10937 checks.go:283] validating the existence of file /etc/kubernetes/pki/ca.crt
I0321 14:57:53.915481   10937 checks.go:433] validating if the connectivity type is via proxy or direct
I0321 14:57:53.915511   10937 join.go:530] [preflight] Discovering cluster-info
I0321 14:57:53.915540   10937 token.go:80] [discovery] Created cluster-info discovery client, requesting info from "10.23.75.240:6443"
I0321 14:57:53.937181   10937 token.go:223] [discovery] The cluster-info ConfigMap does not yet contain a JWS signature for token ID "1z8da4", will try again
I0321 14:57:59.853449   10937 token.go:223] [discovery] The cluster-info ConfigMap does not yet contain a JWS signature for token ID "1z8da4", will try again
```







``` shell
[root@10-23-75-240 k8s-specs]# kubeadm token list
[root@10-23-75-240 k8s-specs]# 

```



``` shell
[root@10-23-75-240 k8s-specs]# kubeadm token create
t982ar.yrfd7pnikofhjb0v
```



``` shell
kubeadm join 10.23.75.240:6443 --token t982ar.yrfd7pnikofhjb0v --discovery-token-ca-cert-hash sha256:e41e47ba0b518ef4a0e1b1e0cd978ebf2fa330c7d14ef921f6ae64b28e217c63 --v=2
```



``` shell
[root@10-23-75-240 k8s-specs]# kubectl top nodes
NAME            CPU(cores)   CPU%        MEMORY(bytes)   MEMORY%     
10-23-184-141   74m          3%          1940Mi          52%         
10-23-245-35    68m          3%          1708Mi          46%         
10-23-75-240    162m         4%          3008Mi          39%         
10-23-130-9     <unknown>    <unknown>   <unknown>       <unknown>   
[root@10-23-75-240 k8s-specs]# kubectl top nodes
NAME            CPU(cores)   CPU%   MEMORY(bytes)   MEMORY%   
10-23-130-9     46m          2%     667Mi           18%       
10-23-184-141   77m          3%     1942Mi          52%       
10-23-245-35    69m          3%     1706Mi          46%       
10-23-75-240    155m         3%     3013Mi          39%   
```



 