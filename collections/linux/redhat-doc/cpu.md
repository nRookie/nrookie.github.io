
The term CPU, which stands for central processing unit, is a misnomer for most systems, since central implies single, whereas most modern systems have more than one processing unit, or core. Physically, CPUs are contained in a package attached to a motherboard in a socket. Each socket on the motherboard has various connections: to other CPU sockets, memory controllers, interrupt controllers, and other peripheral devices. A socket to the operating system is a logical grouping of CPUs and associated resources. This concept is central to most of our discussions on CPU tuning.

Red Hat Enterprise Linux keeps a wealth of statistics about system CPU events; these statistics are useful in planning out a tuning strategy to improve CPU performance. Section 4.1.2, “Tuning CPU Performance” discusses some of the more useful statistics, where to find them, and how to analyze them for performance tuning.


## Topology

Older computers had relatively few CPUs per system, which allowed an architecture known as Symmetric Multi-Processor (SMP). This meant that each CPU in the system had similar (or symmetric) access to available memory. In recent years, CPU count-per-socket has grown to the point that trying to give symmetric access to all RAM in the system has grown to the point that trying to give symmetric access to all RAM in the system has become very expansive. Most high CPU count systems these days have an architecture known as Non-Uniform Memory Access (NUMA) instead of SMP.


AMD processors have had this type of architecture for some time with their Hyper Transport (HT) interconnects, while Intel has begun implementing NUMA in their Quick Path Interconnect (QPI) designs. NUMA and SMP are tuned differently, since you need to account for the topology of the system when allocating resources for an application.


## Threads

Inside the Linux operating system, the unit of execution is known as a thread. Threads have a register context, a stack, and a segment of executable code which they run on a CPU. It is the job of the operating system (OS) to schedule these threads on the available CPUs.
The OS maximizes CPU utilization by load-balancing the threads across available cores. Since the OS is primarily concerned with keeping CPUs busy, it does not make optimal decisions with respect to application performance. Moving an application thread to a CPU on another socket can worsen performance more than simply waiting for the current CPU to become available, since memory access operations can slow drastically across sockets. For high-performance applications, it is usually better for the designer to determine where threads are placed. Section 4.2, “CPU Scheduling” discusses how to best allocate CPUs and memory to best execute application threads.




## Interrupts


One of the less obvious (but nonetheless important)  system events that can impact application performance is the interrupt 
(also known as IRQs in Linux). These events are handled by the operating system, and are used by peripherals to signal the arrival of data or the completion of an operation, such as network write or a time event.


The manner in which the OS or CPU that is executing application code handles an interrupt does not affect the applicaiton's functionality. However, it can impact the performance of the applicaiton. This chapter also discusses tips on preventing interrupts from adversely 
impacting application performance.

