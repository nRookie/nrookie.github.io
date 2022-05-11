- context:
    cluster: cluster.example.com
    namespace: default
    user: aws-user // Changed to aws-user
  name: cluster.example.com

// new entry for the aws user
- name: aws-user
  user:
    exec:
      apiVersion: client.authentication.k8s.io/v1alpha1
      args:
      - token
      - -i
      - cluster.example.com
      command: heptio-authenticator-aws
      env: null

https://github.com/kubernetes/kops/issues/5634


https://github.com/kubernetes-sigs/aws-iam-authenticator


https://docs.aws.amazon.com/eks/latest/userguide/create-kubeconfig.html

https://docs.aws.amazon.com/eks/latest/userguide/troubleshooting_iam.html

https://docs.aws.amazon.com/eks/latest/userguide/security_iam_id-based-policy-examples.html


https://kops.sigs.k8s.io/getting_started/aws/


## finaly solved by

``` shell

aws iam attach-group-policy --policy-arn arn:aws:iam::aws:policy/AmazonEC2FullAccess --group-name kops
aws iam attach-group-policy --policy-arn arn:aws:iam::aws:policy/AmazonRoute53FullAccess --group-name kops
aws iam attach-group-policy --policy-arn arn:aws:iam::aws:policy/AmazonS3FullAccess --group-name kops
aws iam attach-group-policy --policy-arn arn:aws:iam::aws:policy/IAMFullAccess --group-name kops
aws iam attach-group-policy --policy-arn arn:aws:iam::aws:policy/AmazonVPCFullAccess --group-name kops
aws iam attach-group-policy --policy-arn arn:aws:iam::aws:policy/AmazonSQSFullAccess --group-name kops
aws iam attach-group-policy --policy-arn arn:aws:iam::aws:policy/AmazonEventBridgeFullAccess --group-name kops

aws iam create-access-key --user-name kops
aws configure

aws iam list-users


```

### kops cannot edit instance node types

``` shell
 kops get -o yaml 

 kops edit ig --name $NAME nodes-us-east-2a
```

### kops cannot update 

``` shell
goroutine 1 [running]:
k8s.io/kops/upup/pkg/fi/cloudup.PopulateInstanceGroupSpec(0x1400091d900, 0x14000129180, {0x134f433f8, 0x14000c8cfc0}, 0x14000bb8180)
	/private/tmp/kops-20211014-83785-llh589/kops-1.22.0/src/k8s.io/kops/upup/pkg/fi/cloudup/populate_instancegroup_spec.go:164 +0xad8
main.updateInstanceGroup({0x109af22b8, 0x140000560a0}, {0x109b97f98, 0x1400091e230}, 0x14000bb8180, 0x1400091d900, 0x1400023ea80, 0x14000129180)
	/private/tmp/kops-20211014-83785-llh589/kops-1.22.0/src/k8s.io/kops/cmd/kops/edit_instancegroup.go:291 +0x6c
main.RunEditInstanceGroup({0x109af22b8, 0x140000560a0}, 0x14000710900, {0x109a7ad80, 0x14000010018}, 0x140000af590)
	/private/tmp/kops-20211014-83785-llh589/kops-1.22.0/src/k8s.io/kops/cmd/kops/edit_instancegroup.go:268 +0xd90
main.NewCmdEditInstanceGroup.func2(0x14000904a00, {0x140005dfe60, 0x1, 0x3})
	/private/tmp/kops-20211014-83785-llh589/kops-1.22.0/src/k8s.io/kops/cmd/kops/edit_instancegroup.go:107 +0x5c
github.com/spf13/cobra.(*Command).execute(0x14000904a00, {0x140005dfe30, 0x3, 0x3})
	/private/tmp/kops-20211014-83785-llh589/kops-1.22.0/src/k8s.io/kops/vendor/github.com/spf13/cobra/command.go:856 +0x668
github.com/spf13/cobra.(*Command).ExecuteC(0x10c02a620)
	/private/tmp/kops-20211014-83785-llh589/kops-1.22.0/src/k8s.io/kops/vendor/github.com/spf13/cobra/command.go:974 +0x410
github.com/spf13/cobra.(*Command).Execute(...)
	/private/tmp/kops-20211014-83785-llh589/kops-1.22.0/src/k8s.io/kops/vendor/github.com/spf13/cobra/command.go:902
main.Execute()
	/private/tmp/kops-20211014-83785-llh589/kops-1.22.0/src/k8s.io/kops/cmd/kops/root.go:95 +0x90
main.main()
	/private/tmp/kops-20211014-83785-llh589/kops-1.22.0/src/k8s.io/kops/cmd/kops/main.go:20 +0x20
```



  ### recommendation


  https://kops.sigs.k8s.io/getting_started/production/