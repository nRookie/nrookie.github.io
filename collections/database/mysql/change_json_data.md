https://stackoverflow.com/questions/34986223/how-to-update-json-data-type-column-in-mysql-5-7-10



``` mysql
         

MySQL [test_sql]> update t1 set jdoc = JSON_SET(jdoc, "$.key2", "I am ID2");

```

``` sql

Query OK, 3 rows affected (0.00 sec)
Rows matched: 3  Changed: 3  Warnings: 0

MySQL [shared_storage_test]> select short_id, ftp_info from t_fsx;
+-----------------+--------------------------------------------------------------------------------------------------------------------------+
| short_id        | ftp_info                                                                                                                 |
+-----------------+--------------------------------------------------------------------------------------------------------------------------+

3 rows in set (0.00 sec)

```


## remove a field

UPDATE emp
SET emp.col = JSON_REMOVE(emp.col, '$.field1');