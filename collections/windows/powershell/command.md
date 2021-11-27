Option 1: Retrieve general information

``` shell


Get-WmiObject -class win32_logicaldisk

```


Option 2: Retrieve hard drive properties

``` shell
wmic diskdrive get Name,Model,SerialNumber,Size,Status
```


``` shell
Get-Volume
```

``` shell

[System.IO.DriveInfo]::GetDrives()

```

``` shell
get-wmiobject win32_mappedlogicaldisk | select-object deviceid,providername,freespace

```

https://www.techotopia.com/index.php/Working_with_File_Systems_in_Windows_PowerShell_1.0


## search file


``` powershell

Get-Childitem –Path C:\ -Include *what* -Recurse -ErrorAction SilentlyContinue
```