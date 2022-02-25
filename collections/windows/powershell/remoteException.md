## 通过agent 执行简单的http request 可以成功

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

