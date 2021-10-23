You already have a way of inventorying your infrastructure. It might be:

- By mentally remembering hostnames.
- By grouping them by a standard naming convention.
- With a custom database that stores specific properties about the hosts.
- With a full-blown CMDB (configuration management database) that has entries for all the hosts.


Ansible codifies this information by using an inventory.

Inventories describe your infrastructure by listing hosts and groups of hosts. Using an inventory gives you the ability to share standard variables among groups. Inventories are also parsable with host patterns, giving you the ability to broadly or narrowly target your infrastructure with Ansible automation. In its simplest form, an Ansible inventory is a file.

Let’s suppose that, within an infrastructure, there are three web servers and a database server. The Ansible inventory might look like this.


``` shell
[webservers]
web01
web02
web03
[dbservers]
db01

```

Each host or group can be used for pattern matching or to define one or more variables.

``` shell
db01 database_name=customerdb
```

The above command creates a host variable named database_name with the value of customerdb. This variable is scoped to the host db01.



``` shell
[webserver:vars] 
http_port=8080
```

The above snippet creates a group variable named http_port with the value of 8080, and it is shared among all three web servers.

