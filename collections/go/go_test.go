https://stackoverflow.com/questions/14723229/go-test-cant-find-function-in-a-same-package

``` shell
# command-line-arguments [command-line-arguments.test]

./uhost_test.go:30:2: undefined: InitConfig
./uhost_test.go:54:11: undefined: StartUHostInstanceUntilRunning
./uhost_test.go:79:2: undefined: InitConfig
./uhost_test.go:103:11: undefined: StopUHostInstanceUntilStopped
FAIL    command-line-arguments [build failed]
FAIL


FVFF87EFQ6LR :: fsx-recycler/internal/biz ‹init-branch*› » go test -v -run TestStartUHostInstanceUntilRunning                                                                  2 ↵
=== RUN   TestStartUHostInstanceUntilRunning
=== RUN   TestStartUHostInstanceUntilRunning/TestStartUHostInstanceUntilRunning_
--- PASS: TestStartUHostInstanceUntilRunning (2.96s)
    --- PASS: TestStartUHostInstanceUntilRunning/TestStartUHostInstanceUntilRunning_ (2.96s)
PASS
ok      whatever  3.559s
```