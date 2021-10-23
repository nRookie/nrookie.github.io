##  check if relocatable
``` shell
rpm -qpi [rpm package] | head -1
```


## install to other directory

``` shell
rpm -ivh --prefix=/opt  
```

https://www.thegeekdiary.com/how-to-install-an-rpm-package-into-a-different-directory-in-centos-rhel-fedora/
