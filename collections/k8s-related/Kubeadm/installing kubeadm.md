# Installing kubeadm

## Env 

Centos 7.9

## Prepare



make sure br_netfilter is loaded

``` shell
[root@10-23-41-203 ~]# sudo modprobe br_netfilter

[root@10-23-41-203 ~]# lsmod | grep br_netfilter

br_netfilter      24576 0

bridge        192512 1 br_netfilter
```



## set container runtime



### containerd[ ](https://kubernetes.io/docs/setup/production-environment/container-runtimes/#containerd)

#### Prerequisites

``` shell
cat <<EOF | sudo tee /etc/modules-load.d/containerd.conf
overlay
br_netfilter
EOF

sudo modprobe overlay
sudo modprobe br_netfilter

# Setup required sysctl params, these persist across reboots.
cat <<EOF | sudo tee /etc/sysctl.d/99-kubernetes-cri.conf
net.bridge.bridge-nf-call-iptables  = 1
net.ipv4.ip_forward                 = 1
net.bridge.bridge-nf-call-ip6tables = 1
EOF

# Apply sysctl params without reboot
sudo sysctl --system
```



#### 1. Install container.io

setup the docker repository

``` shell
sudo yum install -y yum-utils
sudo yum-config-manager \
--add-repo \
https://download.docker.com/linux/centos/docker-ce.repo
```

install Docker Engine

``` shell
sudo yum install docker-ce docker-ce-cli containerd.io
```



start docker 

``` shell
sudo systemctl start docker
systemctl enable docker
```



Verify

``` shell
sudo docker run hello-world
```



#### 2. Configure containerd

``` shell
sudo mkdir -p /etc/containerd
containerd config default | sudo tee /etc/containerd/config.toml
```



#### 3. Restart containerd:

``` shell
sudo systemctl restart containerd
```





### Using the systemd group driver

To use the system cgroup driver in /etc/containerd/config.toml

with runs



``` shell
      [plugins."io.containerd.grpc.v1.cri".containerd.runtimes]

        [plugins."io.containerd.grpc.v1.cri".containerd.runtimes.runc]
          base_runtime_spec = ""
          container_annotations = []
          pod_annotations = []
          privileged_without_host_devices = false
          runtime_engine = ""
          runtime_root = ""
          runtime_type = "io.containerd.runc.v2"

          [plugins."io.containerd.grpc.v1.cri".containerd.runtimes.runc.options]
            BinaryName = ""
            CriuImagePath = ""
            CriuPath = ""
            CriuWorkPath = ""
            IoGid = 0
            IoUid = 0
            NoNewKeyring = false
            NoPivotRoot = false
            Root = ""
            ShimCgroup = ""
            SystemdCgroup = true # this line
```

restart containerd again



``` shell
sudo systemctl restart containerd
```

When using kubeadm, manually configure the [cgroup driver for kubelet](https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/install-kubeadm/#configure-cgroup-driver-used-by-kubelet-on-control-plane-node).



## Installing kubeadm, kubelet and kubectl



You will install these packages on all of your machines:

- `kubeadm`: the command to bootstrap the cluster.
- `kubelet`: the component that runs on all of the machines in your cluster and does things like starting pods and containers.
- `kubectl`: the command line util to talk to your cluster.



**disable gpgcheck, repo_gpgcheck here, set proxy for kubernetes repo**

``` shell
cat <<EOF | sudo tee /etc/yum.repos.d/kubernetes.repo
[kubernetes]
name=Kubernetes
baseurl=https://packages.cloud.google.com/yum/repos/kubernetes-el7-\$basearch
enabled=1
gpgcheck=0
repo_gpgcheck=0
gpgkey=https://packages.cloud.google.com/yum/doc/yum-key.gpg https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg
exclude=kubelet kubeadm kubectl
proxy=http://localhost:8001
EOF

# Set SELinux in permissive mode (effectively disabling it)
sudo setenforce 0
sudo sed -i 's/^SELINUX=enforcing$/SELINUX=permissive/' /etc/selinux/config

sudo yum install -y kubelet kubeadm kubectl --disableexcludes=kubernetes

sudo systemctl enable --now kubelet
```















## Terminology



### br_netfilter

**This module is required to enable transparent masquerading and to facilitate Virtual Extensible LAN (VxLAN) traffic for communication between Kubernetes pods across the cluster**.







# Appendix



Set proxy



