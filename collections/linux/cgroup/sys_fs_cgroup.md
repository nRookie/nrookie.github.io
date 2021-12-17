
https://www.kernel.org/doc/Documentation/cgroup-v1/cpusets.txt

man cpuset


1. Cpusets
==========

1.1 What are cpusets ?
----------------------

Cpusets provide a mechanism for assigning a set of CPUs and Memory
Nodes to a set of tasks.   In this document "Memory Node" refers to
an on-line node that contains memory.

Cpusets constrain the CPU and Memory placement of tasks to only
the resources within a task's current cpuset.  They form a nested
hierarchy visible in a virtual file system.  These are the essential
hooks, beyond what is already present, required to manage dynamic
job placement on large systems.

Cpusets use the generic cgroup subsystem described in
Documentation/cgroup-v1/cgroups.txt.

Requests by a task, using the sched_setaffinity(2) system call to
include CPUs in its CPU affinity mask, and using the mbind(2) and
set_mempolicy(2) system calls to include Memory Nodes in its memory
policy, are both filtered through that task's cpuset, filtering out any
CPUs or Memory Nodes not in that cpuset.  The scheduler will not
schedule a task on a CPU that is not allowed in its cpus_allowed
vector, and the kernel page allocator will not allocate a page on a
node that is not allowed in the requesting task's mems_allowed vector.

User level code may create and destroy cpusets by name in the cgroup
virtual file system, manage the attributes and permissions of these
cpusets and which CPUs and Memory Nodes are assigned to each cpuset,
specify and query to which cpuset a task is assigned, and list the
task pids assigned to a cpuset.




1.2 Why are cpusets needed ?

The management of large computer systems, with many processors (CPUs),
complex memory cache hierarchies and multiple Memory Nodes having
non-uniform access times (NUMA) presents additional challenges for
the efficient scheduling and memory placement of processes.

Frequently more modest sized systems can be operated with adequate
efficiency just by letting the operating system automatically share
the available CPU and Memory resources amongst the requesting tasks.

But larger systems, which benefit more from careful processor and
memory placement to reduce memory access times and contention,
and which typically represent a larger investment for the customer,
can benefit from explicitly placing jobs on properly sized subsets of
the system.

This can be especially valuable on:

    * Web Servers running multiple instances of the same web application,
    * Servers running different applications (for instance, a web server
      and a database), or
    * NUMA systems running large HPC applications with demanding
      performance characteristics.

These subsets, or "soft partitions" must be able to be dynamically
adjusted, as the job mix changes, without impacting other concurrently
executing jobs. The location of the running jobs pages may also be moved
when the memory locations are changed.

The kernel cpuset patch provides the minimum essential kernel
mechanisms required to efficiently implement such subsets.  It
leverages existing CPU and Memory Placement facilities in the Linux
kernel to avoid any additional impact on the critical scheduler or
memory allocator code.


1.3 How are cpusets implemented ?



Cpusets provide a Linux kernel mechanism to constrain which CPUs and
Memory Nodes are used by a process or set of processes.

The Linux kernel already has a pair of mechanisms to specify on which
CPUs a task may be scheduled (sched_setaffinity) and on which Memory
Nodes it may obtain memory (mbind, set_mempolicy).

Cpusets extends these two mechanisms as follows:

 - Cpusets are sets of allowed CPUs and Memory Nodes, known to the
   kernel.
 - Each task in the system is attached to a cpuset, via a pointer
   in the task structure to a reference counted cgroup structure.
 - Calls to sched_setaffinity are filtered to just those CPUs
   allowed in that task's cpuset.
 - Calls to mbind and set_mempolicy are filtered to just
   those Memory Nodes allowed in that task's cpuset.
 - The root cpuset contains all the systems CPUs and Memory
   Nodes.
 - For any cpuset, one can define child cpusets containing a subset
   of the parents CPU and Memory Node resources.
 - The hierarchy of cpusets can be mounted at /dev/cpuset, for
   browsing and manipulation from user space.
 - A cpuset may be marked exclusive, which ensures that no other
   cpuset (except direct ancestors and descendants) may contain
   any overlapping CPUs or Memory Nodes.
 - You can list all the tasks (by pid) attached to any cpuset.

