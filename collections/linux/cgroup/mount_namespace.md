For an overview of namespaces, see namespaces


Mount namespaces provide isolation of the list of mounts seen by
the processes in each namespace instance.  Thus, the processes in
each of the mount namespace instances will see distinct single-
directory hierarchies.

The views provided by the /proc/[pid]/mounts,
/proc/[pid]/mountinfo, and /proc/[pid]/mountstats files (all
described in proc(5)) correspond to the mount namespace in which
the process with the PID [pid] resides.  (All of the processes
that reside in the same mount namespace will see the same view in
these files.)


A new mount namespace is created using either clone(2) or
unshare(2) with the CLONE_NEWNS flag.  When a new mount namespace
is created, its mount list is initialized as follows:

*  If the namespace is created using clone(2), the mount list of
    the child's namespace is a copy of the mount list in the
    parent process's mount namespace.

*  If the namespace is created using unshare(2), the mount list
    of the new namespace is a copy of the mount list in the
    caller's previous mount namespace.

Subsequent modifications to the mount list (mount(2) and
umount(2)) in either mount namespace will not (by default) affect
the mount list seen in the other namespace (but see the following
discussion of shared subtrees).




## Namespaces
A  namespace  wraps  a global system resource in an abstraction that makes it appear to the processes within the namespace that they have their own isolated instance of the global resource.  Changes to the global resource are visible to other processes that are members of the namespace, but are invisible to other processes.  One use of namespaces is to implement containers.



Linux provides the following namespaces:

Namespace   Constant          Isolates
Cgroup      CLONE_NEWCGROUP   Cgroup root directory
IPC         CLONE_NEWIPC      System V IPC, POSIX message queues
Network     CLONE_NEWNET      Network devices, stacks, ports, etc.
Mount       CLONE_NEWNS       Mount points
PID         CLONE_NEWPID      Process IDs
User        CLONE_NEWUSER     User and group IDs
UTS         CLONE_NEWUTS      Hostname and NIS domain name

 This page describes the various namespaces and the associated /proc files, and summarizes the APIs for working with namespaces.


## cgroup_namespaces

Cgroup namespaces virtualize the view of a process's cgroup (see cgroups(7)) as seen via /proc/[pid]/cgroup and /proc/[pid]/mountinfo.


Each cgroup namespace has its own set of cgruop root directories. These root directorires are the base points for the relative locations displayed in the corresponding records in the /proc/[pid]/cgroup file. When a process creates a new cgroup namespace using clone(2) or ushare(2) with the CLONE_NEWCGROUP flag, it enters a new cgroup namespace in which its current cgroups directories become the cgroup root directories of the new namespaces. 

When viewing /proc/[pid]/cgroup, the pathname shown in the third field of each record will be relative to the reading process's root directory  for  the
corresponding  cgroup  hierarchy.   If the cgroup directory of the target process lies outside the root directory of the reading process's cgroup names‐
pace, then the pathname will show ../ entries for each ancestor level in the cgroup hierarchy.