## generate command
``` shell
mkdir -p /data/folder1/sub_folder1
mkdir -p /data/folder1/sub_folder2
mkdir -p /data/folder2/
mkdir -p /data/folder2/sub_folder1/

touch /data/folder1/sub_folder1/data1
touch /data/folder1/sub_folder1/data2
touch /data/folder1/sub_folder2/data1
touch /data/folder2/sub_folder1/data1
touch /data/data1
touch /data/folder2/data2
touch /data/.hidden1
touch /data/folder1/.hidden2
```


## tar all files into subdirectory

``` shell
shopt -s dotglob;

shopt -u dotglob;
```