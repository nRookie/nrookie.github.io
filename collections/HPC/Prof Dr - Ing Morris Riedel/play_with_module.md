https://modules.readthedocs.io/en/stable/INSTALL.html

#### Enable Modules initialization at shell startup.



1. set module path

2. ``` shell
   ./configure --with-modulepath=/nfs1/opt/modulefiles --prefix=/nfs1/opt/modules-5.0.1
   make
   make install
   ```

3. ``` shell
   ln -s /nfs1/opt/modules-5.0.1/init/profile.sh  /etc/profile.d/modules.sh
   ln -s /nfs1/opt/modules-5.0.1/init/profile.csh   /etc/profile.d/modules.csh

   source PREFIX/init/bash

   [root@primary ~]# module --version
   Modules Release 5.0.1 (2021-10-16)
   [root@node1 opt]# module avail
   [root@node1 opt]#
   ```

4. download gcc

   ![image-20211127182303431](/Users/kestrel/developer/nrookie.github.io/collections/HPC/Prof Dr - Ing Morris Riedel/image-20211127182303431.png)

5. install the gcc(4.8.3, 4.7.2)

6. https://stackoverflow.com/questions/9253695/building-gcc-requires-gmp-4-2-mpfr-2-3-1-and-mpc-0-8-0

7. https://linuxhostsupport.com/blog/how-to-install-gcc-on-centos-7/

8. ``` shell
   cd /nfs1/software/gcc-4.7.2
   ./contrib/download_prerequisites
   ./configure --prefix=/nfs1/opt/gcc/4.7.2/ --disable-multilib --enable-languages=c,c++
   yum install glibc-devel.i686 # need to add this, otherwise no file stub.h will show
   make -j 4
   make install
   ```

9. do the same for the gcc 4.8.2

10. ``` shell
    cd /nfs1/software/gcc-4.8.3
    ./contrib/download_prerequisites
    ./configure --prefix=/nfs1/opt/gcc/4.8.3/ --disable-multilib --enable-languages=c,c++
    yum install glibc-devel.i686 # need to add this, otherwise no file stub.h will show
    make -j 4
    make install
    ```

11. during make error occurred, (do not enable multiline)

12. https://stackoverflow.com/questions/7412548/error-gnu-stubs-32-h-no-such-file-or-directory-while-compiling-nachos-source

13.

14. create modulefiles for gcc

15. ``` shell
    cd /nfs1/opt/
    mkdir modulefiles
    mkdir -p modulefiles/gcc
    vi modulefiles/gcc/4.8.3
    #%Module1.0#####################################################################
    ##
    ## null modulefile
    ##
    ## modulefiles/null.  Generated from null.in by configure.
    ##

    # for Tcl script use only
    set             root                    /nfs1//opt/gcc/4.7.2
    prepend-path    PATH                    $root/bin
    prepend-path    LD_LIBRARY_PATH         $root/lib

    vi modulefiles/gcc/4.7.2



    [root@primary opt]# vi modulefiles/gcc/4.8.3
    [root@primary opt]# vi modulefiles/gcc/4.7.2
    [root@primary opt]# vi  modulefiles/gcc/4.8.3
    [root@primary opt]# module avail
    --------------------------------------------------------- /nfs1/opt/modulefiles ---------------------------------------------------------
    gcc/4.7.2  gcc/4.8.3

    Key:
    modulepath
    ```

16.







### Use some command



1. module avail

``` shell
[root@primary ~]# module avail
---------------------------------------------------- /usr/local/Modules/modulefiles -----------------------------------------------------
dot  module-git  module-info  modules  null  use.own

Key:
modulepath
```



2. module load

   ``` shell
   [root@primary gcc]# module load gcc/4.7.2
   [root@primary gcc]# gcc
   gcc         gcc-ar      gcc-nm      gcc-ranlib
   [root@primary gcc]# gcc --version
   gcc (GCC) 4.7.2
   Copyright © 2012 Free Software Foundation, Inc.
   本程序是自由软件；请参看源代码的版权声明。本软件没有任何担保；
   包括没有适销性和某一专用目的下的适用性担保。
   [root@primary gcc]# module list
   Currently Loaded Modulefiles:
    1) gcc/4.7.2
   [root@primary gcc]# module avail
   ----------------------------------------------------------------------------- /nfs1/opt/modulefiles ------------------------------------------------------------------------------
   gcc/4.7.2  gcc/4.8.3

   Key:
   loaded  modulepath
   ```

3. module unload

4. ``` shell
   [root@primary gcc]# module unload gcc/4.7.2
   [root@primary gcc]# gcc --version
   gcc (GCC) 4.8.5 20150623 (Red Hat 4.8.5-44)
   Copyright © 2015 Free Software Foundation, Inc.
   本程序是自由软件；请参看源代码的版权声明。本软件没有任何担保；
   包括没有适销性和某一专用目的下的适用性担保。
   ```

5. in other nodes

