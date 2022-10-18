# 复盘

KMP 解决什么问题？

字符串匹配。

S是主串， P 是模式串。

找 P在S中出现的位置



## **暴力**

``` go
for i := 0; i < len(S); i ++ {
      flag := true
      for j := 0; j < len(P); j ++ {

            if S[i + j] != P[j] {

                flag = false

                break

            }
			}
      if flag {

        fmt.Printf("found one %d \n", i)

      }
}
```







\```

暴力有什么问题？如果匹配不到元素。

i + j 要回退到 i + 1, j 要回退到 0 重新开始计算。会有很多重复计算的步骤。



我们需要让 i + j,  j 回退的尽可能的少。



![image-20220829045016276](/Users/kestrel/developer/nrookie.github.io/collections/algorithm_recap/image-20220829045016276.png)





![image-20220829051123734](/Users/kestrel/developer/nrookie.github.io/collections/algorithm_recap/image-20220829051123734.png)

勘误， 若 p[j + 1] == p[i] , next[i] = j + 1



``` go


func strStr(s string, p string) int {

    if len(p) == 0 {
        return 0
    }
    
    s = " " + s
    p = " " + p
    n, m := len(s), len(p)
  
    next := make([]int, m)
  
    j := 0 // j + 1 指向的是我们要比较的元素，因此从0开始
    // next[0] 的前缀是 0
    // next[1] 的前缀是 0
    // 2        0
    //  abababab
    // 012345678
    // 000123456

    // 012
    //  ll
    // 001
    //

    // i 从 2 开始就是对应的下标。 // 1 肯定没有前缀。
    for i := 2; i < m; i ++ {
        for j != 0 && p[j + 1] != p[i]  {
            j = next[j]
        }
        if p[j + 1] == p[i] {
            next[i] = j + 1
            //fmt.Printf("%d %d %d \n", i, next[i], j + 1)
            j++ // j 也要往后移动
        }
    }
    // fmt.Println(next)
    // fmt.Printf("%d", n)

    j = 0
    for i := 0; i < n; i ++ {

        for j != 0 && p[j + 1] != s[i]  {
            j = next[j]
        }

        if p[j + 1] == s[i] {
            j ++
        } 
        
        if j == m - 1 { // j 从 0 开始走走到 m -1 走了 m步。 我们用 j + 1 做对比，因此走到 m-1的时候要退出
            return i - j
        }
    }

    return -1
}
```





没考虑到的细节。

构造next数组的时候，如果匹配了。 j要++

