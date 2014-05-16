# rodeo
"rodeo" is a simple [Redis](http://redis.io/) client for golang

[![Build Status](https://travis-ci.org/otiai10/rodeo.svg?branch=master)](https://travis-ci.org/otiai10/rodeo)

# API Samples
## Set & Get
can set and get strings by keys.
```go
vaquero := rodeo.TheVaquero(rodeo.Conf{"localhost","6379"})

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

vaquero := rodeo.TheVaquero(conf)

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
vaqueroA := rodeo.TheVaquero(conf)
go func(){
    for {
        message := <-vaqueroA.Sub("mychan")
        // Hi, this is vaqueroB
    }
}()

vaqueroB := rodeo.TheVaquero(conf)
_ = vaqueroB.Pub("mychan", "Hi, this is vaqueroB")
```
## Tame
can provide active model for 'RANGE' and 'ZRANGE' of Redis
```go
type Member struct {
    Name string
    Age int
}
john := Member{"John",29}
paul := Member{"Paul",31}
george := Member{"George",28}

vaquero := rodeo.TheVaquero(conf)
members := vaquero.Tame(rodo.Z, "members", Member)

members.Push(john)
members.Unshift(paul)
members.Insert(1001, george) // ZRANGE only

mem000 := members.Pop()
// George
mems := members.Range(0, -1)
// Paul, John
```


# Test
```sh
go test ./...
```

# can also support

- [memcached](https://github.com/otiai10/rodeo/tree/master/protocol/memcached)

