## 简介

网络中的防火墙,是一种将内部和外部网络分开的方法,是一种隔离技术。防火墙在内网与外网通信时进行访问控制,依据所设置的规则对数据包作出判断,最大限度地阻止网络中不法分子破坏企业网络,从而加强了企业网络安全。


## 防火墙的分类

硬件防火墙， 如思科的ASA防火墙， H3C的Secpath的防火墙等
软件防火墙， 如iptables, firewalld 等


## Linux 包过滤防火墙简介

1、Linux操作系统中默认内置一个软件防火墙，即iptables防火墙
2、netfilter位于Linux内核中的包过滤功能体系，又称为模块，并且自动加载，是内核很小的一部分称为Linux防火墙的“内核态”，注意，真正生效的是内核态。

3、iptables位于/sbin/iptables,用来管理防火墙规则的工具称为Linux防火墙的“用户态”。仅仅是管理工具，真正起作用的是内核态。



![image-20211030100501666](/Users/kestrel/developer/nrookie.github.io/collections/network/image-20211030100501666.png)



### iptables规则链



规则的作用：对数据包进行过滤或处理

链的作用：容纳各种防火墙规则，相当于容器

链的分类依据：处理数据包的不同时机

系统默认自带的5种规则链：

INPUT：处理入站数据包

OUTPUT：处理出站数据包

FORWARD：处理转发数据包

POSTROUTING：在进行路由选择后处理数据包（出站过滤）

PREROUTING：在进行路由选择前处理数据包（入站过滤）

注意：POSTROUTING、PREROUTING在做NAT时所使用



## iptables规则表

表的作用： 容纳各种规则链

表的划分依据: 防火墙规则的作用相似， 以功能进行划分

默认包括4个规则表:

raw表: 确定是否对该数据包进行状态跟踪

mangle: 为数据包设置标记，标记之后

nat表: 修改数据包中的源、目标IP地址或端口

filter表: 确定是否放行该数据包， 即过滤

注意： 最终规则是存到链里面，最小的容器是链表里面会存放链路



### 五链四表图



![image-20211030111636557](/Users/kestrel/developer/nrookie.github.io/collections/network/route.png)

### iptables 匹配流程



规则表之间的顺序：raw→mangle→nat→filter，即先做状态跟踪→在做标记→在做修改源目IP或端口→在做是否过滤





## 规则链之间的顺序

入站：PREROUTING→INPUT 路由前发现是自己的，直接进站

出站：OUTPUT→POSTROUTING

转发：PREOUTING→FORWARD→POSTROUTING

注意：PREROUTING和POSTROUTING是最外围，规则链是靠时机分的，分为了入站，出站，转发三个时机



## 匹配流程示意图

![image-20211030113554123](/Users/kestrel/developer/nrookie.github.io/collections/network/image-20211030113554123.png)



#### 主机型防火墙：

1、入站：数据包发来，路由前，先做跟踪，再做标记，修改，查看路由，如果是发往本机的直接往上走，进站前标记，然后出站过滤

2、出站：出站和路由后，指的是最上面的路由选择，本机选择之后先经过跟踪→标记→修改→是否过滤，出站之后是mangle表的路由后→nat表的路由后



#### 网络型防火墙：



3、 转发: 数据进来以后， 经过路有前raw、 mangle\ nat, 路由前完成之后进行选择， 发现此数据是需要发到别的地方， 非本地，通过forward, 经过mangle的forward, 还要经过路由后标记、修改IP及端口， 结束。





## iptables命令语法

- 语法构成
  iptables [-I 链名] [-t 表名] [-p 条件] [-j 控制类型]
- 参数详解
  -A：在链的末尾追加一条规则
  -I：在链的开头（或指定序号）插入一条规则
  -L：列出所有的规则条目
  -n：以数字形式显示规则信息（协议解释成数字）
  -v：以更详细的方式显示规则信息
  --line-numbers：查看规则时，显示规则的序号
  -D：删除链内指定序号（或内容）的一条规则
  -F：清空所有的规则
  -P：为指定的链设置默认规则（一条没有匹配上，按照默认规则走）

- 注意事项

  不指定表名时，默认指filter表
  不指定链名时，默认指定表内的所有链
  除非设置链的默认策略，否则必须指定匹配条件
  选项、链名、控制类型使用大写字母，其余均为小写



- 数据包的常见控制类型

  ACCEPT：允许通过
  DROP：直接丢弃，不给出任何回应
  REJECT：拒绝通过，必要时会给出提示
  LOG：记录日志信息，然后传给下一跳规则继续匹配



## 规则的匹配类型



1、通用匹配（可直接使用,不依赖于其他条件或扩展）

常见的通用匹配条件
协议匹配：-p 协议名
地址匹配：-s 源地址、-d 目的地址
接口匹配：-i 入站网卡、-o 出站网卡



2、隐含匹配（要求以特定的协议匹配作为前提）



常见的通用匹配条件
端口匹配：--sport源端口、--dport目的端口
TCP标记匹配：--tcp-flags 检查范围 被设置的标记
ICMP类型匹配：--icmp-type ICMP类型



3、显式匹配（要求以"-m扩展模块”的形式明确指出类型）



常见的通用匹配条件
多端口匹配：-m multiport --sports 源端口列表；-m multiport --dports 目的端口列表
IP范围匹配：-m iprange-src-range IP范围
MAC地址匹配：-m mac --mac-source MAC地址
状态匹配：-m state --state 连接状态



