
**slurmd** is a multi-threaded daemon running on each compute node and can be compared to a remote shell daemon: it reads the common SLURM configuration file and saved state information, notifies the controller that it is active, waits for work, executes the work, returns status, then waits for more work. Because it initi- ates jobs for other users, **it must run as user root**. It also asynchronously exchanges node and job status with . The only job information it has at any given time pertains to its currently executing jobs. has five major components:


## slurmctld

slurmctld can run in either master or standby mode.

slurmctld need not to execute as user root. In fact , it is recommended that a unique user entry be created for executing slurmctld and that user must be identified in the slurm configuration file as Slurmuser. slurmctld has three major components:


1. Node manager: Monitor the state of each node in the cluster.

2. Partition manager: Groups nodes into non-overlapping sets called partitions. Each partition can have associated with it various job limits and access controls. The Partition Manager also allocates nodes to jobs based on node and partition states and configurations. Requests to initiate jobs come from the Job Manager. may be used to administratively alter node and partition
configurations.

3. Job Manager: Accepts user job requests and places pending jobs in a priority- ordered queue. The Job Manager is awakened on a periodic basis and when- ever there is a change in state that might permit a job to begin running, such as job completion, job submission, partition up transition, node up transition, etc. The Job Manager then makes a pass through the priority-ordered job queue. The highest priority jobs for each partition are allocated resources as possible. As soon as an allocation failure occurs for any partition, no lower-priority jobs for that partition are considered for initiation. After completing the schedul- ing cycle, the Job Manager’s scheduling thread sleeps. Once a job has been allocated resources, the Job Manager transfers necessary state information to those nodes, permitting it to commence execution. When the Job Manager detects that all nodes associated with a job have completed their work, it initiates cleanup and performs another scheduling cycle as described above.
