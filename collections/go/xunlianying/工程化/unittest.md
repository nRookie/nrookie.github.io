微软测试指导



![image-20220319134219301](/Users/user/playground/share/nrookie.github.io/collections/go/xunlianying/工程化/image-20220319134219301.png)



![image-20220319134305329](/Users/user/playground/share/nrookie.github.io/collections/go/xunlianying/工程化/image-20220319134305329.png)



不要假定测试之间有顺序。

```
go test -parallel n
```

每个test 代码里面都做一个Once做初始化。

但是做清除的时候不方便。



![image-20220319134839701](/Users/user/playground/share/nrookie.github.io/collections/go/xunlianying/工程化/image-20220319134839701.png)



![image-20220319135022697](/Users/user/playground/share/nrookie.github.io/collections/go/xunlianying/工程化/image-20220319135022697.png)



![image-20220319135411152](/Users/user/playground/share/nrookie.github.io/collections/go/xunlianying/工程化/image-20220319135411152.png)



![image-20220319135450152](/Users/user/playground/share/nrookie.github.io/collections/go/xunlianying/工程化/image-20220319135450152.png)





交到外面的话可以mock掉。



![image-20220319135719433](/Users/user/playground/share/nrookie.github.io/collections/go/xunlianying/工程化/image-20220319135719433.png)





![image-20220319140349064](/Users/user/playground/share/nrookie.github.io/collections/go/xunlianying/工程化/image-20220319140349064.png)



![image-20220319140419403](/Users/user/playground/share/nrookie.github.io/collections/go/xunlianying/工程化/image-20220319140419403.png)

