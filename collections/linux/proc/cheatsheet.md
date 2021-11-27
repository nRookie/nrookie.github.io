## prevent oom killed

``` shell
/proc/[pid]/oom_adj (since Linux 2.6.11)
        This  file  can  be  used  to adjust the score used to select which process should be killed in an out-of-memory (OOM) situation.  The kernel uses this value for a bit-shift operation of the process's oom_score
        value: valid values are in the range -16 to +15, plus the special value -17, which disables OOM-killing altogether for this process.  A positive score increases the likelihood of this process  being  killed  by
        the OOM-killer; a negative score decreases the likelihood.

        The default value for this file is 0; a new process inherits its parent's oom_adj setting.  A process must be privileged (CAP_SYS_RESOURCE) to update this file.

        Since Linux 2.6.36, use of this file is deprecated in favor of /proc/[pid]/oom_score_adj.
```



## read env variable


https://bogdancornianu.com/how-to-get-environment-variables-from-a-running-process-in-linux/

``` shell

``` 