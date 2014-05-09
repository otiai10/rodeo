package rodeo_test

import . "github.com/otiai10/rodeo"

import "fmt"
import "testing"

var conf = Conf{
	Host: "localhost",
	Port: "6379",
}

type tStruct0 struct {
	Foo string
}

func TestTheVaquero(t *testing.T) {

	vaquero, e := TheVaquero(conf, "test")

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

	vaquero, e := TheVaquero(conf, "test")

	e = vaquero.Set("mykey", 12345)
	if e != nil {
		fmt.Println(e)
		t.Fail()
		return
	}

	val := vaquero.Get("mykey")
	if val != "12345" {
		fmt.Println(val)
		t.Fail()
		return
	}
}

func TestVaquero_Store(t *testing.T) {

	vaquero, e := TheVaquero(conf, "test")

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
