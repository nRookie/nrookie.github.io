# Docker and iptables

On Linux, Docker manipulates `iptables` rules to provide network isolation. While this is an implementation detail and you should not modify the rules Docker inserts into your `iptables` policies, it does have some implications on what you need to do if you want to have your own policies in addition to those managed by Docker.



If you’re running Docker on a host that is exposed to the Internet, you will probably want to have iptables policies in place that prevent unauthorized access to containers or other services running on your host. This page describes how to achieve that, and what caveats you need to be aware of.



## Add iptables policies before Docker’s rules



Docker installs two custom iptables chains named `DOCKER-USER` and `DOCKER`, and it ensures that incoming packets are always checked by these two chains first.



All of Docker’s `iptables` rules are added to the `DOCKER` chain. Do not manipulate this chain manually. If you need to add rules which load before Docker’s rules, add them to the `DOCKER-USER` chain. These rules are applied before any rules Docker creates automatically.



Rules added to the FORWARD chain --either manually. If you need to add rules which load before Docker's rules, add them to the DOCKER-USER chain. These rules are applied before any Docker creates automatically.



Rules added to the FORWARD chain --either manually, or by another iptables-based firewall -- are evaluated after these chains. This means that if you expose a port through Docker, this port gets exposed no matter,  



### Restrict connections to the Docker host

By default, all external source IPs are allowed to connect to the Docker host. To allow only a specific IP or network to access the containers, insert a negated rule at the top of the `DOCKER-USER` filter chain. For example, the following rule restricts external access from all IP addresses except `192.168.1.1`:





``` shell
iptables -I DOCKER-USER -i ext_if ! -s 192.168.1.1 -j DROP
```





Please note that you will need to change `ext_if` to correspond with your host’s actual external interface. You could instead allow connections from a source subnet. The following rule only allows access from the subnet `192.168.1.0/24`:

```
 iptables -I DOCKER-USER -i ext_if ! -s 192.168.1.0/24 -j DROP
```



Finally, you can specify a range of IP addresses to accept using `--src-range` (Remember to also add `-m iprange` when using `--src-range` or `--dst-range`):

```
iptables -I DOCKER-USER -m iprange -i ext_if ! --src-range 192.168.1.1-192.168.1.3 -j DROP
```





You can combine `-s` or `--src-range` with `-d` or `--dst-range` to control both the source and destination. For instance, if the Docker daemon listens on both `192.168.1.99` and `10.1.2.3`, you can make rules specific to `10.1.2.3` and leave `192.168.1.99` open.

`iptables` is complicated and more complicated rules are out of scope for this topic. See the [Netfilter.org HOWTO](https://www.netfilter.org/documentation/HOWTO/NAT-HOWTO.html) for a lot more information.