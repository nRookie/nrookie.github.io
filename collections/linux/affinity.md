# Processor affinity
Processor affinity or CPU pinning or "cache affinity", enables the binding and unbinding of a process or a thread to a central processing unit(CPU) or a range of CPUs, so that the process or thread will execute only on the designated CPU or CPUs rather than any CPU. This can be viewed as a modification of the native central queue scheduling algorithm in a symmetric multiprocessing operating system. Each item in the queue has a tag indicating its kin processor. At the time of resource allocation, each task is allocated to its kin processor in preference of others.

Processor affinity takes advantage of the fact that remnants of a process that was run on a given processor may remain in that processor's state (for example, data in cache memory) after another process was run on that processor. Scheduling a CPU intensive process that has few interrupts to execute on the same processor may improve its performance by reducing degrading events such as cache misses, but may slow down ordinary programs because they would need to wait for that CPU to become available again. A practical example of processor affinity is executing multiple instances of a non-thread application, such as some graphics-rendering software.

Scheduling-algorithm implementations vary in adherence to processor affinity, under certain circumstances, some implementations will allow a task to change to another processor if it returns in higher efficiency. For example, when two processor-intensive tasks (A and B) have affinity to one processor while another processor remains unused, many schedulers will shift task B to the second processor in order or maximize processor use. Task B will then acquire affinity with the second processor, while task A will continue to have affinity with the original processor.

## Usage


Processor affinity can effectively reduce cache problems, but it does not reduce the persistent load-balancing problem. Also note that processor affinity becomes more complicated in systems with non-uniform architectures. For example, a system with two dual-core hyper-threadedCPus presents a challenge to a scheduling algorithm.

There is a complete two affinity between two virtual CPUs implemented on the same core via hyper-threading, partial affinity between two cores on the same physical processor (as the core share some, but not all, cache), and no affinity between separate physical processors. As other resources are also shared, processor affinity alone cannot be used as the basis of CPU dispatching. If a process has recently run on one virtual hyper-threaded CPU in a given core, and that virtual CPU is currently busy but its partner is not, cache affinity would suggest that the process should be dispatched to the idle partner CPU. however, the two virtual CPUs compete for essentially all computing, cache, and memory resources. In this situation, it would typically be more efficient to dispatch the process to a different core or CPU, if one is available. This could incur a penalty when process repopulates the cache, but overall performance could be higher as the process would not have to compete for resources within the CPU.[citation needed].


## Specific operating system.

On Linux, the CPU affinity of a process can be altered with the taskset(1) program[3] and the sched_setaffinity(2) system call. The affinity of a thread can be altered with one of the library functions: pthread_setaffinity_np(3) or pthread_attr_setaffinity_np(3).



# CPU Affinity

https://www.linuxjournal.com/article/6799


The ability in Linux to bind one or more processes to one or more processors, called CPU affinity, is a long-requested feature. The idea is to say “always run this process on processor one” or “run these processes on all processors but processor zero”. The scheduler then obeys the order, and the process runs only on the allowed processors.



Other operating systems, such as Windows NT, have long provided a system call to set the CPU affinity for a process. Consequently, demand for such a system call in Linux has been high. Finally, the 2.5 kernel introduced a set of system calls for setting and retrieving the CPU affinity of a process.

In this article, I look at the reasons for introducing a CPU affinity interface to Linux. I then cover how to use the interface in your programs. If you are not a programmer or if you have an existing program you are unable to modify, I cover a simple utility for changing the affinity of a given process using its PID. Finally, we look at the actual implementation of the system call.

Soft vs. Hard CPU Affinity

There are two types of CPU affinity. The first, soft affinity, also called natural affinity, is the tendency of a scheduler to try to keep processes on the same CPU as long as possible. It is merely an attempt; if it is ever infeasible, the processes certainly will migrate to another processor. 
The new O(1) scheduler in 2.5 exhibits excellent natural affinity. On the opposite end, however, is the 2.4 scheduler, which has poor CPU affinity. 
This behavior results in the ping-pong effect. The scheduler bounces processes between multiple processors each time they are scheduled and rescheduled. Table 1 is an example of poor natural affinity; Table 2 shows what good natural affinity looks like.




Hard affinity, on the other hand, is what a CPU affinity system call provides. It is a requirement, and processes must adhere to a specified hard affinity. If a processor is bound to CPU zero, for example, then it can run only on CPU zero.

Why One Needs CPU Affinity
Before we cover the new system calls, let's discuss why anyone would need such a feature. The first benefit of CPU affinity is optimizing cache performance. I said the O(1) scheduler tries hard to keep tasks on the same processor, and it does. But in some performance-critical situations—perhaps a large database or a highly threaded Java server—it makes sense to enforce the affinity as a hard requirement. Multiprocessing computers go through a lot of trouble to keep the processor caches valid. Data can be kept in only one processor's cache at a time. Otherwise, the processor's cache may grow out of sync, leading to the question, who has the data that is the most up-to-date copy of the main memory? Consequently, whenever a processor adds a line of data to its local cache, all the other processors in the system also caching it must invalidate that data. This invalidation is costly and unpleasant. But the real problem comes into play when processes bounce between processors: they constantly cause cache invalidations, and the data they want is never in the cache when they need it. Thus, cache miss rates grow very large. CPU affinity protects against this and improves cache performance.


