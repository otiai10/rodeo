package rodeo_test

import "github.com/otiai10/rodeo"

func ExampleVaquero_Get() {
	vaquero, _ := rodeo.NewVaquero("localhost", "6379")
	v := vaquero.Get("mykey")
	println(v)
}
