

## Goal

*Design a distributed key-value store that is highly available (i.e., reliable),*

*highly scalable, and completely decentralized.*





### *What is Dynamo?*



*Dynamo is a highly available key-value store developed by Amazon for*

*their internal use. Many Amazon services, such as shopping cart, bestseller*

*lists, sales rank, product catalog, etc., need only primary-key access to data. A*

*multi-table relational database system would be an overkill for such services*

*and would also limit scalability and availability. Dynamo provides a flexible*

*design to let applications choose their desired level of availability and*

*consistency.*





**sloppy quorum**

**gossip protocol**

***hinted handoff.***

***vector clocks***

***Merkle trees***



***Data Partitioning***



*The act of distributing data across a set of nodes is called data partitioning.*



*How do we know on which node a particular piece of data will be*

*stored?*



<img src="/Users/kestrel/developer/nrookie.github.io/collections/design_related/grokking advanced system design/image-20221011151633326.png" alt="image-20221011151633326" style="zoom:33%;" />



*we add or remove a server,*

*we have to remap all the keys and move the data based on the new server*

*count, which will be a complete mess!*



*When we add or remove nodes, how do we know what data will be*

*moved from existing nodes to the new nodes? Furthermore, how can we*

*minimize data movement when nodes join or leave?*



***Dynamo uses consistent hashing to solve these problems. The consistent***

***hashing algorithm helps Dynamo map rows to physical nodes and also***

***ensures that only a small set of keys move when servers are added or***

***removed.***



*Consistent hashing represents the data managed by a cluster as a ring. Each*

*node in the ring is assigned a range of data. Dynamo uses the consistent*

*hashing algorithm to determine what row is stored to what node. Here is an*

*example of the consistent hashing ring:*



<img src="/Users/kestrel/developer/nrookie.github.io/collections/design_related/grokking advanced system design/image-20221011151949678.png" alt="image-20221011151949678" style="zoom:33%;" />



*the start*

*of the range is called a token*





*Range start: Token value*

*Range end: Next token value - 1*



<img src="/Users/kestrel/developer/nrookie.github.io/collections/design_related/grokking advanced system design/image-20221011152259691.png" alt="image-20221011152259691" style="zoom:50%;" />



*The consistent hashing scheme described above works great when a node is*

*added or removed from the ring; as only the next node is affected in these* scenarios

However, this scheme can result in non-uniform data and load distribution. Dynamo solves these issues with the help of Virtual nodes.



### Virtual nodes

Adding and removing nodes in any distributed system is quite common. 

Existing nodes can die and may need to be decommissioned. Similarly, new

nodes may be added to an existing cluster to meet growing demands.



***Adding or removing nodes**: Adding or removing nodes will result in*

*recomputing the tokens causing a significant administrative overhead*

*for a large cluster.*



***Hotspots**: Since each node is assigned one large range, if the data is not*

*evenly distributed, some nodes can become hotspots.*





***Node rebuilding***: *Since each node’s data is replicated on a fixed*

*number of nodes (discussed later), when we need to rebuild a node, only*

*its replica nodes can provide the data. This puts a lot of pressure on the*

*replica nodes and can lead to service degradation.*



*To handle these issues, Dynamo introduced a new scheme for distributing*

*the tokens to physical nodes. Instead of assigning a single token to a node,*

*the hash range is divided into multiple smaller ranges, and each physical*

*node is assigned multiple of these smaller ranges. Each of these subranges is*

*called a Vnode. With Vnodes, instead of a node being responsible for just one*

*token, it is responsible for many tokens (or subranges).*





<img src="/Users/kestrel/developer/nrookie.github.io/collections/design_related/grokking advanced system design/image-20221011153055884.png" alt="image-20221011153055884" style="zoom:33%;" />





*Practically, Vnodes are **randomly distributed** across the cluster and are*

*generally non-contiguous so that no two neighboring Vnodes are assigned to*

*the same physical node. Furthermore, nodes do carry replicas of other nodes*

*for fault-tolerance. Also, since there can be heterogeneous machines in the*

*clusters, some servers might hold more Vnodes than others. The figure*

*below shows how physical nodes A, B, C, D, & E are using Vnodes of the*

*Consistent Hash ring. Each physical node is assigned a set of Vnodes and*

*each Vnode is replicated once.*



![image-20221011153304537](/Users/kestrel/developer/nrookie.github.io/collections/design_related/grokking advanced system design/image-20221011153304537.png)







*Vnodes help spread the load more evenly across the physical nodes on*

*the cluster by dividing the hash ranges into smaller subranges. This*

*speeds up the rebalancing process after adding or removing nodes.*

*When a new node is added, it receives many Vnodes from the existing*

