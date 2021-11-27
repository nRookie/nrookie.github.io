

Docker-host

``` shell
[root@10-13-175-37 ~]# ip addr
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host 
       valid_lft forever preferred_lft forever
258: veth11fa44c@if257: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue master docker0 state UP group default 
    link/ether c2:16:91:ef:ed:73 brd ff:ff:ff:ff:ff:ff link-netnsid 1
    inet6 fe80::c016:91ff:feef:ed73/64 scope link 
       valid_lft forever preferred_lft forever
2: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1452 qdisc noqueue state UP group default qlen 1000
    link/ether 52:54:00:da:3f:f6 brd ff:ff:ff:ff:ff:ff
    inet 10.13.175.37/16 brd 10.13.255.255 scope global noprefixroute eth0
       valid_lft forever preferred_lft forever
    inet6 fe80::5054:ff:feda:3ff6/64 scope link 
       valid_lft forever preferred_lft forever

```



Link-local and loopback addresses have link-local scope, which means they are to be used in a directly attached network (link). All other addresses have global (or universal) scope, which means they are globally routable and can be used to connect to addresses with global scope anywhere.

A directly connected network is **a network that is directly attached to one of the router interfaces**.

Docker-container 

``` shell
root@6aa43620d01c:/# ip addr
257: eth0@if258: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default 
    link/ether 02:42:ac:11:00:03 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 172.17.0.3/16 brd 172.17.255.255 scope global eth0
       valid_lft forever preferred_lft forever
    inet6 2001:db8:1::242:ac11:3/64 scope global nodad 
       valid_lft forever preferred_lft forever
    inet6 fe80::42:acff:fe11:3/64 scope link 
       valid_lft forever preferred_lft forever
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host 
       valid_lft forever preferred_lft forever
```

