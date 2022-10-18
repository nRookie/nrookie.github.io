https://wiki.openfoam.com/Tutorials



# Install OpenFoam



https://www.openfoam.com/current-release





## From Source

https://develop.openfoam.com/Development/openfoam/-/blob/master/doc/Build.md



1. download packages

```shell
 wget --no-check-certificate https://sourceforge.net/projects/openfoam/files/v2112/OpenFOAM-v2112.tgz
```

2. download packages

``` shell
wget --no-check-certificate https://sourceforge.net/projects/openfoam/files/v2112/ThirdParty-v2112.tgz
```

3. unpack

   ```shell
   tar -xzvf OpenFOAM-v2112.tgz
   ```

4. check system requirements

   1. https://develop.openfoam.com/Development/openfoam/blob/develop/doc/Requirements.md

5. Install boost

   ``` shell
   wget --no-check-certificate http://sourceforge.net/projects/boost/files/boost/1.48.0/boost_1_48_0.tar.gz
   tar -xzvf boost_1_48_0.tar.gz
   cd boost_1_48_0/
   ./bootstrap.sh
   vi ./tools/build/v2/user-config.jam
   using mpi ;
   ./b2 install --prefix=/nfs1/opt/boost/1.48.0/

   ```

6. Install cmake

   ``` shell
   wget --no-check-certificate https://gitlab.kitware.com/cmake/cmake/-/archive/v3.11.0/cmake-v3.11.0.tar.gz
   tar -xzvf cmake-v3.11.0.tar.gz
   cd cmake-v3.11.0/
   ./bootstrap --prefix=/nfs1/opt/cmake/3.11.0

   ```

7. install fftw

   ``` shell
   wget http://www.fftw.org/fftw-3.3.10.tar.gz
   tar -xzvf fftw-3.3.10.tar.gz
   ./configure --prefix=/nfs1/fftw/3.3.10
   ```

8. install qt

   ``` shell
   wget https://download.qt.io/archive/qt/5.12/5.12.0/single/qt-everywhere-src-5.12.0.tar.xz
   tar -xvf qt-everywhere-src-5.12.0.tar.xz
   cd qt-everywhere-src-5.12.0/
   ./configure --prefix=/nfs1/opt/qt/5.12.0
   yum install openssl -y
   make
   make install
   ```

   Err losing qtx11extra

   ``` shell
   https://forum.qt.io/topic/115827/build-on-linux-qt-xcb-option
   https://forum.qt.io/topic/115827/build-on-linux-qt-xcb-option/20
   ```

   ``` shell
   yum install libxcb-devel libxkbcommon-devel xcb-util-devel xcb-util-image-devel xcb-util-keysyms-devel xcb-util-renderutil-devel xcb-util-wm-devel mesa-libGL-devel -y

   ./configure --prefix=/nfs1/opt/qt/5.12.12 -xcb -xcb-xlib -xcb-xinput
   ```



