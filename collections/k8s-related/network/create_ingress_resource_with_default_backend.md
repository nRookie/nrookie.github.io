### Non-Matching Request



In some cases, we might want to define a default backend. We might want to forward requests that do not match any of the Ingress rules.







Let's take a look at an example.



So far, we have two sets of Ingress rules in our cluster. One accepts all requests with the base path `/demo`. The other forwards all requests coming from the `devopstoolkitseries.com` domain. The request we just sent does not match either of those rules, so the response was once again `404 Not Found`.



## Default Backend Ingress Resource[#](https://www.educative.io/module/lesson/a-practical-guide-to-kubernetes/R1xp45NOorw#Default-Backend-Ingress-Resource)



### Looking into the Definition[#](https://www.educative.io/module/lesson/a-practical-guide-to-kubernetes/R1xp45NOorw#Looking-into-the-Definition)

Let’s imagine that it would be a good idea to forward all requests with the wrong domain to the `devops-toolkit` application. Of course, by “wrong domain”, I mean one of the domains we own, and not one of those that are already included in Ingress rules.



There’s no Deployment, nor is there a Service. This time, we’re creating only an Ingress resource.

The `spec` has no rules, but only a single `backend`.

> When an Ingress `spec` is without rules, it is considered a default backend. As such, it will forward all requests that do not match paths and/or domains set as rules in the other Ingress resources.



We can use the default backend as a substitute for the default `404` pages or for any other occasion that is not covered by other rules.

You’ll notice that the `serviceName` is `devops-toolkit`. The example would be much better if we created a separate application for this purpose but it does not matter for this example. All we want, at the moment, is to see something other than `404 Not Found` response.



``` shell
[root@10-23-75-240 k8s-specs]# kubectl create \
>     -f ingress/default-backend.yml
ingress.networking.k8s.io/default created
[root@10-23-75-240 k8s-specs]# curl -I -H "Host: acme.com" \
>     "http://$IP"
HTTP/1.1 200 OK
Date: Mon, 14 Mar 2022 14:36:47 GMT
Content-Type: text/html
Content-Length: 5316
Connection: keep-alive
Last-Modified: Mon, 07 Mar 2022 10:28:11 GMT
ETag: "6225de3b-14c4"
Accept-Ranges: bytes

```



We explored some of the essential functions of Ingress resources and Controllers. To be more concrete, we examined almost all those that are defined in the Ingress API.

One notable feature we did not explore is TLS configuration. Without it, our services cannot serve HTTPS requests. To enable it, we’d need to configure Ingress to offload SSL certificates.

There are two reasons we did not explore TLS. For one, we do not have a valid SSL certificate. On top of that, we did not yet study Kubernetes Secrets. We’d suggest you to explore SSL setup yourself once you make a decision which Ingress controller to use. Secrets, on the other hand, will be explained soon.

We’ll explore other Ingress Controllers once we move our cluster to “real” servers that we’ll create with one of the hosting vendors. Until then, you might benefit from reading [NGINX Ingress Controller](https://github.com/kubernetes/ingress-nginx/blob/master/README.md) documentation in more detail. Specifically, I suggest you pay close attention to its [annotations](https://github.com/kubernetes/ingress-nginx/blob/master/docs/user-guide/nginx-configuration/annotations.md).



![image-20220314224147300](/Users/user/playground/share/nrookie.github.io/collections/k8s-related/network/image-20220314224147300.png)



## Destroying Everything[#](https://www.educative.io/module/lesson/a-practical-guide-to-kubernetes/gx5oNAoqRKk#Destroying-Everything)

Now that another chapter is finished, we’ll destroy the cluster and let your machine rest for a while. It deserves a break.



