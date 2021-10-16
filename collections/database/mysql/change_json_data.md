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
| ufssmb-5sbguvg1 | {"Addr": "106.75.233.102:22", "User": "ftpAdmin", "Password": "25b52eb7-1c8a-4a57-8e6e-08fa9b60cfe4", "Protocol": "FTP"} |
| ufssmb-stacjyse | {"Addr": "113.31.108.41:22", "User": "ftpAdmin", "Password": "a239c49d-ce53-44ea-8da7-83c98dbd23b7", "Protocol": "FTP"}  |
| ufssmb-5ev1xn1o | {"Addr": "113.31.162.60:22", "User": "ftpAdmin", "Password": "662fa57b-5be1-429f-86e2-154de0f45eba", "Protocol": "FTP"}  |
+-----------------+--------------------------------------------------------------------------------------------------------------------------+
3 rows in set (0.00 sec)

```


## remove a field

UPDATE emp
SET emp.col = JSON_REMOVE(emp.col, '$.field1');