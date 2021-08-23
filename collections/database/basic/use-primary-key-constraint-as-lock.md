

``` golang
type TGlock struct {
	LockID     string `xorm:"varchar(255) lock_id pk not null"`
	Key        string `xorm:"VARCHAR(255) not null"`
	CreateTime int    `xorm:"INT(11)"`
	ModifyTime int    `xorm:"INT(11)"`
}
```


set LockID as the primary_key,
when inserting same LockID in the table.

due to the mysql primary-key constraints mysql will report error

``` 
"Duplicate entry 'recover-working-task' for key 't_glock.PRIMARY'"
```



``` golang
// Init ...
func (gl *GDbLock) Init() (ret bool) {

	t := &store.TGlock{
		LockID: gl.LockID,
		Key:    gl.Key,
	}

	for i := 0; i < InitRetryCount; i++ {
		// 没有错误，说明自己抢先注册成功
		if err := t.Insert(); err == nil {
			ret = true
			return
		}

		// 失败，可能别人已经注册了
		// 确认是否已经有注册了
		if has, err := t.Query(gl.LockID); err == nil && has {
			ret = true
			return
		}

		// 等待5s
		time.Sleep(5 * time.Second)
	}

	return
}

```