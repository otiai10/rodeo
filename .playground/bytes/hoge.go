package main

import "fmt"

func main() {

	str := "single byte string\r\nABC"
	myBytes := []byte(str)

	fmt.Printf("%T\n", str)     // strig
	fmt.Printf("%T\n", myBytes) // []uint8

	for i, v := range myBytes {
		fmt.Printf("%v\t[%v]=[%v]\n", i, v, string(v))
		// 1バイトで表現できるので文字が出て来る
	}

	str = "マルチバイト"
	myBytes = []byte(str)

	for i, v := range myBytes {
		fmt.Printf("%v\t[%v]=[%v]\n", i, v, string(v))
		// 1バイトずつではことばにならない...
	}
}
