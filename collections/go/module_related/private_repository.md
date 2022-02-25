## 1. grab the personal_access_token from github/gitlab

 1. scopes is api.


获取 personal_access_token

登录 gitlab
点击【右上角个人头像】
点击【Settings】
点击左边【Access Tokens】
新建一个 token, scopes 选择 api


创建 ~/.netrc 文件(windows 上为 ~/_netrc);内容如下
## 2. create a ~/.netrc file  (on windows, which is ~/_netrc)

``` 
machine repository url
login <gitlab_user_name>
password <personal_access_token>
```

https://docs.gitlab.com/ee/user/project/#use-your-project-as-a-go-package