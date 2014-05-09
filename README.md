# Usage
```go
package main

import "fmt"
import "github.com/otiai10/rodeo"

type Foo struct {
    Bar string
}

func main() {
    vaquero := rodeo.TheVaquero(conf)

    _ = vaquero.Set("my_key", 12345)
    val, _ := vaquero.Get("my_key")
    fmt.Pritnf("%T %v", val, val)
    // string 12345

    foo := Foo{"bar"}
    _ = vaquero.Set("my_foo", foo, Foo)
    baz := &Foo{}
    _ = vaquero.Cast("my_foo", baz)
    fmt.Pritnf("%T %v", baz, baz)
    // *Foo {"bar"}
}
```

# Test
```sh
go test
go test ./protocol/redis
```
