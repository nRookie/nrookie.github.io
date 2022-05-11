## Deploying Jenkins [#](https://www.educative.io/module/lesson/a-practical-guide-to-kubernetes/JPn98XQX5nJ#Deploying-Jenkins-)

We already deployed Jenkins a few times. Since it is a stateful application, it is an excellent candidate to serve as a playground.

### Looking into the Definition [#](https://www.educative.io/module/lesson/a-practical-guide-to-kubernetes/JPn98XQX5nJ#Looking-into-the-Definition-)

Let’s take a look at a definition stored in the `pv/jenkins-no-pv.yml` file.





The YAML defines the `jenkins` Namespace, an Ingress controller, and a Service. We’re already familiar with those types of resources so we’ll skip explaining them and jump straight to the Deployment definition.

The output of the `cat` command, limited to the `jenkins` Deployment, is as follows.



There’s nothing special about this Deployment. We already used a very similar one. Besides, by now, you’re an expert at Deployment controllers.

The only thing worth mentioning is that there is only one volume mount and it references a secret we’re using to provide Jenkins with the initial administrative user. Jenkins is persisting its state in `/var/jenkins_home`, and we are not mounting that directory.



### Creating the Resources [#](https://www.educative.io/module/lesson/a-practical-guide-to-kubernetes/JPn98XQX5nJ#Creating-the-Resources-)

Let’s create the resources defined in `pv/jenkins-no-pv.yml`.

``` shell
kubectl create \
    -f pv/jenkins-no-pv.yml \
    --record --save-config
```





``` shell
kubectl --namespace jenkins \
    get events
```

``` shell
[root@10-23-75-240 k8s-specs]# kubectl --namespace jenkins \
>     get events
LAST SEEN   TYPE      REASON              OBJECT                         MESSAGE
31s         Normal    Scheduled           pod/jenkins-597b648cd-5nfdq    Successfully assigned jenkins/jenkins-597b648cd-5nfdq to 10-23-245-35
16s         Warning   FailedMount         pod/jenkins-597b648cd-5nfdq    MountVolume.SetUp failed for volume "jenkins-creds" : secret "jenkins-creds" not found
31s         Normal    SuccessfulCreate    replicaset/jenkins-597b648cd   Created pod: jenkins-597b648cd-5nfdq
15s         Normal    Sync                ingress/jenkins                Scheduled for sync
31s         Normal    ScalingReplicaSet   deployment/jenkins             Scaled up replica set jenkins-597b648cd to 1
```



#### Creating the Secret [#](https://www.educative.io/module/lesson/a-practical-guide-to-kubernetes/JPn98XQX5nJ#Creating-the-Secret-)

We can see that the setup of the only volume failed since it could not find the secret referenced as `jenkins-creds`. Let’s create it.



``` shell
kubectl --namespace jenkins \
    create secret \
    generic jenkins-creds \
    --from-literal=jenkins-user=jdoe \
    --from-literal=jenkins-pass=incognito
```





### Verification [#](https://www.educative.io/module/lesson/a-practical-guide-to-kubernetes/JPn98XQX5nJ#Verification-)



Now, with the secret `jenkins-creds` created in the `jenkins` Namespace, we can confirm that the rollout of the Deployment was successful.



``` shell
[root@10-23-75-240 k8s-specs]# kubectl --namespace jenkins \
>     rollout status \
>     deployment jenkins
Waiting for deployment "jenkins" rollout to finish: 0 of 1 updated replicas are available...

deployment "jenkins" successfully rolled out
```





![image-20220317011327091](/Users/user/playground/share/nrookie.github.io/collections/k8s-related/persisting-state/image-20220317011327091.png)





## Creating a Job[#](https://www.educative.io/module/lesson/a-practical-guide-to-kubernetes/xljqDDA05g9#Creating-a-Job)

Please click the *Log in* link, type *jdoe* as the *User*, and *incognito* as the *Password*. When finished, click the *log in* button.

Now that we are authenticated as `jdoe` administrator, we can proceed and create a job. That will generate a state that we can use to explore what happens when a stateful application fails.

Please click the *create new jobs* link, type *my-job* as the item name, select *Pipeline* as the job type, and press the *OK* button.

You’ll be presented with the job configuration screen. There’s no need to do anything here since we are not, at the moment, interested in any specific Pipeline definition. It’s enough to click the *Save* button



