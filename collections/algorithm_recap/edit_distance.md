word1 : horse 

word2 : ros



insert a character

delete a characater

replace a character





horse 

ros 





## 解题思路

### 递归思路

用两个指针i,j分别指向 word1,word2的最后一个元素.

 

当 i 和 j指向的元素相等时,

i和j都向前移动一步.
distance不变

当i 和j指向的元素不相等时,

我们可以进行的操作有三种.





#### 1.插入一个字符.

horse . i

hrse  . j

i指向o j指向h.

在h的后面插入一个字符o,

i向前移动一步， j不变，

distance + 1.

#### 2.删除一个字符.



3.替换一个字符.
将i替换为和j一样的字符,
此时i和j一样了.
distance+1
i向前移动一步,j向前移动一步

确定base case.

1.当j走到-1时, i还没有走完
那么i要删除自己直到长度和j一样.

2.当i走到-1时,j还没有走完.
那么i需要插入一个和j一样的值,直到j走到-1.


自顶向下的递归 , 超时了

优化,
加一个字典,减少重迭项目.

84ms 勉强过

dp保存最小距离
## 动态规划

自底向上.

同样先考虑 base case

j=0时 dp[i][0] = i+1;
i=0时 dp[0][j] = j+1;



最终返回的是dp[m][n]