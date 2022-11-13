## What is a messaging system ?





A messging system is responsible for transferring data among services,

Applications, processes, or servers. Such a system helps decouple different parts of a distributed system by providing an asynchronous way of transferring messaging between the sender and the receiver. Hence, all senders (or producers) and receivers(or cusumers) focus on the data/message without worrying about the mechanism used to share the data.



![image-20221025105619326](/Users/kestrel/developer/nrookie.github.io/collections/design_related/grokking advanced system design/image-20221025105619326.png)





There are two common ways to handle messages: Queuing and Publish-Subscribe.





## Queue 

In the queuing model, messages are stored sequentially in a queue.Producers push messages to the rear of the queue, and consumers extractthe messages from the front of the queue.



![image-20221025105757854](/Users/kestrel/developer/nrookie.github.io/collections/design_related/grokking advanced system design/image-20221025105757854.png)



### publish-subscribe messaging system 





In the pub-sub (short for polish-subscribe) model, messages are divided into topics. A publisher (or a producer) sends a message to a topic that gets stored in the messaging system under that topic. Subscribers (or the consumer) subscribe to a topic to receive every message published to that topic, Unlike the Queuing model, the pub-sub model allows multiple consumers to get the same message; if two consumers subscribe to the same topic, they will receive all messages published to that topic.





![image-20221025110054669](/Users/kestrel/developer/nrookie.github.io/collections/design_related/grokking advanced system design/image-20221025110054669.png)



![image-20221025110121826](/Users/kestrel/developer/nrookie.github.io/collections/design_related/grokking advanced system design/image-20221025110121826.png)



The message broker stores published messages in a queue, and subscribers read them from the queue. Hence, subscribers and publishers do not have to be synchronoized. This loose coupling enables subscribers and publishers to read and write messages at different rates.





The messaging system's ability to store messages provides fault-tolerance. so messages do not get lost between the time they are produced and the time they are consumed.





To summarize , a message system is deployed in an application stack for the following reason.





1. Messaging buffering. *To provide a buffering mechanism in front of*

   *processing (i.e., to deal with temporary incoming message spikes that*

   *are greater than what the processing app can deal with). This enables*

   *the system to safely deal with spikes in workloads by temporarily*

   *storing data until it is ready for processing.*

2. Guarantee of message delivery.  *Allows producers to publish messages*

   *with assurance that the message will eventually be delivered if the*

   *consuming application is unable to receive the message when it is*

   *published.*

3. Providing abstraction.  *A messaging system provides an architectural*

   *separation between the consumers of messages and the applications*

   *producing the messages.*

4. Enabling scale.  *Provides a flexible and highly configurable architecture*

   *that enables many producers to deliver messages to multiple*

   *consumers.*





At a high level, we can call Kafka a distributed Commit Log. A commit Log (also known as a Write-Ahead log or a Transactions log) is an append-only data structure that can persistently store a sequence of records. Records are always appended to the end of the log, and once added, records cannot be deleted or modified. Reading from a commit log always happens from left to right.



![image-20221025110806970](/Users/kestrel/developer/nrookie.github.io/collections/design_related/grokking advanced system design/image-20221025110806970.png)

Kafka stores all of its messages on disk. Since all reads and writes happen insequence, Kafka takes advantage of sequential disk reads (more on thislater).





## Kafka Use cases



1. Metrics:
2. Log Aggregation:
3. Stream processing:
4. Commit Log:
5. Website activity tracking:
6. Product suggestions



### *Kafka common terms*



#### Brokers

*A Kafka server is also called a broker. Brokers are responsible for reliably*

*storing data provided by the producers and making it available to the*

*consumers.*



#### *Records*



*A record is a message or an event that gets stored in Kafka. Essentially, it is*

*the data that travels from producer to consumer through Kafka. A record*

*contains a key, a value, a timestamp, and optional metadata headers.*



![image-20221025111208677](/Users/kestrel/developer/nrookie.github.io/collections/design_related/grokking advanced system design/image-20221025111208677.png)



### Topics #

*Kafka divides its messages into categories called Topics. In simple terms, a*

*topic is like a table in a database, and the messages are the rows in that table.*



*Each message that Kafka receives from a producer is associated with a*

*topic.*