The implementation of cpusets requires a few, simple hooks
into the rest of the kernel, none in performance critical paths:

 - in init/main.c, to initialize the root cpuset at system boot.
 - in fork and exit, to attach and detach a task from its cpuset.
 - in sched_setaffinity, to mask the requested CPUs by what's
   allowed in that task's cpuset.
 - in sched.c migrate_live_tasks(), to keep migrating tasks within
   the CPUs allowed by their cpuset, if possible.
 - in the mbind and set_mempolicy system calls, to mask the requested
   Memory Nodes by what's allowed in that task's cpuset.
 - in page_alloc.c, to restrict memory to allowed nodes.
 - in vmscan.c, to restrict page recovery to the current cpuset.

You should mount the "cgroup" filesystem type in order to enable
browsing and modifying the cpusets presently known to the kernel.  No
new system calls are added for cpusets - all support for querying and
modifying cpusets is via this cpuset file system.

The /proc/<pid>/status file for each task has four added lines,
displaying the task's cpus_allowed (on which CPUs it may be scheduled)
and mems_allowed (on which Memory Nodes it may obtain memory),
in the two formats seen in the following example:

  Cpus_allowed:   ffffffff,ffffffff,ffffffff,ffffffff
  Cpus_allowed_list:      0-127
  Mems_allowed:   ffffffff,ffffffff
  Mems_allowed_list:      0-63

Each cpuset is represented by a directory in the cgroup file system
containing (on top of the standard cgroup files) the following
files describing that cpuset:

 - cpuset.cpus: list of CPUs in that cpuset
 - cpuset.mems: list of Memory Nodes in that cpuset
 - cpuset.memory_migrate flag: if set, move pages to cpusets nodes
 - cpuset.cpu_exclusive flag: is cpu placement exclusive?
 - cpuset.mem_exclusive flag: is memory placement exclusive?
 - cpuset.mem_hardwall flag:  is memory allocation hardwalled
 - cpuset.memory_pressure: measure of how much paging pressure in cpuset
 - cpuset.memory_spread_page flag: if set, spread page cache evenly on allowed nodes
 - cpuset.memory_spread_slab flag: if set, spread slab cache evenly on allowed nodes
 - cpuset.sched_load_balance flag: if set, load balance within CPUs on that cpuset
 - cpuset.sched_relax_domain_level: the searching range when migrating tasks

In addition, only the root cpuset has the following file:
 - cpuset.memory_pressure_enabled flag: compute memory_pressure?

New cpusets are created using the mkdir system call or shell
command.  The properties of a cpuset, such as its flags, allowed
CPUs and Memory Nodes, and attached tasks, are modified by writing
to the appropriate file in that cpusets directory, as listed above.

The named hierarchical structure of nested cpusets allows partitioning
a large system into nested, dynamically changeable, "soft-partitions".

The attachment of each task, automatically inherited at fork by any
children of that task, to a cpuset allows organizing the work load
on a system into related sets of tasks such that each set is constrained
to using the CPUs and Memory Nodes of a particular cpuset.  A task
may be re-attached to any other cpuset, if allowed by the permissions
on the necessary cpuset file system directories.

Such management of a system "in the large" integrates smoothly with
the detailed placement done on individual tasks and memory regions
using the sched_setaffinity, mbind and set_mempolicy system calls.

The following rules apply to each cpuset:

 - Its CPUs and Memory Nodes must be a subset of its parents.
 - It can't be marked exclusive unless its parent is.
 - If its cpu or memory is exclusive, they may not overlap any sibling.

These rules, and the natural hierarchy of cpusets, enable efficient
enforcement of the exclusive guarantee, without having to scan all
cpusets every time any of them change to ensure nothing overlaps a
exclusive cpuset.  Also, the use of a Linux virtual file system (vfs)
to represent the cpuset hierarchy provides for a familiar permission
and name space for cpusets, with a minimum of additional kernel code.

The cpus and mems files in the root (top_cpuset) cpuset are
read-only.  The cpus file automatically tracks the value of
cpu_online_mask using a CPU hotplug notifier, and the mems file
automatically tracks the value of node_states[N_MEMORY]--i.e.,
nodes with memory--using the cpuset_track_online_nodes() hook.

