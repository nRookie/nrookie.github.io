https://medium.com/@zwhitchcox/matchlabels-labels-and-selectors-explained-in-detail-for-beginners-d421bdd05362


https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata


So, in your yaml description file for the deployment, it might look something like this:
kind: Deployment
...
metadata:
  name: nginx
  labels:
    app: nginx
    tier: backend
spec:
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
        tier: backend
...
So, this is kind of confusing. What does the selector, label, and matchLabel do? And why are there multiple nested labels?
So, the first metadata describes the deployment itself. It gives a label for that actual deployment. So, if you want to delete that deployment, you would say kubectl delete -l app=nginx,tier=backend. Simple enough? Ok, so then why is there something called selector, matchlabels, and another labels? Don’t we already have our labels? Well, that’s what I thought too. So it turns out, the second selector, is actually a selector for the deployment to apply to the pod that the deployment is describing. So, that’s kind of hard to understand, mainly, because intuitively, you would think it would be automatic. So let’s break that down.
The template is actually a podTemplate. It describes a pod that is launched. One of the fields for the pod templates is replicas. If we set replicas to 2, it would make 2 pods for that deployment, and the deployment would entail both of those pods. So, that template for the pod has a label. So, this isn’t a label for the deployment anymore, it’s a label for the pod that the deployment is deploying. That’s a subtle, but important distinction.
Then, we have to tell the deployment to match the pods that it’s deploying. Why doesn’t the deployment automatically match the pod it’s deploying? I have no idea. I’m sure someone will tell us in the comments though. Hopefully.
But anyway, the selector: matchLabels tells the resource, whatever it may be, service, deployment, etc, to match the pod, according to that label. So, maybe a potential reason is for uniformity. For instance, when you want to apply a service (to expose the pod to the web or something), you need to apply that service to the pod with a match label, like this:
apiVersion: v1
kind: Service
metadata:
  name: nginx
spec:
  type: LoadBalancer
  ports:
    - port:  80
  selector:
    app: nginx
    tier: frontend
So, matchLabels are not supported by Service, but only certain new resources like Deployment. So, here, you just add the labels you want. You can’t use matchLabel yet, but it means the same thing as if you had specified
selector:
  matchLabels:
    app: nginx
    tier: frontend
So, that was another confusing part, why sometimes it’s matchLabels, and somtimes it’s just selector with a map. Only Job, Deployment, Replica Set, and Daemon Set support matchLabels.
So, that’s it. Now, you know as much as I do. Good luck, and hope this was helpful!
Oh, almost forgot, NodeSelector applies if you add labels to nodes which are vms that your cluster runs on (layman’s explanation).



```  shell
Defining a Service
A Service in Kubernetes is a REST object, similar to a Pod. Like all of the REST objects, you can POST a Service definition to the API server to create a new instance. The name of a Service object must be a valid RFC 1035 label name.

For example, suppose you have a set of Pods where each listens on TCP port 9376 and contains a label app=MyApp:

apiVersion: v1
kind: Service
metadata:
  name: my-service
spec:
  selector:
    app: MyApp
  ports:
    - protocol: TCP
      port: 80
      targetPort: 9376

```
This specification creates a new Service object named "my-service", which targets TCP port 9376 on any Pod with the app=MyApp label.

Kubernetes assigns this Service an IP address (sometimes called the "cluster IP"), which is used by the Service proxies (see Virtual IPs and service proxies below).

The controller for the Service selector continuously scans for Pods that match its selector, and then POSTs any updates to an Endpoint object also named "my-service".

