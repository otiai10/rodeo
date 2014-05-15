# rodeo
"rodeo" is a simple [Redis](http://redis.io/) client for golang

[![Build Status](https://travis-ci.org/otiai10/rodeo.svg?branch=master)](https://travis-ci.org/otiai10/rodeo)
[![Coverage Status](https://coveralls.io/repos/otiai10/rodeo/badge.png)](https://coveralls.io/r/otiai10/rodeo)

# Usage
```go
package main

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
    _ = vaquero.Cast("my_foo", &buz)
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

