行为型模式然后，让我们来看最后一个类别，行为型模式（Behavioral Patterns），

它的特点是关注对象之间的通信。

这一类别的设计模式中，我们会讲到代理模式和选项模式。

代理模式代理模式 (Proxy Pattern)，

可以为另一个对象提供一个替身或者占位符，

以控制对这个对象的访问。以下代码是一个代理模式的实现：





``` go

package proxy

import "fmt"

type Seller interface {
  sell(name string)
}

// 火车站
type Station struct {
  stock int //库存
}

func (station *Station) sell(name string) {
  if station.stock > 0 {
    station.stock--
    fmt.Printf("代理点中：%s买了一张票,剩余：%d \n", name, station.stock)
  } else {
    fmt.Println("票已售空")
  }

}

// 火车代理点
type StationProxy struct {
  station *Station // 持有一个火车站对象
}

func (proxy *StationProxy) sell(name string) {
  if proxy.station.stock > 0 {
    proxy.station.stock--
    fmt.Printf("代理点中：%s买了一张票,剩余：%d \n", name, proxy.station.stock)
  } else {
    fmt.Println("票已售空")
  }
}
```



上述代码中，StationProxy 代理了 Station，代理类中持有被代理类对象，并且和被代理类对象实现了同一接口。





### 选项模式

选项模式（Options Pattern）也是 Go 项目开发中经常使用到的模式，例如，grpc/grpc-go 的NewServer函数，uber-go/zap 包的New函数都用到了选项模式。使用选项模式，我们可以创建一个带有默认值的 struct 变量，并选择性地修改其中一些参数的值。在 Python 语言中，创建一个对象时，可以给参数设置默认值，这样在不传入任何参数时，可以返回携带默认值的对象，并在需要时修改对象的属性。这种特性可以大大简化开发者创建一个对象的成本，尤其是在对象拥有众多属性时。而在 Go 语言中，因为不支持给参数设置默认值，为了既能够创建带默认值的实例，又能够创建自定义参数的实例，不少开发者会通过以下两种方法来实现：第一种方法，我们要分别开发两个用来创建实例的函数，一个可以创建带默认值的实例，一个可以定制化创建实例。





``` go

package options

import (
  "time"
)

const (
  defaultTimeout = 10
  defaultCaching = false
)

type Connection struct {
  addr    string
  cache   bool
  timeout time.Duration
}

// NewConnect creates a connection.
func NewConnect(addr string) (*Connection, error) {
  return &Connection{
    addr:    addr,
    cache:   defaultCaching,
    timeout: defaultTimeout,
  }, nil
}

// NewConnectWithOptions creates a connection with options.
func NewConnectWithOptions(addr string, cache bool, timeout time.Duration) (*Connection, error) {
  return &Connection{
    addr:    addr,
    cache:   cache,
    timeout: timeout,
  }, nil
}
```



使用这种方式，创建同一个 Connection 实例，却要实现两个不同的函数，实现方式很不优雅。另外一种方法相对优雅些。我们需要创建一个带默认值的选项，并用该选项创建实例：





``` go

package options

import (
  "time"
)

const (
  defaultTimeout = 10
  defaultCaching = false
)

type Connection struct {
  addr    string
  cache   bool
  timeout time.Duration
}

type ConnectionOptions struct {
  Caching bool
  Timeout time.Duration
}

func NewDefaultOptions() *ConnectionOptions {
  return &ConnectionOptions{
    Caching: defaultCaching,
    Timeout: defaultTimeout,
  }
}

// NewConnect creates a connection with options.
func NewConnect(addr string, opts *ConnectionOptions) (*Connection, error) {
  return &Connection{
    addr:    addr,
    cache:   opts.Caching,
    timeout: opts.Timeout,
  }, nil
}
```





使用这种方式，虽然只需要实现一个函数来创建实例，但是也有缺点：为了创建 Connection 实例，每次我们都要创建 ConnectionOptions，操作起来比较麻烦。那么有没有更优雅的解决方法呢？答案当然是有的，就是使用选项模式来创建实例。以下代码通过选项模式实现上述功能：



``` go

package options

import (
  "time"
)

type Connection struct {
  addr    string
  cache   bool
  timeout time.Duration
}

const (
  defaultTimeout = 10
  defaultCaching = false
)

type options struct {
  timeout time.Duration
  caching bool
}

// Option overrides behavior of Connect.
type Option interface {
  apply(*options)
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options) {
  f(o)
}

func WithTimeout(t time.Duration) Option {
  return optionFunc(func(o *options) {
    o.timeout = t
  })
}

func WithCaching(cache bool) Option {
  return optionFunc(func(o *options) {
    o.caching = cache
  })
}

// Connect creates a connection.
func NewConnect(addr string, opts ...Option) (*Connection, error) {
  options := options{
    timeout: defaultTimeout,
    caching: defaultCaching,
  }

  for _, o := range opts {
    o.apply(&options)
  }

  return &Connection{
    addr:    addr,
    cache:   options.caching,
    timeout: options.timeout,
  }, nil
}
```

