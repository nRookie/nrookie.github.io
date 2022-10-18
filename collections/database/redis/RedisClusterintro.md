The two major advantages of Redis cluster are:

- The ability to automatically split the dataset among multiple nodes.
- The ability to continue operations when a subset of the nodes are experiencing failures or are unable to communicate with the rest of the cluster.

![image-20220312164059657](/Users/kestrel/developer/nrookie.github.io/collections/database/redis/image-20220312164059657.png)



## Redis cluster working[#](https://www.educative.io/courses/complete-guide-to-redis/B84o59OPnRJ#Redis-cluster-working)

In Redis cluster, each node has two TCP connections open:

1. The normal port which is used for client communication, such as 6379. This connection is available to all clients as well as those cluster nodes that need to use client connection for key migration.
2. The cluster bus port, which is used by other cluster nodes for failure detection, configuration updates, and fail-over authorization. This channel uses less bandwidth by utilizing a binary protocol for data exchange. This port is always 10000 + the normal port. So if the normal port is 6379, then the cluster bus port will be 16379.

### Hash slots[#](https://www.educative.io/courses/complete-guide-to-redis/B84o59OPnRJ#Hash-slots)



Redis cluster uses **hash slots** to shard the keys. A Redis cluster has 16,383 slots. These slots are distributed among the servers. So, if there are 3 servers in a cluster, then:

- Server A contains hash slots from 0 to 5,500.
- Server B contains hash slots from 5,501 to 11,000.
- Server C contains hash slots from 11,001 to 16,383.



> It is not necessary for the slots to be equally distributed among the servers.



Whenever a new server is added or deleted, these slots are redistributed. Moving hash slots from a node to another does not require you to stop operations. Thus, adding and removing nodes, or changing the percentage of hash slots held by nodes does not require any downtime.



When a key is inserted in Redis, the hash slot is found using the CRC-16 hash function to convert the key into an integer. Then modulo 16383 of that integer is calculated to get the **hash slot** for this key.



``` shell
HASH_SLOT = CRC16(key) mod 16383
```



### Hash tags[#](https://www.educative.io/courses/complete-guide-to-redis/B84o59OPnRJ#Hash-tags)





The Redis cluster allows multiple key operations, if the keys are on the same node and have the same hash slot. The user can force multiple keys to be part of the same hash slot by using **hash tags**.

In **hash tags**, a sub-string within {} is added between the key. The hash is then calculated using this sub-string. Thus, all the keys return the same hash. Let’s say we are storing a few keys in Redis, and we want them to end up in the same node. We will insert a random substring(let’s say abcd) in between our keys, e.g., 23464{abcd}2344, relat{abcd}ed, pre34{abcd}pared. When these keys are inserted, the hash slot will be calculated using the string, **abcd**. Thus, all the keys will have the same hash slot.



## Fault tolerance in Redis cluster



The sharding of data won’t help in the case of a server failure. If a server goes down all the keys stored on that server will be lost. To handle this kind of situation, replication is required. Redis Cluster uses a master-slave model, where every hash slot has anywhere from one (the master itself) to N replicas (N-1 additional slave nodes). To be considered healthy, a cluster should have atleast three master nodes and one replica for each master. So, a cluster should have at-least 6 nodes. M1, M2, and M3 are the master nodes, and S1, S2, and S3 are the replicated nodes.



Redis cluster writes data to the slave node asynchronously. This means that when a key is inserted or updated, the master node sends this information to the slave node but does not wait for acknowledgement. If the master node crashes, it is possible that the slave does not contain all the keys that a client has entered. This problem can be solved by storing the data in a slave node synchronously, but doing this will slow down the server. Thus, Redis cluster makes a trade-off between performance and consistency.



## Redis cluster configuration[#](https://www.educative.io/courses/complete-guide-to-redis/B84o59OPnRJ#Redis-cluster-configuration)



By default Redis runs as a single-instance server. If we want to run Redis in cluster mode then we need to provide a new set of directives in the redis.conf file. These directives are defined below.



