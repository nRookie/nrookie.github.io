## error: cannot use promoted field in struct literal of type



https://github.com/golang/go/issues/9859





``` go

package A

type Car struct {
    Color       string
    Make        string
    Model       string
}

```

``` go

package B

type car struct {
    *A.Car
}

func NewCar() car {
    return &car{
        Color: "red",
        Make:  "toyota",
        Model: "prius"}
}

```

shows an error.



## correct way.
``` go
func NewCar() *car {
    return &car{ &A.Car{
        Color: "red",
        Make:  "toyota",
        Model: "prius",
    }}
}
```