*Consumers can subscribe to a topic to get notified when new messages*

*are added to that topic.*

*A topic can have multiple subscribers that read messages from it.*

*In a Kafka cluster, a topic is identified by its name and must be unique.*



![image-20221025111350062](/Users/kestrel/developer/nrookie.github.io/collections/design_related/grokking advanced system design/image-20221025111350062.png)



*Producers*

*Producers are applications that publish (or write) records to Kafka.*



*Consumers*



*Consumers are the applications that subscribe to (read and process) data*

*from Kafka topics. Consumers subscribe to one or more topics and consume*

*published messages by pulling data from the brokers.*

*In Kafka, producers and consumers are fully decoupled and agnostic of each*

*other, which is a key design element to achieve the high scalability that*

*Kafka is known for. For example, producers never need to wait for*

*consumers.*





### *High-level architecture*



*At a high level, applications (producers) send messages to a Kafka broker,*

*and these messages are read by other applications called consumers.*

*Messages get stored in a topic, and consumers subscribe to the topic to*

*receive new messages.*



*Kafka is run as a cluster of one or more servers, where each server is*

*responsible for running one Kafka broker.*







#### *Kafka cluster*



#### *ZooKeeper*

*ZooKeeper is a distributed key-value store and is used for coordination and*

*storing configurations. It is highly optimized for reads. Kafka uses ZooKeeper*

*to coordinate between Kafka brokers; ZooKeeper maintains metadata*

*information about the Kafka cluster. 





![image-20221025111538423](/Users/kestrel/developer/nrookie.github.io/collections/design_related/grokking advanced system design/image-20221025111538423.png)



## *Kafka: Deep Dive*



Kafka is simply a collection of topics. As topics can get quite big, they aresplit into partitions of a smaller size for better performance and scalability.





### *Topic partitions*



*Kafka topics are partitioned, meaning a topic is spread over a number of*

*‘fragments’. Each partition can be placed on a separate Kafka broker. When*

*a new message is published on a topic, it gets appended to one of the topic’s*

*partitions. The producer controls which partition it publishes messages to*

*based on the data. For example, a producer can decide that all messages*

*related to a particular ‘city’ go to the same partition.*



*Essentially, a partition is an ordered sequence of messages. Producers*

*continually append new messages to partitions. Kafka guarantees that all*

*messages inside a partition are stored in the sequence they came in.*



***Ordering of messages is maintained at the partition level, not across the***

***topic.***



![image-20221025111906750](/Users/kestrel/developer/nrookie.github.io/collections/design_related/grokking advanced system design/image-20221025111906750.png)



- A unique sequence ID called an offset gets assigned to every messagethat enters a partition. These numerical offsets are used to identifyevery message’s sequential position within a topic’s partition.

- *Offset sequences are unique only to each partition. This means, to locate*

  *a specific message, we need to know the Topic, Partition, and Offset*

  *number.***

- *Producers can choose to publish a message to any partition. If ordering*

*within a partition is not needed, a round-robin partition strategy can be*

*used, so records get distributed evenly across partitions.*

- *Placing each partition on separate Kafka brokers enables multiple*

  *consumers to read from a topic in parallel. That means, different*

  *consumers can concurrently read different partitions present on*

  *separate brokers.**

- *Placing each partition of a topic on a separate broker also enables a*

*topic to hold more data than the capacity of one server.*

- Messages once written to partitions are immutable and cannot beupdated.
- A producer can add a ‘**key**’ to any message it publishes. Kafkaguarantees that messages with the same key are written to the samepartition.
- *Each broker manages a set of partitions belonging to different topics.*





Kafka follows the principle of **dumb broker** and **smart consumer**.





*This* *means that Kafka does not keep track of what records are read by the*

*consumer. Instead, consumers, themselves, poll Kafka for new messages and*

*say what records they want to read. This allows them to*

*increment/decrement the offset they are at as they wish, thus being able to*

*replay and reprocess messages. Consumers can read messages starting from*

*a specific offset and are allowed to read from any offset they choose. This*

*also enables consumers to join the cluster at any point in time.*



*Every topic can be replicated to multiple Kafka brokers to make the data*

*fault-tolerant and highly available. Each topic partition has one leader*

*broker and multiple replica (follower) brokers.*



