## no such host error in grpc with os.exec

``` shell
{"L":"info","T":"2022-01-24T07:36:56.296Z","C":"agentclient/client.go:138","M":"return res: retCode:RC_FAILED message:\"exit status 1,2022-01-23 23:36:56.254 ERROR Put \\\"http://backup./1234/5678/\\\": dial tcp: lookup backup.: no such host\\n,\" err : <nil> ","service":"whatever/fsx-recycler"}
{"L":"info","T":"2022-01-24T07:36:56.296Z","C":"biz/agentService.go:80","M":"createUS3DirWindows  ","service":"whatever/fsx-recycler"}
```

1. 通过用户网访问，还是不可以访问。因此跟eip没有关系。
2. 云主机直接通过golang exec库执行 us3cli 命令， 可以正常执行。
3. 因此怀疑是grpc 和 golang exec 库搭配使用出现的问题。


``` shell
{"L":"info","T":"2022-01-24T16:03:57.573+0800","C":"agentclient/client.go:114","M":"excute local cmds ","service":"what/fsx-recycler"}
{"L":"info","T":"2022-01-24T16:03:58.665+0800","C":"agentclient/client.go:138","M":"return res: message:\"2022-01-24 16:03:58.729 INFO MakeDir [ 1234/5678/ ] success\\n\" output:\"2022-01-24 16:03:58.729 INFO MakeDir [ 1234/5678/ ] success\\n\" err : <nil> ","service":"whatever/fsx-recycler"}
{"L":"info","T":"2022-01-24T16:03:58.665+0800","C":"biz/agentService.go:126","M":"output is 2022-01-24 16:03:58.729 INFO MakeDir [ 1234/5678/ ] success\n","service":"whatever/fsx-recycler"}
```

linux 系统可以创建成功。
而 Windows系统不可以。


尝试增加异步命令。

用户网不可以访问kun的请求，
因此尝试通过polling的方法去检测任务执行的过程。

短任务直接使用同步的方式下载。 长任务使用异步加 polling的方式下载。

``` shell
function main()
    cmd := exec.Command( "shell.sh" )
    err := cmd.Start()
    if err != nil {
        return err
    }
    pid := cmd.Process.Pid
    // use goroutine waiting, manage process
    // this is important, otherwise the process becomes in S mode
    go func() { 
        err = cmd.Wait()
        fmt.Printf("Command finished with error: %v", err)
    }()
    return nil
}
```

使用polling的方式还是会发生错误

``` shell
PS D:\> cat .\result1.txt                                                                                                        
2022-02-08 18:25:37.972 ERROR Put "http://shared-storage-backup./50000/60000/"
```


// non interactive powershell 

``` powershell
PS D:\> powershell.exe -NoProfile -NonInteractive "D:\\us3cli-windows.exe mkdir us3://shared-storage-backup/133/134  -p > D:\\res
.txt"  

PS D:\> cat .\res.txt                                                                                                            
2022-02-08 18:39:53.892 INFO MakeDir [ 133/134/ ] success                                                                        
```

## 通过agent 执行简单的http request 可以成功但是有报错

``` shell
res, err := smb.executeShellCommand(ctx, "D:\\httpt.exe 2> D:\\httpt.txt ")
```

``` shell
PS D:\> cat .\httpt.txt                                                                                                                    
D:\httpt.exe : 2022/02/08 18:54:01 {                                                                                                       
At line:1 char:1                                                                                                                           
+ D:\httpt.exe 2> D:\httpt.txt                                                                                                             
+ ~~~~~~~~~~~~~~~~~~~~~~~~~~~~                                                                                                             
    + CategoryInfo          : NotSpecified: (2022/02/08 18:54:01 {:String) [], RemoteException                                             
    + FullyQualifiedErrorId : NativeCommandError                                                                                           
                                                                                                                                           
  "userId": 1,                                                                                                                             
  "id": 1,                                                                                                                                 
  "title": "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",                                                   
  "body": "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas                          
totam\nnostrum rerum est autem sunt rem eveniet architecto"                                                                                
}   

```

## 不通过agent直接通过 os.exec

``` golang
package main                                                                                                                     
                                                                                                                                 
import "os/exec"                                                                                                                 
import "fmt"                                                                                                                     
import "bytes"                                                                                                                   
func main() {                                                                                                                    
        ps, _ := exec.LookPath("powershell.exe")                                                                                 
        args := []string{"D:\\us3cli-windows.exe mkdir us3://shared-storage-backup/133/136  -p > D:\\res.txt",}                  
        args = append([]string{"-NoProfile", "-NonInteractive"}, args...)                                                        
        cmd := exec.Command(ps, args...)                                                                                         
                                                                                                                                 
        var stdout bytes.Buffer                                                                                                  
        var stderr bytes.Buffer                                                                                                  
        cmd.Stdout = &stdout                                                                                                     
        cmd.Stderr = &stderr                                                                                                     
                                                                                                                                 
        err := cmd.Run()                                                                                                         
        if err != nil {                                                                                                          
                fmt.Println(err.Error())                                                                                         
        }                                                                                                                        
                                                                                                                                 
        stdOut, stdErr := stdout.String(), stderr.String()                                                                       
        fmt.Println(stdOut)                                                                                                      
        fmt.Println(stdErr)                                                                                                      
                                                                                                                                 
}                  

```

