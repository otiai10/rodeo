package rodeo_test

import . "github.com/otiai10/rodeo"

import "fmt"
import "testing"
import "github.com/robfig/config"

var conf, _ = config.ReadDefault("sample.conf")

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

	key0 := "mykey0"
	obj0 := new(struct {
		Foo bool
	})
	obj0.Foo = true
	e = vaquero.Store(key0, obj0)
	if e != nil {
		fmt.Println(e)
		t.Fail()
		return
	}

	dest0 := new(struct {
		Foo bool
	})
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

	key1 := "mykey1"
	obj1 := new(struct {
		Bar bool
	})
	obj1.Bar = false
	e = vaquero.Store(key1, obj1)
	if e != nil {
		fmt.Println(e)
		t.Fail()
		return
	}

	dest1 := new(struct {
		Bar bool
	})
	e = vaquero.Cast(key1, &dest1)
	if e != nil {
		fmt.Println(e)
		t.Fail()
		return
	}
	if dest1.Bar != false {
		fmt.Println(dest1)
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
