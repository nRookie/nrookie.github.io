A resource controller, also called a cgroup subsystem, represents a single resource, such as CPU time or memory. The Linux kernel provides a range of resource controllers, that are mounted automatically by systemd. Find the list of currently mounted resource controllers in /proc/cgroups, or use the lssubsys monitoring tool. In Red Hat Enterprise Linux 7, systemd mounts the following controllers by default:



Available Controllers in Red Hat Enterprise Linux 7

blkio — sets limits on input/output access to and from block devices;

cpu — uses the CPU scheduler to provide cgroup tasks access to the CPU. It is mounted together with the cpuacct controller on the same mount;

cpuacct — creates automatic reports on CPU resources used by tasks in a cgroup. It is mounted together with the cpu controller on the same mount;


cpuset — assigns individual CPUs (on a multicore system) and memory nodes to tasks in a cgroup;

devices — allows or denies access to devices for tasks in a cgroup;

freezer — suspends or resumes tasks in a cgroup;

memory — sets limits on memory use by tasks in a cgroup and generates automatic reports on memory resources used by those tasks;


net_cls — tags network packets with a class identifier (classid) that allows the Linux traffic controller (the tc command) to identify packets originating from a particular cgroup task. A subsystem of net_cls, the net_filter (iptables) can also use this tag to perform actions on such packets. The net_filter tags network sockets with a firewall identifier (fwid) that allows the Linux firewall (the iptables command) to identify packets (skb->sk) originating from a particular cgroup task;


perf_event — enables monitoring cgroups with the perf tool;

hugetlb — allows to use virtual memory pages of large sizes and to enforce resource limits on these pages.


The Linux kernel exposes a wide range of tunable parameters for resource controllers that can be configured with systemd. See the kernel documentation (list of references in the Controller-Specific Kernel Documentation section) for detailed description of these parameters.


https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/7/html/resource_management_guide/br-resource_controllers_in_linux_kernel#itemlist-Available_Controllers_in_Red_Hat_Enterprise_Linux_7