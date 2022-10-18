``` go
package main



import (

​    "fmt"

​    "log"

​    "net/http"

​    _ "net/http/pprof"

)



func main() {

​    //doDoNotRecommend()

​    Modify2()

}



func doDoNotRecommend() {

​    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

​        fmt.Fprint(w, "Hello,no")

​    })



​    go func() {

​        if err := http.ListenAndServe(":8080", nil); err != nil {

​            log.Fatalf(err.Error())

​        }

​    }()



​    select {}

}



func doDoNotRecommend1() {

​    mux := http.NewServeMux()

​    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

​        fmt.Fprint(w, "Hello,no")

​    })



​    go http.ListenAndServe("127.0.0.1:8001", http.DefaultServeMux)

​    http.ListenAndServe("0.0.0.0:8080", mux)



​    select {}

}



// 每次启动 go routine时 要问自己

// 1. 这个goroutine 什么时候结束？

// 2. 如何让他结束掉？



// 永远不要启动一个不知道何时结束的gorotuine



// 将两个逻辑抽出来

// func serverApp() {

//  mux := http.NewServeMux()

//  mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {

//      fmt.Fprintln(resp, "Hello, QCon!")

//  })

//  http.ListenAndServe("0.0.0.0:8080", mux)

// }



// func serverDebug() {

//  http.ListenAndServe("127.0.0.1:8001", http.DefaultServeMux)

// }



func Modify1() {

​    go serverDebug()

​    serverApp()

}


```



// 如果serverDebug退出了， 仍然感知不到。

``` go
func serverApp() {

​    mux := http.NewServeMux()

​    mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {

​        fmt.Fprintln(resp, "Hello, QCon!")

​    })

​    if err := http.ListenAndServe("0.0.0.0:8080", mux); err != nil {

​        log.Fatal(err)

​    }

}



func serverDebug() {

​    if err := http.ListenAndServe("127.0.0.1:8001", http.DefaultServeMux); err != nil {

​        log.Fatal(err)

​    }

}



func Modify2() {

​    go serverDebug()

​    go serverApp()

​    select {}

}


```





// log.Fatal 调用了 os.Exit , 会无条件终止程序， defers 不会被调用到。

// Only use log.Fatal from main.main or init function.



##  推荐



``` go
// func server(addr string, handler http.Handler, stop <-chan struct{}) error {

//  s := http.Server{

//      Addr:    addr,

//      Handler: handler,

//  }



//  go func() {

//      <-stop // wait for stop signal

//      s.Shutdown(context.Background())

//  }()



//  return s.ListenAndServe()

// }



// func main() {

//  done := make(chan error, 2)

//  stop := make(chan struct{})

//  go func() {

//      done <- serveDebug(stop)

//  }()

//  go func() {

//      done <- serveApp(stop)

//  }()



//  var stopped bool

//  for i := 0; i < cap(done); i ++ {

//      if err := <-done; err != nil {

//          fmt.Println("error: %v", err)

//      }



//      if !stopped {

//          stopped = true

//          close(stop)

//      }

//  }

// }



// 如果启动一个goroutine, 一定要知道它什么时候可以结束
```





## Leave concurrency to the caller

![image-20220317140828607](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/image-20220317140828607.png)



![image-20220317141103892](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/image-20220317141103892.png)



Go routine 泄漏

![image-20220317141340993](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/image-20220317141340993.png)



![image-20220317141649814](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/image-20220317141649814.png)

Leak() 函数执行完成以后，里面的go func（） 永远不会退出。



![image-20220317141843964](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/image-20220317141843964.png)



![image-20220317145217040](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/image-20220317145217040.png)



![image-20220317145414491](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/image-20220317145414491.png)





尽量不要在http请求里面开go routine.

![image-20220317145757428](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/image-20220317145757428.png)



1. 能够控制goroutine什么时候退出。
2. 要搞清楚goroutine什么时候退出。
3. 把并发扔给调用者，调用者决定要在前台执行还是在后台执行。



close 一个channel， 如果还有人在写，一定会panic的。



什么时候可以退出？



1. shutdown 监听退出。

2. 通过context. 超时退出。

3. channel 退出。



![image-20220317154134051](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/image-20220317154134051.png)



![image-20220317155114017](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/image-20220317155114017.png)



![image-20220317155359395](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/image-20220317155359395.png)





![image-20220317155722558](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/image-20220317155722558.png)





![image-20220317160230345](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/image-20220317160230345.png)





https://www.jianshu.com/p/5e44168f47a3
