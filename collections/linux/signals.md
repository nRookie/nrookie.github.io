## SIGHUP

SIGHUP        1       Term    Hangup detected on controlling terminal
                                    or death of controlling process





## Use strace to catch the signal


``` shell
[root@10-23-181-93 ~]# strace -p  2622  -e 'trace=!all'
strace: Process 2622 attached
--- SIGHUP {si_signo=SIGHUP, si_code=SI_USER, si_pid=2622, si_uid=0} ---
```