1.4 What are exclusive cpusets ?
--------------------------------

If a cpuset is cpu or mem exclusive, no other cpuset, other than
a direct ancestor or descendant, may share any of the same CPUs or
Memory Nodes.

A cpuset that is cpuset.mem_exclusive *or* cpuset.mem_hardwall is "hardwalled",
i.e. it restricts kernel allocations for page, buffer and other data
commonly shared by the kernel across multiple users.  All cpusets,
whether hardwalled or not, restrict allocations of memory for user
space.  This enables configuring a system so that several independent
jobs can share common kernel data, such as file system pages, while
isolating each job's user allocation in its own cpuset.  To do this,
construct a large mem_exclusive cpuset to hold all the jobs, and
construct child, non-mem_exclusive cpusets for each individual job.
Only a small amount of typical kernel memory, such as requests from
interrupt handlers, is allowed to be taken outside even a
mem_exclusive cpuset.




Cpusets use the generic cgroup subsystem described in
Documentation/cgroup-v1/cgroups.txt.


Requests by a task, using the sched_setaffinity(2) system call to
include CPUs in its CPU affinity mask, and using the mbind(2) and
set_mempolicy(2) system calls to include Memory Nodes in its memory
policy, are both filtered through that task's cpuset, filtering out any
CPUs or Memory Nodes not in that cpuset.  The scheduler will not
schedule a task on a CPU that is not allowed in its cpus_allowed
vector, and the kernel page allocator will not allocate a page on a
node that is not allowed in the requesting task's mems_allowed vector.




User level code may create and destroy cpusets by name in the cgroup
virtual file system, manage the attributes and permissions of these
cpusets and which CPUs and Memory Nodes are assigned to each cpuset,
specify and query to which cpuset a task is assigned, and list the
task pids assigned to a cpuset.


cgroup.clone_children  cpuset.memory_pressure
cgroup.event_control   cpuset.memory_spread_page
cgroup.procs           cpuset.memory_spread_slab
cpuset.cpu_exclusive   cpuset.mems
cpuset.cpus            cpuset.sched_load_balance
cpuset.mem_exclusive   cpuset.sched_relax_domain_level
cpuset.mem_hardwall    notify_on_release
cpuset.memory_migrate  tasks

Reading them will give you information about the state of this cpuset:
the CPUs and Memory Nodes it can use, the processes that are using
it, its properties.  By writing to these files you can manipulate
the cpuset.

Set some flags:
# /bin/echo 1 > cpuset.cpu_exclusive

Add some cpus:
# /bin/echo 0-7 > cpuset.cpus

Add some mems:
# /bin/echo 0-7 > cpuset.mems

Now attach your shell to this cpuset:
# /bin/echo $$ > tasks

You can also create cpusets inside your cpuset by using mkdir in this
directory.
# mkdir my_sub_cs

To remove a cpuset, just use rmdir:
# rmdir my_sub_cs
This will fail if the cpuset is in use (has cpusets inside, or has
processes attached).

Note that for legacy reasons, the "cpuset" filesystem exists as a
wrapper around the cgroup filesystem.


2.2 Adding/removing cpus
------------------------

This is the syntax to use when writing in the cpus or mems files
in cpuset directories:

# /bin/echo 1-4 > cpuset.cpus		-> set cpus list to cpus 1,2,3,4
# /bin/echo 1,2,3,4 > cpuset.cpus	-> set cpus list to cpus 1,2,3,4


2.4 Attaching processes
-----------------------

# /bin/echo PID > tasks

Note that it is PID, not PIDs. You can only attach ONE task at a time.
If you have several tasks to attach, you have to do it one after another:

# /bin/echo PID1 > tasks
# /bin/echo PID2 > tasks
	...
# /bin/echo PIDn > tasks



3. Questions
============

Q: what's up with this '/bin/echo' ?
A: bash's builtin 'echo' command does not check calls to write() against
   errors. If you use it in the cpuset file system, you won't be
   able to tell whether a command succeeded or failed.

Q: When I attach processes, only the first of the line gets really attached !
A: We can only return one error code per call to write(). So you should also
   put only ONE pid.

4. Contact
==========

Web: http://www.bullopensource.org/cpuset