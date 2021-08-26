For security reasons, in web applications, it is better to use the json.MarshalForHTML() function, which performs an HTMLEscape on the data, so that the text will be safe to embed inside HTML \<script\> tags. The default Go types used in JSON are.


- bool for JSON booleans
- float64 for JSON numbers
- string for JSON strings
- nil for JSON null.

Not everything can be JSON-encoded though; only data structures that can be represented as valid JSON will be encoded:



``` go
map[string]interface{}{
  "Name": "Wednesday",
  "Age": 6,
  "Parents": []interface{}{
    "Gomez",
    "Morticia",
  },
}
```

``` go
m := f.(map[string]interface{})
```


We can then iterate through the map with a range statement and use a type switch to access its values as their concrete types:

``` go
for k, v := range m {
  switch vv := v.(type) {
  case string:
    fmt.Println(k, "is string", vv)
  case int:
    fmt.Println(k, "is int", vv)
  case []interface{}:
    fmt.Println(k, "is an array:")
    for i, u := range vv {
      fmt.Println(i, u)
    }
  default:
    fmt.Println(k, "is of a type I don't know how to handle")
  }
}
```

https://github.com/golang/go/wiki/InterfaceSlice