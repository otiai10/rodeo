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

func assert(t *testing.T, actual interface{}, expected interface{}) {
	if expected != actual {
		fmt.Printf("Expected `%+v`, but Actual `%+v`\n", expected, actual)
		t.Fail()
	}
}

func TestTheVaquero(t *testing.T) {

	vaquero, e := rodeo.TheVaquero(conf, "test")

	if e != nil {
		fmt.Println(e)
		t.Fail()
		return
	}
	if vaquero.Conf.Port != "6379" {
		t.Fail()
		return
	}
}

func TestVaquero_Set(t *testing.T) {

	vaquero, e := rodeo.TheVaquero(conf, "test")

	e = vaquero.Set("mykey", "12345")
	if e != nil {
		fmt.Printf("E?? %+v", e)
		t.Fail()
		return
	}

	val := vaquero.Get("mykey")
	if val != "12345" {
		fmt.Printf("val?? %+v", val)
		t.Fail()
		return
	}

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
	if e != nil {
		fmt.Println(e)
		t.Fail()
		return
	}

	var dest0 tStruct0
	e = vaquero.Cast(key0, &dest0)
	if e != nil {
		fmt.Println(e)
		t.Fail()
		return
	}
	if dest0.Foo != "Hello, rodeo" {
		fmt.Println(dest0)
		t.Fail()
		return
	}
}
