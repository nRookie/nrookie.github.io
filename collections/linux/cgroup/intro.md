how to check if cgroup v2 is installed on my system?

``` shell
grep cgroup /proc/filesystems
nodev   cgroup
nodev   cgroup2

```


docs

https://www.kernel.org/doc/html/latest/admin-guide/cgroup-v2.html



## Systemd Unit Types

All processes running on the system are child processes of the systemd init process.


``` shell
[root@node2 slurm_deployer]# ps -axu | grep " 1 "
root         1  0.0  0.1 191056  5372 ?        Ss   Oct08   0:05 /usr/lib/systemd/systemd --switched-root --system --deserialize 22
root      2644  0.0  0.0 112724  2340 pts/5    S+   17:16   0:00 grep --color=auto  1
```

Systemd provides three unit types that are used for the purpose of resource control 


### Service 
 A process or a group of processes, which systemd started based on a unit configuration file. Services encapsulate the specified processes so that they can be started and stopped as one set. Services are named in the following way:

``` shell
    name.service
```
### Scope

A group of externally created processes. Scopes encapsulate processes that are started and stopped by arbitrary processes through the fork() function and then registered by systemd at runtime. For instance, user sessions, containers, and virtual machines are treated as scopes. Scopes are named as follows:

``` shell
name.scope

```

Here, name stands for the name of the scope.


### Slice 

A group of hierarchically organized units. Slices do not contain processes, they organize a hierarchy in which scopes and services are placed. The actual processes are contained in scopes or in services. In this hierarchical tree, every name of a slice unit corresponds to the path to a location in the hierarchy. The dash ("-") character acts as a separator of the path components. For example, if the name of a slice looks as follows:

``` shell
parent-name.slice
```

it means that a slice called parent-name.slice is a subslice of the parent.slice. This slice can have its own subslice named parent-name-name2.slice, and so on.


There is one root slice denoted as:

``` shell
-.slice
```

Service, scope, and slice units directly map to objects in the cgroup tree. When these units are activated, they map directly to cgroup paths built from the unit names. For example, the ex.service residing in the test-waldo.slice is mapped to the cgroup test.slice/test-waldo.slice/ex.service/.

Services, scopes, and slices are created manually by the system administrator or dynamically by programs. By default, the operating system defines a number of built-in services that are necessary to run the system. Also, there are four slices created by default:


- -.slice - the root slice;
- system.slice -- the default place for all system services;
-  user.slice -- the default place for all user sessions
- machine.slice — the default place for all virtual machines and Linux containers.



## Command to  Recursively show control group contents
``` shell
[root@node2 slurm_deployer]# systemd-cgls 
├─1 /usr/lib/systemd/systemd --switched-root --system --deserialize 22
├─user.slice
│ └─user-0.slice
│   ├─session-692.scope
│   │ ├─ 7313 systemd-cgls
│   │ ├─ 7314 less
│   │ ├─19590 sshd: root@pts/0    
│   │ └─19592 -bash
│   └─session-696.scope
│     ├─ 3110 ssh: /root/.ansible/cp/49632296ca [mux]                                                                                                              
│     ├─ 3113 ssh: /root/.ansible/cp/b07752f855 [mux]                                                                                                              
│     ├─ 3116 ssh: /root/.ansible/cp/6a02093385 [mux]                                                                                                              
│     ├─ 3637 ssh: /root/.ansible/cp/39bb01fa31 [mux]                                                                                                              
│     ├─ 3640 ssh: /root/.ansible/cp/a86ec95546 [mux]                                                                                                              
│     ├─ 4130 /usr/bin/python2 /usr/bin/ansible-playbook -i hosts slurm_install.yaml -v
│     ├─ 7291 /usr/bin/python2 /usr/bin/ansible-playbook -i hosts slurm_install.yaml -v
│     ├─ 7302 sshpass -d8 ssh -C -o ControlMaster=auto -o ControlPersist=60s -o StrictHostKeyChecking=no -o Port=22 -o User="root" -o ConnectTimeout=10 -o ControlP
│     ├─ 7303 ssh -C -o ControlMaster=auto -o ControlPersist=60s -o StrictHostKeyChecking=no -o Port=22 -o User="root" -o ConnectTimeout=10 -o ControlPath=/root/.a
│     ├─30264 sshd: root@pts/5    
│     └─30511 -bash
└─system.slice
  ├─rngd.service
  │ └─709 /sbin/rngd -f
  ├─libstoragemgmt.service
  │ └─702 /usr/bin/lsmd -d
  ├─systemd-udevd.service
  │ └─487 /usr/lib/systemd/systemd-udevd
  ├─system-serial\x2dgetty.slice
  │ └─serial-getty@ttyS0.service
  │   └─1265 /sbin/agetty --keep-baud 115200,38400,9600 ttyS0 vt220
  ├─polkit.service
  │ └─697 /usr/lib/polkit-1/polkitd --no-debug
lines 1-32
```

As you can see, services and scopes contain processes and are placed in slices that do not contain processes of their own. The only exception is PID 1 that is located in the special systemd slice marked as -.slice. Also note that -.slice is not shown as it is implicitly identified with the root of the entire tree.