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

	obj := new(struct {
		Foo bool
	})
	obj.Foo = true
	e = vaquero.Store("mykey", obj)
	if e != nil {
		fmt.Println(e)
		t.Fail()
		return
	}

	dest := new(struct {
		Foo bool
	})
	e = vaquero.Cast("mykey", &dest)
	if e != nil {
		fmt.Println(e)
		t.Fail()
		return
	}
	// fmt.Printf("%T %+v", dest, dest)
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
