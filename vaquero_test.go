package rodeo_test

import . "github.com/otiai10/rodeo"

import "fmt"
import "testing"
import "github.com/robfig/config"

var conf, _ = config.ReadDefault("sample.conf")

type tStruct0 struct {
	Foo bool
}

func TestVaquero_Set(t *testing.T) {

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
	obj0 := tStruct0{true}
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
	if dest0.Foo != true {
		fmt.Println(dest0)
		t.Fail()
		return
	}
}
func TestTheVaqueroFail00(t *testing.T) {
	conf, _ := config.ReadDefault("sample.conf")
	_, e := TheVaquero(conf, "missing")
	if e == nil || e.Error() != "option not found: port" {
		fmt.Println(e)
		t.Fail()
	}
}
func TestTheVaqueroFail01(t *testing.T) {
	conf, _ := config.ReadDefault("sample.conf")
	_, e := TheVaquero(conf, "notfound")
	if e == nil || e.Error() != "dial tcp: invalid port 99999" {
		fmt.Println(e)
		t.Fail()
	}
}
