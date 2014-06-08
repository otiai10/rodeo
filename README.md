# rodeo
"rodeo" is a simple [Redis](http://redis.io/) client for golang

[![Build Status](https://travis-ci.org/otiai10/rodeo.svg?branch=master)](https://travis-ci.org/otiai10/rodeo)

# API Samples
## Set & Get
can set and get strings by keys.
```go
vaquero, _ := rodeo.TheVaquero(rodeo.Conf{"localhost","6379"})

// Set
_ = vaquero.Set("my_key", "12345")

// Get
val := vaquero.Get("my_key")
// string "12345"
```
## Store & Cast
can set and get objects by keys.
```go
type Sample struct {
    Foo string
}

vaquero, _ := rodeo.TheVaquero(conf)

// Store
obj := Sample{"this is foo"}
_ = vaquero.Store("my_key", obj)

// Cast
var dest Sample
_ = vaquero.Cast("my_key", &dest)
// *Sample{"this is foo"}
```
## Pub & Sub
```go
vaqueroA, _ := rodeo.TheVaquero(conf)
go func(){
    for {
        message := <-vaqueroA.Sub("mychan")
        // Hi, this is vaqueroB
    }
}()

vaqueroB, _ := rodeo.TheVaquero(conf)
_ = vaqueroB.Pub("mychan", "Hi, this is vaqueroB")
```


# Test
```sh
go test ./...
```

# can also support

- [memcached](https://github.com/otiai10/rodeo/tree/master/protocol/memcached)

