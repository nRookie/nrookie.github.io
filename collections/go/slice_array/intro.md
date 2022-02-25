``` golang

package main
import "fmt"

func sum(a []int) int {   // function that sums integers
  s := 0
  for i := 0; i < len(a); i++ {
    s += a[i]
    a[i] = 0
  }
  return s
}

func main() {
  var arr = [5]int{0,1,2,3,4} // declare an array
  fmt.Println(sum(arr[:]))    // passing slice to the function

  for i := 0; i < len(arr) ; i ++ {
    fmt.Println(arr[i])
  }
}

```

slice is always passed as reference.

``` shell
10
0
0
0
0
0
```


https://www.educative.io/courses/the-way-to-go/gknxOqxK31j

slice1 := make([]type, len)


In other words, new allocates and make initializes; the following figure illustrates this difference:

