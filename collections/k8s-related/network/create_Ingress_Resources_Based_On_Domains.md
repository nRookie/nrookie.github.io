## Refactoring the definition





``` shell
cat ingress/devops-toolkit-dom.yml

```

When compared with the previous definition, the **only difference** is in the additional entry `host: devopstoolkitseries.com`. Since that will be the only application accessible through that domain, we also removed the `path: /` entry.



``` shell
[root@10-23-75-240 k8s-specs]# kubectl apply \
>   -f ingress/devops-toolkit-dom.yml \
>   --record
Flag --record has been deprecated, --record will be removed in the future
ingress.networking.k8s.io/devops-toolkit configured
deployment.apps/devops-toolkit configured
service/devops-toolkit configured
[root@10-23-75-240 k8s-specs]# curl 10.104.194.180
<html>
<head><title>404 Not Found</title></head>
<body>
<center><h1>404 Not Found</h1></center>
<hr><center>nginx</center>
</body>
</html>
```



There is **no** Ingress resource defined to listen to `/`. The updated Ingress will forward requests only if they come from `devopstoolkitseries.com`.

Since it’s not feasible to give you access to the DNS registry of `devopstoolkitseries.com`. So you cannot configure it with the IP of your Minikube cluster. Therefore, we won’t be able to test it by sending a request to `devopstoolkitseries.com`.

What we can do is to “fake” it by adding that domain to the request header.







``` shell
[root@10-23-75-240 k8s-specs]# curl -I \
> -H "Host: devopstoolkitseries.com" \
> 10.104.194.180
HTTP/1.1 200 OK
Date: Mon, 14 Mar 2022 14:31:03 GMT
Content-Type: text/html
Content-Length: 5316
Connection: keep-alive
Last-Modified: Mon, 07 Mar 2022 10:28:11 GMT
ETag: "6225de3b-14c4"
Accept-Ranges: bytes


```





Now that Ingress received a request that looks like it’s coming from the domain `devopstoolkitseries.com`, it forwarded it to the `devops-toolkit` Service which, in turn, load balanced it to one of the `devops-toolkit` Pods. As a result, we got the response `200 OK`.

Just to be on the safe side, we’ll verify whether `go-demo-2` Ingress still works.



``` shell
[root@10-23-75-240 k8s-specs]# curl -H "Host: acme.com" \
>     "http://$IP/demo/hello"
hello, world!
```



We got the famous `hello, world!` response, thus confirming that both Ingress resources are operational. Even though we “faked” the last request as if it’s coming from `acme.com`, it still worked. Since the `go-demo-2` Ingress does not have any `host` defined, it accepts any request with the `path` starting with `/demo`.

------

We’re still missing a few things. One of those is a setup of a default backend. We’ll go through it in the next lesson.





