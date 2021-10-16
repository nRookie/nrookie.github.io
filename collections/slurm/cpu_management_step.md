# CPU Management Steps performed by Slurm 

https://slurm.schedmd.com/cpu_management.html

Slurm uses four basic steps to manage CPU resources for a job/step:

Step 1: Selection of Nodes
Step 2: Allocation of CPUs from the selected Nodes
Step 3: Distribution of Tasks to the selected Nodes
Step 4: Optional Distribution and Binding of Tasks to CPUs within a Node


## STEP 1: SELECTION OF NODES 

In Step 1, Slurm selects the set of nodes from which CPU resources are to be allocated to a job or job step. Node selection is therefore influenced by many of the configuration and command line options that control the allocation of CPUs (Step 2 below). If SelectType=select/linear is configured, all resources on the selected nodes will be allocated to the job/step. If SelectType is configured to be select/cons_res or select/cons_tres, individual sockets, cores and threads may be allocated from the selected nodes as consumable resources. The consumable resource type is defined by SelectTypeParameters.


Step 1 is performed by slurmctld and the select plugin.

srun/salloc/sbatch command line options that control Step 1


## STEP 2: ALLOCATION OF CPUS FROM THE SELECTED NODES

In Step 2, Slurm allocates CPU resources to a job/step from the set of nodes selected in Step 1. CPU allocation is therefore influenced by the configuration and command line options that relate to node selection. If SelectType=select/linear is configured, all resources on the selected nodes will be allocated to the job/step. If SelectType is configured to be select/cons_res or select/cons_tres, individual sockets, cores and threads may be allocated from the selected nodes as consumable resources. The consumable resource type is defined by SelectTypeParameters.


When using a SelectType of select/cons_res or select/cons_tres, the default allocation method across nodes is block allocation (allocate all available CPUs in a node before using another node). The default allocation method within a node is cyclic allocation (allocate available CPUs in a round-robin fashion across the sockets within a node). Users may override the default behavior using the appropriate command line options described below. The choice of allocation methods may influence which specific CPUs are allocated to the job/step.


Step 2 is performed by slurmctld and the select plugin.


## STEP 3: DISTRIBUTION OF TASKS TO THE SELECTED NODES 

In Step 3, Slurm distributes tasks to the nodes that were selected for the job/step in Step 1. Each task is distributed to only one node, but more than one task may be distributed to each node. Unless overcommitment of CPUs to tasks is specified for the job, the number of tasks distributed to a node is constrained by the number of CPUs allocated on the node and the number of CPUs per task. If consumable resources is configured, or resource sharing is allowed, tasks from more than one job/step may run on the same node concurrently.


Step 3 is performed by slurmctld.


## STEP 4: OPTIONAL DISTRIBUTION AND BINDING OF TASKS TO CPUS WITHIN A NODE 


In optional Step 4, Slurm distributes and binds each task to a specified subset of the allocated CPUs on the node to which the task was distributed in Step 3. Different tasks distributed to the same node may be bound to the same subset of CPUs or to different subsets. This step is known as task affinity or task/CPU binding.


## Additional Notes on CPU Management Steps


For consumable resources, it is important for users to understand the difference between cpu allocation (Step 2) and task affinity/binding (Step 4). Exclusive (unshared) allocation of CPUs as consumable resources limits the number of jobs/steps/tasks that can use a node concurrently. But it does not limit the set of CPUs on the node that each task distributed to the node can use. Unless some form of CPU/task binding is used (e.g., a task or spank plugin), all tasks distributed to a node can use all of the CPUs on the node, including CPUs not allocated to their job/step. This may have unexpected adverse effects on performance, since it allows one job to use CPUs allocated exclusively to another job. For this reason, it may not be advisable to configure consumable resources without also configuring task affinity. Note that task affinity can also be useful when select/linear (whole node allocation) is configured, to improve performance by restricting each task to a particular socket or other subset of CPU resources on a node.




### A NOTE ON CPU NUMBERING 
The number and layout of logical CPUs known to Slurm is described in the node definitions in slurm.conf. This may differ from the physical CPU layout on the actual hardware. For this reason, Slurm generates its own internal, or "abstract", CPU numbers. These numbers may not match the physical, or "machine", CPU numbers known to Linux.


### CPU Management and Slurm Accounting 
CPU management by Slurm users is subject to limits imposed by Slurm Accounting. Accounting limits may be applied on CPU usage at the level of users, groups and clusters. For details, see the sacctmgr man page.
    

## Appendix


A job consists in one or more steps, each consisting in one or more tasks each using one or more CPU.


Jobs are typically created with the sbatch command, steps are created with the srun command, tasks are requested, at the job level with --ntasks or --ntasks-per-node, or at the step level with --ntasks. CPUs are requested per task with --cpus-per-task. Note that jobs submitted with sbatch have one implicit step; the Bash script itself.


``` shell
#SBATCH --nodes 8
#SBATCH --tasks-per-node 8
# The job requests 64 CPUs, on 8 nodes.    

# First step, with a sub-allocation of 8 tasks (one per node) to create a tmp dir. 
# No need for more than one task per node, but it has to run on every node
srun --nodes 8 --ntasks 8 mkdir -p /tmp/$USER/$SLURM_JOBID

# Second step with the full allocation (64 tasks) to run an MPI 
# program on some data to produce some output.
srun process.mpi <input.dat >output.txt

# Third step with a sub allocation of 48 tasks (because for instance 
# that program does not scale as well) to post-process the output and 
# extract meaningful information
srun --ntasks 48 --nodes 6 --exclusive postprocess.mpi <output.txt >result.txt &

# Four step with a sub-allocation on a single node (because maybe 
# it is a multithreaded program that cannot use CPUs on distinct nodes)    
# to compress the raw output. This step runs at the same time as 
# the previous one thanks to the ampersand `&` 
OMP_NUM_THREAD=12 srun --ntasks 12 --nodes 1 --exclusive compress output.txt &

wait


```

Four steps were created and so the accounting information for that job will have 5 lines; one per step plus one for the Bash script itself.


https://stackoverflow.com/questions/46506784/how-do-the-terms-job-task-and-step-relate-to-each-other