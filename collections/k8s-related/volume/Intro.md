## The Volumes[#](https://www.educative.io/module/lesson/a-practical-guide-to-kubernetes/3jP2QljQxZM#The-Volumes)

**Kubernetes Volumes** solve the need to preserve the state across container crashes. In essence, Volumes are references to files and directories made accessible to containers that form a Pod. The significant difference between different types of Kubernetes Volumes is in the way these files and directories are created.

While the primary use-case for Volumes is the preservation of state, there are quite a few others. For example, we might use Volumes to access Docker’s socket running on a host. Or we might use them to access configuration residing in a file on the host file system.



We can describe Volumes as a way to access a file system that might be running on the same host or somewhere else. No matter where that file system is, it is external to the containers that mount volumes. There can be many reasons why someone might mount a Volume, with state preservation being only one of them.



There are over **twenty-five** Volume types supported by Kubernetes. It would take us too much time to go through all of them. Besides, even if we’d like to do that, many Volume types are specific to a hosting vendor. For example, `awsElasticBlockStore` works only with AWS, `azureDisk` and `azureFile` work only with Azure, and so on and so forth.

We’ll limit our exploration to Volume types that can be used within Minikube. You should be able to extrapolate that knowledge to Volume types applicable to your hosting vendor of choice.

Let’s get down to it.



## The hostPath Volume 





`hostPath` allows us to mount a file or a directory from a host to Pods and, through them, to containers. Before we discuss the usefulness of this type, we’ll have a short discussion about use-cases when this is not a good choice.

**Do not** use `hostPath` to store a state of an application. Since it mounts a file or a directory from a host into a Pod, it is not fault-tolerant. If the server fails, Kubernetes will schedule the Pod to a healthy node, and the state will be lost.



For our use case, `hostPath` works just fine. We’re not using it to preserve state, but to gain access to Docker server running on the same host as the Pod.

**Line 15-18:** The `hostPath` type has only **two** fields. The `path` represents the file or a directory we want to mount from the host. Since we want to mount a socket, we set the `type` accordingly. There are other types we could use.



### Types of Mounts in hostPath [#](https://www.educative.io/module/lesson/a-practical-guide-to-kubernetes/N8w9j7qLwY6#Types-of-Mounts-in-hostPath-)

- The `Directory` type will mount a directory from the host. It must exist on the given path. If it doesn’t, we might switch to `DirectoryOrCreate` type which serves the same purpose. The difference is that `DirectoryOrCreate` will create the directory if it does not exist on the host.
- The `File` and `FileOrCreate` are similar to their `Directory` equivalents. The only difference is that this time we’d mount a file, instead of a directory.
- The other supported types are `Socket`, `CharDevice`, and `BlockDevice`. They should be self-explanatory. If you don’t know what character or block devices are, you probably don’t need those types.

These were the types of mounts supported by the hostPath.



`hostPath` is a great solution for accessing host resources like `/var/run/docker.sock`, `/dev/cgroups`, and others. That is, as long as the resource we’re trying to reach is on the same node as the Pod.



