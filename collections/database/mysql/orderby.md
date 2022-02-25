

use explain check execution status of a clause

``` mysql
explain select city, name, age from t where city='what' order by name limit 1000;
```








``` mysql

/* 打开optimizer_trace，只对本线程有效 */
SET optimizer_trace='enabled=on'; 

/* @a保存Innodb_rows_read的初始值 */
select VARIABLE_VALUE into @a from  performance_schema.session_status where variable_name = 'Innodb_rows_read';

/* 执行语句 */
select city, name,age from t where city='杭州' order by name limit 1000; 

/* 查看 OPTIMIZER_TRACE 输出 */
SELECT * FROM `information_schema`.`OPTIMIZER_TRACE`\G

/* @b保存Innodb_rows_read的当前值 */
select VARIABLE_VALUE into @b from performance_schema.session_status where variable_name = 'Innodb_rows_read';

/* 计算Innodb_rows_read差值 */
select @b-@a;

```


``` shell
            ],
            "filesort_summary": {
              "memory_available": 262144,
              "key_size": 8,
              "row_size": 3150,
              "max_rows_per_buffer": 83,
              "num_rows_estimate": 530,
              "num_rows_found": 530,
              "num_initial_chunks_spilled_to_disk": 0,
              "peak_memory_used": 163840,
              "sort_algorithm": "std::stable_sort",
              "sort_mode": "<fixed_sort_key, packed_additional_fields>"
            }


```
这个方法是通过查看 OPTIMIZER_TRACE 的结果来确认的，你可以从 number_of_tmp_files 中看到是否使用了临时文件。


``` shell
接下来，我来修改一个参数，让 MySQL 采用另外一种算法。


SET max_length_for_sort_data = 16;
```


检查数据库版本

``` mysql
SHOW VARIABLES LIKE "%version%";
```




这也就体现了 MySQL 的一个设计思想：如果内存够，就要多利用内存，尽量减少磁盘访问。

对于 InnoDB 表来说，rowid 排序会要求回表多造成磁盘读，因此不会被优先选择。


### rowid 排序


初始化 sort_buffer，确定放入两个字段，即 name 和 id；从索引 city 找到第一个满足 city='杭州’条件的主键 id，也就是图中的 ID_X；到主键 id 索引取出整行，取 name、id 这两个字段，存入 sort_buffer 中；从索引 city 取下一个记录的主键 id；重复步骤 3、4 直到不满足 city='杭州’条件为止，也就是图中的 ID_Y；对 sort_buffer 中的数据按照字段 name 进行排序；遍历排序结果，取前 1000 行，并按照 id 的值回到原表中取出 city、name 和 age 三个字段返回给客户端。



### 全字段排序

初始化 sort_buffer，确定放入 name、city、age 这三个字段；从索引 city 找到第一个满足 city='杭州’条件的主键 id，也就是图中的 ID_X；到主键 id 索引取出整行，取 name、city、age 三个字段的值，存入 sort_buffer 中；从索引 city 取下一个记录的主键 id；重复步骤 3、4 直到 city 的值不满足查询条件为止，对应的主键 id 也就是图中的 ID_Y；对 sort_buffer 中的数据按照字段 name 做快速排序；按照排序结果取前 1000 行返回给客户端。




其实，并不是所有的 order by 语句，都需要排序操作的。从上面分析的执行过程，我们可以看到，MySQL 之所以需要生成临时表，并且在临时表上做排序操作，其原因是原来的数据都是无序的。

所以，我们可以在这个市民表上创建一个 city 和 name 的联合索引，对应的 SQL 语句是：

联合索引可以保证索引中第一个字段相同的情况下，第二个字段的值是有序的


``` mysql
alter table t add index city_user(city, name);
```

在这个索引里面，我们依然可以用树搜索的方式定位到第一个满足 city='杭州’的记录，并且额外确保了，接下来按顺序取“下一条记录”的遍历过程中，只要 city 的值是杭州，name 的值就一定是有序的。

这样整个查询过程的流程就变成了：

1. 从索引 (city,name) 找到第一个满足 city='杭州’条件的主键 id；

2. 到主键 id 索引取出整行，取 name、city、age 三个字段的值，作为结果集的一部分直接返回；

3. 从索引 (city,name) 取下一个记录主键 id；

4. 重复步骤 2、3，直到查到第 1000 条记录，或者是不满足 city='杭州’条件时循环结束。


覆盖索引是指，索引上的信息足够满足查询请求，不需要再回到主键索引上去取数据。

这时，对于 city 字段的值相同的行来说，还是按照 name 字段的值递增排序的，此时的查询语句也就不再需要排序了。这样整个查询语句的执行流程就变成了：从索引 (city,name,age) 找到第一个满足 city='杭州’条件的记录，取出其中的 city、name 和 age 这三个字段的值，作为结果集的一部分直接返回；从索引 (city,name,age) 取下一个记录，同样取出这三个字段的值，作为结果集的一部分直接返回；重复执行步骤 2，直到查到第 1000 条记录，或者是不满足 city='杭州’条件时循环结束。