1. `cluster-enabled`: This property should be **yes** if we need to start the Redis instance in cluster mode. If this property is **no**, then this instance cannot be used in a Redis cluster.
2. `cluster-config-file`: This directive sets the path for the configuration file that stores the changes happening to the cluster. This file is created by Redis and this directive should not be changed by the user. It stores things like the number of nodes in the cluster, their state etc. When a Redis cluster is started, this file is read.
3. `cluster-node-timeout`: This property defines the maximum amount of time in milliseconds, that a node can be unavailable without being considered failed. If a master node is not reachable for more than the specified amount of time, then a slave node is promoted to master. If a slave node is not reachable then it stops accepting requests from clients.
4. `cluster-slave-validity-factor`: If a master is disconnected from the cluster, then a slave tries to fail-over a master. If the value of this property is 0, then a slave will always try to fail-over a master, regardless of the amount of time the link between the master and the slave remains disconnected.





So, if **cluster-node-timeout** is 2 and **cluster-slave-validity-factor** is 10, then the maximum timeout is 20 milliseconds. If a slave is disconnected from a master for more than 20 milliseconds, it is considered unsuitable and cannot be promoted to master.

1. `cluster-migration-barrier`: This directive defines the minimum number of slaves that a master will remain connected with before another slave migrates to a master that is no longer covered by any slave. Let’s say we have two masters, A and B. A has two slaves, called A1 and A2. B has one slave, called B1. If the master B goes down, then B1 is promoted to master. Now master B does not have any slaves. If **cluster-migration-barrier** is set to 2, then B1 will not be able to borrow a slave from A because 2 is the minimum requirement.
2. `cluster-require-full-coverage`: If a master server fails and it does not have any slaves, then the keys that were stored on this server are lost. If this property is set to **yes**, as it is by default, then the entire cluster will become unavailable. If the option is set to no, then the cluster will still serve the queries, but the keys that are lost will return an error.



# Creating Redis cluster



#### Prerequisite

``` shell
yum install gcc
```



#### maybe

``` shell
make distclean
```

``` shell
wget http://download.redis.io/redis-stable.tar.gz
tar xzf redis-stable.tar.gz
cd redis-stable
make
```



``` shell
port 7001
cluster-enabled yes
cluster-config-file nodes.conf
cluster-node-timeout 10000
appendonly yes
```

To create a cluster, run the following command in a new terminal:

``` shell
redis-cli --cluster create 127.0.0.1:7001 127.0.0.1:7002 127.0.0.1:7003 127.0.0.1:7004 127.0.0.1:7005 127.0.0.1:7006 --cluster-replicas 1
```



connecting to master



``` shell
./src/redis-cli -c -p 7001

127.0.0.1:7001> set Alex 54
-> Redirected to slot [15714] located at 127.0.0.1:7003
OK
127.0.0.1:7003> set Zane 32
-> Redirected to slot [9709] located at 127.0.0.1:7002
OK
127.0.0.1:7002> get Alex
-> Redirected to slot [15714] located at 127.0.0.1:7003
"54"
127.0.0.1:7003> get Zane
-> Redirected to slot [9709] located at 127.0.0.1:7002
"32"
127.0.0.1:7002>
```



The **CLUSTER INFO** command can be used to check the cluster status as shown below. It shows that the cluster state is ok, and all the slots are assigned.



``` shell
127.0.0.1:7002> cluster info
cluster_state:ok
cluster_slots_assigned:16384
cluster_slots_ok:16384
cluster_slots_pfail:0
cluster_slots_fail:0
cluster_known_nodes:6
cluster_size:3
cluster_current_epoch:6
cluster_my_epoch:2
cluster_stats_messages_ping_sent:163
cluster_stats_messages_pong_sent:165
cluster_stats_messages_meet_sent:1
cluster_stats_messages_sent:329
cluster_stats_messages_ping_received:165
cluster_stats_messages_pong_received:164
cluster_stats_messages_received:329
127.0.0.1:7002>
```

## Adding nodes to a cluster[#](https://www.educative.io/courses/complete-guide-to-redis/B8KErVGEyr2#Adding-nodes-to-a-cluster)



### Introducing the new instance to the cluster[#](https://www.educative.io/courses/complete-guide-to-redis/B8KErVGEyr2#Introducing-the-new-instance-to-the-cluster)

The **CLUSTER MEET** command is used to introduce a new instance to the cluster. Run the following command in a new terminal to inform the cluster that a new node is available.

``` shell
./redis-stable/src/redis-cli  -c -p 7001 CLUSTER MEET 127.0.0.1 7007
```



### Reshard the hash slots[#](https://www.educative.io/courses/complete-guide-to-redis/B8KErVGEyr2#Reshard-the-hash-slots)

