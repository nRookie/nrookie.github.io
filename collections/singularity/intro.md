# 使用 Singularity

## Build

singularity支持以下几种方式构建镜像：

- URI beginning with **library://** to build from the Container Library
- URI beginning with **docker://** to build from Docker Hub
- URI beginning with **shub://** to build from Singularity Hub
- path to a **existing container** on your local machine
- path to a **directory** to build from a sandbox
- path to a [SingularityCE definition file](https://sylabs.io/guides/3.8/user-guide/definition_files.html#definition-files)

```
singularity build test.sif shub://opensciencegrid/osgvo-tensorflow
singularity build test.sif library://hpc/default/lammps:latest
singularity build --sandbox lammps docker://lammps/lammps
singularity build test.sif /my-local-image.simg
singularity build lammps.sif lammps/
singularity build demo.sif demo.def
```





``` shell
Bootstrap: library
From: ubuntu:18.04
Stage: build
 
%setup
    touch /file1
    touch ${SINGULARITY_ROOTFS}/file2
 
%files
    /file1
    /file1 /opt
 
%environment
    export LISTEN_PORT=12345
    export LC_ALL=C
 
%post
    apt-get update && apt-get install -y netcat
    NOW=`date`
    echo "export NOW=\"${NOW}\"" >> $SINGULARITY_ENVIRONMENT
 
%runscript
    echo "Container was created $NOW"
    echo "Arguments received: $*"
    exec echo "$@"
 
%startscript
    nc -lp $LISTEN_PORT
 
%test
    grep -q NAME=\"Ubuntu\" /etc/os-release
    if [ $? -eq 0 ]; then
        echo "Container base is Ubuntu as expected."
    else
        echo "Container base is not Ubuntu."
        exit 1
    fi
 
%labels
    Author d@sylabs.io
    Version v0.0.1
 
%help
    This is a demo container used to illustrate a def file that uses all
    supported sections.
```





### Example



``` shell
singularity build --sandbox lammps docker://lammps/lammps
singularity shell -w lammps
// do something
exit
singularity build lammps.sif lammps/
```





## Run

``` shell
singularity run shub://GodloveD/lolcow
singularity run docker://library/python
singularity run ./my-local-image.simg
```



## Exec

``` shell

singularity exec path_to_container command_goes_here so_do_parameters
singularity exec GodloveD-lolcow-master-latest.simg date
```

## Shell

``` shell
singularity shell GodloveD-lolcow-master-latest.simg
Singularity GodloveD-lolcow-master-latest.simg:~> date
Singularity GodloveD-lolcow-master-latest.simg:~> df
```



## Quick Start

- Run container from Docker Hub: singularity run [docker://publishing_user/container_name](docker://publishing_user/container_name) 
- Running container from Singularity Hub: singularity run [shub://publishing_user/container_name](shub://publishing_user/container_name) 
- Running a local container image: singularity run path_to_container 
- Passing a command to a container: singularity exec path_to_container command_goes_here so_do_parameters 
- Mounting a directory in a container: singularity run -B path_on_host:path_in_container [shub://publishing_user/container_name](shub://publishing_user/container_name)



## Example

### TensorFlow using Singularity

``` shell
export sdir=/home/support/public/tutorials/singularity
cp $sdir/mnist.py .
singularity pull shub://schanzel/singularity-tensorflow-keras-py3:latest
singularity exec ./schanzel-singularity-tensorflow-keras-py3-master-latest.simg python mnist.py
```







### PyTorch

``` shell
singularity build pytorch.sif docker://pytorch/pytorch:latest
```





ReadOnly file system



https://crc-docs.abudhabi.nyu.edu/hpc/software/singularity_conda.html





https://sylabs.io/guides/3.2/user-guide/definition_files.html?highlight=best%20practice#best-practices-for-build-recipes



https://askubuntu.com/questions/287021/how-to-fix-read-only-file-system-error-when-i-run-something-as-sudo-and-try-to