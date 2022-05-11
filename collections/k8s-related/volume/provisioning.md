## Static Volume Provisioning[#](https://www.educative.io/module/lesson/a-practical-guide-to-kubernetes/BnMLJzKxjWW#Static-Volume-Provisioning)

So far, we used static PersistentVolumes. We had to create both EBS volumes and Kubernetes PersistentVolumes manually. Only after both became available we were able to deploy Pods that are mounting those volumes through PersistentVolumeClaims. We’ll call this process static volume provisioning.



In some cases, static volume provisioning is a necessity. Our infrastructure might not be capable of creating dynamic volumes. That is often the case with on-premise infrastructure with volumes based on Network File System (NFS). Even then, with a few tools, a change in processes, and the right choices for supported volume types, we can often reach the point where volume provisioning is dynamic. Still, that might prove to be a challenge with legacy processes and infrastructure.





Since our cluster is in AWS, we cannot blame legacy infrastructure for provisioning volumes manually. Indeed, we could have jumped straight into this section. After all, AWS is all about dynamic infrastructure management.

However, we felt that it will be easier to understand the processes by exploring manual provisioning first. The knowledge we obtained thus far will help us understand better what’s coming next.

The second reason is that maybe you will run a Kubernetes cluster on infrastructure that has to be static. Even though we’re using AWS for the examples, everything you learned this far can be implemented on static infrastructure. You’ll only have to change EBS with NFS and go through the [NFSVolumeSource](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.10/#nfsvolumesource-v1-core) documentation. There are only **three** NFS-specific fields so you should be up-and-running in no time.

Before we discuss how to enable dynamic persistent volume provisioning, we should understand that it will be used only if none of the static PersistentVolumes match our claims. In other words, Kubernetes will always select statically created PersistentVolumes over dynamic ones.





## Dynamic Volume Provisioning[#](https://www.educative.io/module/lesson/a-practical-guide-to-kubernetes/BnMLJzKxjWW#Dynamic-Volume-Provisioning)



Dynamic volume provisioning allows us to create storage on-demand. Instead of manually pre-provisioning storage, we can provision it automatically when a resource requests it.

We can enable dynamic provisioning through the usage of StorageClasses from the `storage.k8s.io` API group. They allow us to describe the types of storage that can be claimed.

On the one hand, a cluster administrator can create as many StorageClasses as there are storage flavors. On the other hand, the users of the cluster do not have to worry about the details of each available external storage. It’s a win-win situation where the administrators do not have to create PersistentVolumes in advance, and the users can simply claim the storage type they need.



![image-20220321094629198](/Users/user/playground/share/nrookie.github.io/collections/k8s-related/volume/image-20220321094629198.png)





![image-20220321095027739](/Users/user/playground/share/nrookie.github.io/collections/k8s-related/volume/image-20220321095027739.png)