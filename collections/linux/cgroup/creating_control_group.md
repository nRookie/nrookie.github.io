To run a command in a specified cgroup, type as root:


``` shell

systemd-run --unit=name --scope --slice=slice_name command

```



The name stands for the name you want the unit to be known under. If --unit is not specified, a unit name will be generated automatically. It is recommended to choose a descriptive name, since it will represent the unit in the systemctl output. The name has to be unique during runtime of the unit.


Use the optional --scope parameter to create a transient scope unit instead of service unit that is created by default.


With the --slice option, you can make your newly created service or scope unit a member of a specified slice. Replace slice_name with the name of an existing slice (as shown in the output of systemctl -t slice), or create a new slice by passing a unique name. By default, services and scopes are created as members of the system.slice.


Replace command with the command you wish to execute in the service unit. Place this command at the very end of the systemd-run syntax, so that the parameters of this command are not confused for parameters of systemd-run.


Besides the above options, there are several other parameters available for systemd-run. For example, --description creates a description of the unit, --remain-after-exit allows to collect runtime information after terminating the service's process. The --machine option executes the command in a confined container. See the systemd-run(1) manual page to learn more.


## example 

``` shell
systemd-run --unit=toptest --slice=test top -b
Undefined
```

test with systemctl

``` shell
[root@node2 new]# systemctl status toptest
● toptest.service - /usr/bin/top -b
   Loaded: loaded (/run/systemd/system/toptest.service; static; vendor preset: disabled)
  Drop-In: /run/systemd/system/toptest.service.d
           └─50-Description.conf, 50-ExecStart.conf, 50-Slice.conf
   Active: active (running) since Wed 2021-10-13 16:59:30 CST; 43s ago
 Main PID: 1271 (top)
   CGroup: /test.slice/toptest.service
           └─1271 /usr/bin/top -b

Oct 13 17:00:12 node2 top[1271]: 8534 root      20   0  937956  40476  29644 S   0.0  1.0   0:10.60 node
Oct 13 17:00:12 node2 top[1271]: 8593 root      20   0  116456   4660   3200 S   0.0  0.1   0:00.01 bash
Oct 13 17:00:12 node2 top[1271]: 14597 root      20   0  159356   9864   8496 S   0.0  0.3   0:00.09 sshd
Oct 13 17:00:12 node2 top[1271]: 14707 root      20   0  160564  10804   8236 S   0.0  0.3   0:00.95 sshd
Oct 13 17:00:12 node2 top[1271]: 14709 root      20   0  115316   3148   2896 S   0.0  0.1   0:00.00 bash
Oct 13 17:00:12 node2 top[1271]: 14719 root      20   0  113456   3492   2928 S   0.0  0.1   0:00.02 bash
Oct 13 17:00:12 node2 top[1271]: 14767 root      20   0 1026344 105684  38104 S   0.0  2.7   0:08.44 node
Oct 13 17:00:12 node2 top[1271]: 14774 root      20   0  909464  39892  29284 S   0.0  1.0   0:00.59 node
Oct 13 17:00:12 node2 top[1271]: 19514 root      20   0       0      0      0 I   0.0  0.0   0:00.03 kworker/u4+
Oct 13 17:00:12 node2 top[1271]: 26692 rpc       20   0   69344   3868   3248 S   0.0  0.1   0:00.50 rpcbind
[root@node2 new]# 

```

## 2.1.2. Creating Persistent Cgroups

To configure a unit to be started automatically on system boot, execute the systemctl enable command .

Running this command automatically creates a unit file in the /usr/lib/systemd/system/ directory. 

To make persistent changes to the cgroup, add or modify configuration parameters in its unit file. For more information, see Section 2.3.2, “Modifying Unit Files”.


https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/7/html/resource_management_guide/chap-using_control_groups