# Install V2ray

https://github.com/v2fly/fhs-install-v2ray/blob/master/README.zh-Hans-CN.md



1. Install

``` shell
bash <(curl -L https://raw.githubusercontent.com/v2fly/fhs-install-v2ray/master/install-release.sh)
```

2. Setup config: 

   ``` shell
   {
     "inbounds" : [
       {
         "listen" : "127.0.0.1",
         "protocol" : "socks",
         "settings" : {
           "ip" : "127.0.0.1",
           "auth" : "noauth",
           "udp" : false
         },
         "tag" : "socksinbound",
         "port" : 1081
       },
       {
         "listen" : "127.0.0.1",
         "protocol" : "http",
         "settings" : {
           "timeout" : 0
         },
         "tag" : "httpinbound",
         "port" : 8001
       }
     ],  
   "outbounds" : [
       {
         "sendThrough" : "0.0.0.0",
         "mux" : {
           "enabled" : false,
           "concurrency" : 8
         },
         "protocol" : "vmess",
         "settings" : {
           "vnext" : [
             {
               "address" : "yourserver address",
               "users" : [
                 {
                   "id" : "yourid",
                   "alterId" : 0,
                   "security" : "auto",
                   "level" : 0
                 }
               ],
               "port" : yourport
             }
           ]
         },
         "tag" : "j3",
         "streamSettings" : {
           "sockopt" : {
   
           },
           "quicSettings" : {
             "key" : "",
             "header" : {
               "type" : "none"
             },
             "security" : "none"
           },
           "tlsSettings" : {
             "allowInsecure" : false,
             "alpn" : [
               "http\/1.1"
             ],
             "serverName" : "server.cc",
             "allowInsecureCiphers" : false
           },
           "wsSettings" : {
             "path" : "",
             "headers" : {
   
             }
           },
           "httpSettings" : {
             "path" : "",
             "host" : [
               ""
             ]
           },
           "tcpSettings" : {
             "header" : {
               "type" : "none"
             }
           },
           "kcpSettings" : {
             "header" : {
               "type" : "none"
             },
             "mtu" : 1350,
             "congestion" : false,
             "tti" : 20,
             "uplinkCapacity" : 5,
             "writeBufferSize" : 1,
             "readBufferSize" : 1,
             "downlinkCapacity" : 20
           },
           "security" : "none",
           "network" : "tcp"
         }
       }
     ],  
     "routing": {
       "domainStrategy": "IPOnDemand",
       "rules": [{
         "type": "field",
         "ip": ["geoip:private"],
         "outboundTag": "direct"
       }]
     }
   }
   ```

   

3. setup v2rayserver

   ``` shell
   systemctl status v2ray
   systemctl enable v2ray
   systemctl start v2ray
   ```

4.



# setup proxy

### current shell

``` shell
export http_proxy=http://127.0.0.1:8001
export https_proxy=http://127.0.0.1:8001
```

``` shell
/etc/environment
http_proxy=http://127.0.0.1:8001
https_proxy=http://127.0.0.1:8001
```

### yum repo

#### All repo

``` shell
vim /etc/yumconf
proxy=http://127.0.0.1:8001
```

#### Specific repo

``` shell
vi /etc/yum.repos.d/kubernetes.repo 
[kubernetes]
name=Kubernetes
baseurl=https://packages.cloud.google.com/yum/repos/kubernetes-el7-$basearch
enabled=1
gpgcheck=1
repo_gpgcheck=1
gpgkey=https://packages.cloud.google.com/yum/doc/yum-key.gpg https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg
exclude=kubelet kubeadm kubectl
proxy=http://localhost:8001
```



``` shell
rm -fr /var/cache/yum/*
yum clean all
```

### docker

https://docs.docker.com/network/proxy/

https://docs.docker.com/config/daemon/systemd/#httphttps-proxy

``` shell
@mxooc according to what I see in your docker info output, you don't have proxy configured for docker daemon. Please follow documentation at https://docs.docker.com/config/daemon/systemd/#httphttps-proxy to set it correctly, and validate that docker info has following lines similar to:

kad@kad:~> docker info  | grep -i proxy
Http Proxy: http://proxy-chain.example.com:8080
Https Proxy: http://proxy-chain.example.com:8080
No Proxy: localhost,127.0.0.1,.example.com
```



