### network_namespaces - overview of Linux network



Network namespaces provide isolation of the system resources associated with networking; network devices, IPv4 and IPv6 protocol stacks, IP routing tables, firewall rules, the /proc/net directory (which is a symbolic link to /proc/PID/net), the /sys/class/net directory, various files under /proc/sys/net, port numbers (sockets), and so on.



A physical netowrk device can live in exactly one network namespace. When a network namespace is freed (i.e., when the last process in the namespace terminates), its physical network devices are moved back to the initial network namespace ( not to the parent of the process).



A virtual network (veth (4)) device pair provides a pipe-like abstraction that can be used to create tunnels between network namespaces, and can be used to create a bridge to a physical network device in another namespace.



Use of network namespaces requires a kernel that is configured with the CONFIG_NET_NS option.

### veth



veth - Virtual Ethernet Device



The veth devices are virtual Ethernet devices. They can act as tunnels between network namespaces to create a bridge to a physical network device in another namespace, but can also be used as standalone network devices.



​    veth devices are always created in interconnected pairs. A pair can be created using the command:



​      \# ip link add <p1-name> type veth peer name <p2-name>



​    In the above, p1-name and p2-name are the names assigned to the two connected end points.



​    Packets transmitted on one device in the pair are immediately received on the other device. When either devices is down the link state of the pair is down.



​    veth device pairs are useful for combining the network facilities of the kernel together in interesting ways. A particularly interesting use case is to place one end of a veth pair in one network namespace and the other

​    end in another network namespace, thus allowing communication between network namespaces. To do this, one first creates the veth device as above and then moves one side of the pair to the other namespace:



​      \# ip link set <p2-name> netns <p2-namespace>



​    ethtool(8) can be used to find the peer of a veth network interface, using commands something like:



​      \# ip link add ve_A type veth peer name ve_B  # Create veth pair

​      \# ethtool -S ve_A     # Discover interface index of peer

​      NIC statistics:

​        peer_ifindex: 16

​      \# ip link | grep '^16:'  # Look up interface

​      16: ve_B@ve_A: <BROADCAST,MULTICAST,M-DOWN> mtu 1500 qdisc ...





IP-LINK(8)                                                   Linux                                                   IP-LINK(8)



NAME

​    ip-link - network device configuration