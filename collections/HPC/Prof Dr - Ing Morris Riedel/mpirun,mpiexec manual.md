**NAME**

​    orterun, mpirun, mpiexec - Execute serial and parallel jobs in Open MPI. oshrun, shmemrun - Execute serial

​    and parallel jobs in Open SHMEM.





**SYNOPSIS**



Single Process Multiple Data (SPMD) Model:





 **mpirun** [ options ] **<program>** [ <args> ]



 Multiple Instruction Multiple Data (MIMD) Model:



​    Multiple Instruction Multiple Data (MIMD) Model:



​    **mpirun** [ global_options ]

​       [ local_options1 ] **<program1>** [ <args1> ] :

​       [ local_options2 ] **<program2>** [ <args2> ] :

​       ... :

​       [ local_optionsN ] **<programN>** [ <argsN> ]



​    Note that in both models, invoking mpirun via an absolute path name is equivalent to specifying the --prefix

​    option with a <dir> value equivalent to the directory where mpirun resides, minus its last subdirectory.  For

​    example:



​      **%** /usr/local/bin/mpirun ...



​    is equivalent to



​      **%** mpirun --prefix /usr/local





**QUICK** **SUMMARY**

​    If you are simply looking for how to run an MPI application, you probably want to use a command line of the

​    following form:





 **%** mpirun [ -np X ] [ --hostfile <filename> ] <program>







​    This will run X copies of <program> in your current run-time environment (if running under a supported

​    resource manager, Open MPI's mpirun will usually automatically use the corresponding resource manager process

​    starter, as opposed to, for example, rsh or ssh, which require the use of a hostfile, or will default to run‐

​    ning all X copies on the localhost), scheduling (by default) in a round-robin fashion by CPU slot. See the

​    rest of this page for more details.



​    Please note that mpirun automatically binds processes as of the start of the v1.8 series. Three binding pat‐

​    terns are used in the absence of any further directives:



 **Bind** **to** **core:**   when the number of processes is <= 2



**Bind** **to** **socket:**  when the number of processes is > 2



**Bind** **to** **none:**   when oversubscribed



 If your application uses threads, then you probably want to ensure that you are either not bound at all (by

​    specifying --bind-to none), or bound to multiple cores using an appropriate binding level or specific number

​    of processing elements per application process.





**DEFINITION** **OF** **'SLOT'**



The term "slot" is used extensively in the rest of this manual page. A slot is an allocation unit for a

​    process. The number of slots on a node indicate how many processes can potentially execute on that node.  By

​    default, Open MPI will allow one process per slot.



   If Open MPI is not explicitly told how many slots are available on a node (e.g., if a hostfile is used and the

​    number of slots is not specified for a given node), it will determine a maximum number of slots for that node

​    in one of two ways:





​    1. Default behavior

​     By default, Open MPI will attempt to discover the number of processor cores on the node, and use that as

​     the number of slots available.



​    2. When --use-hwthread-cpus is used

​     If --use-hwthread-cpus is specified on the mpirun command line, then Open MPI will attempt to discover the

​     number of hardware threads on the node, and use that as the number of slots available.



This default behavior also occurs when specifying the -host option with a single host. Thus, the command:



  **Slots** **are** **not** **hardware** **resources**

​    Slots are frequently incorrectly conflated with hardware resources. It is important to realize that slots are an entirely different metric

​    than the number (and type) of hardware resources available.



​    Here are some examples that may help illustrate the difference:



​    \1. More processor cores than slots



​     Consider a resource manager job environment that tells Open MPI that there is a single node with 20 processor cores and 2 slots available.

​     By default, Open MPI will only let you run up to 2 processes.



​     Meaning: you run out of slots long before you run out of processor cores.



​    \2. More slots than processor cores



​     Consider a hostfile with a single node listed with a "slots=50" qualification. The node has 20 processor cores. By default, Open MPI will

​     let you run up to 50 processes.



​     Meaning: you can run many more processes than you have processor cores.



**DEFINITION** **OF** **'PROCESSOR** **ELEMENT'**

​    By default, Open MPI defines that a "processing element" is a processor core. However, if --use-hwthread-cpus is specified on the mpirun com‐

​    mand line, then a "processing element" is a hardware thread.