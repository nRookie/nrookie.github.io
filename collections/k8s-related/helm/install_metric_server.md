``` shell
helm install my-release bitnami/metrics-server --namespace metrics

[root@10-23-75-240 k8s]# kubectl -n metrics \
>     rollout status \
>     deployment my-release-metrics-server
Waiting for deployment "my-release-metrics-server" rollout to finish: 0 of 1 updated replicas are available...
```

``` shell
By default 'rollout status' will watch the status of the latest rollout until it's done. If you don't want to wait for the rollout to finish then you can use --watch=false. Note that if a new rollout starts in-between, then 'rollout status' will continue watching the latest revision.
```



change the metrics deployment yaml to

``` yaml
    spec:
      containers:
      - args:
        - --cert-dir=/tmp
        - --secure-port=4443
        - --kubelet-insecure-tls=true
        - --kubelet-preferred-address-types=InternalIP
```





``` shell
│   Normal   Scheduled         6m12s                  default-scheduler  Successfully assigned metrics/my-release-metrics-server-7876fc7996-q8mvv to 10-23-245-35                                   
│   Normal   Pulled            6m11s                  kubelet            Container image "docker.io/bitnami/metrics-server:0.6.1-debian-10-r17" already present on machine                          
│   Normal   Created           6m11s                  kubelet            Created container metrics-server                                                                                           
│   Normal   Started           6m11s                  kubelet            Started container metrics-server                                                                                           
│   Warning  Unhealthy         6m11s                  kubelet            Readiness probe failed: Get "https://10.23.245.35:8443/readyz": dial tcp 10.23.245.35:8443: connect: connection refused    
│   Warning  Unhealthy         5m12s (x8 over 6m10s)  kubelet            Readiness probe failed: HTTP probe failed with statuscode: 500 
```



https://forum.linuxfoundation.org/discussion/860480/exercise-13-3-metrics-server-pod-running-but-not-ready



change again

``` yaml
        - --cert-dir=/tmp
        - --secure-port=8443
        - --kubelet-insecure-tls=true
        - --kubelet-preferred-address-types=InternalIP,ExternalIP,Hostname
        - --kubelet-use-node-status-port
        - --metric-resolution=15s
```



``` shell
│   Type     Reason     Age              From               Message                                                                                                                                 
│   ----     ------     ----             ----               -------                                                                                                                                 
│   Normal   Scheduled  8s               default-scheduler  Successfully assigned metrics/my-release-metrics-server-b88f4598b-smvtx to 10-23-184-141                                                
│   Normal   Pulled     7s               kubelet            Container image "docker.io/bitnami/metrics-server:0.6.1-debian-10-r17" already present on machine                                       
│   Normal   Created    7s               kubelet            Created container metrics-server                                                                                                        
│   Normal   Started    7s               kubelet            Started container metrics-server                                                                                                        
│   Warning  Unhealthy  7s               kubelet            Readiness probe failed: Get "https://10.23.184.141:8443/readyz": dial tcp 10.23.184.141:8443: connect: connection refused               
│   Warning  Unhealthy  5s (x2 over 6s)  kubelet            Readiness probe failed: HTTP probe failed with statuscode: 500 
```



On 8443

``` shell
[root@10-23-184-141 ~]# netstat -an | grep 844
tcp6       0      0 :::8443                 :::*                    LISTEN     
```





## install the metrics servers again (after removing http_proxy)



``` shell
[root@10-23-75-240 k8s]# helm install metrics-server bitnami/metrics-server --namespace kube-system
NAME: metrics-server
LAST DEPLOYED: Sun Mar 13 11:48:33 2022
NAMESPACE: kube-system
STATUS: deployed
REVISION: 1
TEST SUITE: None
NOTES:
CHART NAME: metrics-server
CHART VERSION: 5.11.3
APP VERSION: 0.6.1

** Please be patient while the chart is being deployed **

The metric server has been deployed.

###################################################################################
### ERROR: The metrics.k8s.io/v1beta1 API service is not enabled in the cluster ###
###################################################################################
You have disabled the API service creation for this release. As the Kubernetes version in the cluster
does not have metrics.k8s.io/v1beta1, the metrics API will not work with this release unless:

Option A:

  You complete your metrics-server release by running:

  helm upgrade --namespace kube-system metrics-server bitnami/metrics-server \
    --set apiService.create=true

Option B:

   You configure the metrics API service outside of this Helm chart
[root@10-23-75-240 k8s]# 

```



