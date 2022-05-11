# Injecting Configurations from Environment Files



## Looking into the Definition[#](https://www.educative.io/module/lesson/a-practical-guide-to-kubernetes/xlBAypRr4Dn#Looking-into-the-Definition)

Let’s take a look at the `cm/my-env-file.yml` file.



``` shell
[root@10-23-75-240 k8s-specs]# cat cm/my-env-file.yml
something=else
weather=sunny

```



## Creating the ConfigMap[#](https://www.educative.io/module/lesson/a-practical-guide-to-kubernetes/xlBAypRr4Dn#Creating-the-ConfigMap)

Let’s see what happens if we create a ConfigMap using that file as the source.



We created the ConfigMap using the `--from-env-file` argument, and we retrieved the ConfigMap in `yaml` format.

The **output** of the latter command is as follows

``` shell
>     --from-env-file=cm/my-env-file.yml
configmap/my-config created
[root@10-23-75-240 k8s-specs]# 
[root@10-23-75-240 k8s-specs]# kubectl get cm my-config -o yaml
apiVersion: v1
data:
  something: else
  weather: sunny
kind: ConfigMap
metadata:
  creationTimestamp: "2022-03-16T02:55:23Z"
  name: my-config
  namespace: default
  resourceVersion: "418547"
  uid: 7e4f5189-9004-4386-9167-c44bfc8e6629
```



We can see that there are two entries, each corresponding to key/value pairs from the file. The result is the same as when we created a ConfigMap using `--from-literal` arguments. Two different sources produced the same outcome.

If we used `--from-file` argument, the result would be as follows.



 ``` yaml
 apiVersion: v1
 data:
   my-env-file.yml: |
     something=else
     weather=sunny
 kind: ConfigMap
 ...
 ```



All in all, `--from-file` reads the content of one or more files, and stores it using file names as keys. `--from-env-file`, assumes that content of a file is in key/value format, and stores each as a separate entry.

------

In the next lesson, we will explore how to convert the output of configMap into environment variables.



# Converting ConfigMap Output into Environment Variables



``` shell
[root@10-23-75-240 k8s-specs]# kubectl get cm my-config -o yaml
apiVersion: v1
data:
  something: else
  weather: sunny
kind: ConfigMap
metadata:
  creationTimestamp: "2022-03-16T02:55:23Z"
  name: my-config
  namespace: default
  resourceVersion: "418547"
  uid: 7e4f5189-9004-4386-9167-c44bfc8e6629
[root@10-23-75-240 k8s-specs]# cat cm/alpine-env.yml
apiVersion: v1
kind: Pod
metadata:
  name: alpine-env
spec:
  containers:
  - name: alpine
    image: alpine
    command: ["sleep"]
    args: ["100000"]
    env:
    - name: something
      valueFrom:
        configMapKeyRef:
          name: my-config
          key: something
    - name: weather
      valueFrom:
        configMapKeyRef:
          name: my-config
          key: weather
```





The major difference, when compared with `cm/alpine.yml`, is that `volumeMounts` and `volumes` sections are gone. This time we have an `env` section.



Instead of a `value` field, we have `valueFrom`. Further on, we declared that it should get values from a ConfigMap (`configMapKeyRef`) named `my-config`. Since that ConfigMap has multiple values, we specified the `key` as well.



### Creating the Pod[#](https://www.educative.io/module/lesson/a-practical-guide-to-kubernetes/g738zv5QRwZ#Creating-the-Pod)



``` shell
[root@10-23-75-240 k8s-specs]# kubectl create \
>     -f cm/alpine-env.yml
pod/alpine-env created
[root@10-23-75-240 k8s-specs]# kubectl exec -it alpine-env -- env
error: unable to upgrade connection: container not found ("alpine")
[root@10-23-75-240 k8s-specs]# kubectl exec -it alpine-env -- env
PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
HOSTNAME=alpine-env
TERM=xterm
something=else
weather=sunny
DEVOPS_TOOLKIT_PORT_80_TCP_ADDR=10.111.158.171
JENKINS_SERVICE_PORT=8080
JENKINS_PORT_8080_TCP_ADDR=10.102.25.68
KUBERNETES_PORT=tcp://10.96.0.1:443
KUBERNETES_PORT_443_TCP=tcp://10.96.0.1:443
GO_DEMO_2_DB_SERVICE_PORT=27017
GO_DEMO_2_DB_PORT=tcp://10.105.242.176:27017
GO_DEMO_2_DB_PORT_27017_TCP_PROTO=tcp
GO_DEMO_2_DB_PORT_27017_TCP_PORT=27017
DEVOPS_TOOLKIT_PORT_80_TCP_PORT=80
GO_DEMO_2_API_SERVICE_HOST=10.107.22.121
GO_DEMO_2_DB_PORT_27017_TCP=tcp://10.105.242.176:27017
DEVOPS_TOOLKIT_PORT_80_TCP=tcp://10.111.158.171:80
DEVOPS_TOOLKIT_PORT_80_TCP_PROTO=tcp
JENKINS_PORT=tcp://10.102.25.68:8080
KUBERNETES_SERVICE_PORT_HTTPS=443
KUBERNETES_PORT_443_TCP_PORT=443
GO_DEMO_2_DB_PORT_27017_TCP_ADDR=10.105.242.176
DEVOPS_TOOLKIT_SERVICE_HOST=10.111.158.171
KUBERNETES_PORT_443_TCP_ADDR=10.96.0.1
GO_DEMO_2_DB_SERVICE_HOST=10.105.242.176
GO_DEMO_2_API_PORT=tcp://10.107.22.121:8080
JENKINS_SERVICE_HOST=10.102.25.68
JENKINS_PORT_8080_TCP_PORT=8080
KUBERNETES_SERVICE_HOST=10.96.0.1
KUBERNETES_PORT_443_TCP_PROTO=tcp
GO_DEMO_2_API_PORT_8080_TCP=tcp://10.107.22.121:8080
GO_DEMO_2_API_PORT_8080_TCP_PROTO=tcp
JENKINS_PORT_8080_TCP=tcp://10.102.25.68:8080
JENKINS_PORT_8080_TCP_PROTO=tcp
DEVOPS_TOOLKIT_SERVICE_PORT=80
DEVOPS_TOOLKIT_PORT=tcp://10.111.158.171:80
GO_DEMO_2_API_SERVICE_PORT=8080
GO_DEMO_2_API_PORT_8080_TCP_PORT=8080
KUBERNETES_SERVICE_PORT=443
GO_DEMO_2_API_PORT_8080_TCP_ADDR=10.107.22.121
HOME=/root

```



