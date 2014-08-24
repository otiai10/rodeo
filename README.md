# rodeo [![Build Status](https://travis-ci.org/otiai10/rodeo.svg?branch=master)](https://travis-ci.org/otiai10/rodeo) [![GoDoc](https://godoc.org/github.com/otiai10/rodeo?status.png)](https://godoc.org/github.com/otiai10/rodeo)

"rodeo" is a simple [Redis](http://redis.io/) client for Go.

![rodeo](https://cloud.githubusercontent.com/assets/931554/3240193/73767b3a-f120-11e3-8fea-2ea46ab55cc6.png)

# API Samples
## Set & Get
can set and get strings by keys.
```go
vaquero, _ := rodeo.NewVaquero("localhost","6379")

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

vaquero, _ := rodeo.NewVaquero("localhost","6379")

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
vaqueroA, _ := rodeo.NewVaquero("localhost","6379")
go func(){
    for {
        message := <-vaqueroA.Sub("mychan")
        // Hi, this is vaqueroB
    }
}()

vaqueroB, _ := rodeo.NewVaquero("localhost","6379")
_ = vaqueroB.Pub("mychan", "Hi, this is vaqueroB")
```
## Tame
can provide active model for Sorted Sets of Redis.
```go
type Member struct {
    Name string
}

vaquero, _ := rodeo.NewVaquero("localhost","6379")
// Give representative dummy object in second arg
members, _ := vaquero.Tame("members", Member{})

members.Count() // 0
members.Range().([]Member) // []

members.Add(1020, Member{"John"})
members.Add(1001, Member{"Paul"})
members.Add(1012, Member{"Ringo"})
members.Add(1100, Member{"George"})

// COUNT members -inf +inf
members.Count() // 4

// ZRANGE members -inf +inf
found := members.Find()
found[0].Score() // 1001
found[0].Retrieve().(*Member) // &Member{Name:Paul}
```

# Test
```sh
go test ./...
```