## Simulating and Analyzing Failure[#](https://www.educative.io/module/lesson/a-practical-guide-to-kubernetes/xljqDDA05g9#Simulating-and-Analyzing-Failure)

Next, we’ll simulate a failure by killing `java` process running inside the Pod created by the `jenkins` Deployment. To do that, we need to find out the name of the Pod.

``` shell
kubectl --namespace jenkins \
    get pods \
    --selector=app=jenkins \
    -o json
```

 

We retrieved the Pods from the `jenkins` Namespace, filtered them with the selector `api=jenkins`, and formatted the output as `json`.

The **output**, limited to the relevant parts, as is follows.





``` shell
POD_NAME=$(kubectl \
    --namespace jenkins \
    get pods \
    --selector=app=jenkins \
    -o jsonpath="{.items[*].metadata.name}")

echo $POD_NAME
```



``` shell
kubectl --namespace jenkins \
    exec -it $POD_NAME -- pkill java
```



The container failed once we killed Jenkins process. We already know from experience that a failed container inside a Pod will be recreated. As a result, we had a short downtime, but Jenkins is running once again.

Let’s see what happened to the job we created earlier. I’m sure you know the answer, but we’ll check it anyway.



As expected, *my-job* is nowhere to be found. The container that was hosting `/var/jenkins_home` directory failed, and it was replaced with a new one. The state we created is lost.

We already saw in the [*Volumes*](https://www.educative.io/collection/page/5376908829130752/4742963282313216/4769265662033920) chapter that we can mount a volume in an attempt to preserve state across failures.

However, in the past, we used `emptyDir` which mounts a local volume. Even though that’s better than nothing, such a volume exists only as long as the server it is stored in is up and running. If the server would fail, the state stored in `emptyDir` would be gone.

Such a solution would be only slightly better than not using any volume. By using local disk we would only postpone the inevitable, and, sooner or later, we’d get to the same situation. We’d be left wondering why we lost everything we created in Jenkins. We can do better than that.

------

In the next lesson, we will explore creating AWS Volumes.



### External Storage[#](https://www.educative.io/module/lesson/a-practical-guide-to-kubernetes/RLNA4zmrXwz#External-Storage)

Since we are running our cluster in AWS, we can choose between [S3](https://aws.amazon.com/s3/), [Elastic File System (EFS)](https://aws.amazon.com/efs/), and [Elastic Block Store](https://aws.amazon.com/ebs/).

S3 is meant to be accessed through its API and is not suitable as a local disk replacement. That leaves us with EFS and EBS.

#### Elastic File System (EFS)[#](https://www.educative.io/module/lesson/a-practical-guide-to-kubernetes/RLNA4zmrXwz#Elastic-File-System-(EFS))

Elastic File System (EFS), has a distinct advantage that it can be mounted to multiple EC2 instances spread across multiple availability zones. It is the closest we can get to fault-tolerant storage. Even if a whole zone (datacenter) fails, we’ll still be able to use EFS in the rest of the zones used by our cluster. However, that comes at a cost. EFS introduces a performance penalty. It is, after all, a network file system (NFS), and that entails higher latency.

#### Elastic Block Store (EBS)[#](https://www.educative.io/module/lesson/a-practical-guide-to-kubernetes/RLNA4zmrXwz#Elastic-Block-Store-(EBS))

Elastic Block Store (EBS) is the fastest storage we can use in AWS. Its data access latency is very low thus making it the best choice when performance is the primary concern. The downside is availability. It doesn’t work in multiple availability zones. Failure of one will mean downtime, at least until the zone is restored to its operational state.

#### Our Choice[#](https://www.educative.io/module/lesson/a-practical-guide-to-kubernetes/RLNA4zmrXwz#Our-Choice)

We’ll choose EBS for our storage needs. Jenkins depends heavily on IO, and we need data access to be as fast as possible. However, there is another reason for such a choice. EBS is fully supported by Kubernetes. EFS will come but, at the time of this writing, it is still in the experimental stage. As a bonus advantage, EBS is much cheaper than EFS.

Given the requirements and what Kubernetes offers, the choice is obvious. We’ll use EBS, even though we might run into trouble if the availability zone where our Jenkins will run goes down. In such a case, we’d need to migrate EBS volume to a healthy zone. There’s no such thing as a perfect solution.

We are jumping ahead of ourselves. We’ll leave Kubernetes aside for a while and concentrate on creating an EBS volume.



