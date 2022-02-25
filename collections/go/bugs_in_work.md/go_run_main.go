``` shell
FVFF87EFQ6LR :: ~/developer/design-pattern » go run main.go
fork/exec /var/folders/46/v3z8sbqx5h75tbsysqkrm7lr0000gn/T/go-build3341180393/b001/exe/main: exec format error
```


``` shell
https://stackoverflow.com/questions/13870963/exec-format-error

FVFF87EFQ6LR :: ~/playground/fsx-recycler ‹init-branch*› » go env  | grep GOOS
GOOS="linux"

```


## solved
``` shell
go env -w  GOOS=darwin
```