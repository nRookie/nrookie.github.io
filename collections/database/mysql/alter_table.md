``` shell
ALTER TABLE table_name
MODIFY COLUMN column_name datatype;
```


tidb 
``` shell
mysql> alter table t_fsx
    -> modify column `desc` longtext,
    -> modify column res_ids longtext,
    -> modify column src_net longtext,
    -> modify column tgt_net longtext,
    -> modify column message longtext,
    -> modify column ftp_info  longtext;
ERROR 8200 (HY000): Unsupported multi schema change
```