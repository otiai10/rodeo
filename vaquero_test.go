package rodeo_test

import "github.com/otiai10/rodeo"

import "fmt"
import "testing"
import "os"

import "time"

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
		fmt.Printf("`%+v` expected, but `%+v` actual.\n", expected, actual)
		t.Fail()
		os.Exit(1)
	}
}

func TestTheVaquero(t *testing.T) {

	vaquero, e := rodeo.TheVaquero(conf, "test")
	assert(t, e, nil)
	assert(t, vaquero.Conf.Port, "6379")
}

func TestVaquero_Set(t *testing.T) {

	vaquero, e := rodeo.TheVaquero(conf, "test")

	e = vaquero.Set("mykey", "12345")
	assert(t, e, nil)
	val := vaquero.Get("mykey")
	assert(t, val, "12345")

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
	assert(t, e, nil)

	var dest0 tStruct0
	e = vaquero.Cast(key0, &dest0)
	assert(t, e, nil)
	assert(t, "Hello, rodeo", dest0.Foo)

	key1 := "mykey1"
	obj1 := tStruct1{Foo: Foo{Bar: "This is Bar", Buz: 1001}}
	e = vaquero.Store(key1, obj1)
	assert(t, e, nil)

	var dest1 tStruct1
	e = vaquero.Cast(key1, &dest1)
	assert(t, nil, e)
	assert(t, dest1.Foo.Bar, "This is Bar")
	assert(t, dest1.Foo.Buz, 1001)
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
			continue
		}
	}()

	time.Sleep(1 * time.Second)
	_ = vaqueroB.Pub("mychan", "Hi, this is VaqueroB 000")
	time.Sleep(1 * time.Second)
	_ = vaqueroB.Pub("mychan", "Hi, this is VaqueroB 001")
	time.Sleep(1 * time.Second)
	_ = vaqueroB.Pub("mychan", "Hi, this is VaqueroB 002")

	var count int
	for {
		result := <-fin
		switch count {
		case 0:
			assert(t, result, "Hi, this is VaqueroB 000")
			break
		case 1:
			assert(t, result, "Hi, this is VaqueroB 001")
			break
		case 2:
			assert(t, result, "Hi, this is VaqueroB 002")
			return
		}
		count++
	}
}
