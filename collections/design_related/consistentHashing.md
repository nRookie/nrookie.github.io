Distributed systems can use Consistent Hashing to distribute data across nodes. Consistent Hashing maps data to physical nodes and ensures that only a **small set of keys move when servers are added or removed**.






Consistent Hashing stores the data managed by a distributed system in a ring. Each node in the ring is assigned a range of data. Here is an example of the consistent hash ring:





*1. How do we know on which node a particular piece of data will be*

*stored?*



2. *When we add or remove nodes, how do we know what data will be*

   *moved from existing nodes to the new nodes? Additionally, how can we*

   *minimize data movement when nodes join or leave?*



![image-20221019155415929](/Users/kestrel/developer/nrookie.github.io/collections/design_related/image-20221019155415929.png)
