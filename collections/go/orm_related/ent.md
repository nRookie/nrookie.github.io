ariga.io/entimport

https://golangrepo.com/repo/ariga-entimport-go-database-schema-migration


https://www.codenong.com/cs109066460/


init t_fsx
``` shell
ent init
# copy from existing
go run ariga.io/entimport/cmd/entimport -dsn "mysql://whatever:whatever-123@tcp(127.0.0.1:3306)/shared_storage_backup?parseTime=True"

# 
```


// describe ent

``` shell


go run entgo.io/ent/cmd/ent describe ./ent/schema

```


## add edges in ent

``` golang 
func (TFsx) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("thost", THost.Type),
    }
}
```


https://medium.com/a-journey-with-go/go-ent-graph-based-orm-by-facebook-d9ba6d2290c6


edge.To and edge.From difference?



## what is storageKey?






## tidb related

ent does not have text file, which would cause the changing of datatype from text to longtext.
however TIDB does not support multi schema change 
### alter table

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

``` mysql
## t_fsx
alter table t_fsx modify column `desc` longtext;
alter table t_fsx modify column res_ids longtext;
alter table t_fsx modify column src_net longtext;
alter table t_fsx modify column tgt_net longtext;
alter table t_fsx modify column message longtext;
alter table t_fsx modify column ftp_info  longtext;
alter table t_fsx modify column f_sx_access_info longtext;

## t_host;

alter table t_host modify column network longtext;
```

