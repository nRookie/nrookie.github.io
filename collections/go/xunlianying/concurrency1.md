![image-20220317173652121](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/image-20220317173652121.png)



![image-20220317184756843](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/image-20220317184756843.png)



![image-20220317184936930](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/image-20220317184936930.png)



![image-20220317185218196](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/image-20220317185218196.png)



![image-20220317185419743](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/image-20220317185419743.png)

``` shell
go tool compile -S setlimit-example.go
```



``` shell
go test -race
go build -race
```



``` shell
WARNING: DATA RACE
Write at 0x000104addc48 by goroutine 8:
  main.Routine()
      /Users/user/developer/design-pattern/gonjinjie/concurrency/race_condition_example/main.go:28 +0x68

Previous read at 0x000104addc48 by goroutine 7:
  main.Routine()
      /Users/user/developer/design-pattern/gonjinjie/concurrency/race_condition_example/main.go:25 +0x38

Goroutine 8 (running) created at:
  main.main()
      /Users/user/developer/design-pattern/gonjinjie/concurrency/race_condition_example/main.go:16 +0x60

Goroutine 7 (running) created at:
  main.main()
      /Users/user/developer/design-pattern/gonjinjie/concurrency/race_condition_example/main.go:16 +0x60
==================
Final Counter : 2
Found 1 data race(s)
```



``` golang
package main

import (
	"fmt"
	"sync"
	"time"
)

var Wait sync.WaitGroup

var Counter int = 0

func main() {
	for routine := 1; routine <= 2; routine++ {
		Wait.Add(1)
		go Routine(routine)
	}

	Wait.Wait()
	fmt.Printf("Final Counter : %d \n", Counter)
}

func Routine(id int) {
	for count := 0; count < 2; count++ {
		value := Counter
		time.Sleep(1 * time.Millisecond)
		value++
		Counter = value
	}

	Wait.Done()
}
```



![image-20220317191916356](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/image-20220317191916356.png)



![image-20220317192107725](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/image-20220317192107725.png)





![image-20220317192223134](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/image-20220317192223134.png)



![image-20220317193129022](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/image-20220317193129022.png)



![image-20220317193343884](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/image-20220317193343884.png)





最晚加锁，最早释放

锁里面的代码越短越好。



![image-20220317194213769](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/image-20220317194213769.png)



``` shell
 go test -bench
```



![image-20220317195352267](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/image-20220317195352267.png)



![image-20220317195703499](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/image-20220317195703499.png)







![image-20220317200233670](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/image-20220317200233670.png)





![image-20220317200855509](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/image-20220317200855509.png)



![image-20220317202318263](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/image-20220317202318263.png)



![image-20220317203847762](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/image-20220317203847762.png)





![image-20220317204012795](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/image-20220317204012795.png)



![image-20220317210244421](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/image-20220317210244421.png)





![image-20220317211212396](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/image-20220317211212396.png)



