https://github.com/moby/moby/issues/33775


## docker overlay eventually eat up volumes.
```
[root@10-13-175-37 ~]# df -hT
Filesystem     Type      Size  Used Avail Use% Mounted on
devtmpfs       devtmpfs  1.8G     0  1.8G   0% /dev
tmpfs          tmpfs     1.9G     0  1.9G   0% /dev/shm
tmpfs          tmpfs     1.9G  172M  1.7G  10% /run
tmpfs          tmpfs     1.9G     0  1.9G   0% /sys/fs/cgroup
/dev/vda1      xfs        20G   19G  1.2G  95% /
/dev/vdb       xfs        20G  175M   20G   1% /data
overlay        overlay    20G   19G  1.2G  95% /var/lib/docker/overlay2/3cd99c5ebc79c96c6d75491b90cac4ba337ad40b112301495ce687401deff439/merged
overlay        overlay    20G   19G  1.2G  95% /var/lib/docker/overlay2/3b0368bab7a065aea43684557bd8634679cfbc5bfbd7be788c84c881a350a325/merged
overlay        overlay    20G   19G  1.2G  95% /var/lib/docker/overlay2/46c5b4cd57880ff011bef062c2575c712663384bcfb8f608cb88c1155e2ac944/merged
overlay        overlay    20G   19G  1.2G  95% /var/lib/docker/overlay2/3ac84ab3f81e2c151a6ce35bc9c7967245d71c07ff974e7074d5a19ac40f5bc5/merged
overlay        overlay    20G   19G  1.2G  95% /var/lib/docker/overlay2/fac02b2a0e89c73cec1451eb1411e46e6a713e147335180775ed1b53424965ce/merged
overlay        overlay    20G   19G  1.2G  95% /var/lib/docker/overlay2/922e08a774ca25daf9a4945d24ef00fab2eccdb1ede56a7d42fca07b37684a17/merged
overlay        overlay    20G   19G  1.2G  95% /var/lib/docker/overlay2/64dd667b4d6e56b7aeb37691cf392ce7305e090dba36bba6d2f0f44abcadc8cf/merged
tmpfs          tmpfs     372M     0  372M   0% /run/user/0
overlay        overlay    20G   19G  1.2G  95% /var/lib/docker/overlay2/749a9e5370daff95cb1150c1d400d209aac130396285c353e1f7a369e09963f1/merged
[root@10-13-175-37 ~]# ls
```

stop the container solved the problem but why ?

``` shell
docker kill $(docker ps -q)
```


