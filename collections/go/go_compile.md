``` shell
GOOS=linux GOARCH=amd64 go tool compile -I $GOPATH/pkg/linux_amd64 -S defer_panic.go
```