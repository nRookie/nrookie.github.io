![image-20220331102522186](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/runtime/image-20220331102522186.png)



![image-20220331102737338](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/runtime/image-20220331102737338.png)





![image-20220331102811344](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/runtime/image-20220331102811344.png)



![image-20220331103220453](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/runtime/image-20220331103220453.png)



![image-20220331103330431](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/runtime/image-20220331103330431.png)



![image-20220331103509109](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/runtime/image-20220331103509109.png)

怎么判别一个循环队列是否为空，或为满

![image-20220331103602821](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/runtime/image-20220331103602821.png)





![image-20220331103715568](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/runtime/image-20220331103715568.png)



![image-20220331103846457](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/runtime/image-20220331103846457.png)





![image-20220331104146922](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/runtime/image-20220331104146922.png)



![image-20220331104308229](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/runtime/image-20220331104308229.png)



![image-20220331104500891](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/runtime/image-20220331104500891.png)



![image-20220331104733675](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/runtime/image-20220331104733675.png)



![image-20220331104812104](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/runtime/image-20220331104812104.png)



![image-20220331104844548](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/runtime/image-20220331104844548.png)



![image-20220331105031548](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/runtime/image-20220331105031548.png)



![image-20220331105125613](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/runtime/image-20220331105125613.png)





![image-20220331105404987](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/runtime/image-20220331105404987.png)



![image-20220331105437576](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/runtime/image-20220331105437576.png)

![image-20220331105626678](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/runtime/image-20220331105626678.png)



![image-20220331105703985](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/runtime/image-20220331105703985.png)



![image-20220331105755816](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/runtime/image-20220331105755816.png)



![image-20220331105945981](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/runtime/image-20220331105945981.png)



![image-20220331110140358](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/runtime/image-20220331110140358.png)



![image-20220331110321602](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/runtime/image-20220331110321602.png)



![image-20220331110946971](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/runtime/image-20220331110946971.png)





![image-20220331111025254](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/runtime/image-20220331111025254.png)







![image-20220331111055580](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/runtime/image-20220331111055580.png)







![image-20220331111932955](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/runtime/image-20220331111932955.png)

![image-20220331112013439](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/runtime/image-20220331112013439.png)



6.



### 怎么防止channel关闭后再写



1. 通过两个chan， 一个chan专门用来stop, 可能会导致有一些消息没有被消费，就退出了。
2.





![image-20220331112217791](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/runtime/image-20220331112217791.png)





hchan 分配在堆上。