### Leader #

*A leader is the node responsible for all reads and writes for the given*

*partition. Every partition has one Kafka broker acting as a leader.*



### Follower #

*To handle single point of failure, Kafka can replicate partitions and*

*distribute them across multiple broker servers called followers. Each*

*follower’s responsibility is to replicate the leader’s data to serve as a ‘backup’*

*partition. This also means that any follower can take over the leadership if*

*the leader goes down.*





In the following diagram, we have two partitions and four brokers. Broker1 is the leader of Partition 1 and follower of Partition 2 .



![image-20221025112609318](/Users/kestrel/developer/nrookie.github.io/collections/design_related/grokking advanced system design/image-20221025112609318.png)





*Kafka stores the location of the leader of each partition in ZooKeeper. As all*

*writes/reads happen at/from the leader, producers and consumers directly*

*talk to ZooKeeper to find a partition leader.*







### *In-sync replicas*



An in-sync replica (ISR) is a broker that has the latest data for a givenpartition. A leader is always an in-sync replica. A follower is an in-syncreplica only if it has fully caught up to the partition it is following. In otherwords, ISRs cannot be behind on the latest records for a given partition.Only ISRs are eligible to become partition leaders. Kafka can choose theminimum number of ISRs required before the data becomes available forconsumers to read.





### *High-water mark*



*To ensure data consistency, the leader broker never returns (or exposes)*

*messages which have not been replicated to a minimum set of ISRs. For this,*

 *brokers keep track of the high-water mark,*





**which is the highest offset that all ISRs of a particular partition share. The leader exposes data only up to the high-water mark offset and propagates the high-water mark offset to all followers. Let’s understand this with an example.**



*In the figure below, the leader does not return messages greater than offset*

*‘4’, as it is the highest offset message that has been replicated to all follower*

*brokers.*



![image-20221025112955821](/Users/kestrel/developer/nrookie.github.io/collections/design_related/grokking advanced system design/image-20221025112955821.png)



If a consumer reads the record with offset ‘7’ from the leader (Broker 1), andlater, if the current leader fails, and one of the followers becomes the leaderbefore the record is replicated to the followers, the consumer will not be ableto find that message on the new leader. The client, in this case, willexperience a non-repeatable read. Because of this possibility, Kafka brokersonly return records up to the high-water mark.





### *Consumer Groups*



*Kafka ensures that only a single consumer reads messages from any*

*partition within a consumer group. In other words, topic partitions are a*

*unit of parallelism – only one consumer can work on a partition in a*

*consumer group at a time. If a consumer stops, Kafka spreads partitions*

*across the remaining consumers in the same consumer group. Similarly,*

*every time a consumer is added to or removed from a group, the*

*consumption is rebalanced within the group.*



![image-20221025113441736](/Users/kestrel/developer/nrookie.github.io/collections/design_related/grokking advanced system design/image-20221025113441736.png)



Kafka stores the current offset per consumer group per topic per partition,as it would for a single consumer. This means that unique messages are onlysent to a single consumer in a consumer group, and the load is balancedacross consumers as equally as possible.When the number of consumers exceeds the number of partitions in a topic,all new consumers wait in idle mode until an existing consumerunsubscribes from that partition. Similarly, as new consumers join aconsumer group, Kafka initiates a rebalancing if there are more consumersthan partitions. Kafka uses any unused consumers as failovers.124BackKafka: Deep DiveNextKafka Work owHere is a summary of how Kafka manages the distribution of partitions toconsumers within a consumer group:



![image-20221025113642396](/Users/kestrel/developer/nrookie.github.io/collections/design_related/grokking advanced system design/image-20221025113642396.png)





*A critical dependency of Apache Kafka is Apache ZooKeeper, which is a*

*distributed configuration and synchronization service. ZooKeeper serves as*

*the coordination interface between the Kafka brokers, producers, and*

*consumers. Kafka stores basic metadata in ZooKeeper, such as information*

*about brokers, topics, partitions, partition leader/followers, consumer*

*offsets, etc.*



*As we know, Kafka brokers are stateless; they rely on ZooKeeper to maintain*

*and coordinate brokers, such as notifying consumers and producers of the*

*arrival of a new broker or failure of an existing broker, as well as routing all*

