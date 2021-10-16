https://github.com/kubernetes/

Why kops?#
kops lets us create a production-grade Kubernetes cluster. That means that we can use it not only to create a cluster, but also to upgrade it (without downtime), update it, or destroy it if we don’t need it anymore. A cluster cannot be called “production grade” unless it is highly available and fault tolerant. We should be able to execute it entirely from the command line if we’d like it to be automated. Those and quite a few other things are what kops provides, and what makes it great.

kops follows the same philosophy as Kubernetes. We create a set of JSON or YAML objects which are sent to controllers that create a cluster.

