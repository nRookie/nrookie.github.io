![2021-10-25 23-45-16 的屏幕截图](/Users/user/Downloads/WXWork Files/Caches/Images/2021-10/ac8419e3a7fafc8c57fd44bf1c8c01ab/2021-10-25 23-45-16 的屏幕截图.png)





拿windbg分析了MEMORY.DMP么



![IMG_20211025_231708](/Users/user/Downloads/WXWork Files/Caches/Images/2021-10/e204fa4bda18ab1f89e778f96f027e62_HD/IMG_20211025_231708.jpg)



文件日志。



![image-20211026102321451](/Users/user/Library/Application Support/typora-user-images/image-20211026102321451.png)



clock watchdog timeout 看门狗超时





dump 文件 位置



c/Windows/memory.dmp



```xml
<features>

  [...]

  <hyperv>

​    <relaxed state='on'/>

  </hyperv>

  [...]

</features>


```





![image-20211026142505520](/Users/user/Library/Application Support/typora-user-images/image-20211026142505520.png)



https://libvirt.org/formatdomain.html#hypervisor-features