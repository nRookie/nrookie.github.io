## append slice to a interface

``` golang
infos := res["Infos"]

//wrong way, it will append with infos slice.
infoSlice = append(infoSlice, infos)
// correct way it will extend the slice
infoSlice = append(infoSlice, v...)
```

