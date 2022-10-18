*Cassandra: Introduction*



*Design a distributed and scalable system that can store a huge amount of*

*structured data, which is indexed by a row key where each row can have an*

*unbounded number of columns.*



Cassandra is a distributed, decentralized, scalable, and highly availableNoSQL database. In terms of CAP theorem



AP (i.e.,available and partition tolerant)



*Cassandra can be tuned with replication-factor and consistency* *levels to meet strong consistency requirements, but this comes with a* *performance cost.*



Cassandra uses peer-to-peer architecture, with each node connected to allother nodes. *Each Cassandra node performs all database operations and can* *serve client requests without the need for any leader node.*





#### *Cassandra use cases*



Cassandra is optimized for highthroughput and faster writes, and can be used for collecting big data forperforming real-time analysis. Here are some of its top use cases:



- **Storing key-value data with high availability** - Reddit and Digg useCassandra as a persistent store for their data. Cassandra’s ability to scalelinearly without any downtime makes it very suitable for their growthneeds.



- **Time series data model** - Due to its data model and log-structured storage engine, Cassandra benefits from high-performing write operations. This also makes Cassandra well suited for storing and analyzing sequentially captured metrics (i.e., measurements from sensors, application logs, etc.). Such usages take advantage of the fact that columns in a row are determined by the application, not a predefined schema. Each row in a table can contain a different number of columns, and there is no requirement for the column names to match.
- **Write-heavy applications** - Cassandra is especially suited for write intensive applications such as time-series streaming services, sensor logs, and Internet of Things (IoT) applications.



## *High-level Architecture*





*Cassandra common terms*



Column: A column is a key-value pair and is the most basic unit of datastructure.



Column key: Uniquely identifies a column in a row.Column value: Stores one value or a collection of values.



Row: A row is a container for columns referenced by primary key.

 Cassandra does not store a column that has a null value; this saves a lot of space.

![image-20221017180125712](/Users/kestrel/developer/nrookie.github.io/collections/design_related/grokking advanced system design/image-20221017180125712.png)





Table: A table is a container of rows.



Keyspace: Keyspace is a container for tables that span over one or moreCassandra nodes.



Cluster: Container of Keyspaces is called a cluster.



Node: Node refers to a computer system running an instance of Cassandra. Anode can be a physical host, a machine instance in the cloud, or even aDocker container.



NoSQL: Cassandra is a NoSQL database which means we cannot have joins between tables, there are no foreign keys , and while querying, we cannot add any column in the where clause other than the primary key. These constraints should be kept in mind before deciding to use Cassandra.



### High-level architecture #



#### Data partitioning #



Cassandra uses consistent hashing for data partitioning. Please take a look at Dynamo’s data partitioning



### *Cassandra keys*





### The Primary key uniquely identifies each row of a table. In Cassandra primary key has two parts:

<img src="/Users/kestrel/developer/nrookie.github.io/collections/design_related/grokking advanced system design/image-20221017180345444.png" alt="image-20221017180345444" style="zoom:33%;" />







The partition key decides which node stores the data, and the clustering keydecides how the data is stored within a node. Let’s take the example of atable with PRIMARY KEY ( city_id , employee_id ). This primary key has twoparts represented by the two columns:



1. city_id is the partition key. This means that the data will bepartitioned by the city_id field, that is, all rows with the samecity_id will reside on the same node.
2. employee_id is the clustering key. This means that within each node,the data is stored in sorted order according to the employee_id column.



#### *Clustering keys*



*clustering keys define how the data is stored within a* *node*



![image-20221017180546909](/Users/kestrel/developer/nrookie.github.io/collections/design_related/grokking advanced system design/image-20221017180546909.png)



#### *Partitioner*

*Partitioner is the component responsible for determining how data is* *distributed on the Consistent Hash ring.*



<img src="/Users/kestrel/developer/nrookie.github.io/collections/design_related/grokking advanced system design/image-20221017180630939.png" alt="image-20221017180630939" style="zoom:50%;" />



Murmur3 *hashing function*



*All Cassandra nodes learn about the token assignments of other nodes*

*through gossip (discussed later). This means any node can handle a request* *for any other node’s range.*





*The node receiving the request is called the* **coordinator**, and any node can act in this role.





#### Coordinator node #



*A client may connect to any node in the cluster to initiate a read or write*

*query. This node is known as the coordinator node.*



<img src="/Users/kestrel/developer/nrookie.github.io/collections/design_related/grokking advanced system design/image-20221017180847311.png" alt="image-20221017180847311" style="zoom:33%;" />