*nodes to maintain a balanced cluster. Similarly, when a node needs to*

*be rebuilt, instead of getting data from a fixed number of replicas, many*

*nodes participate in the rebuild process.*



*Vnodes make it easier to maintain a cluster containing*

*heterogeneous machines. This means, with Vnodes, we can assign a*

*high number of ranges to a powerful server and a lower number of*

*ranges to a less powerful server.*



*Since Vnodes help assign smaller ranges to each physical node, the*

*probability of hotspots is much less than the basic Consistent Hashing*

*scheme which uses one big range per node.*





## *Replication*



*Dynamo provides an **eventually* consistent** model. This replication technique is called optimistic

*replication, which means that replicas are not guaranteed to be identical at*

*all times.*



<img src="/Users/kestrel/developer/nrookie.github.io/collections/design_related/grokking advanced system design/image-20221011153624469.png" alt="image-20221011153624469" style="zoom:33%;" />





Dynamo does not enforce strict quorum requirements, andinstead uses something called **sloppy quorum**. With this approach, all read/write operations are performed on the first N healthy nodes from thepreference list, which may not always be the first N nodes encounteredwhile moving clockwise on the consistent hashing ring.

<img src="/Users/kestrel/developer/nrookie.github.io/collections/design_related/grokking advanced system design/image-20221017163847825.png" alt="image-20221017163847825" style="zoom:33%;" />



### *Hinted handoff*

***when a node is unreachable, another node can accept* *writes on its behalf***

*The write is then kept in a local buffer and sent out* *once the destination node is reachable again.*

This makes Dynamo “**always writeable.**



*The main problem is that since a sloppy quorum is not a strict majority, the*

*data can and will diverge, i.e., it is possible for two concurrent writes to the*

*same key to be accepted by non-overlapping sets of nodes. This means that*

*multiple conflicting values against the same key can exist in the system, and*

*we can get stale or conflicting data while reading. Dynamo allows this and*

*resolves these conflicts using **Vector Clocks**.*





## *Vector Clocks and Conflicting Data*



### *What is clock skew?*



On a single machine, all we need to know about is the absolute or wall clocktime: suppose we perform a write to key k with timestamp t1 and thenperform another write to k with timestamp t2 . Since t2 > t1 , the secondwrite must have been newer than the first write, and therefore the databasecan safely overwrite the original value.





In a distributed system, this assumption does not hold. The problem is **clockskew**.

different clocks tend to run at different rates, so we cannotassume that time t on node a happened before time t + 1 on node b .



*The* *most practical techniques that help with synchronizing clocks, like NTP, still*

*do not guarantee that every clock in a distributed system is synchronized at*

*all times. So, without special hardware like GPS units and atomic clocks, just*

*using wall clock timestamps is not enough.*





### What is a vector clock? 

Instead of employing tight synchronization mechanics, Dynamo uses something called **vector clock** in order to capture causality between different versions of the same object.



A **vector clock** is effectively a (node,counter) pair.



*One can determine whether two versions of an* *object are on parallel branches or have a causal ordering by examining their*

*vector clocks.*



1. Server A serves a write  to key k1, with value foo. It assigns it a version of [A:1] . This write gets replicated to server B.
2. Server A. serves a write to key k1, with value bar, It assigns it a version of [A:2]. This write also gets replicated to server B.
3. A network partition occurs. A and B cannot talk to each other.
4. Server A serves a write to key k1, with value bar . It assigns it a version of [A:3] . It cannot replicate it to server B, but it gets stored in a hinted handoff buffer on another server.
5. Server B sees a write to key k1, with value bax. It assigns it a version of [B:1]. It cannot replicate it to server A. but it gets stored in a hinted handoff buffer on another server.

6. The network heals. Server A and B can talk to each other again.

7. Either server gets a read request for key k1. It sees the same key with different versions [A:3] and [A:2][B:1], but it does not know which one is newer. It returns both and tells the client to figure out the version and write the newer version back into the system.



![image-20221017165231430](/Users/kestrel/developer/nrookie.github.io/collections/design_related/grokking advanced system design/image-20221017165231430.png)

![image-20221017165243987](/Users/kestrel/developer/nrookie.github.io/collections/design_related/grokking advanced system design/image-20221017165243987.png)



*semantic reconciliation*,

A typical example of a collapse operation is "merging" different versions of a customer's shopping cart. Using this reconciliation mechanism, an add operation(i.e., adding an item to the cart) is nerver lost. However, deleted items can resurface.



``` shell
Resolving conflicts is similar to how Git works. If Git can merge different versions into one, merging is done automatically. If not, the client(i.e. the developer) has to reconcile conflicts manually
```





Dynamo truncates vector clocks (oldest first) when they grow too large. If Dynamo ends up deleting older vector clocks that are required to reconcile an object's state, Dynamo would not be able to achieve eventual consistency. Dynamo's authors note that this is potential problem but do not specify how this may be addressed. They do mention that this problem has not yet surfaced in any of their production systems.



## Conflict-free replicated data types (CRDTs) #



A more straightforward CRDTs.



*Amazon’s shopping cart is an excellent example of**CRDT*



The idea that any two nodes that have received the same set of updates will see the same end result is called strong eventual consistency.



### Last-write-wins (LWW) #



*Unfortunately, it is not easy to model the data as CRDTs. In many cases, it*

*involves too much effort. Therefore, vector clocks with client-side resolution*

*are considered good enough.*



Instead of vector clocks, Dynamo also offers ways to resolve the conflictsautomatically on the server-side. Dynamo (and Apache Cassandra) often usesa simple conflict resolution policy: last-write-wins (LWW), based on thewall-clock timestamp. LWW can easily end up losing data. For example, iftwo conflicting writes happen simultaneously, it is equivalent to flipping acoin on which write to throw away.



## Strategies for choosing thecoordinator node #





- *Clients can route their requests through a generic load balancer.*

- *Clients can use a partition-aware client library that routes the requests*

  *to the appropriate coordinator nodes with lower latency.*



![image-20221017170253703](/Users/kestrel/developer/nrookie.github.io/collections/design_related/grokking advanced system design/image-20221017170253703.png)





### *Consistency protocol*



Dynamo uses a consistency protocol similar to quorum systems. If R/W isthe minimum number of nodes that must participate in a successfulread/write operation respectively:

- Then R + W > N yields a quorum-like system

- A Common (N , R, W ) configuration used by Dynamo is (3, 2, 2).
  - (3, 3, 1): fast W , slow R, not very durable
  - (3, 1, 3): fast R, slow W , durable
- In this model, the latency of a get() (or put() ) operation depends upon the slowest of the replicas. For this reason, R and W are usually configured to be less than N to provide better latency.
- In general, low values of W and R increase the risk of inconsistency, as write requests are deemed successful and returned to the clients even if a majority of replicas have not processed them. This also introduces a vulnerability window for durability when a write request is successfully returned to the client even though it has been persisted at only a small number of nodes.
- For both Read and Write operations, the requests are forwarded to thefirst ‘N ’ healthy nodes.



## ‘put()’ process #





Dynamo’s put() request will go through the following steps:



1. The coordinator generates a new data version and vector clock*

*component.*

2. Saves new data locally.
3. Sends the write request to N − 1 highest-ranked healthy nodes from the preference list.
4. The put() operation is considered successful after receiving W − 1confirmation.



### *‘get()’ process*

1. The coordinator requests the data version from N − 1 highest-rankedhealthy nodes from the preference list.
2. Waits until R − 1 replies.
3. *Coordinator handles causal data versions through a vector clock.*
4. *Returns all relevant data versions to the caller.*



## *Request handling through state* *machine*



 





# *Anti-entropy Through Merkle Trees*





## *What are Merkle trees?*

A replica can contain a lot of data. Naively splitting up the entire data rangefor checksums is not very feasible; there is simply too much data to be transferred. Therefore, Dynamo uses Merkle trees to compare replicas of a range.



*A Merkle tree is a binary tree of hashes, where each internal node is* *the hash of its two children, and each leaf node is a hash of a portion of the* *original data.*





![image-20221017171246060](/Users/kestrel/developer/nrookie.github.io/collections/design_related/grokking advanced system design/image-20221017171246060.png)



*Comparing Merkle trees is conceptually simple:*



*1. Compare the root hashes of both trees.*

*2. If they are equal, stop.*

*3. Recurse on the left and right children.*



*Ultimately, this means that replicas know precisely which parts of the range*

*are different, and the amount of data exchanged is minimized.*





### *Merits and demerits of Merkle trees*



The principal advantage of using a Merkle tree is that each branch of the tree can be checked independently without requiring nodes to download the entire tree or the whole data set. Hence, Merkle trees minimize the amount of data that needs to be transferred for synchronization and reduce the number of disk reads performed during the anti-entropy process.





The disadvantage of using Merkle trees is that many key ranges can change when a node joins or leaves, and as a result , the trees need to be recalculated.



## *Gossip Protocol*



### *What is gossip protocol?*



In a Dynamo cluster, since we do not have any central node that keeps track of all nodes to know if a node is down or not, how does a node know every other node’s current state? The simplest way to do this is to have every node maintain heartbeats with every other node. When a node goes down, it will stop sending out heartbeats, and everyone else will find out immediately. But then O(N ) messages get sent every tick (N being the number of nodes),which is a ridiculously high amount and not feasible in any sizable cluster.



