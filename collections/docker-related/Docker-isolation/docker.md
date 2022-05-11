# 文件系统



Union FS

- 将不同目录挂载到同一个虚拟文件系统下（ unite several directories into a single virtual filesystem) 的文件系统
- 支持为每一个成员目录（类似Git Branch） 设定readonly、readwrite 和 whiteout-able 权限。
- 文件系统分层， 对readonly权限的branch 可以逻辑上进行修改（增量地， 不影响readonly 部分的）。
- 通常 Union FS 有两个用途， 一方面可以将多个Disk挂到同一个目录下， 另一个更常用的就是将一个readonly的branch和 一个writeable 的branch 联合在一起。

### Docker的文件系统



典型的Linux 文件系统组成：

- Bootfs (boot file system)

  - BootLoader - 引导加载kernel
  - Kernel - 当 kernel 被加载到内存中后umount bootfs

- rootfs （root file system）

  - /dev， /proc， /bin， /etc 等标准目录和文件。
  - 对于不同的linux发行版，bootfs 基本一致的，但rootfs会有差别。

  



# Docker 启动





Linux

- 在启动后， 首先将rootfs 设置为readonly, 进行一系列检查，然后将其切换为 “readwrite” 后供用户使用。



Docker 启动

- 初始化时将rootfs 以readonly 方式加载并检查， 然而接下来利用 union mount的方式将一个readwrite文件系统挂载在readonly的rootfs之上；
- 并且允许再次将下层的FS（fileSystem） 设定为 readonly 并且向上叠加；
- 这样一组readonly 和 一个writable 的结构构成一个container的运行时态， 每一个FS被称作一个FS层。