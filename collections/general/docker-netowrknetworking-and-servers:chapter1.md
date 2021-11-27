Docker networking





<img src="/Users/user/Library/Application Support/typora-user-images/image-20211104181225800.png" alt="image-20211104181225800" style="zoom:50%;" />







### The docker0 bridge

The `docker0` bridge is the heart of default networking. When the Docker service is started, a Linux bridge is created on the host machine. The interfaces on the containers talk to the bridge, and the bridge proxies to the external world. Multiple containers on the same host can talk to each other through the Linux bridge.

`docker0` can be configured via the `--net` flag and has, in general, four modes:





### The --net default mode

In this mode, the default bridge is used as the bridge for containers to connect to each other.

### The --net=none mode

With this mode, the container created is truly isolated and cannot connect to the network.



### The --net=container:$container2 mode

With this flag, the container created shares its network namespace with the container called `$container2`.



### The --net=host mode

With this mode, the container created shares its network namespace with the host.



#### Port mapping in Docker container

In this section, we look at how container ports are mapped to host ports. This mapping can either be done implicitly by Docker Engine or can be specified.





If we create two containers called **Container1** and **Container2**, both of them are assigned an IP address from a private IP address space and also connected to the **docker0** bridge, as shown in the following figure:



As mentioned in the previous section, containers use network namespaces. When the first container is created, a new network namespace is created for the container. A vEthernet link is created between the container and the Linux bridge. Traffic sent from `eth0` of the container reaches the bridge through the vEthernet interface and gets switched thereafter. The following code can be used to show a list of Linux bridges:





