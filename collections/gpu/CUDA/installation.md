

### verify gpu



``` shell
lspci | grep -i nvidia
```



### Download driver run file

``` shell
wget https://us.download.nvidia.com/XFree86/Linux-x86_64/495.46/NVIDIA-Linux-x86_64-495.46.run
```



### run it

![image-20220505143656051](/Users/kestrel/developer/nrookie.github.io/collections/gpu/CUDA/image-20220505143656051.png)





https://askubuntu.com/questions/481414/install-nvidia-driver-instead-of-nouveau

it seems like do not need to install it any more.



## Install Cuda



https://docs.nvidia.com/datacenter/tesla/tesla-installation-notes/index.html





### Download

``` shell
wget https://developer.download.nvidia.com/compute/cuda/11.6.2/local_installers/cuda_11.6.2_510.47.03_linux.run
```



#### Verify the System has the Correct Kernel Headers and Development Packages Installed

https://bugzilla.redhat.com/show_bug.cgi?id=1986132



``` shell
[root@10-9-60-113 gpu]# sudo dnf install kernel-devel-$(uname -r) kernel-headers-$(uname -r)
CentOS-8 - PowerTools                                                                                                                                     81  B/s |  38  B     00:00
Error: Failed to download metadata for repo 'PowerTools': Cannot prepare internal mirrorlist: No URLs in mirrorlist
```

fix : https://techglimpse.com/failed-metadata-repo-appstream-centos-8/

``` shell
   dnf install kernel-devel
   dnf install kernel-headers
```





#### Disabling  nouveau



``` shell
lsmod | grep nouveau
```



1. Create a file at /etc/modprobe.d/blacklist-nouveau.conf with the following contents:

2. ``` shell
   blacklist nouveau
   options nouveau modeset=0
   ```

3. Regenerate the kernel initrams

4. ``` shell
   sudo dracut --force
   ```

5. ``` shell
   reboot
   ```





#### Execute Runfile

``` shell
sh cuda_11.6.2_510.47.03_linux.run
```



#### Errors

``` shell
[root@10-9-60-113 gpu]# sh cuda_11.6.2_510.47.03_linux.run

 Installation failed. See log at /var/log/cuda-installer.log for details.
[root@10-9-60-113 gpu]# cat  /var/log/cuda-installer.log
[INFO]: Driver not installed.
[INFO]: Checking compiler version...
[INFO]: gcc location: /usr/bin/gcc

[INFO]: gcc version: gcc version 8.5.0 20210514 (Red Hat 8.5.0-4) (GCC)
```

####

Install the headers



Execute run file with kernel-source-path

``` shell
sh cuda_11.6.2_510.47.03_linux.run  --kernel-source-path=/usr/src/kernels/4.18.0-240.el8.x86_64
```







Installed



``` shell
[root@10-9-60-113 gpu]# ./cuda_11.6.2_510.47.03_linux.run  --kernel-source-path=/usr/src/kernels/4.18.0-240.el8.x86_64
===========
= Summary =
===========

Driver:   Installed
Toolkit:  Installed in /usr/local/cuda-11.6/

Please make sure that
 -   PATH includes /usr/local/cuda-11.6/bin
 -   LD_LIBRARY_PATH includes /usr/local/cuda-11.6/lib64, or, add /usr/local/cuda-11.6/lib64 to /etc/ld.so.conf and run ldconfig as root

To uninstall the CUDA Toolkit, run cuda-uninstaller in /usr/local/cuda-11.6/bin
To uninstall the NVIDIA Driver, run nvidia-uninstall
Logfile is /var/log/cuda-installer.log

```



### Device Node Verification

``` shell
#!/bin/bash

/sbin/modprobe nvidia

if [ "$?" -eq 0 ]; then
  # Count the number of NVIDIA controllers found.
  NVDEVS=`lspci | grep -i NVIDIA`
  N3D=`echo "$NVDEVS" | grep "3D controller" | wc -l`
  NVGA=`echo "$NVDEVS" | grep "VGA compatible controller" | wc -l`

  N=`expr $N3D + $NVGA - 1`
  for i in `seq 0 $N`; do
    mknod -m 666 /dev/nvidia$i c 195 $i
  done

  mknod -m 666 /dev/nvidiactl c 195 255

else
  exit 1
fi

/sbin/modprobe nvidia-uvm

if [ "$?" -eq 0 ]; then
  # Find out the major device number used by the nvidia-uvm driver
  D=`grep nvidia-uvm /proc/devices | awk '{print $1}'`

  mknod -m 666 /dev/nvidia-uvm c $D 0
else
  exit 1
fi
```















https://docs.nvidia.com/cuda/cuda-installation-guide-linux/index.html