## *Replication*



*Each node in Cassandra serves as a replica for a different range of data.*

*Cassandra stores multiple copies of data and spreads them across various*

*replicas, so that if one node is down, other replicas can respond to queries*

*for that range of data. This process of replicating the data on to different*

*nodes depends upon two factors:*



![image-20221017180939269](/Users/kestrel/developer/nrookie.github.io/collections/design_related/grokking advanced system design/image-20221017180939269.png)







### *Replication factor*



*The replication factor is the number of nodes that will receive the copy of the*

*same data. This means, if a cluster has a replication factor of 3, each row will*

*be stored on three different nodes. Each keyspace in Cassandra can have a*

*different replication factor.*





### *Replication strategy*



*The node that owns the range in which the hash of the partition key falls will*

*be the first replica; all the additional replicas are placed on the consecutive*

*nodes. Cassandra places the subsequent replicas on the next node in a*

*clockwise manner. There are two replication strategies in Cassandra:*





### *Simple replication strategy*



*This strategy is used only for a single data center cluster. Under this strategy,*

*Cassandra places the first replica on a node determined by the partitioner*

*and the subsequent replicas on the next node in a clockwise manner.*



![image-20221017181148961](/Users/kestrel/developer/nrookie.github.io/collections/design_related/grokking advanced system design/image-20221017181148961.png)







### Network topology strategy #



*This strategy is used for multiple data-centers. Under this strategy, we can*

*specify different replication factors for different data-centers. This enables*

*us to specify how many replicas will be placed in each data center.*

*Additional replicas are always placed on the next nodes in a clockwise*

*manner.*





![image-20221017181340609](/Users/kestrel/developer/nrookie.github.io/collections/design_related/grokking advanced system design/image-20221017181340609.png)





### *Cassandra Consistency Levels*





*What are Cassandra’s consistency* *levels?*





*Cassandra’s consistency level is defined as the minimum number of*

*Cassandra nodes that must fulfill a read or write operation before the*

*operation can be considered successful.*



*There is always a tradeoff between consistency and performance.*



### *Write consistency levels*



*For write operations, the consistency level specifies how many replica nodes*

*must respond for the write to be reported as successful to the client. The*

*consistency level is specified per query by the client. Because Cassandra is*

*eventually consistent, updates to other replica nodes may continue in the*

*background. Here are different write consistency levels that Cassandra*

*offers:*



- One or Two or Three: The data must be written to at least the specified number of replica nodes before a write is considered successful.

- Quorum: The data must be written to at least a quorum (or majority) of replica nodes. Quorum is defined as floor(RF /2 + 1), where RF represents the replication factor. For example, in a cluster with a replication factor of five, if three nodes return success, the write is considered successful.

- All: Ensures that the data is written to all replica nodes. This consistency level provides the highest consistency but lowest availability as writes will fail if any replica is down.

- Local_Quoram: Ensures that the data is written to a quorum of nodes in the same data center as the coordinator. It does not wait for the response from the other data-centers.

- Each_Quorum: Ensures that the data is written to a quorum of nodes in each data center.

- Any: The data must be written to at least one node. In the extreme case,when all replica nodes for the given partition key are down, the write can still succeed after a hinted handoff (discussed below) has been written. ‘Any’ consistency level provides the lowest latency and highest availability, however, it comes with the lowest consistency. If all replica nodes are down at write time, an ‘Any’ write is not readable until the replica nodes for that partition have recovered and the latest data is*

  *written on them.*







### *Hinted handoff*



*Depending upon the consistency level, Cassandra can still serve write*

*requests even when nodes are down. For example, if we have the replication*

*factor of three and the client is writing with a quorum consistency level. This*

*means that if one of the nodes is down, Cassandra can still write on the*

*remaining two nodes to fulfill the consistency level, hence, making the write*

*successful.*



![image-20221017181943713](/Users/kestrel/developer/nrookie.github.io/collections/design_related/grokking advanced system design/image-20221017181943713.png)



*When a node is down or does not respond to a write request, the coordinator*

*node writes a hint in a text file on the local disk.* *This hint contains the data*

*itself along with information about which node the data belongs to. When*

*the coordinator node discovers from the Gossiper (will be discussed later)*

*that a node for which it holds hints has recovered, it forwards the write*

*requests for each hint to the target. Furthermore, each node every ten*

*minutes checks to see if the failing node, for which it is holding any hints,*

*has recovered.*



*With consistency level ‘Any,’ if all the replica nodes are down, the*

*coordinator node will write the hints for all the nodes and report success to*

