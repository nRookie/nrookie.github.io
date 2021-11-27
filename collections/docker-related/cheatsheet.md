export a docker file system.


``` shell
(docker export $(docker create busybox) | tar -C rootfs -xvf -)
```