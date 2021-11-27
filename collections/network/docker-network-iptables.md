### iptables



``` shell
iptables -t nat -L -n
Chain PREROUTING (policy ACCEPT)
target     prot opt source               destination         
DOCKER-INGRESS  all  --  0.0.0.0/0            0.0.0.0/0            ADDRTYPE match dst-type LOCAL
DOCKER     all  --  0.0.0.0/0            0.0.0.0/0            ADDRTYPE match dst-type LOCAL

Chain INPUT (policy ACCEPT)
target     prot opt source               destination         

Chain POSTROUTING (policy ACCEPT)
target     prot opt source               destination         
MASQUERADE  all  --  172.25.0.0/16        0.0.0.0/0           
MASQUERADE  all  --  172.24.0.0/16        0.0.0.0/0           
MASQUERADE  all  --  172.23.0.0/16        0.0.0.0/0           
MASQUERADE  all  --  172.22.0.0/16        0.0.0.0/0           
MASQUERADE  all  --  172.21.0.0/16        0.0.0.0/0           
MASQUERADE  all  --  172.19.0.0/16        0.0.0.0/0           
MASQUERADE  all  --  172.17.0.0/16        0.0.0.0/0           
MASQUERADE  all  --  0.0.0.0/0            0.0.0.0/0            ADDRTYPE match src-type LOCAL
MASQUERADE  all  --  172.18.0.0/16        0.0.0.0/0           

Chain OUTPUT (policy ACCEPT)
target     prot opt source               destination         
DOCKER-INGRESS  all  --  0.0.0.0/0            0.0.0.0/0            ADDRTYPE match dst-type LOCAL
DOCKER     all  --  0.0.0.0/0           !127.0.0.0/8          ADDRTYPE match dst-type LOCAL

Chain DOCKER-INGRESS (2 references)
target     prot opt source               destination         
RETURN     all  --  0.0.0.0/0            0.0.0.0/0           

Chain DOCKER (2 references)
target     prot opt source               destination         
RETURN     all  --  0.0.0.0/0            0.0.0.0/0           
RETURN     all  --  0.0.0.0/0            0.0.0.0/0           
RETURN     all  --  0.0.0.0/0            0.0.0.0/0           
RETURN     all  --  0.0.0.0/0            0.0.0.0/0           
RETURN     all  --  0.0.0.0/0            0.0.0.0/0           
RETURN     all  --  0.0.0.0/0            0.0.0.0/0           
RETURN     all  --  0.0.0.0/0            0.0.0.0/0           
RETURN     all  --  0.0.0.0/0            0.0.0.0/0 
```



### ifconfig