*requests to partition leaders.*



*It maintains the last offset position of each consumer group per*

*partition, so that consumers can quickly recover from the last position*

*in case of a failure (although modern clients store offsets in a separate*

*Kafka topic).*



*It tracks the topics, number of partitions assigned to those topics, and*

*leaders’/followers’ location in each partition.*



*It also manages the access control lists (ACLs) to different topics in the*

*cluster. ACLs are used to enforce access or authorization.*





How do producers or consumers find out whothe leader of a partition is? #





*clients fetch*

*metadata information from Kafka brokers directly; brokers talk to*

*ZooKeeper to get the latest metadata. In the diagram below, the producer*

*goes through the following steps before publishing a message:*





*1. The producer connects to any broker and asks for the leader of*

*‘Partition 1’.*

*2. The broker responds with the identification of the leader broker*

*responsible for ‘Partition 1’.*

*3. The producer connects to the leader broker to publish the message.*



![image-20221026125355272](/Users/kestrel/developer/nrookie.github.io/collections/design_related/grokking advanced system design/image-20221026125355272.png)





*All the critical information is stored in the ZooKeeper and ZooKeeper*

*replicates this data across its cluster, therefore, failure of Kafka broker (or*

*ZooKeeper itself) does not affect the state of the Kafka cluster. Upon*

*ZooKeeper failure, Kafka will always be able to restore the state once the* *ZooKeeper restarts after failure. Zookeeper is also responsible for* *coordinating the partition leader election between the Kafka brokers in case*

*of leader failure.*



What is the controller broker? #Within the Kafka cluster, one broker is elected as the Controller. ThisController broker is responsible for admin operations, such ascreating/deleting a topic, adding partitions, assigning leaders to partitions,monitoring broker failures, etc. Furthermore, the Controller periodicallychecks the health of other brokers in the system. In case it does not receive aresponse from a particular broker, it performs a failover to another broker.It also communicates the result of the partition leader election to otherbrokers in the system.



*Split brain*



*When a controller broker dies, Kafka elects a new controller. One of the*

*problems is that we cannot truly know if the leader has stopped for good and*

*has experienced an intermittent failure like a stop-the-world GC pause or a* temporary network disruption. Nevertheless, the cluster has to move on and pick a new controller. If the original Controller had an intermittent failure,the cluster would end up having a so-called zombie controller. A zombiecontroller can be defined as a controller node that had been previouslydeemed dead by the cluster and has come back online. Another broker hastaken its place, but the zombie controller might not know that yet. Thiscommon scenario in distributed systems with two or more active controllers(or central servers) is called split-brain.





*Generation clock*





*Split-brain is commonly solved with a generation clock, which is simply a*

*monotonically increasing number to indicate a server’s generation. In Kafka,*

*the generation clock is implemented through an epoch number. If the old*

*leader had an epoch number of ‘1’, the new one would have ‘2’. This epoch is*

*included in every request that is sent from the Controller to other brokers.*

*This way, brokers can now easily differentiate the real Controller by simply*

*trusting the Controller with the highest number. The Controller with the*

*highest number is undoubtedly the latest one, since the epoch number is*

*always increasing. This epoch number is stored in ZooKeeper.*







*Kafka Delivery Semantics*



*As we know, a producer writes only to the leader broker, and the followers*

*asynchronously replicate the data. How can a producer know that the data is*

*successfully stored at the leader or that the followers are keeping up with the*

*leader? Kafka offers three options to denote the number of brokers that must*

*receive the record before the producer considers the write as successful:*



- Async: Producer sends a message to Kafka and does not wait foracknowledgment from the server. This means that the write isconsidered successful the moment the request is sent out. This fire-andforgetapproach gives the best performance as we can write data toKafka at network speed, but no guarantee can be made that the serverhas received the record in this case.
- Committed to Leader: Producer waits for an acknowledgment from theleader. This ensures that the data is committed at the leader; it will beslower than the ‘Async’ option, as the data has to be written on disk onthe leader. Under this scenario, the leader will respond without waitingfor acknowledgments from the followers. In this case, the record will be lost if the leader crashes immediately after acknowledging the producer but before the followers have replicated it.
- Committed to Leader and Quorum: Producer waits for anacknowledgment from the leader and the quorum. This means theleader will wait for the full set of in-sync replicas to acknowledge therecord. This will be the slowest write but guarantees that the record willnot be lost as long as at least one in-sync replica remains alive. This isthe strongest available guarantee.