*the client. However, this data will not reappear in any subsequent reads* *until one of the replica nodes comes back online, and the coordinator node* *successfully forwards the write requests to it. This is assuming that the*

*coordinator node is up when the replica node comes back. This also means*

*that we can lose our data if the coordinator node dies and never comes back.*

*For this reason, we should avoid using the ‘Any’ consistency level.*



*If a node is offline for some time, the hints can build up considerably on*

*other nodes. Now, when the failed node comes back online, other nodes tend*

*to flood that node with write requests. This can cause issues on the node, as*

*it is already trying to come back after a failure. To address this problem,*

*Cassandra limits the storage of hints to a configurable time window. It is also*

*possible to disable hinted handoff entirely.*





Cassandra, by default, stores hints for three hours. After three hours, olderhints will be removed, which means, if now the failed node recovers, it willhave stale data. Cassandra can fix this stale data while serving a readrequest. Cassandra can issue a Read Repair when it sees stale data; we willgo through this while discussing the read path.





*One thing to remember: When the cluster cannot meet the consistency level*

*specified by the client, Cassandra fails the write request and does not store a*

*hint.*



#### *Read consistency levels*

*The consistency level for read queries specifies how many replica nodes*

*must respond to a read request before returning the data. For example, for a*

*read request with a consistency level of quorum and replication factor of*

*three, the coordinator waits for successful replies from at least two nodes.*



To achieve strong consistency in Cassandra: R + W > RF gives us strong consistency. In this equation, R, W , and RF are the read replica count, the write replica count, and the replication factor, respectively. All client reads will see the most recent write in this scenario, and we will have strong consistency.



**Snitch**: The Snitch is an application that determines the proximity of nodes within the ring and also tells which nodes are faster. Cassandra nodes use this information to route read/write requests efficiently. We will discuss thisin detail later.





![image-20221017182427686](/Users/kestrel/developer/nrookie.github.io/collections/design_related/grokking advanced system design/image-20221017182427686.png)



*How does Cassandra perform a read operation?*



*The coordinator always* *sends the read request to the fastest node. For example, for Quorum=2, the*

*coordinator sends the request to the fastest node and the digest of the data*

*from the second-fastest node. The digest is a checksum of the data and is*

*used to save network bandwidth.*



If the digest does not match, it means some replicas do not have the latestversion of the data. 

In this case, the coordinator reads the data from all thereplicas to determine the latest data. 

The coordinator then returns the latestdata to the client and initiates a read repair request. The read repair operation pushes the newer version of data to nodes with the older version.



<img src="/Users/kestrel/developer/nrookie.github.io/collections/design_related/grokking advanced system design/image-20221017182523974.png" alt="image-20221017182523974" style="zoom:33%;" />



*operation is used as an opportunity to repair inconsistent data across*

*replicas. The latest write-timestamp is used as a marker for the correct*

*version of data. The read repair operation is performed only in a portion of*

*the total reads to avoid performance degradation. Read repairs are*

*opportunistic operations and not a primary operation for anti-entropy.*



Anti-entropy is **a process of comparing the data of all replicas and updating each replica to the newest version**. 



***Read Repair Chance:*** *When the read consistency level is less than ‘All,’*

*Cassandra performs a read repair probabilistically. By default, Cassandra*

*tries to read repair 10% of all requests with DC local read repair. In this case,*

*Cassandra immediately sends a response when the consistency level is met*

*and performs the read repair asynchronously in the background.*



### Snitch #

*Snitch keeps track of the network topology of Cassandra nodes. It determines*

*which data-centers and racks nodes belong to. Cassandra uses this*

*information to route requests efficiently. Here are the two main functions of*

*a snitch in Cassandra:*



- *Snitch determines the proximity of nodes within the ring and also*

*monitors the read latencies to avoid reading from nodes that have*

*slowed down. Each node in Cassandra uses this information to route*

*requests efficiently.*

- *Cassandra’s replication strategy uses the information provided by the*

  *Snitch to spread the replicas across the cluster intelligently. Cassandra*

  *will do its best by not having more than one replica on the same “rack”.*



*To understand Snitch’s role, let’s take the example of Cassandra’s read*

*operation. Let’s assume that the client is performing a read with a quorum*

*consistency level, and the data is replicated on five nodes. To support*

*maximum read speed, Cassandra selects a single replica to query for the full*



*object and asks for the digest of the data from two additional nodes in order*

*to ensure that the latest version of the data is returned. The Snitch helps to*

*identify the fastest replica, and Cassandra asks this replica for the full object.*





#### *Gossiper*



