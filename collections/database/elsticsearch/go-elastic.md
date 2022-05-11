https://developer.okta.com/blog/2021/04/23/elasticsearch-go-developers-guide







``` shell
docker pull docker.elastic.co/elasticsearch/elasticsearch:7.5.2
```



the former one is not working in M1-Chip Apple

``` shell
docker logs be6fd0cab84d
Error: could not find libjava.so
Error: Could not find Java SE Runtime Environment.
```



``` shell
docker pull docker.elastic.co/elasticsearch/elasticsearch:7.14.0-arm64 
```



