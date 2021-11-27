## Linux/Unix 类操作系统可以使用 awk 获取每一列的数据



#### cat 文件

可以把文件的内容输出到终端窗口上

``` shell
FVFF87EFQ6LR :: ~/Downloads » cat my_data.txt 
col1 col2 col3 col4
col11 col12 col13 col14
```



####  |   

管道操作， 如 cat my_data.txt| awk '{print ($1)}'。 my_data.txt的内容会被当作是 awk '{print ($1)}' 的输入

(管道操作需要注意中英文标点符号)

#### awk 'print ($列数)' 

可以输出某个列的数据

``` shell

FVFF87EFQ6LR :: ~/Downloads » cat my_data.txt| awk '{print ($1)}'
col1
col11
FVFF87EFQ6LR :: ~/Downloads » cat my_data.txt| awk '{print ($2)}'
col2
col12
FVFF87EFQ6LR :: ~/Downloads » cat my_data.txt| awk '{print ($3)}'
col3
col13
FVFF87EFQ6LR :: ~/Downloads » cat my_data.txt| awk '{print ($4)}'
col4
col14
```



#### awk 'print($列数1) , ... ($列数x)' 输出某个列



``` shell
## 获取第1，第3行数据
FVFF87EFQ6LR :: ~/Downloads » cat my_data.txt| awk '{print ($1), ($3)}'
col1 col2
col11 col12

## 获取第1， 3，4行，数据，
FVFF87EFQ6LR :: ~/Downloads » cat my_data.txt| awk '{print ($1), ($3), ($4)}'
col1 col3 col4
col11 col13 col14

```

