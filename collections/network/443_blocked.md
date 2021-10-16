https://zhuanlan.zhihu.com/p/115450863

完整报错：

```
curl: (7) Failed to connect to raw.githubusercontent.com port 443: Connection refused
```
网上的报错大多数都是安装 HomeBrew 的时候出现这个错误，Stack Overflow 上给出的解决办法也大多都是重新安装 xcode-select，自己试了下感觉还是不靠谱，而且浪费时间又繁琐。



443 端口连接被拒一般是因为墙的原因，如果你可以科学上网（Virtual Private Network）的话，在命令行键以下命令执行：


``` shell
# 7890 和 789 需要换成你自己的端口
export https_proxy=http://127.0.0.1:7890 http_proxy=http://127.0.0.1:7890 all_proxy=socks5://127.0.0.1:789
```

https://zhuanlan.zhihu.com/p/124199138

设置DNS为114.114.114.114或者8.8.8.8就好了。

