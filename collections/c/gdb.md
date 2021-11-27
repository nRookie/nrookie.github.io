
##  gdb common command

1. show file name

``` shell
info line
```

2. show next line to execute

``` shell
frame
```


3. print name of a symbol

``` shell
info symbol 0x7ffff7dcd790
```


4. variables

``` shell
info args
info variables
info locals
```


5. finish

Continue running until just after function in the selected stack frame returns. Print the returned value (if any). This command can be abbreviated as fin.




6. debugging with fork


 https://stackoverflow.com/questions/15126925/debugging-child-process-after-fork-follow-fork-mode-child-configured

https://sourceware.org/gdb/onlinedocs/gdb/Forks.html


``` shell
set follow-fork-mode child
```



7. display value at break points

3: file_name = 0x0
2: dir_array[i]	= 98 'b'
1: i = 21


8. break point specific point
``` shell

break file.c:linenumber

```
 ## use with tui


``` shell
 gdb -tui
 ```

 https://sourceware.org/gdb/current/onlinedocs/gdb/TUI-Commands.html




 ## gdb error


 (gdb) n段错误(吐核)

