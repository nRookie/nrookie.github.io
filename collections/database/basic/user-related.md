## add a new user

``` sql
CREATE USER 'ok'@'localhost' IDENTIFIED BY 'bad-123';
```


### grant permission
``` sql
GRANT ALL PRIVILEGES ON * . * TO 'ok'@'localhost';
```

``` sql
FLUSH PRIVILEGES;
```