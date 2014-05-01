package rodeo_test

import . "github.com/otiai10/rodeo"

import "testing"
import "github.com/robfig/config"

var conf, _ = config.ReadDefault("sample.conf")

func TestTheVaquero(t *testing.T) {

	vaquero, e := TheVaquero(conf, "test")

	if e != nil {
		t.Fail()
	}
	if vaquero.Conf.Port != "6379" {
		t.Fail()
	}
}
func TestTheVaqueroFail00(t *testing.T) {
	conf, _ := config.ReadDefault("sample.conf")
	_, e := TheVaquero(conf, "missing")
	if e.Error() != "option not found: port" {
		t.Fail()
	}
}
