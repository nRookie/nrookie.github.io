## istioctl kube-inject

kube-inject manually injects the Istio sidecar into Kubernetes workloads. Unsupported resources are left unmodified so it is safe to run kube-inject over a single file that contains multiple Service, ConfigMap, Deployment, etc. definitions for a complex application. When in doubt re-run istioctl kube-inject on deployments to get the most up-to-date changes.

It's best to do kube-inject when the resource is initially created.



``` shell

istioctl kube-inject [flags]

```

## examples

``` shell
  # Update resources on the fly before applying.
  kubectl apply -f <(istioctl kube-inject -f <resource.yaml>)

  # Create a persistent version of the deployment with Istio sidecar injected.
  istioctl kube-inject -f deployment.yaml -o deployment-injected.yaml

  # Update an existing deployment.
  kubectl get deployment -o yaml | istioctl kube-inject -f - | kubectl apply -f -

  # Capture cluster configuration for later use with kube-inject
  kubectl -n istio-system get cm istio-sidecar-injector  -o jsonpath="{.data.config}" > /tmp/inj-template.tmpl
  kubectl -n istio-system get cm istio -o jsonpath="{.data.mesh}" > /tmp/mesh.yaml
  kubectl -n istio-system get cm istio-sidecar-injector -o jsonpath="{.data.values}" > /tmp/values.json

  # Use kube-inject based on captured configuration
  istioctl kube-inject -f samples/bookinfo/platform/kube/bookinfo.yaml \
    --injectConfigFile /tmp/inj-template.tmpl \
    --meshConfigFile /tmp/mesh.yaml \
    --valuesFile /tmp/values.json

```
## Reference

https://istio.io/latest/docs/reference/commands/istioctl/