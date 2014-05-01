package rodeo_test

import . "github.com/otiai10/rodeo"

import "testing"
import "fmt"
import "github.com/robfig/config"

func TestVaquero(t *testing.T) {
	conf, _ := config.ReadDefault("test.conf")
	vaquero, e := TheVaquero(conf)
	fmt.Println(vaquero, e)
}