*As we can see, the above options enable us to configure our preferred tradeoff*

*between durability and performance.*



- *If we would like to be sure that our records are safely stored in Kafka,*

  *we have to go with the last option – Committed to Leader and Quorum.*

- *If we value latency and throughput more than durability, we can choose*

  *one of the first two options. These options will have a greater chance of*

  *losing messages but will have better speed and throughput.*



### *Consumer delivery semantics*



*A consumer can read only those messages that have been written to a set of*

*in-sync replicas. There are three ways of providing consistency to the*

*consumer:*



- At-most-once (Messages may be lost but are never redelivered): In thisoption, a message is delivered a maximum of one time only. Under thisoption, the consumer upon receiving a message, commit (or increment)the offset to the broker. Now, if the consumer crashes before fullyconsuming the message, that message will be lost, as when theconsumer restarts, it will receive the next message from the lastcommitted offset.
- At-least-once (Messages are never lost but maybe redelivered): Underthis option, a message might be delivered more than once, but nomessage should be lost. This scenario occurs when the consumerreceives a message from Kafka, and it does not immediately commit theoffset. Instead, it waits till it completes the processing. So, if theconsumer crashes after processing the message but before committingthe offset, it has to reread the message upon restart. Since, in this case,the consumer never committed the offset to the broker, the broker willredeliver the same message. Thus, duplicate message delivery couldhappen in such a scenario.
- Exactly-once (each message is delivered once and only once): It is veryhard to achieve this unless the consumer is working with a transactional system. Under this option, the consumer puts the message processing and the offset increment in one transaction. This will ensure that the offset increment will happen only if the whole transaction is complete. If the consumer crashes while processing, the transaction will be rolled back, and the offset will not be incremented. When the consumer restarts, it can reread the message as it failed to process it last time. This option leads to no data duplication and no data loss but canlead to decreased throughput.
- 







### *Kafka Characteristics*



### *Storing messages to disks*



*Kafka writes its messages to the local disk and does not keep anything in*

*RAM. Disks storage is important for durability so that the messages will not*

*disappear if the system dies and restarts. Disks are generally considered to*

*be slow. However, there is a huge difference in disk performance between*

*random block access and sequential access. Random block access is slower*

*because of numerous disk seeks, whereas the sequential nature of writing or*

*reading, enables disk operations to be thousands of times faster than*

*random access. Because all writes and reads happen sequentially, Kafka has*

*a very high throughput.*





Writing or reading sequentially from disks are heavily optimized by the OS,via read-ahead (prefetch large block multiples) and write-behind (groupsmall logical writes into big physical writes) techniques.



*Also, modern operating systems cache the disk in free RAM. This is called*

*Pagecache*。*Since Kafka stores*

*messages in a standardized binary format unmodified throughout the whole*

*flow*





*it can make use of the zero-copy*optimization. That is when the* *operating system copies data from the Pagecache directly to a socket,* *effectively bypassing the Kafka broker application entirely.*



"**Zero-copy**" describes [computer](https://en.wikipedia.org/wiki/Computer) operations in which the [CPU](https://en.wikipedia.org/wiki/Central_processing_unit) does not perform the task of copying data from one [memory](https://en.wikipedia.org/wiki/RAM) area to another or in which unnecessary data copies are avoided. This is frequently used to save CPU cycles and memory bandwidth in many time consuming tasks





*Kafka has a protocol that groups messages together. This allows network*

*requests to group messages together and reduces network overhead. The*

*server, in turn, persists chunks of messages in one go, and consumers fetch*

*large linear chunks at once.*



*All of these optimizations allow Kafka to deliver messages at near networkspeed.*





### *Record retention in Kafka*

*By default, Kafka retains records until it runs out of disk space. We can set*

*time-based limits (configurable retention period), size-based limits*

*(configurable based on size), or compaction (keeps the latest version of*

*record using the key). For example, we can set a retention policy of three*

*days, or two weeks, or a month, etc. The records in the topic are available for*

*consumption until discarded by time, size, or compaction.*



