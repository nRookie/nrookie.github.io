### 



# missing go.sum entry after ran go mod tidy



workaround

``` shell
GOFLAGS=-mod=mod go generate ./...
```



 



