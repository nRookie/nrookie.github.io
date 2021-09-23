#  通过 virtual box 官网安装 virtualbox.
vagrant up 失败, 错误信息

``` shell
[oh-my-zsh] plugin 'zsh-syntax-highlighting' not found
fvff87efq6lr :: playground/vagrant/vagrantinit » vagrant up   
Bringing machine 'default' up with 'virtualbox' provider...
==> default: Checking if box 'hashicorp/bionic64' version '1.0.282' is up to date...
==> default: Clearing any previously set forwarded ports...
==> default: Clearing any previously set network interfaces...
==> default: Preparing network interfaces based on configuration...
    default: Adapter 1: nat
==> default: Forwarding ports...
    default: 22 (guest) => 2222 (host) (adapter 1)
==> default: Booting VM...
There was an error while executing `VBoxManage`, a CLI used by Vagrant
for controlling VirtualBox. The command and stderr is shown below.

Command: ["startvm", "dfd1b01e-110d-4e7b-9d08-c78761b4284b", "--type", "headless"]

Stderr: VBoxManage: error: The virtual machine 'vagrantinit_default_1632104004503_81766' has terminated unexpectedly during startup with exit code 1 (0x1)
VBoxManage: error: Details: code NS_ERROR_FAILURE (0x80004005), component MachineWrap, interface IMachine
```
## 在Macos Big sur 安装 virtual box 和 vagrant
 
### 不管用
https://apple.stackexchange.com/questions/408154/macos-big-sur-virtualbox-error-the-virtual-machine-has-terminated-unexpectedly-d


### 通过 homebrew 安装 vagrant 和 virtualbox

```shell
There was an error while executing `VBoxManage`, a CLI used by Vagrant
for controlling VirtualBox. The command and stderr is shown below.

Command: ["startvm", "dfd1b01e-110d-4e7b-9d08-c78761b4284b", "--type", "headless"]

Stderr: VBoxManage: error: The virtual machine 'vagrantinit_default_1632104004503_81766' has terminated unexpectedly during startup with exit code 1 (0x1)
VBoxManage: error: Details: code NS_ERROR_FAILURE (0x80004005), component MachineWrap, interface IMachine
```




## 重来

### 使用 brew-cask



### ref
https://sourabhbajaj.com/mac-setup/Homebrew/Cask.html
https://sourabhbajaj.com/mac-setup/Vagrant/



``` shell

arch -arm64 brew reinstall virtualbox

```

