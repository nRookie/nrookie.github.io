## pods in the cronjob cannot communicate with the service

/ # curl -X POST 7200.access.prj-epc-shared-storage.svc.c3.u4

/ # ping access.prj-epc-shared-storage.svc.c3.u4
PING access.prj-epc-shared-storage.svc.c3.u4 (2002:a40:4e::1): 56 data bytes
--- access.prj-epc-shared-storage.svc.c3.u4 ping statistics ---
2 packets transmitted, 0 packets received, 100% packet loss



### use u4 cannot communicate with the service.

``` shell
apt-get update   

 apt-get install dnsutils

curl  my-nginx.prj-epc-shared-storage.svc.c3.u4:80


```

### use uae instead of u4
``` shell

curl my-nginx.prj-epc-shared-storage.svc.c3.uae:80

```