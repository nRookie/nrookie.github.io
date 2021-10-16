

Transient cgroups are released automatically as soon as the processes they contain finish. By passing the --remain‑after-exit option to systemd-run you can keep the unit running after its processes finished to collect runtime information. To stop the unit gracefully, type:

``` shell
~]# systemctl stop name.service
Undefined
```
Replace name with the name of the service you wish to stop. To terminate one or more of the unit's processes, type as root:

``` shell
~]# systemctl kill name.service --kill-who=PID,... --signal=signal
Undefined
```

Replace name with a name of the unit, for example httpd.service. Use --kill-who to select which processes from the cgroup you wish to terminate. To kill multiple processes at the same time, pass a comma-separated list of PIDs. Replace signal with the type of POSIX signal you wish to send to specified processes. Default is SIGTERM. For more information, see the systemd.kill manual page.


Persistent cgroups are released when the unit is disabled and its configuration file is deleted by running:

``` shell
~]# systemctl disable name.service
Undefined
```

where name stands for the name of the service to be disabled.


## 2.3. MODIFYING CONTROL GROUPS

Each persistent unit supervised by systemd has a unit configuration file in the /usr/lib/systemd/system/ directory. To change parameters of a service unit, modify this configuration file. This can be done either manually or from the command-line interface by using the systemctl set-property command.


### 2.3.1. Setting Parameters from the Command-Line Interface


The systemctl set-property command allows you to persistently change resource control settings during the application runtime. To do so, use the following syntax as root:
``` shell
~]# systemctl set-property name parameter=value
Undefined
```




Replace name with the name of the systemd unit you wish to modify, parameter with a name of the parameter to be changed, and value with a new value you want to assign to this parameter.

Not all unit parameters can be changed at runtime, but most of those related to resource control may, see Section 2.3.2, “Modifying Unit Files” for a complete list. Note that systemctl set-property allows you to change multiple properties at once, which is preferable over setting them individually.

The changes are applied instantly, and written into the unit file so that they are preserved after reboot. You can change this behavior by passing the --runtime option that makes your settings transient:


``` shell

~]# systemctl set-property --runtime name property=value
Undefined

```


Example 2.2. Using systemctl set-property

To limit the CPU and memory usage of httpd.service from the command line, type:


``` shell
~]# systemctl set-property httpd.service CPUShares=600 MemoryLimit=500M
Undefined
```


To make this a temporary change, add the --runtime option:


``` shell
~]# systemctl set-property --runtime httpd.service CPUShares=600 MemoryLimit=500M
Undefined
```



https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/7/html/resource_management_guide/sec-modifying_control_groups