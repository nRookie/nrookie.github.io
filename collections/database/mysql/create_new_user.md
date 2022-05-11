``` shell
CREATE USER 'test'@'localhost' IDENTIFIED BY 'password';
```





``` shell
GRANT ALL  PRIVILEGES  ON *.*  TO 'test'@'localhost' with grant option;
```

``` shell
FLUSH PRIVILEGES;
```





``` shell
ALTER USER 'test'@'localhost' IDENTIFIED WITH mysql_native_password BY 'testme-out';
```





``` shell
SELECT User, Host FROM mysql.user;
```

``` shell
DROP USER 'test'@'127.0.0.1';
```

