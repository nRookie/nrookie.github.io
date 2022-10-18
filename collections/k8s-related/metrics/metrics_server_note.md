![image-20220313135500197](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/metrics/image-20220313135500197.png)



# What is `Metrics Server`? [#](https://www.educative.io/module/lesson/kubernetes-monitoring-logging-auto-scaling/B87wynEVABk#What-is-Metrics-Server?-)

A simple explanation is that it collects information about used resources (memory and CPU) of nodes and Pods. It does not store metrics, so do not think that you can use it to retrieve historical values and predict tendencies. There are other tools for that, and we’ll explore them later. Instead, `Metrics Server's` goal is to provide an API that can be used to retrieve current resource usage. We can use that API through `kubectl` or by sending direct requests with, let’s say, `curl`. In other words, the `Metrics Server` collects cluster-wide metrics and allows us to retrieve them through its API. That, by itself, is very powerful, but it is only part of the story.

I already mentioned extensibility. We can extend the `Metrics Server` to collect metrics from other sources. We’ll get there in due time. For now, we’ll explore what it provides out of the box and how it interacts with some other Kubernetes resources that will help us make our Pods scalable and more resilient.



## Flow of data in `Metrics Server` [#](https://www.educative.io/module/lesson/kubernetes-monitoring-logging-auto-scaling/B87wynEVABk#Flow-of-data-in-Metrics-Server-)

`Metrics Server` will periodically fetch metrics from Kubeletes running on the nodes. Those metrics, for now, contain memory and CPU utilization of the Pods and the nodes. Other entities can request data from the `Metrics Server` through the API Server which has the **Master Metrics API**. An example of those entities is the Scheduler that, once `Metrics Server` is installed, uses its data to make decisions. As you will see soon, the usage of the `Metrics Server` goes beyond the Scheduler but, for now, the explanation should provide an image of the basic flow of data.



![image-20220313135736522](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/metrics/image-20220313135736522.png)





# From Heapster to `Metrics Server` [#](https://www.educative.io/module/lesson/kubernetes-monitoring-logging-auto-scaling/B87wynEVABk#From-Heapster-to-Metrics-Server-)

The critical element in scaling Pods is the **Kubernetes Metrics Server**. You might consider yourself a Kubernetes ninja and yet never heard of the Metrics Server. Don’t be ashamed if that’s the case. You’re not the only one.

If you started observing Kubernetes metrics, you might have used **Heapster**. It’s been around for a long time, and you likely have it running in your cluster, even if you don’t know what it is. Both the Metrics server and Heapster serve the same purpose, with one being deprecated for a while, so let’s clarify things a bit.

Early on, Kubernetes introduced **Heapster** as a tool that enables Container Cluster Monitoring and Performance Analysis for Kubernetes. It’s been around since Kubernetes version 1.0.6. You can say that **Heapster** has been part of Kubernetes’ life since its toddler age. It collects and interprets various metrics like resource usage, events, and so on. **Heapster** has been an integral part of Kubernetes and enabled it to schedule Pods appropriately. Without it, Kubernetes would be blind. It would not know which node has available memory, which Pod is using too much CPU, and so on. But, just as with most other tools that become available early, its design was a “failed experiment”. As Kubernetes continued growing, we (the community around Kubernetes) started realizing that a new, better, and, more importantly, a more extensible design is required. Hence, the `Metrics Server` was born. Right now, even though **Heapster** is still in use, it is considered deprecated, even though today the `Metrics Server` is still in beta state.





## `Metrics Server` for machines [#](https://www.educative.io/module/lesson/kubernetes-monitoring-logging-auto-scaling/m72Y9G9PZmG#Metrics-Server-for-machines-)

There are two important things to note. First of all, it provides current (or short-term) memory and CPU utilization of the containers running inside a cluster. The second and more important note is that we will not use it directly. `Metrics Server` was not designed for humans but for machines. We’ll get there later. For now, remember that there is a thing called `Metrics Server` and that you should not use it directly (once you adopt a tool that will scrape its metrics).

------

Now that we explored `Metrics Server`, we’ll try to put it to good use and learn how to auto-scale our Pods based on resource utilization, in the next lesson.