add these after container args

``` yaml
        - --cert-dir=/tmp
        - --secure-port=8443
        - --kubelet-insecure-tls=true
        - --kubelet-preferred-address-types=InternalIP,ExternalIP,Hostname
        - --kubelet-use-node-status-port
        - --metric-resolution=15s
```



still not working ,jessus

https://github.com/kubernetes-sigs/metrics-server/issues/247

``` note
Edit: this report definitely does not belong here, sorry for contributing to the noise

I was able to resolve it by switching from the bitnami helm chart for metrics server, to kustomized deploy of metrics-server from this repo, with very similar to "test" kustomize manifests. Thank you for providing this.

I am on kubeadm v1.20.2 with a matching kubectl and I had to set apiService.create: true as @debu99 suggested.

This is in conflict with the docs in values.yaml which are maybe incorrect

## API service parameters
##
apiService:
  ## Specifies whether the v1beta1.metrics.k8s.io API service should be created
  ## This should not be necessary in k8s version >= 1.8, but depends on vendors and cloud providers.
  ##
  create: false
Else I ran into error: Metrics API not available

While that's not the subject of this issue report, this issue is one of the top results for "error: Metrics API not available" and it helped me, so I am highlighting it here.

I am not sure if this information belongs here, I'm using bitnami/metrics-server which has a repo of its own in https://github.com/bitnami/charts/ and so I guess the new issue report should go there, if there is a problem.
```



``` shell
if you use bitnami metrics-server, enable this even k8s > 1.18
apiService:
create: true
```



https://www.youtube.com/watch?v=WVxK1k_blPQ





``` shell
helm show values bitnami/metrics-server >/tmp/metrics-server.values

```

**important change as follows**



1. hostnetwork = true
2. api.create= true