Dynamo uses **gossip protocol** that enables each node to keep track of state information about the other nodes in the cluster, like which nodes are reachable, what key ranges they are responsible for, and so on (this isbasically a copy of the hash ring).





Nodes share state information with each other to stay in sync. Gossip protocol is a **peer-to-peer communicationmechanism** in which nodes periodically exchange state information about themselves and other nodes they know about. Each node initiates a gossip round every second to exchange state information about itself and other *nodes with one other random node. This means that any new event will* *eventually propagate through the system, and all nodes quickly learn about* *all other nodes in a cluster.*



<img src="/Users/kestrel/developer/nrookie.github.io/collections/design_related/grokking advanced system design/image-20221017172129586.png" alt="image-20221017172129586" style="zoom:50%;" />





<img src="/Users/kestrel/developer/nrookie.github.io/collections/design_related/grokking advanced system design/image-20221017172208593.png" alt="image-20221017172208593" style="zoom:33%;" />





## External discovery through seed nodes #



*As we know, Dynamo nodes use gossip protocol to find the current state of* *the ring. This can result in a logical partition of the cluster in a particular* *scenario. Let’s understand this with an example:*





An administrator joins node A to the ring and then joins node B to the ring.

Nodes A and B consider themselves part of the ring, yet neither would be immediately aware of each other.

To prevent these logical partitions,Dynamo introduced the concept of **seed nodes**.

*Seed nodes are fully* *functional nodes and can be obtained either from a static configuration or a*

*configuration service.*  *This way, all nodes are aware of seed nodes.*



*Each* *node communicates with seed nodes through gossip protocol to reconcile*

*membership changes; therefore, logical partitions are highly unlikely.*





## *Dynamo Characteristics and Criticism*



Because Dynamo is completely decentralized and does not rely on acentral/leader server (unlike GFS, for example), each node serves three functions:



1. Managing get() and put() requests: A node may act as a coordinator and manage all operations for a particular key or may forward the request to the appropriate node.

2. Keeping track of membership and detecting failures: Every node uses gossip protocol to keep track of other nodes in the system and their associated hash ranges.

3. Local persistent storage: Each node is responsible for being either theprimary or replica store for keys that hash to a specific range of values.These (key, value) pairs are stored within that node using various storage systems depending on application needs. A few examples of such storage systems are:

   *BerkeleyDB Transactional Data Store*

   *MySQL (for large objects)*

   *An in-memory buffer (for best performance) backed by persistent*

   *storage*





*Characteristics of Dynamo*



*Distributed:*



*Decentralized:*



*Scalable:*



*Highly Available:*



*Fault-tolerant and reliable:*



*Tunable consistency:*



*Durable:*



*Eventually Consistent:*







### *Criticism on Dynamo*

*Each Dynamo node contains the entire Dynamo routing table. This is*

*likely to affect the scalability of the system as this routing table will*

*grow larger and larger as nodes are added to the system.*





*Dynamo seems to imply that it strives for symmetry, where every node*

*in the system has the same set of roles and responsibilities, but later, it*

*specifies some nodes as seeds. Seeds are special nodes that are*

*externally discoverable. These are used to help prevent logical*

*partitions in the Dynamo ring. This seems like it may violate Dynamo’s*

*symmetry principle.*



Although security was not a concern as Dynamo was built for internaluse only, DHTs can be susceptible to several different types ofattacks. While Amazon can assume a trusted environment, sometimes abuggy software can act in a manner quite similar to a malicious actor.





Dynamo’s design can be described as a “leaky abstraction,” whereclient applications are often asked to manage inconsistency, and theuser experience is not 100% seamless. For example, inconsistencies inthe shopping cart items may lead users to think that the website isbuggy or unreliable.





*Two of the most famous datastores built on the principles of*

*Dynamo are Riak (https://riak.com/) and Cassandra*





![image-20221017173046149](/Users/kestrel/developer/nrookie.github.io/collections/design_related/grokking advanced system design/image-20221017173046149.png)







Summary #



![image-20221017173206913](/Users/kestrel/developer/nrookie.github.io/collections/design_related/grokking advanced system design/image-20221017173206913.png)



 ![image-20221017173220428](/Users/kestrel/developer/nrookie.github.io/collections/design_related/grokking advanced system design/image-20221017173220428.png)



![image-20221017173253772](/Users/kestrel/developer/nrookie.github.io/collections/design_related/grokking advanced system design/image-20221017173253772.png)





![image-20221017173317611](/Users/kestrel/developer/nrookie.github.io/collections/design_related/grokking advanced system design/image-20221017173317611.png)





