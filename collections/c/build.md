## What are Makefile.am and Makefile.in?


Makefile.am is a programmer-defined file and is used by automake to generate the Makefile.in file (the .am stands for automake). The configure script typically seen in source tarballs will use the Makefile.in to generate a Makefile.

The configure script itself is generated from a programmer-defined file named either configure.ac or configure.in (deprecated). I prefer .ac (for autoconf) since it differentiates it from the generated Makefile.in files and that way I can have rules such as make dist-clean which runs rm -f *.in. Since it is a generated file, it is not typically stored in a revision system such as Git, SVN, Mercurial or CVS, rather the .ac file would be.

Read more on GNU Autotools. Read about make and Makefile first, then learn about automake, autoconf, libtool, etc.



https://stackoverflow.com/questions/2531827/what-are-makefile-am-and-makefile-in

https://en.wikipedia.org/wiki/GNU_Autotools

https://stackoverflow.com/questions/26832264/confused-about-configure-script-and-makefile-in/26832773#26832773



##  The basics of autotools



https://devmanual.gentoo.org/general-concepts/autotools/index.html

