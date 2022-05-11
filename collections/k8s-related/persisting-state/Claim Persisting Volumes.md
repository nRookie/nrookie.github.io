## Usage of Persistent Volumes[#](https://www.educative.io/module/lesson/a-practical-guide-to-kubernetes/RLnX1l0kK1w#Usage-of-Persistent-Volumes)

Kubernetes persistent volumes are useless if no one uses them. They exist only as objects with relation to, in our case, specific EBS volumes. They are waiting for someone to claim them through the PersistentVolumeClaim resource.

Just like Pods which can request specific resources like memory and CPU, PersistentVolumeClaims can request particular sizes and access modes. Both are, in a way, consuming resources, even though of different types. Just as Pods should not specify on which node they should run, PersistentVolumeClaims cannot define which volume they should mount. Instead, Kubernetes scheduler will assign them a volume depending on the claimed resources.



## How to Claim Persistent Volumes?[#](https://www.educative.io/module/lesson/a-practical-guide-to-kubernetes/RLnX1l0kK1w#How-to-Claim-Persistent-Volumes?)



We’ll use `pv/pvc.yml` to explore how we could claim a persistent volume.



``` shell
[root@10-23-75-240 k8s-specs]# vi pv/nfs-pv.yml 

apiVersion: v1
kind: PersistentVolume
metadata:
  name: jenkins
  namespace: jenkins
  labels:
    type: local
    author: qing.na
spec:
  storageClassName: manual
  capacity:
    storage: 40Gi
  accessModes:
    - ReadWriteMany
 # persistentVolumeReclaimPolicy: Retain
  nfs:
    path: /data/
    server: 172.16.0.51
    readOnly: false
```

The YAML file defines a `PersistentVolumeClaim` with the storage class name `manual-ebs`. That is the same class as the persistent volumes `manual-ebs-*` we created earlier. The access mode and the storage request are also matching what we defined for the persistent volume.

Please note that we are not specifying which volume we’d like to use. Instead, this claim specifies a set of attributes (`storageClassName`, `accessModes`, and `storage`). Any of the volumes in the system that match those specifications might be claimed by the `PersistentVolumeClaim` named `jenkins`.

Bear in mind that `resources` do not have to be the exact match. Any volume that has the same or bigger amount of storage is considered a match. A claim for `1Gi` can be translated to *at least `1Gi`*. In our case, a claim for `1Gi` matches all three persistent volumes since they are set to `5Gi`.



``` shell
 vi pv/nfs-pvc.yml 

kind: PersistentVolumeClaim
apiVersion: v1
metadata:
  name: jenkins
  namespace: jenkins
spec:
  storageClassName: manual
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 1Gi


[root@10-23-75-240 k8s-specs]# kubectl create -f pv/nfs-pvc.yml 
persistentvolumeclaim/jenkins created

```



``` shell
│ Name:            jenkins                                                                                                                                      
│ Labels:          author=qing.na                                                                                                                               
│                  type=local                                                                                                                                   
│ Annotations:     <none>                                                                                                                                       
│ Finalizers:      [kubernetes.io/pv-protection]                                                                                                                
│ StorageClass:    manual                                                                                                                                       
│ Status:          Available                                                                                                                                    
│ Claim:                                                                                                                                                        
│ Reclaim Policy:  Retain                                                                                                                                       
│ Access Modes:    RWX                                                                                                                                          
│ VolumeMode:      Filesystem                                                                                                                                   
│ Capacity:        40Gi                                                                                                                                         
│ Node Affinity:   <none>                                                                                                                                       
│ Message:                                                                                                                                                      
│ Source:                                                                                                                                                       
│     Type:      NFS (an NFS mount that lasts the lifetime of a pod)                                                                                            
│     Server:    172.16.0.51                                                                                                                                    
│     Path:      /data/                                                                                                                                         
│     ReadOnly:  false                                                                                                                                          
│ Events:        <none>                                                                                                                                         
│                                                                                                                                                               
│                              
```



``` shell
Used By:       <none>                                                                                                                                         
│ Events:                                                                                                                                                       
│   Type     Reason              Age   From                         Message                                                                                     
│   ----     ------              ----  ----                         -------                                                                                     
│   Warning  ProvisioningFailed  5s    persistentvolume-controller  storageclass.storage.k8s.io "manual" not found    
```

https://stackoverflow.com/questions/49174300/storageclass-storage-k8s-io-standard-not-found-for-pvc-on-bare-metal-kubernete



``` shell
---
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: manual
provisioner: kubernetes.io/gce-pd
parameters:
  type: pd-standard
```



``` shel
[root@10-23-75-240 k8s-specs]# kubectl apply -f pv/nfs-pv.yml 
persistentvolume/jenkins configured
storageclass.storage.k8s.io/manual created
```



``` shell
│   Type     Reason                Age                   From                         Message                                                                   
│   ----     ------                ----                  ----                         -------                                                                   
│   Warning  ProvisioningFailed    25s (x11 over 2m53s)  persistentvolume-controller  storageclass.storage.k8s.io "manual" not found                            
│   Normal   ExternalProvisioning  10s (x2 over 10s)     persistentvolume-controller  waiting for a volume to be created, either by external provisioner "pd.c  
│ si.storage.gke.io" or manually created by system administrator   
```



https://www.linuxtechi.com/configure-nfs-persistent-volume-kubernetes/



