## 4.1 CPU Topology

### 4.1.1 CPU and NUMA Topology

The first computer processors were uniprocessors, meaning that the system had a single CPU. The illusion of executing processes in parallel was done by the operating system rapidly switching the single CPU from one thread of execution (process) to another. In the quest for increasing system performance, designers noted that increasing the clock rate to execute instructions faster only worked up to a point (usually the limitations on creating a stable clock waveform with the current technology). In an effort to get more overall system performance, designers added another CPU to the system, allowing two parallel streams of execution. This trend of adding processors has continued over time. In an effort to get more overall system performance, designers added another CPU to the system, allowing two parallel streams of execution. This trend of adding processors has continued over time.


Most early multiprocessor systems were designed so that each CPU had the same logical path to each memory location (usually a parallel bus). This let each CPU access any memory location in the same amount of time as any other CPU in the system. This type of architecture is known as a Symmetric Multi-Processor (SMP) system. SMP is fine for a small number of CPUs, but once the CPU count gets above a certain point (8 or 16), the number of parallel traces required to allow equal access to memory uses too much of the available board real estate, leaving less room for peripherals.


Two new concepts combined to allow for a higher number of CPUs in a system:

1. Serial buses
2. NUMA topologies

A serial bus is a single-wire communication path with a very high clock rate, which transfers data as packetized bursts. Hardware designers began to use serial buses as high-speed interconnects between CPUs, and between  CPUs and memory controllers and other peripherals. This means that instead of requiring between 32 and 64 traces on the board from each CPU to the memory subsystem, there was now one trace, substantially reducing the amount of space required on the board.


At the same time, hardware designers were packing more transistor into the same space by reducing die sizes, 
Instead of putting individual CPUs directly onto the main board, they started packing them into a processor package as multi-core processors. Then, instead of trying to provide equal access to memory from each processor package, designers resorted to a Non-Uniform Memory Access (NUMA) strategy, where each package package/socket combination has one or more dedicated memory area for high speed access. Each socket also has an interconnect to other sockets for slower access to the other socket's memory.


As a simple NUMA example, suppose we have a two-socket motherboard, where each socket has been populated with a quad-core package. This means the total number of CPUs in the system is eight; four in each socket. Each socket also has an attached memory bank with four gigabytes of RAM, for a total system memory of eight gigabytes. For the purposes of this example, CPUs 0-3 are in socket 0, and CPUs 4-7 are in socket 1. Each socket in this example also corresponds to a NUMA node.


It might take three clock cycles for CPU 0 to access memory from bank 0: a cycle to present the address to the memory controller, a cycle to set up access to the memory location, and a cycle to read or write to the location. 


However, it might take six clock cycles for CPU 4 to access memory from the same location; because it is on a separate socket, it must go through two memory controllers: the local memory controller on socket 1, and then the remote memory controller on socket 0. If memory is contested on that location (that is, if more than one CPU is attempting to access the same location simultaneously), memory controllers need to arbitrate and serialize access to the memory, so memory access will take longer. Adding cache consistency (ensuring that local CPU caches contain the same data for the same memory location) complicates the process further.

The latest high-end processors from both Intel (Xeon) and AMD (Opteron) have NUMA topologies. The AMD processors use an interconnect known as HyperTransport™ or HT, while Intel uses one named QuickPath Interconnect™ or QPI. The interconnects differ in how they physically connect to other interconnects, memory, or peripheral devices, but in effect they are a switch that allows transparent access to one connected device from another connected device. In this case, transparent refers to the fact that there is no special programming API required to use the interconnect, not a "no cost" option.



Because system architectures are so diverse, it is impractical to specifically characterize the performance penalty imposed by accessing non-local memory. We can say that each hop across an interconnect imposes at least some relatively constant performance penalty per hop, so referencing a memory location that is two interconnects from the current CPU imposes at least 2N + memory cycle time units to access time, where N is the penalty per hop.


Given this performance penalty, performance-sensitive applications should avoid regularly accessing remote memory in a NUMA topology system. The application should be set up so that it stays on a particular node and allocates memory from that node.


1. To do this, there are a few things that applications will need to know:

1. What is the topology of the system?
2. Where is the application currently executing?
3. Where is the closest memory bank?


## 4.1.2  Tuning CPU Performance


Read this section to understand how to tune for better CPU performance, and for an introduction to several tools that aid in the process.


NUMA was originally used to connect a single processor to multiple memory banks. As CPU manufacturers refined their processes and die sizes shrank, multiple CPU cores could be included in one package. These CPU cores were clustered so that each had equal access time to a local memory bank, and cache could be shared between the cores; however, each 'hop' across an interconnect between core, memory, and cache involves a small performance penalty.



