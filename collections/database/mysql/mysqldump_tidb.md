## pingcap 
https://docs.pingcap.com/tidb/v2.1/backup-and-restore


``` shell
./bin/mydumper -h 172.31.145.198 -P 20106 -u root -B what_pub --skip-tz-utc -o what_pub
```

在管理网使用不管用


``` shell
** (mydumper:2732): CRITICAL **: 10:49:26.079: Error connecting to database: Can't connect to local MySQL server through socket '/var/lib/mysql/mysql.sock' (2)
```


### 正确使用方法

``` shell
./bin/mydumper  --host 172.31.145.198 -P 20106 -u root -B fsx_pub --skip-tz-utc -o fsx_pub -p {密码必须填写}  --no-locks
```


## mysqldump


### dump

``` shell

[qing.na@JumperServer-Shanghai db]$ mysqldump  -h 1 -P 201 -u root -p what_pub > what_pub.sql
Enter password: 
mysqldump: Error: 'Expression #6 of SELECT list is not in GROUP BY clause and contains nonaggregated column 'INFORMATION_SCHEMA.FILES.EXTRA' which is not functionally dependent on columns in GROUP BY clause; this is incompatible with sql_mode=only_full_group_by' when trying to dump tablespaces
```

### import

``` shell

mysql> source what_pub.sql
Query OK, 0 rows affected (0.00 sec)

Query OK, 0 rows affected (0.00 sec)

Query OK, 0 rows affected (0.00 sec)

Query OK, 0 rows affected (0.00 sec)

Query OK, 0 rows affected (0.00 sec)

Query OK, 0 rows affected (0.00 sec)

Query OK, 0 rows affected (0.00 sec)

Query OK, 0 rows affected (0.01 sec)

Query OK, 0 rows affected (0.00 sec)

Query OK, 0 rows affected (0.00 sec)

Query OK, 0 rows affected, 1 warning (0.00 sec)

Query OK, 0 rows affected (0.00 sec)

Query OK, 0 rows affected (0.00 sec)

ERROR 2013 (HY000): Lost connection to MySQL server during query
ERROR 2006 (HY000): MySQL server has gone away
No connection. Trying to reconnect...
Connection id:    21
Current database: what_pub

Query OK, 0 rows affected (0.01 sec)

ERROR 1146 (42S02): Table 'what_pub.t_fsx' doesn't exist
ERROR 1146 (42S02): Table 'what_pub.t_fsx' doesn't exist
ERROR 1146 (42S02): Table 'what_pub.t_fsx' doesn't exist
ERROR 1146 (42S02): Table 'what_pub.t_fsx' doesn't exist
Query OK, 0 rows affected (0.00 sec)
```

https://stackoverflow.com/questions/10474922/error-2006-hy000-mysql-server-has-gone-away

https://docs.pingcap.com/tidb-data-migration/stable/task-configuration-file-full


### production db info

``` shell
+-------------------------+--------------------------------------------------------------------------+
| Variable_name           | Value                                                                    |
+-------------------------+--------------------------------------------------------------------------+
| innodb_version          | 5.6.25                                                                   |
| protocol_version        | 10                                                                       |
| tidb_row_format_version | 1                                                                        |
| tls_version             | TLSv1,TLSv1.1,TLSv1.2                                                    |
| version                 | 5.7.25-TiDB-v4.0.8-dirty                                                 |
| version_comment         | TiDB Server (Apache License 2.0) Community Edition, MySQL 5.7 compatible |
| version_compile_machine | x86_64                                                                   |
| version_compile_os      | osx10.8                                                                  |

```


#### test db info
| Variable_name           | Value                                                                    |
+-------------------------+--------------------------------------------------------------------------+
| innodb_version          | 5.6.25                                                                   |
| protocol_version        | 10                                                                       |
| tidb_row_format_version | 2                                                                        |
| tls_version             | TLSv1,TLSv1.1,TLSv1.2                                                    |
| version                 | 5.7.25-TiDB-v4.0.14                                                      |
| version_comment         | TiDB Server (Apache License 2.0) Community Edition, MySQL 5.7 compatible |
| version_compile_machine | x86_64                                                                   |
| version_compile_os      | osx10.8                                                                  |
```` shell
```

