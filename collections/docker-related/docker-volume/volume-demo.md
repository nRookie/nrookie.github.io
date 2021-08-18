# Demonstrating Volumes with Containers and Services

Let's see how to use volumes with containers and services.

The examples will be from a system with no pre-existing volumes, and everything we demonstrate applies to both Linux and Windows.


## Mount volumes #

Use the following command to create a new standalone container that mounts a volume called bizvol.

Linux example:
``` shell
docker container run -dit --name voltainer --mount source=bizvol,target=/vol alpine
```





The command uses the --mount flag to mount a volume called “bizvol” into the container at either /vol or c:\vol. The command completes successfully despite the fact there is no volume on the system called bizvol. This raises some interesting points:

- If you specify an existing volume, Docker will use the existing volume.
- If you specify a volume that doesn’t exist, Docker will create it for you.

In this case, bizvol didn’t exist, so Docker created it and mounted it into the new container. This means you’ll be able to see it with docker volume ls.

``` shell
ubuntu@10-13-63-31:~$ sudo docker container run -dit --name voltainer --mount source=bizvol,target=/vol alpine
70a4ed5153a69cd9ebdd958ceed5eaa160d09071f73bc56c5ed859880e47bacd
ubuntu@10-13-63-31:~$ docker volume ls
Got permission denied while trying to connect to the Docker daemon socket at unix:///var/run/docker.sock: Get "http://%2Fvar%2Frun%2Fdocker.sock/v1.24/volumes": dial unix /var/run/docker.sock: connect: permission denied
ubuntu@10-13-63-31:~$ sudo docker volume ls
DRIVER    VOLUME NAME
local     bizvol

```

## Containers and volume lifecycles #

Although containers and volumes have separate lifecycles, you cannot delete a volume that is in use by a container. Try it.

``` shell
docker volume rm bizvol
```
The volume is brand new, so it doesn’t have any data. Let’s exec onto the container and write some data to it. The example cited is Linux, if you’re following along on Windows just replace sh with pwsh.exe at the end of the command. All other commands will work on Linux and Windows.

## Check volume mount point #

Because the volume still exists, you can look at its mount point on the host to check if the data is still there.

Run the following commands from the terminal of your Docker host. The first one will show that the file still exists, the second will show the contents of the file.

Be sure to use the C:\ProgramData\Docker\volumes\bizvol\_data directory if you’re following along on Windows.

``` shell
$ ls -l /var/lib/docker/volumes/bizvol/_data/
```


## Mounting volumes to a new service or container #

It’s even possible to mount the bizvol volume into a new service or container. The following command creates a new Docker service, called hellcat, and mounts bizvol into the service replica at /vol. You’ll need to be running in swarm mode for this command to work. If you’re running in single-engine mode you can use a docker container run command instead.




``` shell
docker service create \
  --name hellcat \
  --mount source=bizvol,target=/vol \
  alpine sleep 1d
```



### Replicas

We didn’t specify the --replicas flag, so only a single service replica was deployed. Find which node in the Swarm it’s running on.

``` shell
ubuntu@10-13-63-31:~$ sudo docker service ps hellcat
ID             NAME        IMAGE           NODE          DESIRED STATE   CURRENT STATE                ERROR     PORTS
7v6tot7j4ggg   hellcat.1   alpine:latest   10-13-63-31   Running         Running about a minute ago   

```



