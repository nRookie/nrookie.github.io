## Defining Underlay and Overlay Networking

Underlay networks refer to the physical network infrastructure: DWDM equipment (in the case of wide-area networks), eithernet swtiches and routers (from vendors like Arista , Cisco, Juniper, and Nokia), and the cable plant physical infrastructure such as fiber optic cabling that connects all these network devices into a network topology.


Underlay networks can be Layer 2 or Layer 3 networks. Layer 2 underlay networks today are typically based on Ethernet, with segmentation accomplished via VLANs. The Internet is an example of a Layer 3 underlay network, where Autonomous System run control planes based on interiror gateway protocols (IGPs) such as OSPF and ISIS , and BGP serves as the Internet-wide routing protocol. And Multi-Protocol Label Switched (MPLS) networks are a legacy underlay WAN technology that falls between Layer 2 and Layer 3.

By contrast, overlay networks implement network virtualization concepts. A virtualized network consists of overlay nodes (e.g., routers) where Layer 2 and Layer 3 tunneling encapsulation (VXLAN, GRE, and IPSec) serves as the transport overlay protocol sometimes referred to as OTV (Overlay Transport Virtualization).



There are two prominent examples of virtual network overlays. The first and best known are SD-WAN architecture that rely heavily on VPN functionality to replace MPLS circuits, making it less costly and easier to connect various branch offices, retail locations, and other remote sites to a WAN. The other example is cloud-native networking, where encapsulating traffic with VPN tunnels is the preferred method of connecting VPCs to enterprise locations.


## Overlay Network Achilles Heel: The Internet

Overlay networks offer notable benefits, including software-driven network automation and VPN privacy between tunnel endpoints. Recently, providers of multi-cloud networking have created solutions that further abstract the per CSP networking logic so that it’s easier to manage overlay connectivity between clouds. These are all good things.

However, overlay networks can’t escape the gravitational pull of the Internet as an underlying network. VPN tunnels or not, the Internet isn’t private, is rife with security threats, and exacts a significant latency tax on traffic flows due to its shared, collective nature. Even cloud provider backbones are shared network services that function as extensions of the Internet.

Furthermore, there are workflows that just don’t belong in a VPN tunnel, such as when you’re:


Building a WAN between colocation data centers
Connecting a digital operations backbone to reach edge locations
Moving significant numbers of user flows to a critical enterprise cloud or SaaS application
Transporting significant volumes of application, transaction, or data replication traffic on a hybrid or multi-cloud basis
Trying to do anything latency-sensitive