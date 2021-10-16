rpmbuild  is  used  to build both binary and source software packages.  A package consists of an archive of files and meta-data used to install  and  erase  the  archive  files.  The  meta-data  includes  helper  scripts,  file attributes, and descriptive information about the package.  Packages come in two varieties: binary packages, used to encapsulate software to be installed, and source packages, containing the source code and recipe necessary  to produce binary packages.

One  of the following basic modes must be selected: Build Package, Build Package from Tarball, Recompile Package,
Show Configuration.



``` shell

rpmbuild {-ta|-tb|-tp|-tc|-ti|-tl|-ts} [rpmbuild-options] TARBALL ...

```



uninstall package

``` shell

rpm -e 

```


install package
``` shell

rpm -ivh 

```