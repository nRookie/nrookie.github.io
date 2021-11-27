#### Configuring a DNS server



Docker provides hostname and DNS configurations for each container without us having to build a custom image. It overlays the /etc folder inside the container with virtual files, in whcih it can write new information.





This can be seen by running the mount command inside the container. Containers receive the same resolve.conf file as that of the host machine when they are created initially. If a host's resolv.conf  file is modified, this will be reflected in the container's /resolv.conf file only when the container is restarted.



In Docker you can set DNS options in two ways:



- Using docker run --dns=<ip-address>
- Adding DOCKER_OPTS="--dns ip-address" to the docker daemon file.



You can also specify the search domain using --dns-search=<DOMAIN>.

The following figure shows a nameserver being configured in a container using the DOCKER_OPTS setting in the Docker daemon file:

<img src="/Users/user/Library/Application Support/typora-user-images/image-20211105112316829.png" alt="image-20211105112316829" style="zoom:50%;" />



The main DNS files are as follows: 

- /etc/hostname 
- /etc/resolv.conf 
- /etc/hosts 

The following is the command to add a DNS server:



``` shell
docker run --dns=8.8.8.8 --net="bridge" -t -i ubuntu:latest /bin/bash
```



``` shell
docker run --dns=8.8.8.8 --hostname=docker-vm1 -t -i ubuntu:latest
/bin/bash
```





## Communication between containers and external networks



Packets can only pass between containers if the ip_forward parameter is set to 1. Usually, you will simply leave the Docker server at its default setting, --ip-forward=true, and Docker will set ip_forward to 1 for you when the server starts up.



To check the settings or to turn IP forwarding on manually, use these commands:





``` shell
# cat /proc/sys/net/ipv4/ip_forward
0
# echo 1 > /proc/sys/net/ipv4/ip_forward
# cat /proc/sys/net/ipv4/ip_forward
1
```





By enabling ip_forward, users can make communication possible between containers and the external world; it will also be required for inter-container communication if you are in a multiple-bridge setup. The following figure shows how ip_forward = false forwards all the packets to/from the container from/to the external network:



<img src="/Users/user/Library/Application Support/typora-user-images/image-20211105112910429.png" style="zoom:50%;" />



Docker will not delete or modify any pre-existing rules from the Docker filter chain. This allows users to create rules to restrict access to containers.



Docker uses the docker0 bridge for packet flow between all the containers on a single host. It adds a rule to forward the chain using IPTables in order for the packets to flow between two containers. Setting --icc=false will drop all the packets.



When the Docker daemon is configured with both --icc=false and --iptables=true and docker run is invoked with the --link option, the Docker server will insert a pair of IPTables accept rules for new containers to connect to the ports exposed by the other containers, which will be the ports that have been mentioned in the exposed lines of its Dockerfile. The following figure shows how ip_forward = false drops all the packets to/from the container from/to the external network:



<img src="/Users/user/Library/Application Support/typora-user-images/image-20211105113123517.png" alt="image-20211105113123517" style="zoom:50%;" />



By default, Docker’s forward rule permits all external IPs. To allow only a specific IP or network to access the containers, insert a negated rule at the top of the Docker filter chain. For example, using the following command, you can restrict external access such that only the source IP 10.10.10.10 can access the containers:





``` shell
#iptables –I DOCKER –i ext_if ! –s 10.10.10.10 –j DROP
```







### Overlay networks and underlay networks



An overlay is a virtual network that is built on top of underlying network infrastructure (the underlay). The purpose is to implement a network service that is not available in the physical network.



Network overlay dramatically increases the number of virtual subnets that can be created on top of the physical network, which in turn supports multi-tenancy and virtualization.



Every container in Docker is assigned an IP address, which is used for communication with other containers. If a container has to communicate with the external network, you set up networking in the host system and expose or map the port from the container to the host machine. With this, applications running inside containers will not be able to advertise their external IP and ports, as the information will not be available to them.



The solution is to somehow assign unique IPs to each Docker container across all hosts and have some networking product that routes traffic between hosts.



There are different projects to deal with Docker networking, as follows:



- Flannel

- Weave

- Open vSwitch





Flannel provides a solution by giving each container an IP that can be used for containerto-container communication. Using packet encapsulation, it creates a virtual overlay network over the host network. By default, Flannel provides a /24 subnet to hosts, from which the Docker daemon allocates IPs to containers. The following figure shows the communication between containers using Flannel:





<img src="/Users/user/Library/Application Support/typora-user-images/image-20211105141840261.png" alt="image-20211105141840261" style="zoom:50%;" />





Flannel runs an agent flannels, on each host and is responsible for allocating a subnet lease out of a preconfigured 

address space. Flannel uses etcd to store the network configuration, allocated subnets, and auxiliary data (Such as the host' IP).



Flannel uses the universal TUN/TAP device and creates an overlay network using UDP to encapsulate IP packets. Subnet allocation is done with the help of etcd, which maintains the overlay subnet-to-host mappings.



Weave creates a virtual network that connects Docker containers deployed across hosts/VMs and enables their automatic discovery. The following figure shows a Weave network:



<img src="/Users/user/Library/Application Support/typora-user-images/image-20211105142039963.png" alt="image-20211105142039963" style="zoom:50%;" />



Weave can traverse firewalls and operate in partially connected networks. Traffic can be optionally encrypted, allowing hosts/VMs to be connected across an untrusted network.



Weave augments Docker’s existing (single host) networking capabilities, such as the docker0 bridge, so these can continue to be used by containers.



Open vSwitch is an open source OpenFlow-capable virtual switch that is typically used with hypervisors to interconnect virtual machines within a host and between different hosts across networks. Overlay networks need to create a virtual datapath using supported tunneling encapsulations, such as VXLAN and GRE.



The overlay datapath is provisioned between tunnel endpoints residing in the Docker host, which gives the appearance of all hosts within a given provider segment being directly connected to one another.



As a new container comes online, the prefix is updated in the routing protocol, announcing its location via a tunnel endpoint. As the other Docker hosts receive the updates, the forwarding rule is installed into the OVS for the tunnel endpoint that the host resides on. When the host is de-provisioned, a similar process occurs and tunnel endpoint Docker hosts remove the forwarding entry for the de-provisioned container. The following figure shows the communication between containers running on multiple hosts through OVSbased VXLAN tunnels:



<img src="/Users/user/Library/Application Support/typora-user-images/image-20211105142627106.png" alt="image-20211105142627106" style="zoom:50%;" />





