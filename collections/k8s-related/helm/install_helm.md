https://helm.sh/docs/intro/install/



``` shell
 wget https://get.helm.sh/helm-v3.8.1-linux-amd64.tar.gz
 tar -xvzf helm-v3.8.1-linux-amd64.tar.gz 
 mv linux-amd64/helm /usr/local/bin/
```



### add repo

``` shell
helm repo add bitnami https://charts.bitnami.com/bitnami
```



### install example chart 



``` shell
helm search repo bitnami
 helm install bitnami/mysql --generate-name
```

## Learn about release



``` shell

[root@10-23-75-240 k8s]# helm list
NAME            	NAMESPACE	REVISION	UPDATED                                	STATUS  	CHART       	APP VERSION
mysql-1647094512	default  	1       	2022-03-12 22:15:15.114071124 +0800 CST	deployed	mysql-8.8.26	8.0.28 
```





## Uninstall a Release



```fallback
helm uninstall mysql-1612624192
release "mysql-1612624192" uninstalled
```