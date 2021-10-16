
-------------------------------------------------------------------------------
Test set: com.ewolff.microservice.order.kafka.OrderKafkaTest
-------------------------------------------------------------------------------
Tests run: 1, Failures: 0, Errors: 1, Skipped: 0, Time elapsed: 10.178 s <<< FAILURE! - in com.ewolff.microservice.order.kafka.OrderKafkaTest
com.ewolff.microservice.order.kafka.OrderKafkaTest  Time elapsed: 10.178 s  <<< ERROR!
org.testcontainers.containers.ContainerLaunchException: Container startup failed
Caused by: org.testcontainers.containers.ContainerFetchException: Can't get Docker image: RemoteDockerImage(imageName=confluentinc/cp-kafka:5.2.1, imagePullPolicy=DefaultPullPolicy())
Caused by: java.lang.IllegalStateException: Could not find a valid Docker environment. Please see logs and check configuration



use sudo with root environment.
``` shell
sudo ./mvnw clean package
```




## difference between the JRE and Java SE
https://www.java.com/en/download/help/techinfo.html



## Mac os Big Sur | No compiler is provided in this environment. Perhaps you are running on a JRE rather than a JDK?

最上面返回的是jre，他需要jdk

将jre的文件夹copy到其他地方

```   shell
FVFF87EFQ6LR :: /Library/PreferencePanes » /usr/libexec/java_home -V
Matching Java Virtual Machines (2):
    13.0.8 (arm64) "Azul Systems, Inc." - "Zulu JRE 13.42.17" /Library/Java/JavaVirtualMachines/zulu-13.jre/Contents/Home
    13.0.8 (arm64) "Azul Systems, Inc." - "Zulu 13.42.17" /Library/Java/JavaVirtualMachines/zulu-13.jdk/Contents/Home
/Library/Java/JavaVirtualMachines/zulu-13.jre/Contents/Home
FVFF87EFQ6LR :: /Library/PreferencePanes » cd /Library/Java/JavaV
```

``` shell
FVFF87EFQ6LR :: /Library/PreferencePanes » cd /Library/Java/JavaVirtualMachines
FVFF87EFQ6LR :: /Library/Java/JavaVirtualMachines » ls
zulu-13.jdk zulu-13.jre
FVFF87EFQ6LR :: /Library/Java/JavaVirtualMachines » cp zulu-13.jre ~/zulu-13.jre.backup
cp: zulu-13.jre is a directory (not copied).
FVFF87EFQ6LR :: /Library/Java/JavaVirtualMachines » ls                                                                1 ↵
zulu-13.jdk zulu-13.jre
FVFF87EFQ6LR :: /Library/Java/JavaVirtualMachines » mv  zulu-13.jre ~/zulu-13.jre.backup
mv: rename zulu-13.jre to /Users/user/zulu-13.jre.backup: Permission denied
FVFF87EFQ6LR :: /Library/Java/JavaVirtualMachines » mv  zulu-13.jre ~/zulu-13.jre.backup                              1 ↵
mv: rename zulu-13.jre to /Users/user/zulu-13.jre.backup: Permission denied
FVFF87EFQ6LR :: /Library/Java/JavaVirtualMachines » sudo mv  zulu-13.jre ~/zulu-13.jre.backup                         1 ↵
Password:
FVFF87EFQ6LR :: /Library/Java/JavaVirtualMachines » ls
zulu-13.jdk
FVFF87EFQ6LR :: /Library/Java/JavaVirtualMachines » /usr/libexec/java_home -V                
Matching Java Virtual Machines (1):
    13.0.8 (arm64) "Azul Systems, Inc." - "Zulu 13.42.17" /Library/Java/JavaVirtualMachines/zulu-13.jdk/Contents/Home
/Library/Java/JavaVirtualMachines/zulu-13.jdk/Contents/Home
FVFF87EFQ6LR :: /Library/Java/JavaVirtualMachines » 


```