``` shell
PS D:\> cat .\res.txt                                                                                                            
2022-02-08 18:43:42.475 INFO MakeDir [ 133/136/ ] success  
```
可以成功执行。

## 通过agent 执行 mkdir.exe文件


``` golang
	res, err := smb.executeShellCommand(ctx, "D:\\mkdir.exe") 
```

``` shell

PS D:\> cat .\res.txt                                                                                                                      2022-02-08 18:48:45.303 ERROR Put "http://shared-storage-backup./133/136/": dial tcp: lookup shared-storage-backup.: no such host 

```



## 

``` golang

res, err := smb.executeShellCommand(ctx, "D:\\mkdir.ps1")
```
``` shell
PS D:\> cat .\mkdir.ps1                                                                                                                    
powershell.exe -NoProfile -NonInteractive "D:\\us3cli-windows.exe mkdir us3://shared-storage-backup/133/134  -p > D:\\res.txt"   

PS D:\> cat .\res.txt                                                                                                                      
2022-02-08 18:58:08.622 ERROR Put "http://shared-storage-backup./133/134/": dial tcp: lookup shared-storage-backup.: no such host 
```





## profile related

https://docs.microsoft.com/en-us/powershell/module/microsoft.powershell.core/about/about_profiles?view=powershell-7.2

### echo home
``` remote

res, err := smb.executeShellCommand(ctx, "echo $HOME > D:\\home.txt") //D:\\pingwhatever.exe   > D:\\pingwhatever ")
C:\Windows\system32\config\systemprofile 

```

```  shell
echo $HOME                                                                                                                         
C:\Users\Administrator
```

### echo psHome

remote 
``` shell
res, err := smb.executeShellCommand(ctx, "echo $PSHOME > D:\\pshome.txt")
PS D:\> cat .\pshome.txt                                                                                                                   
C:\Windows\System32\WindowsPowerShell\v1.0   
```

local

``` shell
PS D:\> echo $PSHOME                                                                                                                       
C:\Windows\System32\WindowsPowerShell\v1.0  
```


### username

``` shell
res, err := smb.executeShellCommand(ctx, "$env:UserName  > D:\\username.txt")
cat .\username.txt                                                                                                                 
WIN-0VHMAQ8Q4SF$ 
```


### login as admin

``` shell
PS D:\> cat .\admin.ps1                                                                                                                    
$username = 'administrator'                                                                                                                
$password = 'password'                                                                                                                   
                                                                                                                                           
$securePassword = ConvertTo-SecureString $password -AsPlainText -Force                                                                     
$credential = New-Object System.Management.Automation.PSCredential $username, $securePassword                                              
                                                                                                                                           
$env:UserName  > D:\\username.txt   
```



### env

local

``` 
PS D:\> dir env:                                                                                                                           
                                                                                                                                           
Name                           Value                                                                                                       
----                           -----                                                                                                       
ALLUSERSPROFILE                C:\ProgramData                                                                                              
APPDATA                        C:\Users\Administrator\AppData\Roaming                                                                      
c28fc6f98a2c44abbbd89d6a303... d:\ftpfiles                                                                                                 
ChocolateyInstall              C:\ProgramData\chocolatey                                                                                   
ChocolateyLastPathUpdate       132784213802437083                                                                                          
ChocolateyToolsLocation        C:\tools                                                                                                    
CommonProgramFiles             C:\Program Files\Common Files                                                                               
CommonProgramFiles(x86)        C:\Program Files (x86)\Common Files                                                                         
CommonProgramW6432             C:\Program Files\Common Files                                                                               
COMPUTERNAME                   WIN-0VHMAQ8Q4SF                                                                                             
ComSpec                        C:\Windows\system32\cmd.exe                                                                                 
DriverData                     C:\Windows\System32\Drivers\DriverData                                                                      
HOME                           C:\Users\Administrator                                                                                      
HOMEDRIVE                      C:                                                                                                          
HOMEPATH                       \Users\Administrator                                                                                        
LOCALAPPDATA                   C:\Users\Administrator\AppData\Local                                                                        
LOGNAME                        administrator                                                                                               
NUMBER_OF_PROCESSORS           2                                                                                                           
OS                             Windows_NT                                                                                                  
Path                           C:\Windows\system32;C:\Windows;C:\Windows\System32\Wbem;C:\Windows\System32\WindowsPowerShell\v1.0\;C:\W... 
PATHEXT                        .COM;.EXE;.BAT;.CMD;.VBS;.VBE;.JS;.JSE;.WSF;.WSH;.MSC;.CPL                                                  
PROCESSOR_ARCHITECTURE         AMD64                                                                                                       
PROCESSOR_IDENTIFIER           Intel64 Family 6 Model 85 Stepping 6, GenuineIntel                                                          
PROCESSOR_LEVEL                6                                                                                                           
PROCESSOR_REVISION             5506                                                                                                        
ProgramData                    C:\ProgramData                                                                                              
ProgramFiles                   C:\Program Files                                                                                            
ProgramFiles(x86)              C:\Program Files (x86)                                                                                      
ProgramW6432                   C:\Program Files                                                                                            
PROMPT                         administrator@WIN-0VHMAQ8Q4SF $P$G                                                                          
PSModulePath                   C:\Users\Administrator\Documents\WindowsPowerShell\Modules;C:\Program Files\WindowsPowerShell\Modules;C:... 
PUBLIC                         C:\Users\Public                                                                                             
SHELL                          c:\windows\system32\windowspowershell\v1.0\powershell.exe                                                   
SSH_CLIENT                     106.75.220.2 61736 22                                                                                       
SSH_CONNECTION                 106.75.220.2 61736 172.16.0.126 22                                                                          
SSH_TTY                        windows-pty                                                                                                 
SystemDrive                    C:                                                                                                          
SystemRoot                     C:\Windows                                                                                                  
TEMP                           C:\Users\Administrator\AppData\Local\Temp                                                                   
TERM                           xterm-256color                                                                                              
TMP                            C:\Users\Administrator\AppData\Local\Temp                                                                   
USER                           administrator                                                                                               
USERDOMAIN                     WORKGROUP                                                                                                   
USERNAME                       administrator                                                                                               
USERPROFILE                    C:\Users\Administrator                                                                                      
windir                         C:\Windows 
```


