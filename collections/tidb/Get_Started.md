# Q/A

## Q1: cannot run helm on macos big sur after installation

``` shell
sudo spctl --master-disable
``` 





# following tutorial

``` shell
kubectl apply -f https://raw.githubusercontent.com/pingcap/tidb-operator/{tidb-operator version}/manifests/crd.yaml
helm repo add pingcap https://charts.pingcap.org/
helm install tidb-operator pingcap/tidb-operator --version {​tidb-operator version}

```

## 1. install kind

``` shell
GO111MODULE="on" go get sigs.k8s.io/kind@v0.11.1 && kind create cluster
```

https://kind.sigs.k8s.io/


### 1.1 commonly used command

``` shell
kind create cluster

kind delete cluster
```
To create a cluster from Kubernetes source:

- ensure that Kubernetes is cloned in $(go env GOPATH)/src/k8s.io/kubernetes
- build a node image and create a cluster with

``` shell
kind build node-image
kind create cluster --image kindest/node:latest
```

### 1.2 check whether the cluster is successfully created
``` shell
FVFF87EFQ6LR :: ~/playground » kubectl cluster-info
Kubernetes control plane is running at https://127.0.0.1:61972
CoreDNS is running at https://127.0.0.1:61972/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy

To further debug and diagnose cluster problems, use 'kubectl cluster-info dump'.

```

##  Install TiDB Operator CRDs


``` shell
kubectl apply -f https://raw.githubusercontent.com/pingcap/tidb-operator/master/manifests/crd.yaml
```

#### the above command maybe not working when the network condition is very bad.

try copy the content of  https://raw.githubusercontent.com/pingcap/tidb-operator/master/manifests/crd.yaml into a local file crd.yaml

``` shell
kubectl apply -f crd.yaml                                                           130 ↵
Warning: apiextensions.k8s.io/v1beta1 CustomResourceDefinition is deprecated in v1.16+, unavailable in v1.22+; use apiextensions.k8s.io/v1 CustomResourceDefinition
customresourcedefinition.apiextensions.k8s.io/tidbclusters.pingcap.com created
customresourcedefinition.apiextensions.k8s.io/dmclusters.pingcap.com created
customresourcedefinition.apiextensions.k8s.io/backups.pingcap.com created
customresourcedefinition.apiextensions.k8s.io/restores.pingcap.com created
customresourcedefinition.apiextensions.k8s.io/backupschedules.pingcap.com created
customresourcedefinition.apiextensions.k8s.io/tidbmonitors.pingcap.com created
customresourcedefinition.apiextensions.k8s.io/tidbinitializers.pingcap.com created
customresourcedefinition.apiextensions.k8s.io/tidbclusterautoscalers.pingcap.com created
```

## Deploy TiDB Operator

### 1. Add the PingCAP repository

``` shell
helm repo add pingcap https://charts.pingcap.org/
``` 

### 2. Create a namespace for TiDB Operator

``` shell
kubectl create namespace tidb-admin

```

### 3. install TiDB Operator

``` shell
helm install --namespace tidb-admin tidb-operator pingcap/tidb-operator --version v1.2.0
```

#### expected Output

``` shell
NAME: tidb-operator
LAST DEPLOYED: Mon Jun  1 12:31:43 2020
NAMESPACE: tidb-admin
STATUS: deployed
REVISION: 1
TEST SUITE: None
NOTES:
Make sure tidb-operator components are running:

    kubectl get pods --namespace tidb-admin -l app.kubernetes.io/instance=tidb-operator

```

### 4. confirm that the TiDB Operator components are running

``` shell
kubectl get pods --namespace tidb-admin -l app.kubernetes.io/instance=tidb-operator
``` 


#### Expected output

``` shell

NAME                                       READY   STATUS    RESTARTS   AGE
tidb-controller-manager-6d8d5c6d64-b8lv4   1/1     Running   0          2m22s
tidb-scheduler-644d59b46f-4f6sb            2/2     Running   0          2m22s
```


##  Deploy a TiDB cluster and its monitoring services

``` shell
kubectl create namespace tidb-cluster && \
    kubectl -n tidb-cluster apply -f https://raw.githubusercontent.com/pingcap/tidb-operator/master/examples/basic/tidb-cluster.yaml
```


#### expected output

``` shell
namespace/tidb-cluster created
tidbcluster.pingcap.com/basic created
```

### Deploy TiDB monitoring services

``` shell
curl -LO https://raw.githubusercontent.com/pingcap/tidb-operator/master/examples/basic/tidb-monitor.yaml && \
kubectl -n tidb-cluster apply -f tidb-monitor.yaml
``` 

## View the pod status
``` shell
kubectl get po -n tidb-cluster
```

