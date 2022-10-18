![image-20220317221952940](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/image-20220317221952940.png)





![image-20220317222005534](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/image-20220317222005534.png)



![image-20220317222159933](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/image-20220317222159933.png)



![image-20220317222356480](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/image-20220317222356480.png)





![image-20220317222423563](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/image-20220317222423563.png)





![image-20220317222716557](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/image-20220317222716557.png)





![image-20220317223225244](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/image-20220317223225244.png)





channel 最大的坑是 一定是交给发送着来close channel。



![image-20220317223609700](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/image-20220317223609700.png)



吞吐是靠多个goroutine 来消费它。



``` go
type g struct {
	stack       stack
	stackguard0 uintptr
}
```

其中 `stack` 字段描述了当前 Goroutine 的栈内存范围 [stack.lo, stack.hi)，另一个字段 `stackguard0` 可以用于调度器抢占式调度。除了 `stackguard0` 之外，Goroutine 中还包含另外三个与抢占密切相关的字段：



```go
type g struct {
	preempt       bool // 抢占信号
	preemptStop   bool // 抢占时将状态修改成 `_Gpreempted`
	preemptShrink bool // 在同步安全点收缩栈
}
```

Goroutine 与我们在前面章节提到的 `defer` 和 `panic` 也有千丝万缕的联系，每一个 Goroutine 上都持有两个分别存储 `defer` 和 `panic` 对应结构体的链表：''





```go
type g struct {
	_panic       *_panic // 最内侧的 panic 结构体
	_defer       *_defer // 最内侧的延迟函数结构体
}
```



最后，我们再节选一些作者认为比较有趣或者重要的字段：

```go
type g struct {
	m              *m
	sched          gobuf
	atomicstatus   uint32
	goid           int64
}
```