``` shell
vi /tmp/metric-server.values
[root@10-23-75-240 metrics-server]# cat /tmp/metric-server.values 
## @section Global parameters
## Global Docker image parameters
## Please, note that this will override the image parameters, including dependencies, configured to use the global value
## Current available global Docker image parameters: imageRegistry, imagePullSecrets and storageClass

## @param global.imageRegistry Global Docker image registry
## @param global.imagePullSecrets Global Docker registry secret names as an array
##
global:
  imageRegistry: ""
  ## E.g.
  ## imagePullSecrets:
  ##   - myRegistryKeySecretName
  ##
  imagePullSecrets: []

## @section Common parameters

## @param nameOverride String to partially override common.names.fullname template (will maintain the release name)
##
nameOverride: ""
## @param fullnameOverride String to fully override common.names.fullname template
##
fullnameOverride: ""
## @param commonLabels Add labels to all the deployed resources
##
commonLabels: {}
## @param commonAnnotations Add annotations to all the deployed resources
##
commonAnnotations: {}
## @param extraDeploy Array of extra objects to deploy with the release
##
extraDeploy: []

## @section Metrics Server parameters

## Bitnami Metrics Server image version
## ref: https://hub.docker.com/r/bitnami/metrics-server/tags/
## @param image.registry Metrics Server image registry
## @param image.repository Metrics Server image repository
## @param image.tag Metrics Server image tag (immutable tags are recommended)
## @param image.pullPolicy Metrics Server image pull policy
## @param image.pullSecrets Metrics Server image pull secrets
##
image:
  registry: docker.io
  repository: bitnami/metrics-server
  tag: 0.6.1-debian-10-r17
  ## Specify a imagePullPolicy
  ## Defaults to 'Always' if image tag is 'latest', else set to 'IfNotPresent'
  ## ref: https://kubernetes.io/docs/user-guide/images/#pre-pulling-images
  ##
  pullPolicy: IfNotPresent
  ## Optionally specify an array of imagePullSecrets.
  ## Secrets must be manually created in the namespace.
  ## ref: https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/
  ## e.g:
  ## pullSecrets:
  ##   - myRegistryKeySecretName
  ##
  pullSecrets: []

## @param hostAliases Add deployment host aliases
## https://kubernetes.io/docs/concepts/services-networking/add-entries-to-pod-etc-hosts-with-host-aliases/
##
hostAliases: []
## @param replicas Number of metrics-server nodes to deploy
##
replicas: 1
## @param updateStrategy.type Set up update strategy for metrics-server installation.
## Set to Recreate if you use persistent volume that cannot be mounted by more than one pods to make sure the pods is destroyed first.
## ref: https://kubernetes.io/docs/concepts/workloads/controllers/deployment/#strategy
## Example:
## updateStrategy:
##  type: RollingUpdate
##  rollingUpdate:
##    maxSurge: 25%
##    maxUnavailable: 25%
##
updateStrategy:
  type: RollingUpdate
## Role Based Access
## ref: https://kubernetes.io/docs/admin/authorization/rbac/
##
rbac:
  ## @param rbac.create Enable RBAC authentication
  ##
  create: true
## Pods Service Account
## ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/
##
serviceAccount:
  ## @param serviceAccount.create Specifies whether a ServiceAccount should be created
  ##
  create: true
  ## @param serviceAccount.name The name of the ServiceAccount to create
  ## If not set and create is true, a name is generated using the common.names.fullname template
  name: ""
  ## @param serviceAccount.automountServiceAccountToken Automount API credentials for a service account
  ## ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/#use-the-default-service-account-to-access-the-api-server
  ##
  automountServiceAccountToken: true
## API service parameters
##
apiService:
  ## @param apiService.create Specifies whether the v1beta1.metrics.k8s.io API service should be created. You can check if it is needed with `kubectl get --raw "/apis/metrics.k8s.io/v1beta1/nodes"`.
  ## This is still necessary up to at least k8s version >= 1.21, but depends on vendors and cloud providers.
  ##
  create: true
  ## @param apiService.insecureSkipTLSVerify Specifies whether to skip self-verifying self-signed TLS certificates. Set to "false" if you are providing your own certificates.
  ## Note that "false" MUST be in quotation marks (cf. https://github.com/helm/helm/issues/3308), since false without quotation marks will render to true
  insecureSkipTLSVerify: true
  ## @param apiService.caBundle A base64-encoded string of concatenated certificates for the CA chain for the APIService.
  caBundle: ""
## @param securePort Port where metrics-server will be running
##
securePort: 8443
## @param hostNetwork Enable hostNetwork mode
## You would require this enabled if you use alternate overlay networking for pods and
## API server unable to communicate with metrics-server. As an example, this is required
## if you use Weave network on EKS
##
hostNetwork: true
## @param dnsPolicy Default dnsPolicy setting
## If you enable hostNetwork then you may need to set your dnsPolicy to something other
## than "ClusterFirst" depending on your requirements.
dnsPolicy: "ClusterFirst"
## @param command Override default container command (useful when using custom images)
##
command: ["metrics-server"]
## @param extraArgs Extra arguments to pass to metrics-server on start up
## ref: https://github.com/kubernetes-incubator/metrics-server/blob/master/README.md#flags
##
## extraArgs:
##   kubelet-insecure-tls: true
##   kubelet-preferred-address-types: InternalIP
##
extraArgs: {}
## @param podLabels Pod labels
## ref: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/
##
podLabels: {}
## @param podAnnotations Pod annotations
## ref: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/
##
podAnnotations: {}
## @param priorityClassName Priority class for pod scheduling
## ref: https://kubernetes.io/docs/concepts/configuration/pod-priority-preemption/#priorityclass
priorityClassName: ""
## @param podAffinityPreset Pod affinity preset. Ignored if `affinity` is set. Allowed values: `soft` or `hard`
## ref: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#inter-pod-affinity-and-anti-affinity
##
podAffinityPreset: ""
## @param podAntiAffinityPreset Pod anti-affinity preset. Ignored if `affinity` is set. Allowed values: `soft` or `hard`
## ref: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#inter-pod-affinity-and-anti-affinity
##
podAntiAffinityPreset: soft
## Pod disruption budget
## ref: https://kubernetes.io/docs/concepts/workloads/pods/disruptions/#pod-disruption-budgets
## @param podDisruptionBudget.enabled Create a PodDisruptionBudget
## @param podDisruptionBudget.minAvailable Minimum available instances
## @param podDisruptionBudget.maxUnavailable Maximum unavailable instances
##
podDisruptionBudget:
  enabled: false
  minAvailable: ""
  maxUnavailable: ""
## Node affinity preset
## ref: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#node-affinity
##
nodeAffinityPreset:
  ## @param nodeAffinityPreset.type Node affinity preset type. Ignored if `affinity` is set. Allowed values: `soft` or `hard`
  ##
  type: ""
  ## @param nodeAffinityPreset.key Node label key to match. Ignored if `affinity` is set.
  ## E.g.
  ## key: "kubernetes.io/e2e-az-name"
  ##
  key: ""
  ## @param nodeAffinityPreset.values Node label values to match. Ignored if `affinity` is set.
  ## E.g.
  ## values:
  ##   - e2e-az1
  ##   - e2e-az2
  ##
  values: []
## @param affinity Affinity for pod assignment
## ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#affinity-and-anti-affinity
## Note: podAffinityPreset, podAntiAffinityPreset, and nodeAffinityPreset will be ignored when it's set
##
affinity: {}
## @param topologySpreadConstraints Topology spread constraints for pod
## ref: https://kubernetes.io/docs/concepts/workloads/pods/pod-topology-spread-constraints
##
topologySpreadConstraints: []
## @param nodeSelector Node labels for pod assignment
## ref: https://kubernetes.io/docs/user-guide/node-selection/
##
nodeSelector: {}
## @param tolerations Tolerations for pod assignment
## ref: https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/
##
tolerations: []
##  Metrics Server K8s svc properties
##
service:
  ## @param service.type Kubernetes Service type
  ##
  type: ClusterIP
  ## @param service.port Kubernetes Service port
  ##
  port: 443
  ## @param service.nodePort Kubernetes Service port
  ## ref: https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport
  ## e.g:
  ## nodePort: 30001
  ##
  nodePort: ""
  ## @param service.loadBalancerIP LoadBalancer IP if Service type is `LoadBalancer`
  ## Set the LoadBalancer service type to internal only.
  ## ref: https://kubernetes.io/docs/concepts/services-networking/service/#internal-load-balancer
  ##
  loadBalancerIP: ""
  ## @param service.annotations Annotations for the Service
  ## set the LoadBalancer service type to internal only.
  ## ref: https://kubernetes.io/docs/concepts/services-networking/service/#internal-load-balancer
  ##
  annotations: {}
  ## @param service.labels Labels for the Service
  ## have metrics-server show up in `kubectl cluster-info`
  ##  kubernetes.io/cluster-service: "true"
  ##  kubernetes.io/name: "Metrics-server"
  ##
  labels: {}
## Metric Server containers' resource requests and limits
## ref: https://kubernetes.io/docs/user-guide/compute-resources/
## We usually recommend not to specify default resources and to leave this as a conscious
## choice for the user. This also increases chances charts run on environments with little
## resources, such as Minikube. If you do want to specify resources, uncomment the following
## lines, adjust them as necessary, and remove the curly braces after 'resources:'.
## @param resources.limits The resources limits for the container
## @param resources.requests The requested resources for the container
##
resources:
  ## Example:
  ## limits:
  ##    cpu: 250m
  ##    memory: 256Mi
  limits: {}
  ## Examples:
  ## requests:
  ##    cpu: 250m
  ##    memory: 256Mi
  requests: {}
## Configure extra options for liveness probe
## ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-probes/#configure-probes
## @param livenessProbe.enabled Enable livenessProbe
## @param livenessProbe.httpGet.path Request path for livenessProbe
## @param livenessProbe.httpGet.port Port for livenessProbe
## @param livenessProbe.httpGet.scheme Scheme for livenessProbe
## @param livenessProbe.periodSeconds Period seconds for livenessProbe
## @param livenessProbe.failureThreshold Failure threshold for livenessProbe
##
livenessProbe:
  enabled: true
  failureThreshold: 3
  httpGet:
    path: /livez
    port: https
    scheme: HTTPS
  periodSeconds: 10
## Configure extra options for readiness probe
## ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-probes/#configure-probes
## @param readinessProbe.enabled Enable readinessProbe
## @param readinessProbe.httpGet.path Request path for readinessProbe
## @param readinessProbe.httpGet.port Port for readinessProbe
## @param readinessProbe.httpGet.scheme Scheme for livenessProbe
## @param readinessProbe.periodSeconds Period seconds for readinessProbe
## @param readinessProbe.failureThreshold Failure threshold for readinessProbe
##
readinessProbe:
  enabled: true
  failureThreshold: 3
  httpGet:
    path: /readyz
    port: https
    scheme: HTTPS
  periodSeconds: 10
## @param customLivenessProbe Custom Liveness probes for metrics-server
##
customLivenessProbe: {}
## @param customReadinessProbe Custom Readiness probes metrics-server
##
customReadinessProbe: {}
## Container security context
## ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/#set-the-security-context-for-a-container
## @param containerSecurityContext.enabled Enable Container security context
## @param containerSecurityContext.readOnlyRootFilesystem ReadOnlyRootFilesystem for the container
## @param containerSecurityContext.runAsNonRoot Run containers as non-root users
##
containerSecurityContext:
  enabled: true
  readOnlyRootFilesystem: false
  runAsNonRoot: true
## Pod security context
## ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/#set-the-security-context-for-a-pod
## @param podSecurityContext.enabled Pod security context
##
podSecurityContext:
  enabled: false
## Extra volumes to mount
## @param extraVolumes Extra volumes
## @param extraVolumeMounts Mount extra volume(s)
## Example Use Case: mount an `emptyDir` to allow running with a `readOnlyRootFilesystem: true`
##  extraVolumes:
##  - name: tmpdir
##    emptyDir: {}
##
extraVolumes: []
##  extraVolumeMounts:
##  - name: tmpdir
##    mountPath: /tmp
##
extraVolumeMounts: []
## @param extraContainers Extra containers to run within the pod
##
extraContainers: {}
```



