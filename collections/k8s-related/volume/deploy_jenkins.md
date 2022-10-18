

```yaml
[root@10-23-75-240 k8s-specs]# cat volume/jenkins.yml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: jenkins
  annotations:
    kubernetes.io/ingress.class: "nginx"
    ingress.kubernetes.io/ssl-redirect: "false"
    nginx.ingress.kubernetes.io/ssl-redirect: "false"
spec:
  rules:
  - http:
      paths:
      - path: /jenkins
        pathType: ImplementationSpecific
        backend:
          service:
            name: jenkins
            port:
              number: 8080

---

apiVersion: apps/v1
kind: Deployment
metadata:
  name: jenkins
spec:
  selector:
    matchLabels:
      type: master
      service: jenkins
  strategy:
    type: Recreate
  template:
    metadata:
      labels:
        type: master
        service: jenkins
    spec:
      containers:
      - name: jenkins
        image: vfarcic/jenkins
        env:
        - name: JENKINS_OPTS
          value: --prefix=/jenkins

---

apiVersion: v1
kind: Service
metadata:
  name: jenkins
spec:
  ports:
  - port: 8080
  selector:
    type: master
    service: jenkins
```





``` shell
kubectl create \
    -f volume/jenkins.yml \
    --record --save-config
```





![image-20220316103101208](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/volume/image-20220316103101208.png)



``` shell
POD_NAME=$(kubectl get pods \
    -l service=jenkins,type=master \
    -o jsonpath="{.items[*].metadata.name}")
```



``` shell
[root@10-23-75-240 k8s-specs]# kubectl exec -it $POD_NAME -- kill 1
```

``` shell
[root@10-23-75-240 k8s-specs]# kubectl get pods
NAME                              READY   STATUS    RESTARTS      AGE
devops-toolkit-74845c5869-4f72v   1/1     Running   1 (10h ago)   36h
devops-toolkit-74845c5869-ctrkk   1/1     Running   1 (10h ago)   36h
devops-toolkit-74845c5869-mrqtl   1/1     Running   1 (10h ago)   36h
go-demo-2-api-796b4dd46c-5hhgd    1/1     Running   1 (10h ago)   36h
go-demo-2-api-796b4dd46c-dxgtw    1/1     Running   1 (10h ago)   36h
go-demo-2-api-796b4dd46c-w9cfm    1/1     Running   1 (10h ago)   36h
go-demo-2-db-844c5cff45-jsn9x     1/1     Running   1 (10h ago)   36h
jenkins-755896546-98kww           1/1     Running   2 (19s ago)   5m26s
```



We can see that a container is running. Since we killed the main process and, with it, the first container, the number of restarts was increased to one.

Let’s go back to Jenkins UI and check what happened to the job. I’m sure you already know the answer, but we’ll double check it anyways.





As expected, the job we created is gone. When Kubernetes recreated the failed container, it created a new one from the same image. Everything we generated inside the running container is no more. We reset to the initial state.



## Updating the Jenkins Deployment Definition



``` shell
cat volume/jenkins-empty-dir.yml

...
kind: Deployment
...
spec:
  ...
  template:
    ...
    spec:
      containers:
        ...
        volumeMounts:
        - mountPath: /var/jenkins_home
          name: jenkins-home
      volumes:
      - emptyDir: {}
        name: jenkins-home
...
```





We added a mount that references the `jenkins-home` Volume. The Volume type is, this time, `emptyDir`. We’ll discuss the new Volume type soon. But, before we dive into explanations, we’ll try to experience its effects.



``` shell
kubectl apply \
    -f volume/jenkins-empty-dir.yml

kubectl rollout status deploy jenkins
```



``` shell
POD_NAME=$(kubectl get pods \
    -l service=jenkins,type=master \
    -o jsonpath="{.items[*].metadata.name}")

    kubectl exec -it $POD_NAME -- kill 1

    kubectl get pods

    jenkins-65cd66b7dc-6tpr2          1/1     Running   1 (8s ago)    99s
```

## Persisting State[#](https://www.educative.io/module/lesson/a-practical-guide-to-kubernetes/39ML7WLOL0O#Persisting-State)

Finally, let’s open Jenkins’ Home screen one more time.





This time, the `test` job is there. The state of the application was preserved even when the container failed, and Kubernetes created a new one.



## The emptyDir Volume[#](https://www.educative.io/module/lesson/a-practical-guide-to-kubernetes/39ML7WLOL0O#The-emptyDir-Volume)



Now let’s talk about the `emptyDir` Volume. It is considerably different from those we explored thus far.

> An `emptyDir` Volume is created when a Pod is assigned to a node. It will exist for as long as the Pod continues running on that server.





What that means is that `emptyDir` can survive container failures. When a container crashes, a Pod is not removed from the node. Instead, Kubernetes will recreate the failed container inside the same Pod and, thus, preserve the `emptyDir` Volume. All in all, this Volume type is only partially fault-tolerant.



If `emptyDir` is not entirely fault-tolerant, you might be wondering why we are discussing it in the first place.

The `emptyDir` Volume type is closest we can get to fault-tolerant volumes without using a network drive. Since we do not have any, we had to resort to `emptyDir` as the-closest-we-can-get-to-fault-tolerant-persistence type of Volume.



As you start deploying third-party applications, you’ll discover that many of them come with the recommended YAML definition. If you pay closer attention, you’ll notice that many are using `emptyDir` Volume type. It’s not that `emptyDir` is the best choice, but that it all depends on your needs, your hosting provider, your infrastructure, and quite a few other things.



There is no one-size-fits-all type of persistent and fault-tolerant Volume type. On the other hand, `emptyDir` always works. Since it has no external dependencies, it is safe to put it as an example, with the assumption that people will change to whichever type fits them better.





There is an unwritten assumption that `emptyDir` is used for testing purposes, and will be changed to something else before it reaches production.



As long as we’re using Minikube to create a Kubernetes cluster, we’ll use `emptyDir` as a solution for persistent volumes. Do not despair. Later on, once we move into a “more serious” cluster setup, we’ll explore better options for persisting state.

------

In the next lesson, we will test your understanding of Volumes with the help of a quick quiz.





![image-20220316104507679](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/volume/image-20220316104507679.png)



The next chapter is dedicated to the `configMap` Volume type. It will, hopefully, solve a few problems and provide better solutions to some use-cases than those we employed in this chapter. ConfigMaps deserve a full chapter, so they’re getting one.
