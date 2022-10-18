https://github.com/kairen/kubeadm-ansible

## Install

``` shell
yum install ansible
```

### disable host key checking

``` shell
sed -i 's/#host_key_checking/host_key_checking/g' /etc/ansible/ansible.cfg
```

### add password info in inventory

``` shell
[kube-cluster:children]
master
node
[all:vars]
ansible_connection=ssh
ansible_user=root
ansible_ssh_pass=pass
```

## setup v2ray

``` shell
following v2ray docs
```



### Hangs

![image-20220402000736453](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/Kubeadm/image-20220402000736453.png)

``` shell
Apr  2 00:07:08 10-23-21-70 v2ray: 2022/04/02 00:07:08 [Warning] [3502202686] app/proxyman/inbound: connection ends > proxy/http: failed to read http request > read tcp 127.0.0.1:8001->127.0.0.1:64116: read: connection reset by peer
Apr  2 00:07:08 10-23-21-70 v2ray: 2022/04/02 00:07:08 [Warning] [1617189173] app/dispatcher: non existing tag: direct
Apr  2 00:07:17 10-23-21-70 v2ray: 2022/04/02 00:07:17 [Warning] [1617189173] proxy/http: failed to read response from mirrors.ucloud.cn > unexpected EOF
Apr  2 00:07:17 10-23-21-70 v2ray: 2022/04/02 00:07:17 [Warning] [1216240453] app/dispatcher: non existing tag: direct
Apr  2 00:07:23 10-23-21-70 v2ray: 2022/04/02 00:07:23 [Warning] [1216240453] proxy/http: failed to read response from mirrors.ucloud.cn > unexpected EOF
Apr  2 00:07:23 10-23-21-70 v2ray: 2022/04/02 00:07:23 [Warning] [914545191] app/dispatcher: non existing tag: direct
Apr  2 00:07:32 10-23-21-70 v2ray: 2022/04/02 00:07:32 [Warning] [914545191] proxy/http: failed to read response from mirrors.ucloud.cn > unexpected EOF
Apr  2 00:07:32 10-23-21-70 v2ray: 2022/04/02 00:07:32 [Warning] [2071160009] app/dispatcher: non existing tag: direct
Apr  2 00:07:38 10-23-21-70 v2ray: 2022/04/02 00:07:38 [Warning] [2071160009] proxy/http: failed to read response from mirrors.ucloud.cn > unexpected EOF
Apr  2 00:07:38 10-23-21-70 v2ray: 2022/04/02 00:07:38 [Warning] [301832536] app/dispatcher: non existing tag: direct
```







### Yum install error

``` shell
rm -f /var/lib/rpm/__*
rpm --rebuilddb -v -v

yum clean dbcache
yum clean metadata
yum clean rpmdb
yum clean headers
yum clean all

rm -rf /var/cache/yum/timedhosts.txt
rm -rf /var/cache/yum/*

yum makecache

```



## NO_PROXY



it seems like only capital NO_PROXY works

