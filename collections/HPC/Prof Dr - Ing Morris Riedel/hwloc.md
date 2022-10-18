[toc]



https://www.open-mpi.org/projects/hwloc/



The Portable Hardware Locality (hwloc) software package provides a **portable abstraction** (across OS, versions, architectures, ...) of the **hierarchical topology of modern architectures**, including NUMA memory nodes, sockets, shared caches, cores and simultaneous multithreading. It also gathers various system attributes such as cache and memory information as well as the locality of I/O devices such as network interfaces, InfiniBand HCAs or GPUs.





[![Sample hwloc output](https://www.open-mpi.org/projects/hwloc/lstopo.png)





hwloc primarily aims at helping applications with **gathering information about increasingly complex parallel computing platforms so as to exploit them accordingly and efficiently**. For instance, two tasks that tightly cooperate should probably be placed onto cores sharing a cache. However, two independent memory-intensive tasks should better be spread out onto different sockets so as to maximize their memory throughput. As described in [this paper](http://hal.inria.fr/inria-00496295), OpenMP threads have to be placed according to their affinities and to the hardware characteristics. MPI implementations apply similar techniques while also adapting their communication strategies to the network locality as described in [this paper](http://hal.inria.fr/inria-00486178) or [this one](http://hal.inria.fr/inria-00566246).



# Portability and support



hwloc supports the following operating systems:



- Linux (including old kernels not having sysfs topology information, with knowledge of cgroups, offline CPUs, ScaleMP vSMP, and NumaScale NumaConnect) on all supported hardware, including Intel Xeon Phi.
- Solaris, AIX and HP-UX
- NetBSD, FreeBSD and kFreeBSD/GNU
- Darwin / OS X
- Microsoft Windows (either using MinGW, or Cygwin, or a native Visual Studio solution)
- IBM BlueGene/Q Compute Node Kernel (CNK)









Additionally hwloc can detect the locality PCI devices as well as OpenCL, CUDA and Xeon Phi accelerators, network and InfiniBand interfaces, etc. See the [Best of lstopo](https://www.open-mpi.org/projects/hwloc/lstopo) for more examples of supported platforms. The topologies of many existing platforms are also available in the [XML topology database](https://hwloc.gitlabpages.inria.fr/xmls/) for testing your software on architectures you don't have access to.

hwloc may display the topology in multiple convenient formats (see [v2.6.0 examples](https://www.open-mpi.org/projects/hwloc/doc/v2.6.0/a00358.php#cli_examples) and the [Best of lstopo](https://www.open-mpi.org/projects/hwloc/lstopo)). It also offers a powerful programming interface to gather information about the hardware, bind processes, and much more.

Since it uses standard Operating System information, hwloc's support is almost always independent from the processor type (x86, powerpc, ia64, ...), and just relies on the Operating System support. Whenever the OS does not support topology information (e.g. some BSDs), hwloc uses an x86-only CPUID-based backend.

To check whether hwloc works on a particular machine, just try to build it and run `lstopo` or `lstopo-no-graphics`. If some things do not look right (e.g. bogus or missing cache information), see Questions and bugs below











# what-is-meant-by-the-terms-cpu-core-die-and-package



https://superuser.com/questions/324284/what-is-meant-by-the-terms-cpu-core-die-and-package



"Core 2 Duo" is Intel's trademark name for some of its processor.



**(Physical) processor core** is an independent execution unit that can run one program thread at a time in parallel with other cores.





**Processor die** is a single continuous piece of semiconductor material (usually silicon). A die can contain any number of cores. Up to 15 are available on the Intel product line. Processor die is where the transistors making up the CPU actually reside.





**Processor package** is what you get when you buy a single processor. It contains one or more dies, plastic/ceramic housing for dies and gold-plated contacts that match those on your motherboard.





Do note that you always have at least one core, one die and one package. For processor to make sense, it has to have an unit that can execute commands, a piece of silicon physically containing the transistors implementing the processor, and the package that attaches that silicon to contacts that mate to motherboard and IO.



**Dual-core processor** is a processor *package* that has two physical cores inside. It can be either on one die or two dies. Often the first generation multi-core processors used several dies on single package, while modern designs put them to same die, which gives advantages like being able to share on-die cache.





The term **"CPU"** can be ambiguous. When people buy "a CPU", they buy a CPU package. When they inspect "CPU scaling", they talk about logical cores. The reason for this is that for most practical purposes dual-core processor behaves like two processor system, ie. system that has two CPU sockets and two CPU single core packages installed to them, so when talking about scaling, it makes most sense to count the cores available; how they are installed to dies, packages and motherboard is less important.



The term **"package"** also has several meanings: Here CPU "package" means the piece of plastic, ceramic and metal that contain the CPU. Each CPU socket on motherboard can accept exactly one package; package is the unit that's plugged to the socket.

You can see example of two-die quad-core processor [here](http://www.anandtech.com/show/2112).







https://unix.stackexchange.com/questions/113544/interpret-the-output-of-lstopo



PU P#" = Processing Unit Processor #. These are processing elements within the cores of the CPU. On my laptop (Intel i5) I have 2 cores that each have 2 processing elements, for a total of 4. But in actuality I have only 2 physical cores.



1. L#i = Instruction Cache, L#d = Data Cache. L1 = a Level 1 cache.



![ss of my laptop](https://i.stack.imgur.com/VSRgt.png)



1. In the Intel architectures the instruction & data get mixed as you move down from L1 → L2 → L3.

1. "Socket P#" is that there are 2 physical sockets on the motherboard, there are 2 physically discrete CPUs in this setup.
2. In multiple CPU architectures the RAM is usually split so that a portion of it is assigned to each core. If CPU0 needs data from CPU1's RAM, then it needs to "request" this data through CPU1. There are a number of reasons why this is done, too many to elaborate here. Read up on [NUMA style memory architectures](https://en.wikipedia.org/wiki/Non-uniform_memory_access) if you're really curious.
3. ![ss of numa](https://i.stack.imgur.com/1Acrw.png)
4. The drawing is showing 4 cores (with 1 Processing Unit in each) that are in 2 physical CPU packages. Each physical CPU has "isolated" access to 16 GB of RAM
5. No, there is no shared memory among all the CPUs. The 2 CPUs have to interact with the other's RAM through the CPU. Again see the [NUMA Wikipage for more on the Non Uniform Memory Architecture](https://en.wikipedia.org/wiki/Non-uniform_memory_access).
6. Yes, the system has a total of 32 GB of RAM. But only 1/2 of the RAM is accessible by either physical CPU directly.









# Play with lstopo

``` shell
**-i** <specification>, **--input** <specification>
Simulate a fake hierarchy (instead of discovering the topology on the local machine). If <specification> is "node:2 pu:3", the topology will contain two NUMA nodes with 3 processing units in each of them. The <specification> string must end with a number of PUs.
```



``` shell
lstopo --output-format pdf --no-io > topo.pdf
```





![image-20211202134909574](/Users/kestrel/developer/nrookie.github.io/collections/HPC/Prof Dr - Ing Morris Riedel/image-20211202134909574.png)





``` shell
 lstopo --output-format pdf --no-io > topo.pdf
```



![image-20211202135132242](/Users/kestrel/developer/nrookie.github.io/collections/HPC/Prof Dr - Ing Morris Riedel/image-20211202135132242.png)





``` shell
lstopo --input "node:1 socket:1 cache:1 cache:2 cache:1 core:1 pu:2" --output-format pdf > topo.pdf
```

![image-20211202135430160](/Users/kestrel/developer/nrookie.github.io/collections/HPC/Prof Dr - Ing Morris Riedel/image-20211202135430160.png)



``` shell
lstopo --input "node:1 socket:2 cache:1 cache:2 cache:1 core:1 pu:2" --output-format pdf > topo.pdf
```

![image-20211202141848104](/Users/kestrel/developer/nrookie.github.io/collections/HPC/Prof Dr - Ing Morris Riedel/image-20211202141848104.png)





``` shell
lstopo --input "node:2 socket:2 cache:1 cache:2 cache:1 core:1 pu:2" --output-format pdf > topo.pdf
```

![image-20211202141933626](/Users/kestrel/developer/nrookie.github.io/collections/HPC/Prof Dr - Ing Morris Riedel/image-20211202141933626.png)





``` shell
lstopo --input "node:1 socket:2 cache:1 cache:2 cache:1 core:1 pu:4" --output-format pdf > topo.pdf
```





![image-20211202142156115](/Users/kestrel/developer/nrookie.github.io/collections/HPC/Prof Dr - Ing Morris Riedel/image-20211202142156115.png)



``` shell
lstopo --input "node:1 socket:2 cache:1 cache:2 cache:1 core:2 pu:2" --output-format pdf > topo.pdf
```





![image-20211202142513762](/Users/kestrel/developer/nrookie.github.io/collections/HPC/Prof Dr - Ing Morris Riedel/image-20211202142513762.png)



``` shell
lstopo --input "node:1 socket:2 cache:1 cache:2 cache:1 core:2 pu:1" --output-format pdf > topo.pdf
```



![image-20211202142642980](/Users/kestrel/developer/nrookie.github.io/collections/HPC/Prof Dr - Ing Morris Riedel/image-20211202142642980.png)





``` shell
lstopo --input "node:1 socket:1 cache:1 cache:2 cache:1 core:2 pu:1" --output-format pdf > topo.pdf
```

![image-20211202142736961](/Users/kestrel/developer/nrookie.github.io/collections/HPC/Prof Dr - Ing Morris Riedel/image-20211202142736961.png)





``` shell
 lstopo --input "node:1 socket:1 cache:1 cache:2 cache:1 core:1 pu:1" --output-format pdf > topo.pdf
```



![image-20211202142812951](/Users/kestrel/developer/nrookie.github.io/collections/HPC/Prof Dr - Ing Morris Riedel/image-20211202142812951.png)





``` shell
lstopo --input "node:1 socket:1 cache:1 cache:1 cache:1 core:1 pu:1" --output-format pdf > topo.pdf
```



![image-20211202143102061](/Users/kestrel/developer/nrookie.github.io/collections/HPC/Prof Dr - Ing Morris Riedel/image-20211202143102061.png)







``` shell
 lstopo --input "node:1 socket:1  cache:1 core:1 pu:1" --output-format pdf > topo.pdf
```

![image-20211202143211649](/Users/kestrel/developer/nrookie.github.io/collections/HPC/Prof Dr - Ing Morris Riedel/image-20211202143211649.png)



​

``` shell
lstopo --input "node:1 socket:1  cache:2 core:1 pu:1" --output-format pdf > topo.pdf
```

![image-20211202143325571](/Users/kestrel/developer/nrookie.github.io/collections/HPC/Prof Dr - Ing Morris Riedel/image-20211202143325571.png)



``` shell
lstopo --input "node:1 socket:1 cache:1 cache:1 core:1 pu:1" --output-format pdf > topo.pdf
```



![image-20211202143405974](/Users/kestrel/developer/nrookie.github.io/collections/HPC/Prof Dr - Ing Morris Riedel/image-20211202143405974.png)





``` shell
lstopo --input "node:1 socket:1 cache:1 cache:2 cache:2 core:1 pu:1" --output-format pdf > topo.pdf
```



![image-20211202143604257](/Users/kestrel/developer/nrookie.github.io/collections/HPC/Prof Dr - Ing Morris Riedel/image-20211202143604257.png)



``` shell
lstopo --input "node:1 socket:1 cache:1 cache:1 cache:2 core:1 pu:1" --output-format pdf > topo.pdf
```





![image-20211202143639819](/Users/kestrel/developer/nrookie.github.io/collections/HPC/Prof Dr - Ing Morris Riedel/image-20211202143639819.png)



# hwloc



Refereing hwloc Index

## hwloc object



Objects in tuples can be any of the following strings (listed from "biggest" to "smallest"):

###

### machine

A set of processors and memory.



### numanode

A NUMA node; a set of processors around memory which the processors can directly access. If **hbm** is

 used instead of **numanode** in locations, command-line tools only consider high-bandwidth memory nodes

 such as KNL's MCDRAM.



### package

Typically a physical package or chip, that goes into a package, it is a grouping of one or more processors.





### cache

A cache memory. If several kinds of caches exist in the system, a specific one may be identified by its level (e.g. **l1cache**) and optionally by its type (e.g. **l1icache**).



### **core**



A single, physical processing unit which may still contain multiple logical processors, such as

 hardware threads.

### pu

Short for processor unit (not process!). The smallest physical execution unit that hwloc recognizes. For example, there may be multiple PUs on a core (e.g., hardware threads).







###  **hwloc** **Indexes**



Indexes are integer values that uniquely specify a given object of a specific type. Indexes can be expressed

either as logical values or physical values. Most hwloc utilities accept logical indexes by default. Passing **--physical** switches to physical/OS indexes. Both logical and physical indexes are described on this man page.



Logical indexes are relative to the object order in the output from the lstopo command. They always start

with 0 and increment by 1 for each successive object.





Physical indexes are how the operating system refers to objects. Note that while physical indexes are non-negative integer values, the hardware and/or operating system may choose arbitrary values -- they may not start with 0, and successive objects may not have consecutive values.



For example, if the first few lines of lstopo -p output are the following:





```
         Machine (47GB)
           NUMANode P#0 (24GB) + Package P#0 + L3 (12MB)
             L2 (256KB) + L1 (32KB) + Core P#0 + PU P#0
             L2 (256KB) + L1 (32KB) + Core P#1 + PU P#0
             L2 (256KB) + L1 (32KB) + Core P#2 + PU P#0
             L2 (256KB) + L1 (32KB) + Core P#8 + PU P#0
             L2 (256KB) + L1 (32KB) + Core P#9 + PU P#0
             L2 (256KB) + L1 (32KB) + Core P#10 + PU P#0
           NUMANode P#1 (24GB) + Package P#1 + L3 (12MB)
             L2 (256KB) + L1 (32KB) + Core P#0 + PU P#0
             L2 (256KB) + L1 (32KB) + Core P#1 + PU P#0
             L2 (256KB) + L1 (32KB) + Core P#2 + PU P#0
             L2 (256KB) + L1 (32KB) + Core P#8 + PU P#0
             L2 (256KB) + L1 (32KB) + Core P#9 + PU P#0
             L2 (256KB) + L1 (32KB) + Core P#10 + PU P#0
```



In this example, the first core on the second package is logically number 6 (i.e., logically the 7th core, starting from 0). Its physical index is 0, but note that another core also has a physical index of 0. Hence, physical indexes may only be relevant within the scope of their parent (or set of ancestors). In this example, to uniquely identify logical core 6 with physical indexes, you must specify (at a minimum) both a package and a core: package 1, core 0.



​    Index values, regardless of whether they are logical or physical, can be expressed in several different forms

​    (where X, Y, and N are positive integers):





![image-20211202143806566](/Users/kestrel/developer/nrookie.github.io/collections/HPC/Prof Dr - Ing Morris Riedel/image-20211202143806566.png)
