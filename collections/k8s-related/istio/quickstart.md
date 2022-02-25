## curl by apigateway and 
``` shell
[qing.na@JumperServer-Shanghai istio]$ curl -g http://apigateway.prj-uhost-ussg.svc.c3.uae:12344/helloworld
Hello World
[qing.na@JumperServer-Shanghai istio]$ curl -g http://helloworld.prj-uhost-ussg.svc.c3.uae:12345
Hello World
```



### download istioctl


``` shell
curl -L https://git.io/getLatestIstio | sh -
```



### install istioctl

``` shell
install istio-1.13.1/bin/istioctl /usr/local/bin
```

### istio command list

https://istio.io/latest/docs/reference/commands/istioctl/




### enable istio


``` shell
kubectl apply -f <(istioctl kube-inject -f basic.yaml)
```


## update a new version


