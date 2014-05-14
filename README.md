# rodeo
"rodeo" is a simple [Redis](http://redis.io/) client.

# Usage
```go
package main

import "fmt"
import "github.com/otiai10/rodeo"

type Foo struct {
    Bar string
}

func main() {

    vaquero := rodeo.TheVaquero(rodeo.Conf{"localhost","6379"})

    // Set & Get
    _ = vaquero.Set("my_key", "12345")
    val := vaquero.Get("my_key")
    // string "12345"

    // Store & Cast
    foo := Foo{"bar"}
    _ = vaquero.Store("my_foo", foo)
    var buz Foo
    _ = vaquero.Cast("my_foo", buz)
    // *Foo {"bar"}

    // Pub & Sub
    go func(){
        mess := <-vaquero.Sub("mychan")
        // "Hello, pub/sub!!"
    }
    vaquero.Pub("mychan", "Hello, pub/sub!!")
}
```

# Test
```sh
go test ./...
```

# can also support

- [memcached](https://github.com/otiai10/rodeo/tree/master/protocol/memcached)

