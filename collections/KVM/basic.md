## 概念

虚拟化， 即为在一物理机上， 同时运行多个独立的操作系统。由 hypervisor（虚拟机管理程序） 创建一抽象层以完成对硬件的控制与分离， 并为 guest（客户机）提供对硬件的访问途径。 KVM （Kernel-based Virtual Machine）是一种Linux完全虚拟化的解决方案， 并借由处理器硬件的特性， 为客户机提供底层的实体抽象， 以至于客户及并不知自己身处虚拟化的环境中。

![7a39b80d59f5865e4120dc990922ccf6.png](https://img-blog.csdnimg.cn/20190915075150821.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L1NhaW50eXl1,size_16,color_FFFFFF,t_70)

裸金属架构 = BareMetal = Type-I

宿主型架构 = Host = Type-II

## 准备

使用的是物理云主机。
使用的操作系统是 Centos7

## 安装 KVM

### Centos 7 安装 KVM

1. 查看系统版本

```shell
cat /etc/centos-release
CentOS Linux release 7.6.1810 (Core)
```

1. 查看CPU是否支持虚拟化

```shell
cat /proc/cpuinfo | egrep 'vmx|svm'
lscpu | grep Virtualization
```

1. 查看是否加载KVM

```shell
lsmod | grep kvm
kvm_intel             299008  0
kvm                   851968  1 kvm_intel
irqbypass              16384  1 kvm
```

> **lsmod** is a trivial program which nicely formats the contents of the /proc/modules, showing what kernel modules are currently loaded.

如上显示的结果说明已经加载， 如果没有加载请执行如下命令：

```shell
modprobe kvm
```

> **modprobe** intelligently adds or removes a module from the Linux kernel:

1. 关闭selinux

```shell
setenforce 0
vim /etc/sysconfig/selinux
SELINUX=disabled
```

> SELinux - NSA Security-Enhanced Linux

1. 安装KVM相关软件包

```shell
yum install qemu-kvm qemu-img \
virt-manager libvirt libvirt-python virt-manager \
libvirt-client virt-install virt-viewer -y
```

1. 启动libvirt并设置开机自启动

```shell
systemctl start libvirtd
systemctl enable libvirtd
```

1. 查看机器的存储

```shell
df -hT
```

1. 创建物理桥接设备 (桥接后用virt-install遇到了很多问题，后面使用了NAT， 8 这一步骤应该跳过）

> 默认的情况下, 基于 dhcpd的网桥已经被libvirtd配置了. 这种情况下， 所有的 VMs（guest machine） 只能和在同一个宿主上的其他VM通信，如果要使他们被同一个LAN内的其他服务器访问，需要给他们配置一个网桥。

查看网卡信息

```shell
ifconfig
```

关闭NetworkManager服务

```shell
chkconfig NetworkManager off
service NetworkManager stop
```

桥接设备关联网卡

> 不清楚生成桥接失败是什么意思，

```shell
[root@ ~]# virsh iface-bridge net1 br0
使用附加设备 br0 生成桥接 net1 失败
已启动桥接接口 br0
```

命令说明

```shell
iface-bridge interface bridge [--no-stp] [delay] [--no-start]

Create a bridge device named bridge, and attach the existing network device interface to the new bridge. The new bridge defaults to starting immediately, with STP enabled and a delay of 0; these settings can be altered with --no-stp, --no-start, and an integer number of seconds for delay. All IP address configuration of interface will be moved to the new bridge device.

iface-unbridge for undoing this operation
```

查看是否成功

```shell
[root@ ~]# brctl show
bridge name	bridge id		STP enabled	interfaces
br-426c8cc14f36		8000.02421c9f1d37	no		veth5c49b7f
docker0		8000.0242b1675603	no		
virbr0		8000.52540049967e	yes		virbr0-nic
```

1. 下载Centos 镜像到物理云主机的/root/kvm-learn/iso/目录下

```shell
[root@10-9-19-59 iso]# wget https://mirrors.tuna.tsinghua.edu.cn/centos/7.9.2009/isos/x86_64/CentOS-7-x86_64-Minimal-2009.iso
```

等待下载完成。

1. 下载 sha256sum.txt 并查看是否成功下载iso 文件

```shell
wget sha256sum.txt
[root@iso]# ls
CentOS-7-x86_64-Minimal-2009.iso  index.html  sha256sum.txt
[root@iso]# sha256sum -c sha256sum.txt 
sha256sum: CentOS-7-x86_64-Everything-2009.iso: 没有那个文件或目录
CentOS-7-x86_64-Everything-2009.iso：打开或读取失败
sha256sum: CentOS-7-x86_64-NetInstall-2009.iso: 没有那个文件或目录
CentOS-7-x86_64-NetInstall-2009.iso：打开或读取失败
CentOS-7-x86_64-Minimal-2009.iso: 确定
sha256sum: CentOS-7-x86_64-DVD-2009.iso: 没有那个文件或目录
CentOS-7-x86_64-DVD-2009.iso：打开或读取失败
sha256sum: 警告：3 个列出的文件无法读取
```

可以看到 CentOS-7-x86_64-Minimal-2009.iso: 确定

至此准备阶段的内容已完成。

## 创建虚拟机

### 虚拟机磁盘 （disk）

创建一块qcow2格式的虚拟磁盘， 容量 20GB， 用做虚拟机的系统盘

```shell
qemu-img create -f qcow2 centos7.demo.img 20G
```

### 虚拟机网络（network)

宿主机上创建一个虚拟网络，配置为NAT模式（桥接模式遇到了问题）。

```shell
<network>
	<name>net.demo</name>
	<bridge name="virbr1"/>
	<forward mode="nat"/>
	<ip address="192.168.123.1" netmask="255.255.255.0">
		<dhcp>
		<range start="192.168.123.100" end="192.168.123.199"/>
		<host mac='00:00:00:00:00:11' name='vm1' ip='192.168.123.11' />
		<host mac='00:00:00:00:00:12' name='vm2' ip='192.168.123.12' />
		<host mac='00:00:00:00:00:13' name='vm3' ip='192.168.123.13' />
		</dhcp>
	</ip>
</network>
```

定义网络

```shell
virsh net-define net.demo.xml
```

移除网络

```shell
virsh net-undefine net.demo
```

启动网络

```shell
virsh net-start net.demo
```

停止网络

```shell
virsh net-destroy net.demo
```

这样

在宿主机上创建了一个虚拟内部网络：192.168.123.0/24
宿主机上会自动创建一个新的网卡，并绑定地址 192.168.123.1
这个网络内部由libvirt支持的隐式dhcp服务，地址范围：192.168.123.100-199，dns=宿主机（即192.168.123.1）
（可选）示意了简单的 dhcp 静态地址绑定功能，创建虚拟机时，指定了mac地址后，即可与静态地址绑定

注意：本文所述的“虚拟网络”，仅指代此处通过libvirt定义的虚拟的网络，即 net.demo

### 虚拟机定义 (domain) （通过定义xml配置文件）

定义配置如下的文件为 vm.demo1.xml

```xml
<domain type='kvm'>
	<name>vm.demo1</name>
	<memory unit='GiB'>8</memory>
	<vcpu placement='static'>4</vcpu>
	<os>
	<type arch='x86_64' machine='pc'>hvm</type>
	<boot dev='hd'></boot>
	<boot dev='cdrom'></boot>
	<bootmenu enable='yes'></bootmenu>
	</os>
	<features>
	<acpi></acpi>
	<apic></apic>
	<pae></pae>
	</features>
	<cpu mode='host-passthrough'>
	</cpu>
	<clock offset='utc'>
	<timer name='pit' tickpolicy='delay'></timer>
	</clock>
	<on_poweroff>destroy</on_poweroff>
	<on_reboot>restart</on_reboot>
	<on_crash>restart</on_crash>
	<devices>
		<serial type='pty'>
			<target port='0'></target>
		</serial>
		<console type='pty'>
			<target type='serial' port='0'></target>
		</console>
		<input type='tablet' bus='usb'/>
		<input type='mouse' bus='ps2' />
		<graphics type='vnc' port='-1' autoport='yes'></graphics>
		<video>
			<model type='cirrus' vram='9216' heads='1'></model>
		</video>
		<memballoon model='virtio'>
		</memballoon>
		<disk type='file' device='disk'>
			<driver name='qemu' type='qcow2' cache='writethrough' io='native'/>
			<source file='/root/kvm-learn/images/vm.centos7.img'/>
			<target dev='vda' bus='virtio'/>
			<serial>UCLOUD_DISK_VDA</serial>
		</disk>
		<disk type='file' device='cdrom'>
			<source file='root/kvm-learn/CentOS-7-x86_64-Minimal-2009.iso'/>
			<target dev='hdd' bus='ide'/>
			<readonly/>
		</disk>
		<interface type="network">
			<source network="net.demo"/>
			<mac address='00:00:00:00:00:11'/>
		</interface>
	</devices>
</domain>
```

定义domain

```shell
virsh define vm.demo1.xml
```

移除domain

```shell
virsh undefine vm.demo1.xml
```

启动domain

```shell
virsh start vm.demo1
```

停止domain

```shell
virsh destroy vm.demo1
```

### 遇到的问题

当配置文件中

```xml
    <disk type='file' device='disk'>
      <driver name='qemu' type='qcow2' cache='writethrough' io='native'/>
      <source file='/root/kvm-learn/images/vm.centos7.img'/>
      <target dev='vda' bus='virtio'/>
      <serial>UCLOUD_DISK_VDA</serial>
    </disk>
```

会报错

```shell
virsh start --domain vm.demo2 
错误：开始域 vm.demo2 失败
错误：unsupported configuration: native I/O needs either no disk cache or directsync cache mode, QEMU will fallback to aio=threads
```

查看了 virt-installer 的manual 并 将相应的cache 更改为 directsync 就可以启动了

```xml
    <disk type='file' device='disk'>
      <driver name='qemu' type='qcow2' cache='directsync' io='native'/>
      <source file='/root/kvm-learn/images/vm.centos7.img'/>
      <target dev='vda' bus='virtio'/>
      <serial>UCLOUD_DISK_VDA</serial>
    </disk>
```

1. 创建虚拟机，命名 vm.demo1
2. CPU x 4，RAM x 8GB，开机启动顺序：优先光驱，其次硬盘
3. 启用VNC控制台，且自动绑定一个端口（默认tcp 5900起步，根据需求可以配置静态端口）
4. 挂载一块虚拟磁盘，虚拟磁盘文件路径为：/root/kvm-learn/vm.demo1.img，该磁盘由 qemu-img create 指令创建
   虚拟光驱中挂载镜像文件：/root/kvm-learn/CentOS-7-x86_64-Minimal-2009.iso
   挂载虚拟网卡并关联虚拟网络 net.demo，同时为此网卡指定mac地址为 00:00:00:00:00:11（启动后将通过dhcp获得预配置的IP地址 192.168.123.11/24）

### 使用 virt-install 创建

### virt-install 使用手册

```shell
virt-install is a command line tool for creating new KVM, Xen, or Linux
container guests using the "libvirt" hypervisor management library.  See the
EXAMPLES section at the end of this document to quickly get started.

virt-install tool supports graphical installations using (for example) VNC
or SPICE, as well as text mode installs over serial console. The guest can
be configured to use one or more virtual disks, network interfaces, audio
devices, physical USB or PCI devices, among others.
```

1. 

```shell
virt-install \
--virt-type=kvm \
--name centos7 \
--memory 2048 \
--vcpus=1 \
--os-variant=centos7.0 \
--cdrom=/root/kvm-learn/iso/CentOS-7-x86_64-Minimal-2009.iso \
--network=bridge=br0,model=virtio \
--graphics vnc \
--disk path=/var/lib/libvirt/images/centos7.qcow2,size=40,bus=virtio,format=qcow2

# options
-- virt-type
The hypervisor to install on. Example choices are kvm, qemu, or xen.
Available options are listed via 'virsh capabilities' in the <domain> tags.

This deprecates the --accelerate option, which is now the default behavior. To install a plain QEMU guest, use '--virt-type qemu'

/

Name of the new guest virtual machine instance. This must be unique amongst
all guests known to the hypervisor on the connection, including those not
currently active. To re-define an existing guest, use the virsh(1) tool to
shut it down ('virsh shutdown') & delete ('virsh undefine') it prior to
running "virt-install".



--memory OPTIONS

Memory to allocate for the guest, in MiB. This deprecates the -r/--ram option.
Sub options are available, like 'maxmemory', 'hugepages', 'hotplugmemorymax'
and 'hotplugmemoryslots'.  The memory parameter is mapped to <currentMemory>
element, the 'maxmemory' sub-option is mapped to <memory> element and
'hotplugmemorymax' and 'hotplugmemoryslots' are mapped to <maxMemory> element.

To configure memory modules which can be hotunplugged see --memdev
description.

Use --memory=? to see a list of all available sub options. Complete details at
<http://libvirt.org/formatdomain.html#elementsMemoryAllocation>


--vcpus OPTIONS

Number of virtual cpus to configure for the guest. If 'maxvcpus' is specified,
the guest will be able to hotplug up to MAX vcpus while the guest is running,
but will startup with VCPUS.

CPU topology can additionally be specified with sockets, cores, and threads.
If values are omitted, the rest will be autofilled preferring sockets over
cores over threads.

'cpuset' sets which physical cpus the guest can use. "CPUSET" is a comma
separated list of numbers, which can also be specified in ranges or cpus to
exclude. Example:

0,2,3,5     : Use processors 0,2,3 and 5
1-5,^3,8    : Use processors 1,2,4,5 and 8

If the value 'auto' is passed, virt-install attempts to automatically
determine an optimal cpu pinning using NUMA data, if available.

Use --vcpus=? to see a list of all available sub options. Complete details at
<http://libvirt.org/formatdomain.html#elementsCPUAllocation>

--os-variant OS_VARIANT

Optimize the guest configuration for a specific operating system (ex.
'fedora18', 'rhel7', 'winxp'). While not required, specifying this options is
HIGHLY RECOMMENDED, as it can greatly increase performance by specifying
virtio among other guest tweaks.

By default, virt-install will attempt to auto detect this value from the
install media (currently only supported for URL installs). Autodetection can
be disabled with the special value 'none'. Autodetection can be forced with
the special value 'auto'.

Use the command "osinfo-query os" to get the list of the accepted OS variants.

--cdrom OPTIONS


File or device used as a virtual CD-ROM device.  It can be path to an ISO
image or a URL from which to fetch/access a minimal boot ISO image. The URLs
take the same format as described for the "--location" argument. If a cdrom
has been specified via the "--disk" option, and neither "--cdrom" nor any
other install option is specified, the "--disk" cdrom is used as the install
media.


--network OPTIONS

Connect the guest to the host network. The value for "NETWORK" can take one of
4 formats:

bridge=BRIDGE
    Connect to a bridge device in the host called "BRIDGE". Use this option if
    the host has static networking config & the guest requires full outbound
    and inbound connectivity  to/from the LAN. Also use this if live migration
    will be used with this guest.

network=NAME
    Connect to a virtual network in the host called "NAME". Virtual networks
    can be listed, created, deleted using the "virsh" command line tool. In an
    unmodified install of "libvirt" there is usually a virtual network with a
    name of "default". Use a virtual network if the host has dynamic
    networking (eg NetworkManager), or using wireless. The guest will be NATed
    to the LAN by whichever connection is active.

Other available options are:

model
    Network device model as seen by the guest. Value can be any nic model
    supported by the hypervisor, e.g.: 'e1000', 'rtl8139', 'virtio', ...

mac 
    Fixed MAC address for the guest; If this parameter is omitted, or the
    value "RANDOM" is specified a suitable address will be randomly generated.
    For Xen virtual machines it is required that the first 3 pairs in the MAC
    address be the sequence '00:16:3e', while for QEMU or KVM virtual machines
    it must be '52:54:00'.


--graphics TYPE,opt1=arg1,opt2=arg2,...
   Specifies the graphical display configuration. This does not configure any
   virtual hardware, just how the guest's graphical display can be accessed.
   Typically the user does not need to specify this option, virt-install will try
   and choose a useful default, and launch a suitable connection.
  


--disk OPTIONS
Specifies media to use as storage for the guest, with various options. The
general format of a disk string is

   --disk opt1=val1,opt2=val2,...

The simplest invocation to create a new 10G disk image and associated disk
device:

   --disk size=10

virt-install will generate a path name, and place it in the default image
location for the hypervisor. To specify media, the command can either be:

   --disk /some/storage/path[,opt1=val1]...

or explicitly specify one of the following arguments:

path
   A path to some storage media to use, existing or not. Existing media can
   be a file or block device.

   Specifying a non-existent path implies attempting to create the new
   storage, and will require specifying a 'size' value. Even for remote
   hosts, virt-install will try to use libvirt storage APIs to automatically
   create the given path.
 

Other available options:

size
   size (in GiB) to use if creating new storage

bus 
    Disk bus type. Value can be 'ide', 'sata', 'scsi', 'usb', 'virtio' or
    'xen'.  The default is hypervisor dependent since not all hypervisors
    support all bus types.

format
    Disk image format. For file volumes, this can be 'raw', 'qcow2', 'vmdk',
    etc. See format types in <http://libvirt.org/storage.html> for possible
    values. This is often mapped to the driver_type value as well.

    If not specified when creating file images, this will default to 'qcow2'.

    If creating storage, this will be the format of the new image. If using an
    existing image, this overrides libvirt's format auto-detection.
```

## 步骤

1. 以下命令使用 virt-install 建立一个 2GB 内存，2核， 1个网卡 ， 40GB 大小的 Centos7 虚拟机

```shell
virt-install \
--virt-type=kvm \
--name centos7 \
--memory 2048 \
--vcpus=1 \
--os-variant=centos7.0 \
--cdrom=/home/iso/CentOS-7-x86_64-Minimal-2009.iso \
--network=bridge=br0,model=virtio \
--graphics vnc \
--disk path=/var/lib/libvirt/images/centos7.qcow2,size=40,bus=virtio,format=qcow2
```

> **遇到了错误，还没有确定原因**

```shell
WARNING  需要图形显示，但未设置 DISPLAY。不能运行 virt-viewer。
WARNING  没有控制台用于启动客户机，默认为 --wait -1

开始安装......
正在分配 'centos7.qcow2'                                         |  40 GB  00:00:00     
ERROR    unsupported format character '?' (0xffffffe7) at index 47
域安装失败，您可以运行下列命令重启您的域：
'virsh start virsh --connect qemu:///system start centos7-demo'
否则请重新开始安装。
```

> 再次更新，virt-install 运行的shell 会block住， 这时 可以用vnc连接到虚拟机里面安装操作系统，安装完毕后退出就可以完成安装。

## 连接虚拟机

### 使用vnc连接到虚拟机

1. 使用vnc的自动端口绑定功能 （通常都是从5900 线性递增， 如果只有一台虚拟机，vnc端口基本就是5900）

```shell
    ps ax | grep vm.demo2   # 找出虚拟机对应qemu-kvm 进程id
    ss -nlpt | grep 进程id， # 结合进程id 得到 vncserver 的绑定端口
```

> 如果是公有云，需要开启外网防火墙

使用macos 自带的屏幕共享（打开访达，点击键盘 commnand + k）会停留在 连接界面
使用vnc viewer 可以连接到虚拟机，可以操作键盘，但是不可以操作鼠标（触控板）

> 引申的问题，如何给虚拟机的vnc server 设置密码？ 请参考附录

> 不可以操作鼠标的问题没有解决，时间原因先跳过了
> 不确定以下几个点
> 1.kvm里面的 vnc server 权限设置的问题吗？
> 2.是 macos big sur 鼠标权限设置的问题吗？
> 3.是不是 vnc server 那边无法识别到 macbook 的trackpad， 从而认为没有接入鼠标？ 需要拿一个鼠标试一试
>
> [参考解决方案](https://help.realvnc.com/hc/en-us/articles/360002712837-Known-Issues-when-connecting-to-macOS-Mojave-Catalina-and-Big-Sur-10-14-10-15-and-11-0-#screen-recording-fixes-black-screen-catalina-big-sur-only--0-6)
> 实际使用，装了vnc viewer以后 根本就没有vncagent这个进程, 所以没起作用

> 下午又重新装了一个虚拟机vm.demo3 可以使用鼠标了。 没找到root cause

> 装了多次虚拟机，又碰到了鼠标不管用的问题，目前认为是在vnc viewer连接的过程中切换到其他应用程序的原因

### 如何ssh到虚拟机内部

> 因为虚拟机宿主机通过虚拟机网络 net.demo 关联在了同一个网络下，因此，在宿主机上直接ssh虚拟机的IP地址即可
> 如果在虚拟网络中配置了静态IP地址绑定，那么虚拟机的IP地址也就是可以预期的了
> 否则，通过控制台的方式拿到虚拟机的IP地址再ssh登

1. 第一次安装虚拟机centos7， 无法ping 通虚拟机，无法启用ens3接口, 也没有查到租约信息

```shell
 virsh net-dhcp-leases net.demo
```

通过vnc连接到虚拟机在命令行输入打开网口命令

```shell
ifup ens3
Error: Connection activation failed: IP configuration could not be reserved (no available address, timeout, etc.)
```

1. 安装一个新的虚拟机centos7，尝试在安装界面打开以太网口，可以成功获取到IP地址，并可以ping 通

```shell
 virsh net-dhcp-leases net.demo
Expiry Time          MAC address        Protocol  IP address                Hostname        Client ID or DUID
-------------------------------------------------------------------------------------------------------------------
 2021-06-01 18:25:30  00:00:00:00:00:13  ipv4      192.168.123.119/24 
 
[root@ images]# ping 192.168.123.119
PING 192.168.123.119 (192.168.123.119) 56(84) bytes of data.
64 bytes from 192.168.123.119: icmp_seq=1 ttl=64 time=0.331 ms
64 bytes from 192.168.123.119: icmp_seq=2 ttl=64 time=0.242 ms
64 bytes from 192.168.123.119: icmp_seq=3 ttl=64 time=0.208 ms
64 bytes from 192.168.123.119: icmp_seq=4 ttl=64 time=0.224 ms
```

> 

安装完成以后现象和1一样，无法ping 通 无法打开网口。

#### 解决不能ssh到虚拟机的问题(NAT 模式）

1. 更改对应guest机器的配置

```shell
virsh edit vm.demo
```

1. 在编辑器里面查找 interface， 找到相应的type=network ，

修改前

```xml
   <interface type='network'>
      <mac address='00:00:00:00:00:11'/>
      <source network='net.demo'/>
      <model type='rtl8139'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x03' function='0x0'/>
  </interface>
```

修改后

```xml
<interface type='network'>
      <mac address='52:54:00:73:f3:38'/>
      <source network='net.demo'/>
      <model type='virtio'/>
      <driver name='vhost' txmode='iothread' ioeventfd='on' event_idx='off'/>
  </interface>
```

1. 重启vm.demo 虚拟机

```xml
    virsh shutdown vm.demo
    virsh start vm.demo
```

1. 查看virt dhcp server 分配的IP地址

```shell
[root@10-9-19-59 ~]# virsh net-dhcp-leases net.demo
setlocale: No such file or directory
 Expiry Time          MAC address        Protocol  IP address                Hostname        Client ID or DUID
-------------------------------------------------------------------------------------------------------------------
 2021-06-02 13:34:24  00:00:00:00:00:13  ipv4      192.168.123.13/24         vm3             -
 2021-06-02 13:42:06  52:54:00:73:f3:38  ipv4      192.168.123.122/24        vm2             -
```

可以ping 通

也可以ssh到虚拟机里面

```shell
[root@10-9-19-59 ~]# ping -c 2 192.168.123.13
PING 192.168.123.13 (192.168.123.13) 56(84) bytes of data.
64 bytes from 192.168.123.13: icmp_seq=1 ttl=64 time=0.218 ms
64 bytes from 192.168.123.13: icmp_seq=2 ttl=64 time=0.132 ms

--- 192.168.123.13 ping statistics ---
2 packets transmitted, 2 received, 0% packet loss, time 1040ms
rtt min/avg/max/mdev = 0.132/0.175/0.218/0.043 ms
[root@10-9-19-59 ~]# ssh root@192.168.123.13
The authenticity of host '192.168.123.13 (192.168.123.13)' can't be established.
ECDSA key fingerprint is SHA256:QyDOe8BOz2ZPZcbRcvqKIOy/Wre2bFDKjnhc+sVZ6Rg.
ECDSA key fingerprint is MD5:0e:b1:c9:f0:9b:7d:3a:9e:29:2b:12:34:f8:f0:87:9f.
Are you sure you want to continue connecting (yes/no)? yes
Warning: Permanently added '192.168.123.13' (ECDSA) to the list of known hosts.
root@192.168.123.13's password: 
Last login: Tue Jun  1 23:47:30 2021
[root@vm3 ~]# 
```

二重ssh

因为虚拟机宿主机通过虚拟网络 net.demo关联在了同一个网络下，因此宿主机上直接ssh虚拟机的IP地址即可

## 附录

### virsh 常用的操作

1. 使用virt-install 安装虚拟机失败，再次安装显示name 已经被占用， 如何查看已经安装的虚拟机？

```shell
virsh list
virsh list --all
[root iso]# virsh list
 Id    名称                         状态
----------------------------------------------------
 1     centos7                        running
```

1. 查看某个guest的对应的存储文件

```shell
[root@iso]# virsh dumpxml --domain centos7 | grep 'source file'
      <source file='/var/lib/libvirt/images/centos7.qcow2'/>
      <source file='/home/iso/CentOS-7-x86_64-Minimal-2009.iso'/>
```

1. 关闭客户机

```shell
virsh shutdown VM_NAME
virsh shutdown --domain VM_NAME
```

也可以使用 destory 来强制关闭

```shell
# virsh destroy VM_NAME
# virsh destroy --domain VM_NAME
```

4 删除客户机

```shell
# virsh undefine VM_NAME
# virsh undefine --domain VM_NAME
```

加上选项 --remove-all-storage 可以删除对应的卷

```shell
[root@ iso]# virsh undefine centos7 --remove-all-storage
域 centos7 已经被取消定义
已删除卷 'vda'(/var/lib/libvirt/images/centos7.qcow2)。
```

1. 为某个客户机的vnc server 设置密码

边界客户机域的xml文件

```shell
virsh edit vm.demo2
```

找到 graphics type = vnc 这一个选项, 增加一个属性passwd 对应的是连接要用到的密码

```xml
<graphics type='vnc' port='-1' autoport='yes' passwd='vnc'>
      <listen type='address'/>
    </graphics>
```

> 这里报了错误， 不确定是不是由于我没有关闭就更改域的文件

```shell
错误：XML document failed to validate against schema: Unable to validate doc against /usr/share/libvirt/schemas/domain.rng
Extra element devices in interleave
Element domain failed to validate content

失败的。 Try again? [y,n,i,f,?]: 
错误：XML document failed to validate against schema: Unable to validate doc against /usr/share/libvirt/schemas/domain.rng
Extra element devices in interleave
Element domain failed to validate content
失败的。 Try again? [y,n,i,f,?]: 
```

键盘输入 i 忽略校验错误。

```
编辑了域 vm.demo2 XML 配置。
```

重启 vm.demo2

```
virsh shutdown vm.demo2
virsh start vm.demo2
```

至此可以使用密码连接到vnc了

### 关于虚拟网络（libvirt的） 和 虚拟机的网卡地址的关系

1. 域配置文件vm.demo.xml 里面的 mac地址就是虚拟机里面默认网卡的mac地址

```xml
   <interface type='network'>
      <mac address='00:00:00:00:00:12'/>
      <source network='net.demo'/>
      <model type='rtl8139'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x03' function='0x0'/>
```

## 参考连接

[kvm 虚拟机demo](https://ushare.ucloudadmin.com/pages/viewpage.action?pageId=7770354)

[kvm 虚拟机 学习](https://ushare.ucloudadmin.com/pages/viewpage.action?pageId=70676402)

[How to Install kvm on centos 7](https://www.cyberciti.biz/faq/how-to-install-kvm-on-centos-7-rhel-7-headless-server/)

[REMOVING AND DELETING A VIRTUAL MACHINE](https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/7/html/virtualization_deployment_and_administration_guide/sect-virsh-delete)

[set virtio drivers for nic device](https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/7/html/virtualization_deployment_and_administration_guide/sect-kvm_para_virtualized_virtio_drivers-using_kvm_virtio_drivers_for_nic_devices)https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/7/html/virtualization_deployment_and_administration_guide/sect-kvm_para_virtualized_virtio_drivers-using_kvm_virtio_drivers_for_nic_devices)