We will move the hash slot, **9709**, which is currently on the node running at port 7002. To move the hash slot, we will first need to find out the source and destination node id. This can be done using the **CLUSTER NODES** command.



``` shell
[root@10-23-184-141 play_redis]# ./redis-stable/src/redis-cli  -c -p 7001 CLUSTER MEET 127.0.0.1 7007
OK
[root@10-23-184-141 play_redis]# 6344:M 12 Mar 2022 17:30:58.284 # IP address for this node updated to 127.0.0.1
6344:M 12 Mar 2022 17:30:58.929 # Configuration change detected. Reconfiguring myself as a replica of 597858662217229ac7204fade6bbc2a07eec7c5a
6344:S 12 Mar 2022 17:30:58.929 * Before turning into a replica, using my own master parameters to synthesize a cached master: I may be able to synchronize with the new master with just a partial transfer.
6344:S 12 Mar 2022 17:30:58.929 * Connecting to MASTER 127.0.0.1:7002
6344:S 12 Mar 2022 17:30:58.929 * MASTER <-> REPLICA sync started
6344:S 12 Mar 2022 17:30:58.930 * Non blocking connect for SYNC fired the event.
6344:S 12 Mar 2022 17:30:58.930 * Master replied to PING, replication can continue...
6344:S 12 Mar 2022 17:30:58.930 * Trying a partial resynchronization (request 34dca00f477b9886568138a92736126c6b6c2553:1).
6344:S 12 Mar 2022 17:30:58.931 * Full resync from master: bd396ba0d92be140375254b5499e6573343e97eb:432
6344:S 12 Mar 2022 17:30:58.931 * Discarding previously cached master state.
6344:S 12 Mar 2022 17:30:58.987 * MASTER <-> REPLICA sync: receiving 189 bytes from master to disk
6344:S 12 Mar 2022 17:30:58.987 * MASTER <-> REPLICA sync: Flushing old data
6344:S 12 Mar 2022 17:30:58.987 * MASTER <-> REPLICA sync: Loading DB in memory
6344:S 12 Mar 2022 17:30:58.988 * Loading RDB produced by version 6.2.6
6344:S 12 Mar 2022 17:30:58.988 * RDB age 0 seconds
6344:S 12 Mar 2022 17:30:58.988 * RDB memory usage when created 2.56 Mb
6344:S 12 Mar 2022 17:30:58.988 # Done loading RDB, keys loaded: 1, keys expired: 0.
6344:S 12 Mar 2022 17:30:58.988 * MASTER <-> REPLICA sync: Finished with success
6344:S 12 Mar 2022 17:30:58.989 * Background append only file rewriting started by pid 6814
6344:S 12 Mar 2022 17:30:59.012 * AOF rewrite child asks to stop sending diffs.
6814:C 12 Mar 2022 17:30:59.012 * Parent agreed to stop sending diffs. Finalizing AOF...
6814:C 12 Mar 2022 17:30:59.012 * Concatenating 0.00 MB of AOF diff received from parent.
6814:C 12 Mar 2022 17:30:59.012 * SYNC append only file rewrite performed
6814:C 12 Mar 2022 17:30:59.013 * AOF rewrite: 4 MB of memory used by copy-on-write
6344:S 12 Mar 2022 17:30:59.029 * Background AOF rewrite terminated with success
6344:S 12 Mar 2022 17:30:59.029 * Residual parent diff successfully flushed to the rewritten AOF (0.00 MB)
6344:S 12 Mar 2022 17:30:59.029 * Background AOF rewrite finished successfully
6344:S 12 Mar 2022 17:30:59.031 # Cluster state changed: ok

[root@10-23-184-141 play_redis]# ./redis-stable/src/redis-cli -c -p 7001 CLUSTER NODEs
597858662217229ac7204fade6bbc2a07eec7c5a 127.0.0.1:7002@17002 master - 0 1647077528210 2 connected 5461-10922
c163cbab99c768cc744b6585e88242f16e36d0c8 127.0.0.1:7007@17007 slave 597858662217229ac7204fade6bbc2a07eec7c5a 0 1647077530216 2 connected
fd2c7194b0eac3a41c3c0e62797d610a94cf3396 127.0.0.1:7004@17004 slave ad81b6954916d2fafe938749802dfa7f2780dc35 0 1647077530000 3 connected
cbc3b3f9e87e40807b9f1c65b8e19326fb407ebc 127.0.0.1:7001@17001 myself,master - 0 1647077529000 1 connected 0-5460
191d65a58f10317c8087f7d8c5779d69136fc9bb 127.0.0.1:7006@17006 slave 597858662217229ac7204fade6bbc2a07eec7c5a 0 1647077531219 2 connected
08d06061000a2fb7fe7b857c7c4b591a1ae311b3 127.0.0.1:7005@17005 slave cbc3b3f9e87e40807b9f1c65b8e19326fb407ebc 0 1647077532222 1 connected
ad81b6954916d2fafe938749802dfa7f2780dc35 127.0.0.1:7003@17003 master - 0 1647077531000 3 connected 10923-16383

```



