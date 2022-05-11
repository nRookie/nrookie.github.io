#### 结构型模式

我已经向你介绍了单例模式、工厂模式这两种创建型模式，接下来我们来看结构型模式（Structural Patterns），它的特点是关注类和对象的组合。这一类型里，我想详细讲讲策略模式和模板模式。



### 策略模式

策略模式（Strategy Pattern）定义一组算法，将每个算法都封装起来，并且使它们之间可以互换。



在什么时候，我们需要用到策略模式呢？在项目开发中，我们经常要根据不同的场景，采取不同的措施，也就是不同的策略。比如，假设我们需要对 a、b 这两个整数进行计算，根据条件的不同，需要执行不同的计算方式。我们可以把所有的操作都封装在同一个函数中，然后通过 if ... else ... 的形式来调用不同的计算方式，这种方式称之为硬编码





``` go

package strategy

// 策略模式

// 定义一个策略类
type IStrategy interface {
  do(int, int) int
}

// 策略实现：加
type add struct{}

func (*add) do(a, b int) int {
  return a + b
}

// 策略实现：减
type reduce struct{}

func (*reduce) do(a, b int) int {
  return a - b
}

// 具体策略的执行者
type Operator struct {
  strategy IStrategy
}

// 设置策略
func (operator *Operator) setStrategy(strategy IStrategy) {
  operator.strategy = strategy
}

// 调用策略中的方法
func (operator *Operator) calculate(a, b int) int {
  return operator.strategy.do(a, b)
}
```



### 模版模式



模版模式 (Template Pattern) 定义一个操作中算法的骨架，而将一些步骤延迟到子类中。这种方法让子类在不改变一个算法结构的情况下，就能重新定义该算法的某些特定步骤。简单来说，模板模式就是将一个类中能够公共使用的方法放置在抽象类中实现，将不能公共使用的方法作为抽象方法，强制子类去实现，这样就做到了将一个类作为一个模板，让开发者去填充需要填充的地方。以下是模板模式的一个实现：





``` go

package template

import "fmt"

type Cooker interface {
  fire()
  cooke()
  outfire()
}

// 类似于一个抽象类
type CookMenu struct {
}

func (CookMenu) fire() {
  fmt.Println("开火")
}

// 做菜，交给具体的子类实现
func (CookMenu) cooke() {
}

func (CookMenu) outfire() {
  fmt.Println("关火")
}

// 封装具体步骤
func doCook(cook Cooker) {
  cook.fire()
  cook.cooke()
  cook.outfire()
}

type XiHongShi struct {
  CookMenu
}

func (*XiHongShi) cooke() {
  fmt.Println("做西红柿")
}

type ChaoJiDan struct {
  CookMenu
}

func (ChaoJiDan) cooke() {
  fmt.Println("做炒鸡蛋")
}
```





``` go
func TestTemplate(t *testing.T) { 
  // 做西红柿 
  xihongshi := &XiHongShi{} 
  doCook(xihongshi) fmt.Println("\n=====> 做另外一道菜") 
  // 做炒鸡蛋 
  chaojidan := &ChaoJiDan{} 
  doCook(chaojidan)}
```



