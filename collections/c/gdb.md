
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