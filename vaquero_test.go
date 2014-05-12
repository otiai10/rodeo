package rodeo_test

import "github.com/otiai10/rodeo"

import "fmt"
import "testing"

var conf = rodeo.Conf{
	Host: "localhost",
	Port: "6379",
}

type tStruct0 struct {
	Foo string
}
type tStruct1 struct {
	Foo Foo
}
type Foo struct {
	Bar string
	Buz int
}

func assert(t *testing.T, actual interface{}, expected interface{}) {
	if expected != actual {
		fmt.Printf("Expected `%+v`, but Actual `%+v`\n", expected, actual)
		t.Fail()
	}
}

func TestTheVaquero(t *testing.T) {

	vaquero, e := rodeo.TheVaquero(conf, "test")
	assert(t, nil, e)
	assert(t, "6379", vaquero.Conf.Port)
}

func TestVaquero_Set(t *testing.T) {

	vaquero, e := rodeo.TheVaquero(conf, "test")

	e = vaquero.Set("mykey", "12345")
	assert(t, nil, e)
	val := vaquero.Get("mykey")
	assert(t, "12345", val)

	e = vaquero.Set("mykey", "67890")
	assert(t, e, nil)

	val = vaquero.Get("mykey")
	assert(t, val, "67890")
}

func TestVaquero_Store(t *testing.T) {

	vaquero, e := rodeo.TheVaquero(conf, "test")

	key0 := "mykey0"
	obj0 := tStruct0{"Hello, rodeo"}
	e = vaquero.Store(key0, obj0)
	assert(t, nil, e)

	var dest0 tStruct0
	e = vaquero.Cast(key0, &dest0)
	assert(t, nil, e)
	assert(t, "Hello, rodeo", dest0.Foo)

	key1 := "mykey1"
	obj1 := tStruct1{Foo: Foo{Bar: "This is Bar", Buz: 1001}}
	e = vaquero.Store(key1, obj1)
	assert(t, nil, e)

	var dest1 tStruct1
	e = vaquero.Cast(key1, &dest1)
	assert(t, nil, e)
	assert(t, "This is Bar", dest1.Foo.Bar)
	assert(t, 1001, dest1.Foo.Buz)
}

func TestVaquero_PubSub(t *testing.T) {

	fin := make(chan string)

	vaqueroA, _ := rodeo.TheVaquero(conf, "test")
	vaqueroB, _ := rodeo.TheVaquero(conf, "test")

	subscriber := vaqueroA.Sub("mychan")

	go func() {
		for {
			message := <-subscriber
			fin <- message
		}
	}()

	_ = vaqueroB.Pub("mychan", "Hi, this is VaqueroB")

	for {
		result := <-fin
		assert(t, "Hi, this is VaqueroB", result)
		return
	}
}
