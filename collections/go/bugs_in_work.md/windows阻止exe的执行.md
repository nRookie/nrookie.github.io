windows阻止exe的执行



``` shell
GOOS=windows GOARCH=amd64 go build  -o  -win-debug  main.go  
```

直接去windows 执行会显示 (prevent)
file contains a virus or potentially unwanted softwareAt powershell， 并删除相关的文件

``` shell

mv : Cannot find path 'D:\us3debug\us3cli-win-debug' because it does not exist.                                                                                                                             
At line:1 char:1                                                                                                                                                                                            
+ mv us3cli-win-debug us3cli-win-debug.exe                                                                                                                                                                  
+ ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~                                                                                                                                                                  
    + CategoryInfo          : ObjectNotFound: (D:\us3debug\us3cli-win-debug:String) [Move-Item], ItemNotFoundException                                                                                      
    + FullyQualifiedErrorId : PathNotFound,Microsoft.PowerShell.Commands.MoveItemCommand                                                                                                                    
                                                                                         

```         
解决方法, 增加一些 buildflag
``` shell 
GO           		?= go
LDFLAGS			?= -X $(VER_PKG).GitCommit=$(COMMIT_SHA1) 
LDFLAGS			+= -X $(VER_PKG).GitBranch=$(shell git symbolic-ref --short HEAD) 
LDFLAGS			+= -X $(VER_PKG).GitSummary=build-by-$(shell id -u -n) 
LDFLAGS			+= -X $(VER_PKG).BuildTime=$(TIME) 
LDFLAGS			+= -X $(VER_PKG).Version=$(LAST_TAG)

BUILD		 	?= $(GO) build -ldflags "${LDFLAGS}"

.PHONY: build
build:
	GOOS=windows GOARCH=amd64 $(BUILD) -o test.exe main.go
```