## Authorization Methods

Just like almost everything else in Kubernetes, authorization is modular. We can choose to use Node, ABAC, Webhook, or RBAC authorization.

Node: Node authorization grants permissions to kubelets based on the Pods they are scheduled to run.

ABAC: Attribute-based access control (ABAC) is based on attributes combined with policies and is considered deprecated in favor of RBAC.

Webhooks: Webhooks are used for event notifications through HTTP POST requests.

RBAC: Role-based access control (RBAC) grants (or denies) access to resources based on roles of individual users or groups.


What can we do with RBAC?

We can use it to secure the cluster by allowing access only to authorized users.

We can define roles that would grant different levels of access to users and groups. Some could have god-like permissions that would allow them to do almost anything, while others could be limited only to basic non-destructive operations. There can be many other roles in between.

We can combine RBAC with Namespaces and allow users to operate only within specific segments of a cluster.

There are many other combinations we could apply depending on particular use-cases.

We’ll leave the rest for later and explore details through a few examples. As you might already suspect, we’ll kick it off with a new Minikube cluster.



https://kops.sigs.k8s.io/getting_started/aws/