## add a new user

``` sql
CREATE USER 'ucloud'@'localhost' IDENTIFIED BY 'ucloud-123';
```


### grant permission
``` sql
GRANT ALL PRIVILEGES ON * . * TO 'ucloud'@'localhost';
```

``` sql
FLUSH PRIVILEGES;
```