9. Install Paraview

   ref: https://gitlab.kitware.com/paraview/paraview/blob/v5.6.1/Documentation/dev/build.md

   ``` shell
   git clone https://gitlab.kitware.com/paraview/paraview.git
   mkdir paraview_build
   cd paraview
   git checkout tag
   git submodule update --init --recursive
   cd ../paraview_build
   module load cmake/3.15.0
   export cc=path to set the gcc version of cmake
   cmake -DCMAKE_PREFIX_PATH=/nfs1/opt/paraview/5.6.0 -GNinja -DPARAVIEW_USE_PYTHON=ON -DPARAVIEW_USE_MPI=ON -DVTK_SMP_IMPLEMENTATION_TYPE=TBB -DCMAKE_BUILD_TYPE=Release ../paraview
   https://discourse.paraview.org/t/how-to-install-vtk-as-a-standalone-package/930
   alias ninja=ninja-build
   ninja
   ```

   err

   ``` shell
   https://www.cfd-online.com/Forums/openfoam-installation/197626-paraview-cmake-error.html
   https://forum.qt.io/topic/102066/building-5-12-3-on-ubuntu-19-04-results-in-missing-x11extras-module/4
   https://stackoverflow.com/questions/17275348/how-to-specify-new-gcc-path-for-cmake
   ```

   Fix

   ``` shell
    export LD_LIBRARY_PATH="/usr/local/lib64/:$LD_LIBRARY_PATH"
   ```

   Threading Building Blocks 4.1.1 is not compatible with gcc 7.3 gcc8.3

   ``` shell
   cmake -DCMAKE_PREFIX_PATH=/nfs1/opt/paraview/5.6.0 -GNinja -DPARAVIEW_USE_PYTHON=ON -DPARAVIEW_USE_MPI=ON  -DCMAKE_BUILD_TYPE=Release ../paraview
   ```

   cannot forward x11

   ``` shell
   yum install xorg-x11-xauth
   vi /etc/ssh/sshd_config
   Set `X11UseLocalhost no`
   ```

   successfully run the paraview, but have errors

   ![image-20220114142945682](/Users/kestrel/developer/nrookie.github.io/collections/HPC/openfoam/image-20220114142945682.png)

   ![image-20220114143119562](/Users/kestrel/developer/nrookie.github.io/collections/HPC/openfoam/image-20220114143119562.png)

   ``` shell
   [root@primary paraview_build]# ./bin/vtkProbeOpenGLVersion
   libGL error: unable to load driver: swrast_dri.so
   libGL error: failed to load driver: swrast
   X Error of failed request:  BadValue (integer parameter out of range for operation)
     Major opcode of failed request:  149 (GLX)
     Minor opcode of failed request:  3 (X_GLXCreateContext)
     Value in failed request:  0x0
     Serial number of failed request:  55
     Current serial number in output stream:  56
   [root@primary paraview_build]# yum install -y mesa-libGLw-devel.x86_64
   yum install mesa-dri-drivers

   [root@primary paraview_build]# ./bin/vtkProbeOpenGLVersion
   libGL error: No matching fbConfigs or visuals found
   libGL error: failed to load driver: swrast
   X Error of failed request:  BadValue (integer parameter out of range for operation)
     Major opcode of failed request:  149 (GLX)
     Minor opcode of failed request:  3 (X_GLXCreateContext)
     Value in failed request:  0x0
     Serial number of failed request:  61
     Current serial number in output stream:  62

   [root@primary paraview_build]# glxinfo
   name of display: localhost:10.0
   libGL: OpenDriver: trying /usr/lib64/dri/tls/swrast_dri.so
   libGL: OpenDriver: trying /usr/lib64/dri/swrast_dri.so
   libGL: Can't open configuration file /etc/drirc: No such file or directory.
   libGL: Can't open configuration file /root/.drirc: No such file or directory.
   libGL: Can't open configuration file /etc/drirc: No such file or directory.
   libGL: Can't open configuration file /root/.drirc: No such file or directory.
   libGL error: No matching fbConfigs or visuals found
   libGL error: failed to load driver: swrast
   X Error of failed request:  GLXBadContext
     Major opcode of failed request:  149 (GLX)
     Minor opcode of failed request:  6 (X_GLXIsDirect)
     Serial number of failed request:  31
     Current serial number in output stream:  30
   [root@primary paraview_build]# yum remove  mesa-libGL
   [root@primary paraview_build]# ./bin/
   paraview                    pvdataserver                smTestDriver                vtkH5make_libsettings       vtkParseJava-pv5.6          vtkWrapHierarchy-pv5.6      vtkWrapPython-pv5.6
   paraview-config             pvrenderserver              TestingDemo                 vtkkwProcessXML-pv5.6       vtkProbeOpenGLVersion       vtkWrapJava-pv5.6
   protoc                      pvserver                    vtkH5detect                 vtkLegacyColorMapXMLToJSON  vtkWrapClientServer-pv5.6   vtkWrapPythonInit-pv5.6
   [root@primary paraview_build]# ./bin/vtk
   vtkH5detect                 vtkkwProcessXML-pv5.6       vtkParseJava-pv5.6          vtkWrapClientServer-pv5.6   vtkWrapJava-pv5.6           vtkWrapPython-pv5.6
   vtkH5make_libsettings       vtkLegacyColorMapXMLToJSON  vtkProbeOpenGLVersion       vtkWrapHierarchy-pv5.6      vtkWrapPythonInit-pv5.6
   [root@primary paraview_build]# ./bin/vtkProbeOpenGLVersion
   ./bin/vtkProbeOpenGLVersion: error while loading shared libraries: libGLX.so.0: cannot open shared object file: No such file or directory
   ```

   ``` shell
   yum install  glxinfo
   ```