#### *How does Cassandra use gossip protocol?*





Cassandra uses gossip protocol that allows each node to keep track of stateinformation about the other nodes in the cluster. Nodes share stateinformation with each other to stay in sync. Gossip protocol is a peer-to-peercommunication mechanism in which nodes periodically exchange stateinformation about themselves and other nodes they know about. Each nodeinitiates a gossip round every second to exchange state information aboutthemselves (and other nodes) with one to three other random nodes. Thisway, all nodes quickly learn about all other nodes in a cluster.



*Each gossip message has a version associated with it, so that during a gossip*

*exchange, older information is overwritten with the most current state for a*

*particular node.*





***Generation number**:*  *In Cassandra, each node stores a generation number*

*which is incremented every time a node restarts*. *The* *generation number remains the same while the node is alive and is*

*incremented each time the node restarts.* *If the generation number in the gossip*

*message is higher, it knows that the node was restarted.*





***Seed nodes:***  *The seed node designation* *has no purpose other than bootstrapping the gossip process for new nodes*

*joining the cluster. Thus, seed nodes are not a single point of failure, nor do*

*they have any other special purpose in cluster operations other than the*

*bootstrapping of nodes.*







![image-20221017183237625](/Users/kestrel/developer/nrookie.github.io/collections/design_related/grokking advanced system design/image-20221017183237625.png)

![image-20221017183247538](/Users/kestrel/developer/nrookie.github.io/collections/design_related/grokking advanced system design/image-20221017183247538.png)







### *Node failure detection*



*Heartbeating uses a fixed timeout, and if there is*

*no heartbeat from a server, the system, after the timeout, assumes that the*

*server has crashed. Here the value of the timeout is critical. If we keep the*

*timeout short, the system will be able to detect failures quickly but with*

*many false positives due to slow machines or faulty networks. On the other*

*hand, if we keep the timeout long, the false positives will be reduced, but the*

*system will not perform efficiently for being slow in detecting failures.*





Cassandra uses an adaptive failure detection mechanism as described by **PhiAccrual Failure Detector**







*outputs the*

*suspicion level about a server; a higher suspicion level means there are*

*higher chances that the server is down. Using Phi Accrual Failure Detector, if*

*a node does not respond, its suspicion level is increased and could be*

*declared dead later. As a node’s suspicion level increases, the system can*

*gradually decide to stop sending new requests to it. Phi Accrual Failure*

*Detector makes a distributed system efficient as it takes into account*

*fluctuations in the network environment and other intermittent server*

*issues before declaring a system completely dead.*





***Anatomy of Cassandra's Write Operation***



*Cassandra stores data both in memory and on disk to provide both high*

*performance and durability. Every write includes a timestamp. Write path*

*involves a lot of components, here is the summary of Cassandra’s write path:*





1. Each write is appended to a commit log, which is stored on disk.
2. Then it is written to MemTable in memory.
3. Periodically, MemTables are flushed to SSTables on the disk.
4. *Periodically, compaction runs to merge SSTables.*



#### *Commit log*



*When a node receives a write request, it immediately writes the data to a*

*commit log. The commit log is a write-ahead log and is stored on disk. It is*

*used as a crash-recovery mechanism to support Cassandra’s durability goals.*

*A write will not be considered successful on the node until it’s written to the*

*commit log; this ensures that if a write operation does not make it to the in-*

memory store (the MemTable, discussed in a moment), it will still be possible to recover the data. If we shut down the node or it crashes unexpectedly, the commit log can ensure that data is not lost. That’s because if the node restarts, the commit log gets replayed. 



![image-20221017183702541](/Users/kestrel/developer/nrookie.github.io/collections/design_related/grokking advanced system design/image-20221017183702541.png)



### MemTable #

*After it is written to the commit log, the data is written to a memory-resident*

*data structure called the MemTable.*



- *Each node has a MemTable in memory for each Cassandra table.*

- *Each MemTable contains data for a specific Cassandra table, and it*

  *resembles that table in memory.*

- *Each MemTable accrues writes and provides reads for data not yet*

  *flushed to disk.*

- *Commit log stores all the writes in sequential order, with each new*

  *write appended to the end, whereas MemTable stores data in the sorted*

  *order of partition key and clustering columns.*

- *After writing data to the Commit Log and MemTable, the node sends an*

  *acknowledgment to the coordinator that the data has been successfully*

  *written.*

- 





![image-20221017183838449](/Users/kestrel/developer/nrookie.github.io/collections/design_related/grokking advanced system design/image-20221017183838449.png)



### SStable #



