## The gitRepo Volume Type [#](https://www.educative.io/module/lesson/a-practical-guide-to-kubernetes/7XBqRQ1DKpj#The-gitRepo-Volume-Type-)





The `gitRepo` Volume type is probably not going to be on your list of top three Volume types. Or, maybe it will. It all depends on your use cases. We like it since it demonstrates how a concept of a Volume can be extended to a new and innovative solution.



``` shell
cat volume/github.yml

apiVersion: v1
kind: Pod
metadata:
  name: github
spec:
  containers:
  - name: github
    image: docker:17.11
    command: ["sleep"]
    args: ["100000"]
    volumeMounts:
    - mountPath: /var/run/docker.sock
      name: docker-socket
    - mountPath: /src
      name: github
  volumes:
  - name: docker-socket
    hostPath:
      path: /var/run/docker.sock
      type: Socket
  - name: github
    gitRepo:
      repository: https://github.com/vfarcic/go-demo-2.git
      directory: .
```



This Pod definition is very similar to `volume/docker.yml`. The only significant difference is that we added the second `volumeMount`. It will mount the directory `/src` inside the container, and will use the Volume named `github`. The Volume definition is straightforward. The `gitRepo` type defines the Git `repository` and the `directory`. If we skipped the latter, we’d get the repository mounted as `/src/go-demo-2`.



``` shell
The gitRepo Volume type allows a third field which we haven’t used. We could have set a specific revision of the repository. But, for demo purposes, the HEAD should do.
```





### Creating the Pod [#](https://www.educative.io/module/lesson/a-practical-guide-to-kubernetes/7XBqRQ1DKpj#Creating-the-Pod-)



``` shell
[root@10-23-75-240 k8s-specs]# kubectl create \
>     -f volume/github.yml
Warning: spec.volumes[1].gitRepo: deprecated in v1.11
pod/github created
```



Now that we created the Pod, we’ll enter its only container, and check whether `gitRepo` indeed works as expected.



``` shell
kubectl exec -it github sh
cd /src
ls -l
```



``` shell
│ Events:                                                                                                                                                                                           
│   Type     Reason       Age               From               Message                                                                                                                              
│   ----     ------       ----              ----               -------                                                                                                                              
│   Normal   Scheduled    69s               default-scheduler  Successfully assigned default/github to 10-23-184-141                                                                                
│   Warning  FailedMount  5s (x8 over 69s)  kubelet            MountVolume.SetUp failed for volume "github" : failed to exec 'git clone -- https://github.com/vfarcic/go-demo-2.git .': : executab  
│ le file not found in $PATH   
```





