# Installation

## Install lua

### download
``` shell
wget https://sourceforge.net/projects/lmod/files/lua-5.1.4.9.tar.bz2
```

### install
``` shell
tar xf lua-X.Y.Z.tar.bz2
cd lua-X.Y.Z
./configure --prefix=/opt/apps/lua/X.Y.Z
make; make install
cd /opt/apps/lua; ln -s X.Y.Z lua
mkdir /usr/local/bin; ln s /opt/apps/lua/lua/bin/lua /usr/local/bin
```



## Why does Lmod install differently?¶

Lmod automatically creates a version directory for itself. So, for example, if the installation prefix is set to /apps, and the current version is X.Y.Z, installation will create /apps/lmod and /apps/lmod/X.Y.Z. This way of configuring is different from most packages. There are two reasons for this:

Lmod is designed to have just one version of it running at one time. Users will not be switching version during the course of their interaction in a shell.
By making the symbolic link the startup scripts in /etc/profile.d do not have to change. They just refer to /apps/lmod/lmod/... and not /apps/lmod/X.Y.Z/...