### expected output 
``` shell
FVFF87EFQ6LR :: ~/playground/install_tidb » kubectl get po -n tidb-cluster
NAME                               READY   STATUS            RESTARTS   AGE
basic-discovery-68d7b985cd-lzn6c   1/1     Running           0          67s
basic-monitor-0                    0/3     PodInitializing   0          31s
basic-pd-0                         1/1     Running           0          67s
basic-tidb-0                       2/2     Running           0          24s
basic-tikv-0                       1/1     Running           0          54s
```

Wait until all Pods for all services are started. As soon as you see Pods of each type (-pd, -tikv, and -tidb) are in the "Running" state, you can press Ctrl+C to get back to the command line and go on to connect to your TiDB cluster.

``` shell
NAME                              READY   STATUS    RESTARTS   AGE
basic-discovery-6bb656bfd-xl5pb   1/1     Running   0          9m9s
basic-monitor-5fc8589c89-gvgjj    3/3     Running   0          8m58s
basic-pd-0                        1/1     Running   0          9m8s
basic-tidb-0                      2/2     Running   0          7m14s
basic-tikv-0                      1/1     Running   0          8m13s

```


## Connect to TiDB

Because TiDB supports the MySQL protocol and most of its syntax, you can connect to TiDB using the MySQL client.

### Install the MySQL client

To connect to TiDB, you need a MySQL-compatible client installed on the host where kubectl is installed.
This can be the mysql executable from an installation of MySQL Server, MariaDB Server, Percona Server, or a standalone client executable from your operating system's package repository.


### Forward port 4000

You can connect to TiDB by first forwarding a port from the local host to the TiDB service in Kubernetes.

First, get a list of services in the tidb-cluster namespace:


``` shell
kubectl get svc -n tidb-cluster
```

``` shell
FVFF87EFQ6LR :: ~/playground/install_tidb » kubectl get svc -n tidb-cluster
NAME                     TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)               AGE
basic-discovery          ClusterIP   10.96.70.71     <none>        10261/TCP,10262/TCP   4m45s
basic-grafana            ClusterIP   10.96.246.108   <none>        3000/TCP              4m9s
basic-monitor-reloader   ClusterIP   10.96.193.23    <none>        9089/TCP              4m9s
basic-pd                 ClusterIP   10.96.79.163    <none>        2379/TCP              4m45s
basic-pd-peer            ClusterIP   None            <none>        2380/TCP              4m45s
basic-prometheus         ClusterIP   10.96.99.169    <none>        9090/TCP              4m10s
basic-tidb               ClusterIP   10.96.174.10    <none>        4000/TCP,10080/TCP    4m2s
basic-tidb-peer          ClusterIP   None            <none>        10080/TCP             4m2s
basic-tikv-peer          ClusterIP   None            <none>        20160/TCP             4m32s
```
In this case, the TiDB service is called basic-tidb. Run the following command to forward this port from the local host to the cluster:

``` shell
kubectl port-forward -n tidb-cluster svc/basic-tidb 4000 > pf4000.out &
```

### connect to TiDB service 
``` shell
mysql -h 127.0.0.1 -P 4000 -u root
```


#### MySQL

https://hub.docker.com/_/mysql

This image can also be used as a client for non-Docker or remote instances:


``` shell
 
docker run -it --rm mysql mysql -h 127.0.0.1 --port 4000 -u root
```

**M1 docker**
``` shell

docker run --platform linux/x86_64
```

the solution with docker is not working yet.


https://tech-cookbook.com/2021/03/10/using-mysql-workbench-on-macos-big-sur-m1/

download on my macos pc

mysql shell.

``` shell
FVFF87EFQ6LR :: ~/playground/install_tidb » mysqlsh -h 127.0.0.1 -P 4000 -u root --auth-method=mysql_native_password                           130 ↵
Please provide the password for 'root@127.0.0.1:4000': 
Save password for 'root@127.0.0.1:4000'? [Y]es/[N]o/Ne[v]er (default No): 
MySQL Shell 8.0.26

Copyright (c) 2016, 2021, Oracle and/or its affiliates.
Oracle is a registered trademark of Oracle Corporation and/or its affiliates.
Other names may be trademarks of their respective owners.

Type '\help' or '\?' for help; '\quit' to exit.
Creating a session to 'root@127.0.0.1:4000?auth-method=mysql_native_password'
Fetching schema names for autocompletion... Press ^C to stop.
Your MySQL connection id is 653
Server version: 5.7.25-TiDB-v5.1.0 TiDB Server (Apache License 2.0) Community Edition, MySQL 5.7 compatible
No default schema selected; type \use <schema> to set one.
 MySQL  127.0.0.1:4000  JS > 
 MySQL  127.0.0.1:4000  JS > 
 MySQL  127.0.0.1:4000  JS > ls
ReferenceError: ls is not defined
 MySQL  127.0.0.1:4000  JS > 
 ```
 at least this one is working


 