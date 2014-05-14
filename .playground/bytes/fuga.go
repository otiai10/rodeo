package main

import "fmt"

func main() {
	var b byte
	b = 'a'
	fmt.Printf("%T %v", b, b)

	// b = 'ab'
	// fmt.Printf("%T %v", b, b)

	var bs []byte
	bs = []byte("This is a string")
	fmt.Printf("%T %v", bs, bs)
}
