``` golang
// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func sendCash(buyer chan int) {
	var i int
	for i = 0; i <= 3; i++ {
		buyer <- i
	}
	close(buyer)
	buyer <- i
}

func main() {
	money := make(chan int)

	go sendCash(money)

	for seller := range money {
		fmt.Println(seller)
	}
}

```



```
0
1
2
3
panic: send on closed channel
```



![image-20220314230340231](/Users/user/playground/share/nrookie.github.io/collections/go/bugs_in_work.md/image-20220314230340231.png)

![image-20220314235812104](/Users/user/playground/share/nrookie.github.io/collections/go/bugs_in_work.md/image-20220314235812104.png)





``` golang
panic: sync: WaitGroup is reused before previous Wait has returned
```



sequential



``` golang
package main
import "time"

type Node struct {
	Data int 
	Sleep time.Duration
	Left *Node
	Right *Node
}

var treeTraversal []int 

func (n *Node) TreeTraversalS() {

	if n == nil {
		return
	}

	n.Left.TreeTraversalS()
	n.ProcessNode()
	n.Right.TreeTraversalS()
}

```



Concurrent

``` golang
var wg sync.WaitGroup 

func (n *Node) ProcessNodeParallel() {

	defer wg.Done()
    
	for i := 0; i < 10000; i++ {
		time.Sleep(n.Sleep)
		
	}
    treeTraversal = append(treeTraversal, n.Data)

}


func (n *Node) TreeTraversalParallel() {
	
	defer wg.Done()

	if n == nil {
		return
	}
	//Write your code here
	go func() {
		wg.Add(1)
		n.Left.TreeTraversalParallel()
	}()
	
	wg.Add(1)
	n.ProcessNodeParallel()
	go func()	 {
		wg.Add(1)
		n.Right.TreeTraversalParallel()
	}()
}
```