- `m` — 当前 Goroutine 占用的线程，可能为空；
- `atomicstatus` — Goroutine 的状态；
- `sched` — 存储 Goroutine 的调度相关的数据；
- `goid` — Goroutine 的 ID，该字段对开发者不可见，Go 团队认为引入 ID 会让部分 Goroutine 变得更特殊，从而限制语言的并发能力[10](https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-goroutine/#fn:10)；



上述四个字段中，我们需要展开介绍 `sched` 字段的 [`runtime.gobuf`](https://draveness.me/golang/tree/runtime.gobuf) 结构体中包含哪些内容：



``` golang
type gobuf struct {
	sp   uintptr
	pc   uintptr
	g    guintptr
	ret  sys.Uintreg
	...
}
```



- `sp` — 栈指针；
- `pc` — 程序计数器；
- `g` — 持有 [`runtime.gobuf`](https://draveness.me/golang/tree/runtime.gobuf) 的 Goroutine；
- `ret` — 系统调用的返回值；







这些内容会在调度器保存或者恢复上下文的时候用到，其中的栈指针和程序计数器会用来存储或者恢复寄存器中的值，改变程序即将执行的代码。



结构体 [`runtime.g`](https://draveness.me/golang/tree/runtime.g) 的 `atomicstatus` 字段存储了当前 Goroutine 的状态。除了几个已经不被使用的以及与 GC 相关的状态之外，Goroutine 可能处于以下 9 种状态：

| 状态          | 描述                                                         |
| ------------- | ------------------------------------------------------------ |
| `_Gidle`      | 刚刚被分配并且还没有被初始化                                 |
| `_Grunnable`  | 没有执行代码，没有栈的所有权，存储在运行队列中               |
| `_Grunning`   | 可以执行代码，拥有栈的所有权，被赋予了内核线程 M 和处理器 P  |
| `_Gsyscall`   | 正在执行系统调用，拥有栈的所有权，没有执行用户代码，被赋予了内核线程 M 但是不在运行队列上 |
| `_Gwaiting`   | 由于运行时而被阻塞，没有执行用户代码并且不在运行队列上，但是可能存在于 Channel 的等待队列上 |
| `_Gdead`      | 没有被使用，没有执行代码，可能有分配的栈                     |
| `_Gcopystack` | 栈正在被拷贝，没有执行代码，不在运行队列上                   |
| `_Gpreempted` | 由于抢占而被阻塞，没有执行用户代码并且不在运行队列上，等待唤醒 |
| `_Gscan`      | GC 正在扫描栈空间，没有执行代码，可以与其他状态同时存在      |



上述状态中比较常见是 `_Grunnable`、`_Grunning`、`_Gsyscall`、`_Gwaiting` 和 `_Gpreempted` 五个状态，这里会重点介绍这几个状态。Goroutine 的状态迁移是个复杂的过程，触发 Goroutine 状态迁移的方法也很多，在这里我们也没有办法介绍全部的迁移路线，只会从中选择一些介绍。



![image-20220317225913332](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/image-20220317225913332.png)





虽然 Goroutine 在运行时中定义的状态非常多而且复杂，但是我们可以将这些不同的状态聚合成三种：等待中、可运行、运行中，运行期间会在这三种状态来回切换：

- 等待中：Goroutine 正在等待某些条件满足，例如：系统调用结束等，包括 `_Gwaiting`、`_Gsyscall` 和 `_Gpreempted` 几个状态；
- 可运行：Goroutine 已经准备就绪，可以在线程运行，如果当前程序中有非常多的 Goroutine，每个 Goroutine 就可能会等待更多的时间，即 `_Grunnable`；
- 运行中：Goroutine 正在某个线程上运行，即 `_Grunning`；



![image-20220317230000571](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/image-20220317230000571.png)



上图展示了 Goroutine 状态迁移的常见路径，其中包括创建 Goroutine 到 Goroutine 被执行、触发系统调用或者抢占式调度器的状态迁移过程。



### M [#](https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-goroutine/#m)

Go 语言并发模型中的 M 是操作系统线程。调度器最多可以创建 10000 个线程，但是其中大多数的线程都不会执行用户代码（可能陷入系统调用），最多只会有 `GOMAXPROCS` 个活跃线程能够正常运行。

在默认情况下，运行时会将 `GOMAXPROCS` 设置成当前机器的核数，我们也可以在程序中使用 [`runtime.GOMAXPROCS`](https://draveness.me/golang/tree/runtime.GOMAXPROCS) 来改变最大的活跃线程数。





![image-20220317230134739](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/image-20220317230134739.png)



在默认情况下，一个四核机器会创建四个活跃的操作系统线程，每一个线程都对应一个运行时中的 [`runtime.m`](https://draveness.me/golang/tree/runtime.m) 结构体。

在大多数情况下，我们都会使用 Go 的默认设置，也就是线程数等于 CPU 数，默认的设置不会频繁触发操作系统的线程调度和上下文切换，所有的调度都会发生在用户态，由 Go 语言调度器触发，能够减少很多额外开销。

Go 语言会使用私有结构体 [`runtime.m`](https://draveness.me/golang/tree/runtime.m) 表示操作系统线程，这个结构体也包含了几十个字段，这里先来了解几个与 Goroutine 相关的字段：



``` go
type m struct {
	g0   *g
	curg *g
	...
}
```



![image-20220317230229336](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/image-20220317230229336.png)



g0 是一个运行时中比较特殊的 Goroutine，它会深度参与运行时的调度过程，包括 Goroutine 的创建、大内存分配和 CGO 函数的执行。在后面的小节中，我们会经常看到 g0 的身影。

[`runtime.m`](https://draveness.me/golang/tree/runtime.m) 结构体中还存在三个与处理器相关的字段，它们分别表示正在运行代码的处理器 `p`、暂存的处理器 `nextp` 和执行系统调用之前使用线程的处理器 `oldp`：

``` go
type m struct {
	p             puintptr
	nextp         puintptr
	oldp          puintptr
}
```

除了在上面介绍的字段之外，[`runtime.m`](https://draveness.me/golang/tree/runtime.m) 还包含大量与线程状态、锁、调度、系统调用有关的字段，我们会在分析调度过程时详细介绍它们。



### P [#](https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-goroutine/#p)

调度器中的处理器 P 是线程和 Goroutine 的中间层，它能提供线程需要的上下文环境，也会负责调度线程上的等待队列，通过处理器 P 的调度，每一个内核线程都能够执行多个 Goroutine，它能在 Goroutine 进行一些 I/O 操作时及时让出计算资源，提高线程的利用率。

因为调度器在启动时就会创建 `GOMAXPROCS` 个处理器，所以 Go 语言程序的处理器数量一定会等于 `GOMAXPROCS`，这些处理器会绑定到不同的内核线程上。

[`runtime.p`](https://draveness.me/golang/tree/runtime.p) 是处理器的运行时表示，作为调度器的内部实现，它包含的字段也非常多，其中包括与性能追踪、垃圾回收和计时器相关的字段，这些字段也非常重要，但是在这里就不展示了，我们主要关注处理器中的线程和运行队列：







