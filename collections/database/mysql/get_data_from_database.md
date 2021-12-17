

``` shell
mysqlsh --uri root@ip:port/xxx -p  --sql --execute  "select * from table " | column -t   >  filename
```