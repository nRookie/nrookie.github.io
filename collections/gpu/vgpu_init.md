

### check if vgpu is installed

``` shell
ulowrisk:~$ lsmod | grep 'nvidia_vgpu_vfio'
nvidia_vgpu_vfio       57344  0 
nvidia              34041856  10 nvidia_vgpu_vfio
mdev                   20480  2 vfio_mdev,nvidia_vgpu_vfio
vfio                   32768  5 vfio_mdev,nvidia_vgpu_vfio,vfio_iommu_type1,vfio_pci
```



### check gpu type

``` shell
lspci | grep NVIDIA
82:00.0 3D controller: NVIDIA Corporation TU104GL [Tesla T4] (rev a1)
```



### check vgpu type



``` shell
 cd /sys/class/mdev_bus/0000:82:00.0/mdev_supported_types/ (0000:82:00.0 by gpu pci)
 
 /sys/devices/pci0000:80/0000:80:01.2/0000:82:00.0/mdev_supported_types$ grep -l "T4" nvidia-*/name
nvidia-222/name
nvidia-223/name
nvidia-224/name
nvidia-225/name
nvidia-226/name
nvidia-227/name
nvidia-228/name
nvidia-229/name
nvidia-230/name
nvidia-231/name
nvidia-232/name
nvidia-233/name
nvidia-234/name
nvidia-252/name
nvidia-319/name
nvidia-320/name
nvidia-321/name
```





### Confirm the vgpu type to be initialized, number of instance to intialized.



``` shell
/sys/devices/pci0000:80/0000:80:01.2/0000:82:00.0/mdev_supported_types$ cat nvidia-222/name 
GRID T4-1B
/sys/devices/pci0000:80/0000:80:01.2/0000:82:00.0/mdev_supported_types$ cat nvidia-222/available_instances 
0
```



VGPU type , for example, GRID T4-1B (1B is 1GB gpu memory)

available instance is the maximum number of vgpu a physical GPU can   initialized. in here it is 0 because the host machine has already initialized.a physical gpu card can only initialize same type of vgpu, if there are more than one physical gpu card on the host, then the host can be intialized with different types of gpu.





``` shell
### initialize 
echo "87d76e10-0c90-49dd-981f-8cd03a1c2865" > nvidia-222/create

###  check if initialization is succeed
ls -l /sys/bus/mdev/devices

### deleteing vgpu
echo 1 >/sys/bus/mdev/devices/87d76e10-0c90-49dd-981f-8cd03a1c2865/remove

```



 

## startup vgpu



``` shell
#!/bin/bash

PCI_LIST=()
dev="$1"

Prefix_1="00688158"
Prefix_="-"
gen=0
##uuidgen模版: 00688158-0001-0004-0001-010066144074
##ucloud标志
##预留4位数 物理显卡的数量 count
##预留4位数 nvidia-* 可初始化的数量 ava_count
##预留4位数 vgpu的序号0001,0002,...
##预留12位  宿主机ip

ip2int() {

    ip=$1
    if [[ "x$ip" == "x" ]];then
        echo "hostip get failed"
        exit 1;
    fi
    a=$(echo $ip | awk -F'.' '{print $1}')
    b=$(echo $ip | awk -F'.' '{print $2}')
    c=$(echo $ip | awk -F'.' '{print $3}')
    d=$(echo $ip | awk -F'.' '{print $4}')

    a=$(seq -f "%03g" $a | tail -1)
    b=$(seq -f "%03g" $b | tail -1)
    c=$(seq -f "%03g" $c | tail -1)
    d=$(seq -f "%03g" $d | tail -1)
    ip=$a$b$c$d
}
main() {

##以宿主机ip为标志之一
hostip=$(ip route get 172.16.0.0 | grep -oP '(?<=src ).*' | awk '{print $1}')
ip2int $hostip
prefix_ip=$Prefix_$ip

nvidia=$(lspci | grep NVIDIA | wc -l)
if [ $nvidia -eq 0 ];then
  echo "not have pgpu"
  exit 1;
fi
##保险起见执行 理论上是vgpu驱动加载的时候会自动加载。如果没有加载会缺少iommu_group文件导致机器开机失败
modprobe vfio_mdev
##找到物理显卡的pci
i=0
for addr in `lspci | grep NVIDIA | grep -E 'VGA|3D controller' | awk '{print $1}'`
do
  PCI_LIST[$i]="0000:$addr"
  i=$(( $i + 1 ))
done
count=${#PCI_LIST[@]}
prefix_count=$(seq -f "${Prefix_}%04g" $count | tail -1)

for ((i=0;i<count;i=i+1))
do
  file="/sys/class/mdev_bus/${PCI_LIST[i]}/mdev_supported_types"
  ls /sys/class/mdev_bus/${PCI_LIST[i]}/mdev_supported_types/$dev >/dev/null 2>/dev/null
  echo "files: ${PCI_LIST[i]}"
  if [ $? -ne 0 ];then
    echo "nvidia: $dev is not exist"
    exit 1;
  fi
  ##此类型 nvidia-* 可虚拟vgpu的个数为:ava_count
  ava_count=$(grep [0-9] $file/$dev/available_instances)
  prefix_ava=$(seq -f "${Prefix_}%04g" $ava_count | tail -1)
  for ((j=0;j<ava_count;j=j+1))
  do
    ((gen++))
    prefix_id=$(seq -f "${Prefix_}%04g" $gen | tail -1)
    id=$Prefix_1$prefix_count$prefix_ava$prefix_id$prefix_ip
    echo "$id: $file/$dev/create"
    ##开始初始化vgpu
    echo "$id" > $file/$dev/create
    if [ $? -ne 0 ];then
    echo "vgpu initialization failed"
    exit 1;
    fi
  done
done
echo "vgpu initialization done!!!"
}
main
exit 0;
```





