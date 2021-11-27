Virtual IP address

A virtual IP address (VIP or VIPA) is an IP address that doesn't correspond to an actual physical network interface. Uses for VIPs include network address translation (especially, one-to-many NAT), fault-tolerance, and mobility.


## Usage


For one-to-many NAT, a VIP address is advertised from the NAT device (often a router), and incoming data packets destined to that VIP address are routed to different actual IP addresses (with address translation). These VIP addresses have several variations and implementation scenarios, including Common Address Redundancy Protocol (CARP) and Proxy ARP.[1] In addition, if there are multiple actual IP addresses, load balancing can be performed as part of NAT.

VIP addresses are also used for connection redundancy by providing alternative fail-over options for one machine. For this to work, the host has to run an interior gateway protocol like Open Shortest Path First (OSPF), and appear as a router to the rest of the network. It advertises virtual links connected via itself to all of its actual network interfaces. If one network interface fails, normal OSPF topology reconvergence will cause traffic to be sent via another interface.[2][3]

A VIP address can be used to provide nearly unlimited mobility. For example, if an application has an IP address on a physical subnet, that application can be moved only to a host on that same subnet. VIP addresses can be advertised on their own subnet,[a] so its application can be moved anywhere on the reachable network without changing addresses.[2]



https://en.wikipedia.org/wiki/Virtual_IP_address