*When the number of objects stored in the MemTable reaches a threshold, the*

*contents of the MemTable are flushed to disk in a file called SSTable. At this*

*point, a new MemTable is created to store subsequent data. This flushing is a *

*non-blocking operation; multiple MemTables may exist for a single table, one*

*current, and the rest waiting to be flushed. Each SStable contains data for a* *specific table.*



*When the MemTable is flushed to SStables, corresponding entries in the*

*Commit Log are removed.*



### *Why are they called ‘SSTables’?*

*The term ‘SSTables’ is short for ‘**Sorted* *String Table’** and first appeared in Google’s **Bigtable** which is also a storage*

*system. Cassandra borrowed this term even though **it does not store data as***

***strings on the disk.***



*Once a MemTable is flushed to disk as an SSTable, it is immutable and cannot*

*be changed by the application. If we are not allowed to update SSTables, how*

*do we delete or update a column? In Cassandra, **each delete or update is***

***considered a new write operation.** We will look into this in detail while*

*discussing **Tombstones**.*





*The current data state of a Cassandra table consists of its MemTables in*

*memory and SSTables on the disk. Therefore, on reads, Cassandra will read*

*both SSTables and MemTables to find data values, as the MemTable may*

*contain values that have not yet been flushed to the disk. The MemTable*

*works like a write-back cache that Cassandra looks up by key.*





**Generation number** is an index number that is incremented every time anew SSTable is created for a table and is used to uniquely identify SSTables.Here is the summary of Cassandra’s write path:





![image-20221017184150413](/Users/kestrel/developer/nrookie.github.io/collections/design_related/grokking advanced system design/image-20221017184150413.png)







### *Anatomy of Cassandra's Read Operation*





#### Caching #





*To boost read performance, Cassandra provides three optional forms of*

*caching:*



1. **Row cache**: The row cache, caches frequently read (or hot) rows. It stores a complete data row, which can be returned directly to the client if requested by a read operation. This can significantly speed up read access for frequently accessed rows, at the cost of more memory usage.

2. **Key cache**: Key cache stores a map of recently read partition keys to their SSTable offsets. This facilitates faster read access into SSTables stored on disk and improves the read performance but could slow down*

   *the writes, as we **have to update the Key cache for every write**.*

3. ***Chunk cache:*** *Chunk cache is used to store uncompressed chunks of*

   *data read from SSTable files that are accessed frequently.*

   

### Reading from MemTable #

![image-20221017184526737](/Users/kestrel/developer/nrookie.github.io/collections/design_related/grokking advanced system design/image-20221017184526737.png)



![image-20221017184550608](/Users/kestrel/developer/nrookie.github.io/collections/design_related/grokking advanced system design/image-20221017184550608.png)



#### Reading from SSTable #



#### Bloom filters #



*Each SStable has a Bloom filter associated with it, which tells if a particular*

*key is present in it or not. Bloom filters are used to boost the performance of*

*read operations. Bloom filters are very fast, non-deterministic algorithms for*

*testing whether an element is a member of a set. They are non-deterministic*

*because it is possible to get a false-positive read from a Bloom filter,*  *but*

*false-negative is not possible. Bloom filters work by mapping the values in a*

*data set into a bit array and condensing a larger data set into a digest string*

*using a hash function.* *The digest, by definition, uses a much smaller amount*

*of memory than the original data would. The filters are stored in memory*

*and are used to improve performance by reducing the need for disk access*

*on key lookups. Disk access is typically much slower than memory access. So,*

*in a way, a Bloom filter is a special kind of key cache.*



*Cassandra maintains a Bloom filter for each SSTable. When a query is*

*performed, the Bloom filter is checked first before accessing the disk.*

*Because false negatives are not possible, if the filter indicates that the*

*element does not exist in the set, it certainly does not; but if the filter thinks*

*that the element is in the set, the disk is accessed to make sure.*







### How are SSTables stored on the disk? #



*Each SSTable consists of two files:*



1. **Data File**: Actual data is stored in a data file. It has partitions and rows associated with those partitions. The partitions are in sorted order.
2. **Partition Index file**: Stored on disk, partition index file stores thesorted partition keys mapped to their SSTable offsets. It enables locatinga partition exactly in an SSTable rather than scanning data.

![image-20221017184841858](/Users/kestrel/developer/nrookie.github.io/collections/design_related/grokking advanced system design/image-20221017184841858.png)





#### *Partition index summary file*





![image-20221017184925547](/Users/kestrel/developer/nrookie.github.io/collections/design_related/grokking advanced system design/image-20221017184925547.png)





