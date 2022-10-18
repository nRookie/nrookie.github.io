



![企业微信截图_16379191144608](/Users/kestrel/developer/nrookie.github.io/collections/HPC/Prof Dr - Ing Morris Riedel/企业微信截图_16379191144608.png)







numa绑核就是你上次说的，两侧的内存绑到socket上，让socket内的cpu不需要通过QPI去访问另一侧的内存导致性能下降


然后绑核就是将vCPU都绑到物理CPU上，让客户的业务都跑在固定的cpu上，防止在不同的cpu上反复切换，性能会有很大的提高但不能超售

### what is QPI



The Intel QuickPath Interconnect (QPI) is **a point-to-point processor interconnect developed by Intel** which replaced the front-side bus (FSB) in Xeon, Itanium, and certain desktop platforms starting in 2008.













### sysfs meaning

[https://man7.org/linux/man-pages/man5/sysfs.5.html#:~:text=sys%2Fclass%2Fnet%20Each%20of,%2Fsys%2Fdevices%20directory.%20%2F](https://man7.org/linux/man-pages/man5/sysfs.5.html#:~:text=sys%2Fclass%2Fnet Each of,%2Fsys%2Fdevices directory. %2F)



```
/sys/class
              This subdirectory contains a single layer of further
              subdirectories for each of the device classes that have
              been registered on the system (e.g., terminals, network
              devices, block devices, graphics devices, sound devices,
              and so on).  Inside each of these subdirectories are
              symbolic links for each of the devices in this class.
              These symbolic links refer to entries in the /sys/devices
              directory.
       /sys/class/net
              Each of the entries in this directory is a symbolic link
              representing one of the real or virtual networking devices
              that are visible in the network namespace of the process
              that is accessing the directory.  Each of these symbolic
              links refers to entries in the /sys/devices directory.
```



https://www.kernel.org/doc/Documentation/ABI/testing/sysfs-class-net



### Network card numa information

``` shell
[root@hd08-uhost-148-240 ~]# cat /sys/class/net/net0/device/numa_node
0
[root@hd08-uhost-148-240 ~]# cat /sys/class/net/net1/device/numa_node
0
[root@hd08-uhost-148-240 ~]# cat /sys/class/net/net2/device/numa_node
0
[root@hd08-uhost-148-240 ~]# cat /sys/class/net/net3/device/numa_node
0
```









Check CPU per information

``` shell
 mpstat -A 5
```

