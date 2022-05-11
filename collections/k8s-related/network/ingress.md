## Why Use Ingress Objects?[#](https://www.educative.io/module/lesson/a-practical-guide-to-kubernetes/N732V7qGWo2#Why-Use-Ingress-Objects?)

Applications that are not accessible to users are useless. Kubernetes Services provide accessibility with a usability cost. Each application can be reached through a different port. We **cannot** expect users to know the port of each service in our cluster.



Ingress objects manage external access to the applications running inside a Kubernetes cluster.



While, at first glance, it might seem that we already accomplished that through Kubernetes Services, they do not make the applications truly accessible. We still need forwarding rules based on paths and domains, SSL termination and a number of other features.

In a more traditional setup, we’d probably use an external proxy and a load balancer. Ingress provides an API that allows us to accomplish these things, in addition to a few other features we expect from a dynamic cluster.

We’ll explore the problems and the solutions through examples. For now, we need first to create a cluster.





## Why Ingress Controllers are Required?[#](https://www.educative.io/module/lesson/a-practical-guide-to-kubernetes/YMm9noR2nVM#Why-Ingress-Controllers-are-Required?)

We need a mechanism that will accept requests on pre-defined ports (e.g., `80` and `443`) and forward them to Kubernetes Services. It should be able to distinguish requests based on paths and domains as well as to be able to perform SSL offloading.

Kubernetes itself does not have a ready-to-go solution for this. Unlike other types of Controllers that are typically part of the `kube-controller-manager` binary, Ingress Controller needs to be installed separately. Instead of a Controller, `kube-controller-manager` offers *Ingress resource* that other third-party solutions can utilize to provide requests forwarding and SSL features. In other words, Kubernetes only provides an *API*, and we need to set up a Controller that will use it.

Fortunately, the community already built a myriad of Ingress Controllers. We won’t evaluate all of the available options since that would require a lot of space, and it would mostly depend on your needs and your hosting vendor. Instead, we’ll explore how Ingress Controllers work through the one that is already available in Minikube.





### Enable ingress controller



``` shell
[root@10-23-75-240 k8s-specs]# helm install nginx-ingress-controller bitnami/nginx-ingress-controller -n network
```







#### create resource



``` shell
kubectl create \
    -f ingress/go-demo-2-ingress.yml
    
    
[root@10-23-75-240 k8s-specs]# kubectl get \
>     -f ingress/go-demo-2-ingress.yml
NAME        CLASS    HOSTS   ADDRESS         PORTS   AGE
go-demo-2   <none>   *       10.23.184.141   80      17s
```









## Break down



Let’s see, through a sequence diagram, what happened when we created the Ingress resource.

1. The Kubernetes client (`kubectl`) sent a request to the API server requesting the creation of the Ingress resource defined in the `ingress/go-demo-2.yml` file.
2. The ingress controller is watching the API server for new events. It detected that there is a new Ingress resource.
3. The ingress controller configured the load balancer. In this case, it is nginx which was enabled by `minikube addons enable ingress` command. It modified `nginx.conf` with the values of all `go-demo-2-api` endpoints.

![image-20220313153441602](/Users/user/playground/share/nrookie.github.io/collections/k8s-related/network/image-20220313153441602.png)







Ingress is a (kind of) Service that runs on all nodes of a cluster. A user can send requests to any and, as long as they match one of the rules, they will be forwarded to the appropriate Service.





![image-20220313153635314](/Users/user/playground/share/nrookie.github.io/collections/k8s-related/network/image-20220313153635314.png)



![image-20220313154135475](/Users/user/playground/share/nrookie.github.io/collections/k8s-related/network/image-20220313154135475.png)