## remote


``` shell
                                                                                                                                           
PS D:\> cat .\envremote.txt                                                                                                                
                                                                                                                                           
Name                           Value                                                                                                       
----                           -----                                                                                                       
ALLUSERSPROFILE                C:\ProgramData                                                                                              
APPDATA                        C:\Windows\system32\config\systemprofile\AppData\Roaming                                                    
ChocolateyInstall              C:\ProgramData\chocolatey                                                                                   
CommonProgramFiles             C:\Program Files\Common Files                                                                               
CommonProgramFiles(x86)        C:\Program Files (x86)\Common Files                                                                         
CommonProgramW6432             C:\Program Files\Common Files                                                                               
COMPUTERNAME                   WIN-0VHMAQ8Q4SF                                                                                             
ComSpec                        C:\Windows\system32\cmd.exe                                                                                 
DriverData                     C:\Windows\System32\Drivers\DriverData                                                                      
GOROOT                         C:\go1.16.7                                                                                                 
LOCALAPPDATA                   C:\Windows\system32\config\systemprofile\AppData\Local                                                      
NUMBER_OF_PROCESSORS           2                                                                                                           
OS                             Windows_NT                                                                                                  
Path                           C:\Windows\system32;C:\Windows;C:\Windows\System32\Wbem;C:\Windows\System32\WindowsPo...                    
PATHEXT                        .COM;.EXE;.BAT;.CMD;.VBS;.VBE;.JS;.JSE;.WSF;.WSH;.MSC;.CPL                                                  
PROCESSOR_ARCHITECTURE         AMD64                                                                                                       
PROCESSOR_IDENTIFIER           Intel64 Family 6 Model 85 Stepping 6, GenuineIntel                                                          
PROCESSOR_LEVEL                6                                                                                                           
PROCESSOR_REVISION             5506                                                                                                        
ProgramData                    C:\ProgramData                                                                                              
ProgramFiles                   C:\Program Files                                                                                            
ProgramFiles(x86)              C:\Program Files (x86)                                                                                      
ProgramW6432                   C:\Program Files                                                                                            
PSModulePath                   WindowsPowerShell\Modules;C:\Program Files\WindowsPowerShell\Modules;C:\Windows\syste...                    
PUBLIC                         C:\Users\Public                                                                                             
SystemDrive                    C:                                                                                                          
SystemRoot                     C:\Windows                                                                                                  
TEMP                           C:\Windows\TEMP                                                                                             
TMP                            C:\Windows\TEMP                                                                                             
USERDOMAIN                     WORKGROUP                                                                                                   
USERNAME                       WIN-0VHMAQ8Q4SF$                                                                                            
USERPROFILE                    C:\Windows\system32\config\systemprofile                                                                    
windir                         C:\Windows 
```

### ssh execute a command

``` 
FVFF87EFQ6LR :: ~ » ssh administrator@106.75.237.12 "D:\\us3cli-windows.exe mkdir us3://shared-storage-backup/345/345  -p > D:\\result1.txt"
administrator@106.75.237.12's password:

PS D:\> cat .\result1.txt                                                                                                                  
2022-02-08 22:11:41.110 INFO MakeDir [ 345/345/ ] success

```

``` golang
ssh.ExecOutput(context.Background(), "D:\\us3cli-windows.exe mkdir us3://shared-storage-backup/456/456  -p > D:\\sshresult1.txt")
```

use ssh execute the command.
