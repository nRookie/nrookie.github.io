## some tips about bind function

``` golang

	if err := c.ShouldBindJSON(action.req); err != nil {
		log.Error(ctx, "bind req err", log.Fields{
			"err": err,
		})
		return nil, errs.ErrToResponser(actionName, errs.ErrParamParse, err.Error())
	}
```

the arguments of the API have to be same as the name defined in the structure field.

``` golang
type request struct {
RegionID              int64  `json:"az_group"  binding:"required"`
}

```

use az_group  instead of  region_id
``` json
{
    "Action": "some action",
    "az_group" : xxx
}
```

otherwise it will report 

``` shell
Field validation for 'RegionID' failed on the 'required' tag ]"}
```