6. ``` shell
   [root@primary gcc]# ssh node1
   Last login: Sat Nov 27 21:13:20 2021 from 10.23.131.240
   [root@node1 ~]# gcc --version
   gcc (GCC) 4.8.5 20150623 (Red Hat 4.8.5-44)
   Copyright © 2015 Free Software Foundation, Inc.
   本程序是自由软件；请参看源代码的版权声明。本软件没有任何担保；
   包括没有适销性和某一专用目的下的适用性担保。
   [root@node1 ~]# module
   add           --color       display       initadd       is-loaded     --no-pager    refresh       -s            show          -T            unuse         -w
   aliases       --color=      edit          initclear     is-saved      --paginate    reload        save          sh-to-mod     test          use           whatis
   append-path   config        -h            initlist      is-used       path          remove        savelist      --silent      --trace       -v            --width
   apropos       -D            help          initprepend   keyword       paths         remove-path   saverm        source        try-add       -V            --width=
   avail         --debug       --help        initrm        list          prepend-path  restore       saveshow      swap          try-load      --verbose
   clear         del           info-loaded   is-avail      load          purge         rm            search        switch        unload        --version
   [root@node1 ~]# module load gcc/4.
   gcc/4.7.2  gcc/4.8.3
   [root@node1 ~]# module load gcc/4.7.2
   [root@node1 ~]# gcc --version
   gcc (GCC) 4.7.2
   Copyright © 2012 Free Software Foundation, Inc.
   本程序是自由软件；请参看源代码的版权声明。本软件没有任何担保；
   包括没有适销性和某一专用目的下的适用性担保。
   ```

7.



https://www.admin-magazine.com/HPC/Articles/Environment-Modules



https://www.youtube.com/watch?v=LLJ_a8hS1GU







### Install OpenMPI



1. download, unzip the OpenMPI

2. ``` shell
    wget --no-check-certificate https://download.open-mpi.org/release/open-mpi/v3.0/openmpi-3.0.3.tar.bz2
    tar -xvf openmpi-3.0.3.tar.bz2
   ```

3. ``` shell
    ./configure --prefix=/nfs1/opt/openmpi/3.0.3/ CC=/nfs1/opt/gcc/4.7.2/bin/gcc CXX=/nfs1/opt/gcc/4.7.2/bin/g++ FC=/nfs1/opt/gcc/4.7.2/bin/gfortran
   ```

4. 报错

5. ![image-20211127221559759](/Users/kestrel/developer/nrookie.github.io/collections/HPC/Prof Dr - Ing Morris Riedel/image-20211127221559759.png)

6. ``` shell
    module unload gcc/4.7.2
   ```

7. ``` shell
   ./configure --prefix=/nfs1/opt/openmpi/3.0.3/
   Open MPI configuration:
   -----------------------
   Version: 3.0.3
   Build MPI C bindings: yes
   Build MPI C++ bindings (deprecated): no
   Build MPI Fortran bindings: mpif.h, use mpi
   MPI Build Java bindings (experimental): no
   Build Open SHMEM support: yes
   Debug build: no
   Platform file: (none)

   Miscellaneous
   -----------------------
   CUDA support: no
   PMIx support: internal

   Transports
   -----------------------
   Cray uGNI (Gemini/Aries): no
   Intel Omnipath (PSM2): no
   Intel SCIF: no
   Intel TrueScale (PSM): no
   Mellanox MXM: no
   Open UCX: no
   OpenFabrics Libfabric: no
   OpenFabrics Verbs: no
   Portals4: no
   Shared memory/copy in+copy out: yes
   Shared memory/Linux CMA: yes
   Shared memory/Linux KNEM: no
   Shared memory/XPMEM: no
   TCP: yes

   Resource Managers
   -----------------------
   Cray Alps: no
   Grid Engine: no
   LSF: no
   Moab: no
   Slurm: yes
   ssh/rsh: yes
   Torque: no

   OMPIO File Systems
   -----------------------
   Generic Unix FS: yes
   Lustre: no
   PVFS2/OrangeFS: no
   ```

8. ``` shell
   make -j 4
   make install
   ```

9. ``` shell
   [root@primary 3.0.3]# module show openmpi/3.0.3
   -------------------------------------------------------------------
   /nfs1/opt/modulefiles/openmpi/3.0.3:

   prepend-path    PATH /nfs1/opt/openmpi/3.0.3/bin
   prepend-path    LD_LIBRARY_PATH /nfs1/opt/openmpi/3.0.3/lib
   -------------------------------------------------------------------
   ```

10. ``` shell
    [root@primary 3.0.3]# module list
    Currently Loaded Modulefiles:
     1) openmpi/3.0.3
    [root@primary 3.0.3]# mpirun --version
    mpirun (Open MPI) 3.0.3

    Report bugs to http://www.open-mpi.org/community/help/
    ```



### 使用MPIRUN 执行作业

1. 编写脚本

2. ``` shell
   #!/bin/bash
   #SBATCH -J hello-example
   #SBATCH -N 3
   #SBATCH --mail-user=whatever@google.com
   #SBATCH --mail-type=end
   module load openmpi/3.0.3

   srun /bin/hostname
   #mpirun  --allow-run-as-root  /bin/hostname
   ```

3. 执行结果

4. ![image-20211128102004061](/Users/kestrel/developer/nrookie.github.io/collections/HPC/Prof Dr - Ing Morris Riedel/image-20211128102004061.png)

5. **为什么MPIRUN会在一个节点上执行两次hostname？**









# Environment settings



https://docs.fedoraproject.org/en-US/packaging-guidelines/EnvironmentModules/



``` shell


\#%Module1.0#####################################################################

\##

\## null modulefile

\##

\## modulefiles/null. Generated from null.in by configure.

\##



\# for Tcl script use only

set       root          /nfs1/opt/openmpi/3.0.3

prepend-path  PATH          $root/bin

prepend-path  LD_LIBRARY_PATH     $root/lib

prepend-path  MANPATH         $root/share/man

setenv     MPI_SYSCONFIG      $root/etc

setenv     MPI_INCLUDE       $root/include

setenv     MPI_LIB         $root/lib
```