The example system in Figure 4.1, “Local and Remote Memory Access in NUMA Topology” contains two NUMA nodes. Each node has four CPUs, a memory bank, and a memory controller. Any CPU on a node has direct access to the memory bank on that node. Following the arrows on Node 1, the steps are as follows:


1. A CPU (any of 0-3) presents the memory address to the local memory controller.
2. The memory controller sets up access to the memory address.
3. The CPU performs read or write operations on that memory address.



However, if a CPU on one node needs to access code that resides on the memory bank of a different NUMA node, the path it has to take is less direct:

A CPU (any of 0-3) presents the remote memory address to the local memory controller.


The CPU's request for that remote memory address is passed to a remote memory controller, local to the node containing that memory address.

The remote memory controller sets up access to the remote memory address.

The CPU performs read or write operations on that remote memory address.


Every action needs to pass through multiple memory controllers, so access can take more than twice as long when attempting to access remote memory addresses. The primary performance concern in a multi-core system is therefore to ensure that information travels as efficiently as possible, via the shortest, or fastest, path.



To configure an application for optimal CPU performance, you need to know:

- the topology of the system (how its components are connected),
- the core on which the application executes, and
- the location of the closest memory bank.




### Tools


#### taskset
Setting CPU Affinity with taskset.

taskset retrieves and sets the CPU affinity of a running process (by process ID). It can also be used to launch a process with a given CPU affinity, which binds the specified process to a specified CPU or set of CPUs. However, taskset will not guarantee local memory allocation. If you require the additional performance benefits of local memory allocation, we recommend numactl over taskset; see Section 4.1.2.2, “Controlling NUMA Policy with numactl” for further details.

CPU affinity is represented as a bitmask. The lowest-order bit corresponds to the first logical CPU, and the highest-order bit corresponds to the last logical CPU. These masks are typically given in hexadecimal, so that 0x00000001 represents processor 0, and 0x00000003 represents processors 0 and 1.


To set the CPU affinity of a running process, execute the following command, replacing mask with the mask of the processor or processors you want the process bound to, and pid with the process ID of the process whose affinity you wish to change.

``` shell
 taskset -p mask pid
```


To launch a process with a given affinity, run the following command, replacing mask with the mask of the processor or processors you want the process bound to, and program with the program, options, and arguments of the program you want to run.

``` shell
    taskset mask -- program
```


``` shell
# taskset -c 0,5,7-9 -- myprogram
```
Further information about taskset is available from the man page: man taskset.


#### 4.1.2.2 Controlling NUMA Policy with numactl

numactl runs processes with a specified scheduling or memory placement policy. The selected policy is set for that process and all of its children. numactl can also set a persistent policy for shared memory segments or files, and set the CPU affinity and memory affinity of a process. It uses the /sys file system to determine system topology.


The /sys file system contains information about how CPUs, memory, and peripheral devices are connected via NUMA interconnects. Specifically, the /sys/devices/system/cpu directory contains information about how a system's CPUs are connected to one another. The /sys/devices/system/node directory contains information about the NUMA nodes in the system, and the relative distances between those nodes.


In a NUMA system, the greater the distance between a processor and a memory bank, the slower the processor's access to that memory bank. Performance-sensitive applications should therefore be configured so that they allocate memory from the closest possible memory bank.



Performance-sensitive applications should also be configured to execute on a set number of cores, particularly in the case of multi-threaded applications. Because first-level caches are usually small, if multiple threads execute on one core, each thread will potentially evict cached data accessed by a previous thread. When the operating system attempts to multitask between these threads, and the threads continue to evict each other's cached data, a large percentage of their execution time is spent on cache line replacement. This issue is referred to as cache thrashing. It is therefore recommended to bind a multi-threaded application to a node rather than a single core, since this allows the threads to share cache lines on multiple levels (first-, second-, and last-level cache) and minimizes the need for cache fill operations. However, binding an application to a single core may be performant if all threads are accessing the same cached data.

numactl allows you to bind an application to a particular core or NUMA node, and to allocate the memory associated with a core or set of cores to that application. Some useful options provided by numactl are:

--show
Display the NUMA policy settings of the current process. This parameter does not require further parameters, and can be used like so: numactl --show.



### Dynamic Resource Affinity on Power Architecture.

On Power Architecture Platform Reference systems that support logical partitions (LPARs), processing may be transparently moved to either unused CPU or memory resources. The most common causes of this are either new resources being added, or existing resources being taken out of service. When this occurs, the new memory or CPU may be in a different NUMA domain and this may result in memory affinity which is not optimal because the Linux kernel is unaware of the change.




## terminology

die size

The die size of a specific chip is the physical dimensions of a bare die. In other words, the length and width of the integrated circuit