The steps to migrate a hash slot from one node to another are listed below

####  Import a hash slot[#](https://www.educative.io/courses/complete-guide-to-redis/B8KErVGEyr2#a)-Import-a-hash-slot)

The receiving node will run a command to inform the cluster that it wishes to import a hash slot. The command for this is

``` shell
CLUSTER SETSLOT <hash-slot> IMPORTING <source-id>
```

Here, **hash-slot** is the slot that needs to be moved, and **source-id** is the node id of the node on which this hash slot is currently residing. In our case, we need to import a hash slot from the node running on port 7002. So, we will be running the following command from the node that is going to receive the hash slot.





#### Migrate the hash slot[#](https://www.educative.io/courses/complete-guide-to-redis/B8KErVGEyr2#b)-Migrate-the-hash-slot)



Once the receiving node has run the import command, the source node will run the migrate command. The command for migrating a hash slot is:

``` shell
CLUSTER SETSLOT <hash-slot> MIGRATING <destination-id>:
```



Here, **destination-id** is the node id of the destination server.

Since we are moving the slot to a node running on port 7007, we will run the following command:



``` shell
CLUSTER SETSLOT 9707 MIGRATING 2206481813a1cbe55043f92b05726290737adb7b
```

#### c) Migrating the keys[#](https://www.educative.io/courses/complete-guide-to-redis/B8KErVGEyr2#c)-Migrating-the-keys)



We have migrated the hash slot, 9709, to the new node. Now it’s time to move the keys that were stored on this slot to the new node. The following command will tell us how many keys are in a particular slot.

```shell
CLUSTER COUNTKEYSINSLOT 9707
```



To migrate the key, we should know the key name as well. This can be found using the following command:



``` shell
CLUSTER GETKEYSINSLOT <slot> <amount>
```



Here, amount is the number of key names we want to get.



Once we know the key name, we can run the following command to migrate the key.





### Informing all the nodes about hash slot migration[#](https://www.educative.io/courses/complete-guide-to-redis/B8KErVGEyr2#Informing-all-the-nodes-about-hash-slot-migration)

All the nodes in the cluster should be informed of the new location of the hash slot. This is required so that other nodes can look for the key at the correct location. The command for this operation is:

``` shell
CLUSTER SETSLOT <hash-slot> NODE <owner-id>:


```

Here, **owner-id** is the node id of the instance where the hash slot has been moved.



This command should only be run for master nodes.



``` shell
$ redis-cli -c -p 7001 CLUSTER SETSLOT 9709 NODE
2206481813a1cbe55043f92b05726290737adb7b
$ redis-cli -c -p 7002 CLUSTER SETSLOT 9709 NODE
2206481813a1cbe55043f92b05726290737adb7b
$ redis-cli -c -p 7003 CLUSTER SETSLOT 9709 NODE
2206481813a1cbe55043f92b05726290737adb7b
```



## Removing nodes from a cluster[#](https://www.educative.io/courses/complete-guide-to-redis/B8KErVGEyr2#Removing-nodes-from-a-cluster)





To remove a node, first, we need to migrate all its hash slots to different nodes. A node can only be removed from a cluster when there are no hash slots assigned to it.

After all the slots are migrated from a node, then the following command should be run on all the master nodes.





```
CLUSTER FORGET <node-id>
```



As soon as CLUSTER FORGET is executed, the node is added to a ban list. This ban list exists in order to avoid adding the node back to the cluster when nodes exchange messages. The expiration time for the ban list is 60 seconds, so the above command should be executed in all of the master servers within 60 seconds.
