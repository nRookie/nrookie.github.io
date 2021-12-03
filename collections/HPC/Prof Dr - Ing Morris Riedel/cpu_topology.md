# hwloc







 hwloc - General information about hwloc ("hardware locality").



**DESCRIPTION**

​    hwloc provides command line tools and a C API to obtain the hierarchical map of key computing elements, such

​    as: NUMA memory nodes, shared caches, processor packages, processor cores, and processor "threads".  hwloc

​    also gathers various attributes such as cache and memory information, and is portable across a variety of dif‐

​    ferent operating systems and platforms.



**Definitions**

​    hwloc has some specific definitions for terms that are used in this man page and other hwloc documentation.



​    **hwloc** **CPU** **set:**

​      A set of processors included in an hwloc object, expressed as a bitmask indexed by the physical numbers

​      of the CPUs (as announced by the OS). The hwloc definition of "CPU set" does not carry any the same con‐

​      notations as Linux's "CPU set" (e.g., process affinity, etc.).



​    **hwloc** **node** **set:**

​      A set of NUMA memory nodes near an hwloc object, expressed as a bitmask indexed by the physical numbers

​      of the NUMA nodes (as announced by the OS).



​    **Linux** **CPU** **set:**

​      See http://www.mjmwired.net/kernel/Documentation/cpusets.txt for a discussion of Linux CPU sets. A

​      super-short-ignoring-many-details description (taken from that page) is:



​       "Cpusets provide a mechanism for assigning a set of CPUs and Memory Nodes to a set of tasks."



​    **Linux** **Cgroup:**

​      See http://www.mjmwired.net/kernel/Documentation/cgroups.txt for a discussion of Linux control groups. A

​      super-short-ignoring-many-details description (taken from that page) is:



​       "Control Groups provide a mechanism for aggregating/partitioning sets of tasks, and all their future

​      children, into hierarchical groups with specialized behaviour."



​    To be clear, hwloc supports all of the above concepts. It is simply worth noting that they are different

​    things.









# lstopo











# References



https://www.open-mpi.org/projects/hwloc/tutorials/20120702-POA-hwloc-tutorial.html