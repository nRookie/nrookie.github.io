点操作
package main

import (
    "fmt"
    . "foo/bar/baz"
)

func main() {
    fmt.Println(Hello(), World()) // 直接使用包内的方法即可 不需要显式使用包名
}
. 导入可以让包内的方法注册到当前包的上下文中，直接调用方法名即可，不需要再加包前缀。


## cannot use local import error

https://www.jianshu.com/p/246ffe580ebd