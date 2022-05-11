https://www.vmware.com/content/dam/digitalmarketing/vmware/en/pdf/products/vmware-deploy-antrea-k8s-networking.pdf



![image-20220312150721632](/Users/user/playground/share/nrookie.github.io/collections/k8s-related/network/image-20220312150721632.png)



The Antrea Controller watches NetworkPolicy, pod and namespace resources from the Kubernetes API, computes NetworkPolicies and distributes the computed policies to all Antrea Agents.



### Antrea Agent



The Antrea Agent manages the OVS bridge and pod interfaces, and implements pod networking with OVS on every Kubernetes node.



The Antrea Agent exposes a gRPC service (CNI service), which is invoked by the antrea-cni binary to perform CNI operations. For each new pod to be created on the node, after getting the CNI ADD call from antrea-cni, the Agent creates the pod’s network interface, allocates an IP address, connects the interface to the OVS bridge, and installs the necessary flows in OVS. To learn more about the OVS flows, read the OVS pipeline doc.



The Antrea Agent includes two Kubernetes controllers: • The node controller watches the Kubernetes API server for new nodes and creates an OVS (Geneve/VXLAN/GRE/STT) tunnel to each remote node. • The NetworkPolicy controller watches the computed NetworkPolicies from the Antrea Controller API and installs OVS flows to implement the NetworkPolicies for the local pods. The Antrea Agent also exposes a REST API on a local HTTP endpoint for antctl.



### OVS daemons 

The two OVS daemons—ovsds-server and ovs-vswitched—run in a separate container, called antrea-ovs, of the Antrea Agent DaemonSet.

#### antrea-cni 

antrea-cni is the CNI plug-in binary of Antrea. It is executed by kubelet for each CNI command. It is a simple gRPC client that issues an RPC to the Antrea Agent for each CNI command. The Agent performs the actual work (sets up networking for the pod) and returns the result or an error to antrea-cni.



#### antctl

antctl is a command-line tool for Antrea. At the moment, it can show basic runtime information for the Antrea Controller and the Antrea Agent for debugging purposes. When accessing the Controller, antctl invokes the Controller API to query the required information. As previously described, antctl can reach the Controller API through the Kubernetes API, and have the Kubernetes API authenticate, authorize and proxy the API requests to the Controller. antctl also can be executed through kubectl as a kubectl plug-in. When accessing the Agent, antctl connects to the Agent’s local REST endpoint and can only be executed locally in the Agent’s container.



#### Octant UI plug-in

Antrea also implements an Octant plug-in, which can show the Controller and Agent’s health and basic runtime information in the Octant UI. The Octant plug-in gets the Controller and Agent’s information from the AntreaControllerInfo and AntreaAgentInfo Custom Resource Definition (CRD) in the Kubernetes API. A CRD is created by the Antrea Controller and each Antrea Agent to populate their health and runtime information.





### Pod networking

#### Pod interface configuration and IP address management (IPAM)

On every node, the Antrea Agent creates an OVS bridge (named br-int by default), and creates a veth pair for each pod, with one end being in the pod’s network namespace and the other connected to the OVS bridge. On the OVS bridge, the Antrea Agent also creates an internal port—antrea-gw0 by default—to be the gateway of the node’s subnet, and a tunnel port—antrea-tun0—for creating overlay tunnels to other nodes.





![image-20220312151642378](/Users/user/playground/share/nrookie.github.io/collections/k8s-related/network/image-20220312151642378.png)

### 

Each node is assigned a single subnet, and all pods on the node get an IP from the subnet. Antrea leverages Kubernetes’ NodeIPAMController for the node subnet allocation, which sets the podCIDR field of the Kubernetes node spec to the allocated subnet. The Antrea Agent retrieves the subnets of nodes from the podCIDR field. It reserves the first IP of the local node’s subnet to be the gateway IP and assigns it to the antrea-gw0 port, and invokes the host-local IPAM plug-in to allocate IPs from the subnet to all local pods. A local pod is assigned an IP when the CNI ADD command is received for that pod.

![image-20220312152023377](/Users/user/playground/share/nrookie.github.io/collections/k8s-related/network/image-20220312152023377.png)





For every remote node, the Antrea Agent adds an OVS flow to send the traffic to that node through the appropriate tunnel. The flow matches the packet’s destination IP against each node’s subnet:



-  Intra-node traffic – Packets between two local pods will be forwarded by the OVS bridge directly. 

- Inter-node traffic – Packets to a pod on another node will be first forwarded to the antrea-tun0 port, encapsulated and sent to the destination node through the tunnel. Then, they will be decapsulated, injected through the antrea-tun0 port to the OVS bridge and, finally, forwarded to the destination pod.

- Pod to external traffic – Packets sent to an external IP or the node’s network will be forwarded to the antrea-gw0 port (as it is the gateway of the local pod subnet), routed (based on routes configured on the node) to the appropriate network interface of the node (e.g., a physical network interface for a bare-metal node), and sent out to the node network from there. The Antrea Agent creates an iptables (MASQUERADE) rule to perform SNAT on the packets from pods, so their source IP will be rewritten to the node’s IP before going out.





![image-20220312152240280](/Users/user/playground/share/nrookie.github.io/collections/k8s-related/network/image-20220312152240280.png)



At the moment, Antrea leverages kube-proxy to serve traffic for ClusterIP and NodePort type services. The packets from a pod to a service’s ClusterIP will be forwarded through the antrea-gw0 port, then kube-proxy will select one service back-end pod to be the connection’s destination and DNAT the packets to the pod’s IP and port. If the destination pod is on the local node, the packets will be forwarded to the pod directly. If it is on another node, the packets will be sent to that node via the tunnel. 



kube-proxy can be used in any supported mode: user-space iptables or IPVS. See the Kubernetes Service documentation for more details.





