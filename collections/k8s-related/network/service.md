## Defining a Service

A Service in Kubernetes is a REST object, similar to a Pod. Like all of the REST objects, you can `POST` a Service definition to the API server to create a new instance. The name of a Service object must be a valid [RFC 1035 label name](https://kubernetes.io/docs/concepts/overview/working-with-objects/names#rfc-1035-label-names).

For example, suppose you have a set of Pods where each listens on TCP port 9376 and contains a label `app=MyApp`:



```yaml
apiVersion: v1
kind: Service
metadata:
  name: my-service
spec:
  selector:
    app: MyApp
  ports:
    - protocol: TCP
      port: 80
      targetPort: 9376
```



This specification creates a new Service object named "my-service", which targets TCP port 9376 on any Pod with the `app=MyApp` label.



Kubernetes assigns this Service an IP address (sometimes called the "cluster IP"), which is used by the Service proxies (see [Virtual IPs and service proxies](https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies) below).

The controller for the Service selector continuously scans for Pods that match its selector, and then POSTs any updates to an Endpoint object also named "my-service".

**Note:** A Service can map *any* incoming `port` to a `targetPort`. By default and for convenience, the `targetPort` is set to the same value as the `port` field.



## Virtual IPs and service proxies



Every node in a Kubernetes cluster runs a `kube-proxy`. `kube-proxy` is responsible for implementing a form of virtual IP for `Services` of type other than [`ExternalName`](https://kubernetes.io/docs/concepts/services-networking/service/#externalname).



### Configuration

Note that the kube-proxy starts up in different modes, which are determined by its configuration.

- The kube-proxy's configuration is done via a ConfigMap, and the ConfigMap for kube-proxy effectively deprecates the behaviour for almost all of the flags for the kube-proxy.
- The ConfigMap for the kube-proxy does not support live reloading of configuration.
- The ConfigMap parameters for the kube-proxy cannot all be validated and verified on startup. For example, if your operating system doesn't allow you to run iptables commands, the standard kernel kube-proxy implementation will not work. Likewise, if you have an operating system which doesn't support `netsh`, it will not run in Windows userspace mode.

### User space proxy mode

In this (legacy) mode, kube-proxy watches the Kubernetes control plane for the addition and removal of Service and Endpoint objects. For each Service it opens a port (randomly chosen) on the local node. Any connections to this "proxy port" are proxied to one of the Service's backend Pods (as reported via Endpoints). kube-proxy takes the `SessionAffinity` setting of the Service into account when deciding which backend Pod to use.

Lastly, the user-space proxy installs iptables rules which capture traffic to the Service's `clusterIP` (which is virtual) and `port`. The rules redirect that traffic to the proxy port which proxies the backend Pod.

By default, kube-proxy in userspace mode chooses a backend via a round-robin algorithm.

![image-20220312101807132](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/network/image-20220312101807132.png)



### `iptables` proxy mode



In this mode, kube-proxy watches the Kubernetes control plane for the addition and removal of Service and Endpoint objects. For each Service, it installs iptables rules, which capture traffic to the Service's `clusterIP` and `port`, and redirect that traffic to one of the Service's backend sets. For each Endpoint object, it installs iptables rules which select a backend Pod.

By default, kube-proxy in iptables mode chooses a backend at random.

Using iptables to handle traffic has a lower system overhead, because traffic is handled by Linux netfilter without the need to switch between userspace and the kernel space. This approach is also likely to be more reliable.

If kube-proxy is running in iptables mode and the first Pod that's selected does not respond, the connection fails. This is different from userspace mode: in that scenario, kube-proxy would detect that the connection to the first Pod had failed and would automatically retry with a different backend Pod.

You can use Pod [readiness probes](https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#container-probes) to verify that backend Pods are working OK, so that kube-proxy in iptables mode only sees backends that test out as healthy. Doing this means you avoid having traffic sent via kube-proxy to a Pod that's known to have failed.





![image-20220312102055378](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/network/image-20220312102055378.png)



...



## Publishing Services (ServiceTypes)

For some parts of your application (for example, frontends) you may want to expose a Service onto an external IP address, that's outside of your cluster.

Kubernetes `ServiceTypes` allow you to specify what kind of Service you want. The default is `ClusterIP`.

`Type` values and their behaviors are:

- `ClusterIP`: Exposes the Service on a cluster-internal IP. Choosing this value makes the Service only reachable from within the cluster. This is the default `ServiceType`.
- [`NodePort`](https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport): Exposes the Service on each Node's IP at a static port (the `NodePort`). A `ClusterIP` Service, to which the `NodePort` Service routes, is automatically created. You'll be able to contact the `NodePort` Service, from outside the cluster, by requesting `<NodeIP>:<NodePort>`.
- [`LoadBalancer`](https://kubernetes.io/docs/concepts/services-networking/service/#loadbalancer): Exposes the Service externally using a cloud provider's load balancer. `NodePort` and `ClusterIP` Services, to which the external load balancer routes, are automatically created.
- [`ExternalName`](https://kubernetes.io/docs/concepts/services-networking/service/#externalname): Maps the Service to the contents of the `externalName` field (e.g. `foo.bar.example.com`), by returning a `CNAME` record with its value. No proxying of any kind is set up.



You can also use [Ingress](https://kubernetes.io/docs/concepts/services-networking/ingress/) to expose your Service. Ingress is not a Service type, but it acts as the entry point for your cluster. It lets you consolidate your routing rules into a single resource as it can expose multiple services under the same IP address.



