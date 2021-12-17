https://www.kernel.org/doc/Documentation/cgroup-v1/memory.txt


Memory Resource Controller

NOTE: This document is hopelessly outdated and it asks for a complete
      rewrite. It still contains a useful information so we are keeping it
      here but make sure to check the current code if you need a deeper
      understanding.

NOTE: The Memory Resource Controller has generically been referred to as the
      memory controller in this document. Do not confuse memory controller
      used here with the memory controller that is used in hardware.



(For editors)
In this document:
      When we mention a cgroup (cgroupfs's directory) with memory controller,
      we call it "memory cgroup". When you see git-log and source code, you'll
      see patch's title and function names tend to use "memcg".
      In this document, we avoid using it.


Benefits and Purpose of the memory controller

The memory controller isolates the memory behaviour of a group of tasks
from the rest of the system. The article on LWN [12] mentions some probable
uses of the memory controller. The memory controller can be used to

a. Isolate an application or a group of applications
   Memory-hungry applications can be isolated and limited to a smaller
   amount of memory.
b. Create a cgroup with a limited amount of memory; this can be used
   as a good alternative to booting with mem=XXXX.
c. Virtualization solutions can control the amount of memory they want
   to assign to a virtual machine instance.
d. A CD/DVD burner could control the amount of memory used by the
   rest of the system to ensure that burning does not fail due to lack
   of available memory.
e. There are several other use cases; find one or use the controller just
   for fun (to learn and hack on the VM subsystem).

Current Status: linux-2.6.34-mmotm(development version of 2010/April)



Brief summary of control files.

 tasks				 # attach a task(thread) and show list of threads
 cgroup.procs			 # show list of processes
 cgroup.event_control		 # an interface for event_fd()
 memory.usage_in_bytes		 # show current usage for memory
				 (See 5.5 for details)
 memory.memsw.usage_in_bytes	 # show current usage for memory+Swap
				 (See 5.5 for details)
 memory.limit_in_bytes		 # set/show limit of memory usage
 memory.memsw.limit_in_bytes	 # set/show limit of memory+Swap usage
 memory.failcnt			 # show the number of memory usage hits limits
 memory.memsw.failcnt		 # show the number of memory+Swap hits limits
 memory.max_usage_in_bytes	 # show max memory usage recorded
 memory.memsw.max_usage_in_bytes # show max memory+Swap usage recorded
 memory.soft_limit_in_bytes	 # set/show soft limit of memory usage
 memory.stat			 # show various statistics
 memory.use_hierarchy		 # set/show hierarchical account enabled
 memory.force_empty		 # trigger forced page reclaim
 memory.pressure_level		 # set memory pressure notifications
 memory.swappiness		 # set/show swappiness parameter of vmscan
				 (See sysctl's vm.swappiness)
 memory.move_charge_at_immigrate # set/show controls of moving charges
 memory.oom_control		 # set/show oom controls.
 memory.numa_stat		 # show the number of memory usage per numa node

 memory.kmem.limit_in_bytes      # set/show hard limit for kernel memory
 memory.kmem.usage_in_bytes      # show current kernel memory allocation
 memory.kmem.failcnt             # show the number of kernel memory usage hits limits
 memory.kmem.max_usage_in_bytes  # show max kernel memory usage recorded

 memory.kmem.tcp.limit_in_bytes  # set/show hard limit for tcp buf memory
 memory.kmem.tcp.usage_in_bytes  # show current tcp buf memory allocation
 memory.kmem.tcp.failcnt            # show the number of tcp buf memory usage hits limits
 memory.kmem.tcp.max_usage_in_bytes # show max tcp buf memory usage recorded


 2. Memory Control

Memory is a unique resource in the sense that it is present in a limited
amount. If a task requires a lot of CPU processing, the task can spread
its processing over a period of hours, days, months or years, but with
memory, the same physical memory needs to be reused to accomplish the task.

The memory controller implementation has been divided into phases. These
are:

1. Memory controller
2. mlock(2) controller
3. Kernel user memory accounting and slab control
4. user mappings length controller

The memory controller is the first controller developed.


2.1. Design

The core of the design is a counter called the page_counter. The
page_counter tracks the current memory usage and limit of the group of
processes associated with the controller. Each cgroup has a memory controller
specific data structure (mem_cgroup) associated with it.



2.2. Accounting

		+--------------------+
		|  mem_cgroup        |
		|  (page_counter)    |
		+--------------------+
		 /            ^      \
		/             |       \
           +---------------+  |        +---------------+
           | mm_struct     |  |....    | mm_struct     |
           |               |  |        |               |
           +---------------+  |        +---------------+
                              |
                              + --------------+
                                              |
           +---------------+           +------+--------+
           | page          +---------->  page_cgroup|
           |               |           |               |
           +---------------+           +---------------+

             (Figure 1: Hierarchy of Accounting)


Figure 1 shows the important aspects of the controller

1. Accounting happens per cgroup
2. Each mm_struct knows about which cgroup it belongs to
3. Each page has a pointer to the page_cgroup, which in turn knows the
   cgroup it belongs to

The accounting is done as follows: mem_cgroup_charge_common() is invoked to
set up the necessary data structures and check if the cgroup that is being
charged is over its limit. If it is, then reclaim is invoked on the cgroup.
More details can be found in the reclaim section of this document.
If everything goes well, a page meta-data-structure called page_cgroup is
updated. page_cgroup has its own LRU on cgroup.
(*) page_cgroup structure is allocated at boot/memory-hotplug time.


2.2.1 Accounting details

All mapped anon pages (RSS) and cache pages (Page Cache) are accounted.
Some pages which are never reclaimable and will not be on the LRU
are not accounted. We just account pages under usual VM management.

RSS pages are accounted at page_fault unless they've already been accounted
for earlier. A file page will be accounted for as Page Cache when it's
inserted into inode (radix-tree). While it's mapped into the page tables of
processes, duplicate accounting is carefully avoided.

An RSS page is unaccounted when it's fully unmapped. A PageCache page is
unaccounted when it's removed from radix-tree. Even if RSS pages are fully
unmapped (by kswapd), they may exist as SwapCache in the system until they
are really freed. Such SwapCaches are also accounted.
A swapped-in page is not accounted until it's mapped.

Note: The kernel does swapin-readahead and reads multiple swaps at once.
This means swapped-in pages may contain pages for other tasks than a task
causing page fault. So, we avoid accounting at swap-in I/O.

At page migration, accounting information is kept.

Note: we just account pages-on-LRU because our purpose is to control amount
of used pages; not-on-LRU pages tend to be out-of-control from VM view.

2.3 Shared Page Accounting

Shared pages are accounted on the basis of the first touch approach. The
cgroup that first touches a page is accounted for the page. The principle
behind this approach is that a cgroup that aggressively uses a shared
page will eventually get charged for it (once it is uncharged from
the cgroup that brought it in -- this will happen on memory pressure).

But see section 8.2: when moving a task to another cgroup, its pages may
be recharged to the new cgroup, if move_charge_at_immigrate has been chosen.

Exception: If CONFIG_MEMCG_SWAP is not used.
When you do swapoff and make swapped-out pages of shmem(tmpfs) to
be backed into memory in force, charges for pages are accounted against the
caller of swapoff rather than the users of shmem.'




https://www.kernel.org/doc/Documentation/cgroup-v1/memory.txt