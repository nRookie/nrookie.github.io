https://kubernetes.io/docs/tasks/administer-cluster/migrating-from-dockershim/find-out-runtime-you-use/







``` shell
/var/lib/kubelet/kubeadm-flags.env and add the containerd runtime to the flags. --container-runtime=remote and --container-runtime-endpoint=unix:///run/containerd/containerd.sock
```







kubectl get nodes -o wide



https://kubernetes.io/docs/tasks/administer-cluster/migrating-from-dockershim/change-runtime-containerd/







kubectl drain <node-to-drain> --ignore-daemonsets





```shell
systemctl stop kubelet
systemctl disable docker.service --now
```

 



https://kubernetes.io/docs/tasks/debug/debug-cluster/crictl/#:~:text=crictl%20is%20a%20command%2Dline,in%20the%20cri%2Dtools%20repository.



``` shell
vi /etc/crictl.yaml
runtime-endpoint: unix:///var/run/containerd/containerd.sock
image-endpoint: unix:///var/run/containerd/containerd.sock
timeout: 10
debug: true
```

