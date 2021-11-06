## 1. grpc: error while marshaling: string field contains invalid UTF-8","level":"error","msg":"The 0 time failure from GRPC server","req":{"excuteStr":"\n\nStart-Service sshd\n\n$identity = \"Everyone\"\n## 设置目录权限继承，这样FtpAdmin的用户的文件，才可以被SMB的Everyone用户访问\n## see https://4sysops.com/archives/install-and-configure-an-ftp-server-with-powershell/\n$NewAcl = Get-Acl

发生的原因不确定