``` shell
helm install  metrics-server bitnami/metrics-server --namespace kube-system --values /tmp/metrics-server.values
```



``` shell
kubectl edit deployment metrics-server -n kube-system
        - --secure-port=8443
        - --kubelet-insecure-tls=true
        - --kubelet-preferred-address-types=InternalIP,ExternalIP,Hostname
        - --kubelet-use-node-status-port
        - --metric-resolution=15s
```



its worked! cheers

``` shell
[root@10-23-75-240 metrics-server]# kubectl top node
NAME            CPU(cores)   CPU%   MEMORY(bytes)   MEMORY%   
10-23-184-141   65m          3%     1410Mi          38%       
10-23-245-35    63m          3%     1381Mi          37%       
10-23-75-240    123m         6%     2066Mi          55% 

[root@10-23-75-240 metrics-server]#  kubectl get --raw "/apis/metrics.k8s.io/v1beta1/nodes"
{"kind":"NodeMetricsList","apiVersion":"metrics.k8s.io/v1beta1","metadata":{},"items":[{"metadata":{"name":"10-23-184-141","creationTimestamp":"2022-03-13T04:47:13Z","labels":{"beta.kubernetes.io/arch":"amd64","beta.kubernetes.io/os":"linux","kubernetes.io/arch":"amd64","kubernetes.io/hostname":"10-23-184-141","kubernetes.io/os":"linux"}},"timestamp":"2022-03-13T04:47:04Z","window":"20.021s","usage":{"cpu":"67330954n","memory":"1446908Ki"}},{"metadata":{"name":"10-23-245-35","creationTimestamp":"2022-03-13T04:47:13Z","labels":{"beta.kubernetes.io/arch":"amd64","beta.kubernetes.io/os":"linux","kubernetes.io/arch":"amd64","kubernetes.io/hostname":"10-23-245-35","kubernetes.io/os":"linux"}},"timestamp":"2022-03-13T04:47:04Z","window":"20.022s","usage":{"cpu":"56302577n","memory":"1411724Ki"}},{"metadata":{"name":"10-23-75-240","creationTimestamp":"2022-03-13T04:47:13Z","labels":{"beta.kubernetes.io/arch":"amd64","beta.kubernetes.io/os":"linux","kubernetes.io/arch":"amd64","kubernetes.io/hostname":"10-23-75-240","kubernetes.io/os":"linux","node-role.kubernetes.io/control-plane":"","node-role.kubernetes.io/master":"","node.kubernetes.io/exclude-from-external-load-balancers":""}},"timestamp":"2022-03-13T04:47:02Z","window":"10.012s","usage":{"cpu":"104636674n","memory":"2116604Ki"}}]}
```





### have no idea why it is not working if I change the namespace to the metrics from kube-system