There’s another, often more useful way to specify environment variables from a ConfigMap. Before we try it, we’ll remove the currently running Pod.



``` shell
kubectl delete \
    -f cm/alpine-env.yml
```



``` shell
[root@10-23-75-240 k8s-specs]# cat cm/alpine-env-all.yml
apiVersion: v1
kind: Pod
metadata:
  name: alpine-env
spec:
  containers:
  - name: alpine
    image: alpine
    command: ["sleep"]
    args: ["100000"]
    envFrom:
    - configMapRef:
        name: my-config

```

The difference is only in the way environment variables are defined.



This time, the syntax is much shorter. We have `envFrom`, instead of the `env` section. It can be either `configMapRef` or `secretRef`. Since we did not yet explore Secrets, we’ll stick with the prior. Inside `configMapRef` is the `name` reference to the `my-config` ConfigMap.



``` shell
[root@10-23-75-240 k8s-specs]# kubectl create \
>     -f cm/alpine-env-all.yml
pod/alpine-env created
[root@10-23-75-240 k8s-specs]# kubectl exec -it alpine-env -- env
PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
HOSTNAME=alpine-env
TERM=xterm
weather=sunny
something=else
GO_DEMO_2_DB_PORT_27017_TCP_ADDR=10.105.242.176
GO_DEMO_2_API_SERVICE_PORT=8080
GO_DEMO_2_API_PORT=tcp://10.107.22.121:8080
JENKINS_PORT_8080_TCP=tcp://10.102.25.68:8080
JENKINS_PORT_8080_TCP_ADDR=10.102.25.68
KUBERNETES_SERVICE_PORT_HTTPS=443
KUBERNETES_PORT_443_TCP=tcp://10.96.0.1:443
GO_DEMO_2_DB_SERVICE_HOST=10.105.242.176
JENKINS_PORT=tcp://10.102.25.68:8080
DEVOPS_TOOLKIT_SERVICE_PORT=80
DEVOPS_TOOLKIT_PORT_80_TCP_PORT=80
KUBERNETES_PORT_443_TCP_PORT=443
KUBERNETES_PORT_443_TCP_ADDR=10.96.0.1
GO_DEMO_2_DB_PORT=tcp://10.105.242.176:27017
GO_DEMO_2_API_SERVICE_HOST=10.107.22.121
GO_DEMO_2_API_PORT_8080_TCP=tcp://10.107.22.121:8080
GO_DEMO_2_API_PORT_8080_TCP_ADDR=10.107.22.121
JENKINS_PORT_8080_TCP_PORT=8080
KUBERNETES_SERVICE_PORT=443
KUBERNETES_PORT=tcp://10.96.0.1:443
KUBERNETES_PORT_443_TCP_PROTO=tcp
GO_DEMO_2_DB_SERVICE_PORT=27017
GO_DEMO_2_API_PORT_8080_TCP_PROTO=tcp
JENKINS_PORT_8080_TCP_PROTO=tcp
GO_DEMO_2_DB_PORT_27017_TCP=tcp://10.105.242.176:27017
JENKINS_SERVICE_HOST=10.102.25.68
DEVOPS_TOOLKIT_PORT=tcp://10.111.158.171:80
GO_DEMO_2_DB_PORT_27017_TCP_PROTO=tcp
DEVOPS_TOOLKIT_SERVICE_HOST=10.111.158.171
DEVOPS_TOOLKIT_PORT_80_TCP=tcp://10.111.158.171:80
GO_DEMO_2_API_PORT_8080_TCP_PORT=8080
JENKINS_SERVICE_PORT=8080
KUBERNETES_SERVICE_HOST=10.96.0.1
GO_DEMO_2_DB_PORT_27017_TCP_PORT=27017
DEVOPS_TOOLKIT_PORT_80_TCP_PROTO=tcp
DEVOPS_TOOLKIT_PORT_80_TCP_ADDR=10.111.158.171
HOME=/root

```





We created the Pod and retrieved all the environment variables from inside its only container. The output of the latter command, limited to the relevant parts, is as follows.



The result is the **same** as before. The difference is only in the way we define environment variables.



With `env.valueFrom.configMapKeyRef` syntax, we need to specify each ConfigMap key separately. That gives us control over the scope and the relation with the names of container variables.

The `envFrom.configMapRef` converts all ConfigMap’s data into environment variables. That is often a better and simpler option if you don’t need to use different names between ConfigMap and environment variable keys. The syntax is short, and we don’t need to worry whether we forgot to include one of the ConfigMap’s keys.