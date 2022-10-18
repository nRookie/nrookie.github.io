



show taint of a node 

``` shell
kubectl describe node instance-4	
```



 



delete a taint

``` shell
kubectl taint instance-4  key1=example-key:NoSchedule-
```





add a taint

``` shell
kubectl taint instance-4  key1=example-key:NoSchedule
```



 