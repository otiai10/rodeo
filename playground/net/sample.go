package main

import "net"

func main() {
	println("01")
	conn, e := net.Dial("tcp", "localhost:6379")
	println("02", conn, e)
	var buff = make([]byte, 1024)
	println("03", conn, e)

	var cmd string = "*2\r\n$9\r\nSUBSCRIBE\r\n$6\r\nmychan\r\n"
	n, e := conn.Write([]byte(cmd))
	println(n, e, string(buff))

	for {
		n, e = conn.Read(buff)
		println("04", conn, e)
		println(n, e, string(buff))
	}
}