``` shell
[root@10-13-175-37 ~]# ifconfig
br-12591d3d0a93: flags=4099<UP,BROADCAST,MULTICAST>  mtu 1500
        inet 172.24.0.1  netmask 255.255.0.0  broadcast 172.24.255.255
        ether 02:42:26:35:d0:28  txqueuelen 0  (Ethernet)
        RX packets 1779  bytes 93074 (90.8 KiB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 905  bytes 39230 (38.3 KiB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0

br-346d9079ec7f: flags=4099<UP,BROADCAST,MULTICAST>  mtu 1500
        inet 172.22.0.1  netmask 255.255.0.0  broadcast 172.22.255.255
        ether 02:42:9f:ad:16:46  txqueuelen 0  (Ethernet)
        RX packets 68369  bytes 4275555 (4.0 MiB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 106453  bytes 307233226 (293.0 MiB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0

br-475d2a265046: flags=4099<UP,BROADCAST,MULTICAST>  mtu 1500
        inet 172.19.0.1  netmask 255.255.0.0  broadcast 172.19.255.255
        inet6 fe80::42:8bff:fec5:b8c6  prefixlen 64  scopeid 0x20<link>
        ether 02:42:8b:c5:b8:c6  txqueuelen 0  (Ethernet)
        RX packets 1779  bytes 93074 (90.8 KiB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 905  bytes 39230 (38.3 KiB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0

br-b5fd829df118: flags=4099<UP,BROADCAST,MULTICAST>  mtu 1500
        inet 172.25.0.1  netmask 255.255.0.0  broadcast 172.25.255.255
        ether 02:42:78:95:42:a4  txqueuelen 0  (Ethernet)
        RX packets 720  bytes 75386 (73.6 KiB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 720  bytes 75386 (73.6 KiB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0

br-d7de7ade4a03: flags=4099<UP,BROADCAST,MULTICAST>  mtu 1500
        inet 172.23.0.1  netmask 255.255.0.0  broadcast 172.23.255.255
        ether 02:42:e7:27:56:fe  txqueuelen 0  (Ethernet)
        RX packets 33440568  bytes 7292472227 (6.7 GiB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 33661538  bytes 4608691280 (4.2 GiB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0

br-e63f7ce0b933: flags=4099<UP,BROADCAST,MULTICAST>  mtu 1500
        inet 172.21.0.1  netmask 255.255.0.0  broadcast 172.21.255.255
        inet6 fe80::42:70ff:feb7:f122  prefixlen 64  scopeid 0x20<link>
        ether 02:42:70:b7:f1:22  txqueuelen 0  (Ethernet)
        RX packets 68369  bytes 4275555 (4.0 MiB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 106453  bytes 307233226 (293.0 MiB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0

docker0: flags=4099<UP,BROADCAST,MULTICAST>  mtu 1500
        inet 172.17.0.1  netmask 255.255.0.0  broadcast 172.17.255.255
        inet6 fe80::42:49ff:fe12:58a3  prefixlen 64  scopeid 0x20<link>
        ether 02:42:49:12:58:a3  txqueuelen 0  (Ethernet)
        RX packets 68369  bytes 4275555 (4.0 MiB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 106453  bytes 307233226 (293.0 MiB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0

docker_gwbridge: flags=4099<UP,BROADCAST,MULTICAST>  mtu 1500
        inet 172.18.0.1  netmask 255.255.0.0  broadcast 172.18.255.255
        inet6 fe80::42:b5ff:fe51:4e21  prefixlen 64  scopeid 0x20<link>
        ether 02:42:b5:51:4e:21  txqueuelen 0  (Ethernet)
        RX packets 68369  bytes 4275555 (4.0 MiB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 106453  bytes 307233226 (293.0 MiB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0

eth0: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1452
        inet 10.13.175.37  netmask 255.255.0.0  broadcast 10.13.255.255
        inet6 fe80::5054:ff:feda:3ff6  prefixlen 64  scopeid 0x20<link>
        ether 52:54:00:da:3f:f6  txqueuelen 1000  (Ethernet)
        RX packets 33440568  bytes 7292472227 (6.7 GiB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 33661542  bytes 4608692536 (4.2 GiB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0

eth1: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1452
        ether 52:54:00:da:3f:f6  txqueuelen 1000  (Ethernet)
        RX packets 1779  bytes 93074 (90.8 KiB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 905  bytes 39230 (38.3 KiB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0

eth2: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1452
        ether 52:54:00:da:3f:f6  txqueuelen 1000  (Ethernet)
        RX packets 33438789  bytes 7292379153 (6.7 GiB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 33660637  bytes 4608653306 (4.2 GiB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0

lo: flags=73<UP,LOOPBACK,RUNNING>  mtu 65536
        inet 127.0.0.1  netmask 255.0.0.0
        inet6 ::1  prefixlen 128  scopeid 0x10<host>
        loop  txqueuelen 1000  (Local Loopback)
        RX packets 720  bytes 75386 (73.6 KiB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 720  bytes 75386 (73.6 KiB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0

```



### docker network

``` shell
docker network ls
d7de7ade4a03   atsea-sample-shop-app_back-tier    bridge    local
346d9079ec7f   atsea-sample-shop-app_default      bridge    local
12591d3d0a93   atsea-sample-shop-app_front-tier   bridge    local
be58ad34c4f2   bridge                             bridge    local
e63f7ce0b933   counter-app_counter-net            bridge    local
433dbce26e5a   docker_gwbridge                    bridge    local
ccb8a6d5ae5e   host                               host      local
b5fd829df118   ms_default                         bridge    local
475d2a265046   mskafka_default                    bridge    local
266bd93eac1c   none                               null      local
```



![image-20211105102821064](/Users/user/Library/Application Support/typora-user-images/image-20211105102821064.png)



