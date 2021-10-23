
An Ansible ad-hoc command uses the command-line tool ansible to automate a single task on one or more nodes. An ad-hoc command looks like this:

``` shell
ansible [pattern] -m [module] -a "[module options]"
```

Ad-hoc commands offer a quick and easy way to execute automation, but they are not built for reusability. Their advantages are that they leverage Ansible modules which are idempotent, and the knowledge gained by using them directly ports over to the Ansible playbook language.


By default, what happens when a task fails in a playbook?

A)
The host on which the task failed is removed


Running ad-hoc commands on the localhost and containers is straightforward. It gets a little more involved, however, when you target remote hosts.

Ansible relies on an inventory to match a pattern against. The pattern localhost creates an implicit inventory when you do not mention an inventory.

When you target remote hosts, you need to specify an inventory source. You pass an inventory to the ansible command with the -i option. One method for defining an inventory is to create a host_list.

You can create a host_list bypassing DNS names or IP addresses to the -i option. A comma separates each DNS name or IP address. By default, Ansible attempts to parse the input of the -i option as an inventory file. To bypass that, place a comma , at the end of the host_list.



``` shell

# Replace the <Public Ip Address> with the actual
# Linux Instance or VM IP Address
ansible all -i <Public Ip Address>, -m ping

```


Accepting this per remote machine isn’t ideal. One solution is to disable host key checking.

You can disable host key checking by setting the variable ansible_ssh_common_args to -o StrictHostKeyChecking=no and re-run the Ansible command.

``` shell
# Replace the <Public Ip Address> with the actual 
# Linux instance or VM IP Address
ansible all -i <Public Ip Address>, -m ping -e "ansible_ssh_common_args='-o StrictHostKeyChecking=no'"
```

The ping command fails because no credentials have been provided. You pass credentials to the ansible command by using two more variables, ansible_user and ansible_password.


``` shell
# Replace the <Public Ip Address> with the actual 
# Linux instance or VM IP Address and 
# <Password> with your actual password
ansible all -i <Public Ip Address>, -m ping -e "ansible_user=ansible ansible_password=<Password> ansible_ssh_common_args='-o StrictHostKeyChecking=no'"
```


The ad-hoc Ansible command is becoming somewhat ridiculous. Having to specify the variables every time makes the command way too long. It isn’t feasible to type it out every time.

In this lesson, you connected to your Ansible development environment and explored the different variables and modules required for that purpose.