``` shell
│ Events:                                                                                                                                                       
│   Type     Reason              Age   From                         Message                                                                                     
│   ----     ------              ----  ----                         -------                                                                                     
│   Warning  ProvisioningFailed  4s    persistentvolume-controller  storageclass.storage.k8s.io "nfs" not found
```

### NFS

```yaml
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: example-nfs
provisioner: example.com/external-nfs
parameters:
  server: nfs-server.example.com
  path: /share
  readOnly: "false"
```

- `server`: Server is the hostname or IP address of the NFS server.
- `path`: Path that is exported by the NFS server.
- `readOnly`: A flag indicating whether the storage will be mounted as read only (default false).

- `server`: Server is the hostname or IP address of the NFS server.
- `path`: Path that is exported by the NFS server.
- `readOnly`: A flag indicating whether the storage will be mounted as read only (default false).

Kubernetes doesn't include an internal NFS provisioner. You need to use an external provisioner to create a StorageClass for NFS. Here are some examples:

- [NFS Ganesha server and external provisioner](https://github.com/kubernetes-sigs/nfs-ganesha-server-and-external-provisioner)
- [NFS subdir external provisioner](https://github.com/kubernetes-sigs/nfs-subdir-external-provisioner)

https://kubernetes.io/docs/concepts/storage/storage-classes/#nfs



``` shell
   Type     Reason                Age                   From                         Message                                                                   
│   ----     ------                ----                  ----                         -------                                                                   
│   Warning  ProvisioningFailed    15s (x19 over 4m41s)  persistentvolume-controller  storageclass.storage.k8s.io "nfs" not found                               
│   Normal   ExternalProvisioning  0s (x2 over 0s)       persistentvolume-controller  waiting for a volume to be created, either by external provisioner "exam  
│ ple.com/external-nfs" or manually created by system administrator                                                                                             
│                                                                                 
```



https://kubernetes.io/docs/concepts/storage/persistent-volumes/





``` shell
$ helm repo add nfs-subdir-external-provisioner https://kubernetes-sigs.github.io/nfs-subdir-external-provisioner/
$ helm install nfs-subdir-external-provisioner nfs-subdir-external-provisioner/nfs-subdir-external-provisioner \
    --set nfs.server=172.16.0.51\
    --set nfs.path=/data
    NAME: nfs-subdir-external-provisioner
LAST DEPLOYED: Thu Mar 17 02:09:32 2022
NAMESPACE: default
STATUS: deployed
REVISION: 1
TEST SUITE: None

```



``` shell
[root@10-23-75-240 k8s-specs]# helm install nfs-subdir-external-provisioner nfs-subdir-external-provisioner/nfs-subdir-external-provisioner  --namespace jenkins    --set nfs.server=172.16.0.51    --set nfs.path=/data 
NAME: nfs-subdir-external-provisioner
LAST DEPLOYED: Thu Mar 17 02:11:33 2022
NAMESPACE: jenkins
STATUS: deployed
REVISION: 1
TEST SUITE: None

```



![image-20220317022304265](/Users/user/playground/share/nrookie.github.io/collections/k8s-related/persisting-state/image-20220317022304265.png)

``` shell
[root@10-23-75-240 k8s-specs]# yum -y install nfs-utils
```





it worked . Cheers!



### bad news not working anymore



https://github.com/kubernetes-sigs/nfs-subdir-external-provisioner/issues/107



``` shell
helm install nfs-subdir-external-provisioner nfs-subdir-external-provisioner/nfs-subdir-external-provisioner  --namespace jenkins    --set nfs.server=172.16.0.51    --set nfs.path=/data --set storageClass.provisionerName=k8s-sigs.io/nfs-subdir-external-provisioner
```

still not working





![image-20220321091408451](/Users/user/playground/share/nrookie.github.io/collections/k8s-related/persisting-state/image-20220321091408451.png)





``` yaml
[root@10-23-75-240 k8s-specs]# cat pv/nfs-pvc.yml 
kind: PersistentVolumeClaim
apiVersion: v1
metadata:
  name: jenkins
  namespace: jenkins
spec:
  storageClassName: nfs-client
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 10Gi
[root@10-23-75-240 k8s-specs]# cat pv/nfs-pv
nfs-pvc.yml  nfs-pv.yml   
[root@10-23-75-240 k8s-specs]# cat pv/nfs-pv.yml 
apiVersion: v1
kind: PersistentVolume
metadata:
  name: jenkins
  namespace: jenkins
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
[root@10-23-75-240 k8s-specs]# 

```





``` yaml
[root@10-23-75-240 k8s-specs]# cat pv/nfs-pvc.yml 
kind: PersistentVolumeClaim
apiVersion: v1
metadata:
  name: jenkins
  namespace: jenkins
spec:
  storageClassName: nfs-client
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 10Gi
```





``` shell
[root@10-23-75-240 k8s-specs]# cat pv/strclass.yml 
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: nfs
  namespace: jenkins
provisioner: nfs-subdir-external-provisioner
parameters:
  server: 172.16.0.51
  path: /data
  readOnly: "false"
```





``` shell
helm install nfs-subdir-external-provisioner nfs-subdir-external-provisioner/nfs-subdir-external-provisioner  --namespace jenkins    --set nfs.server=172.16.0.51    --set nfs.path=/data --set storageClass.provisionerName=k8s-sigs.io/nfs-subdir-external-provisioner
```



It worked
