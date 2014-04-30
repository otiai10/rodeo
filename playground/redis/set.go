package main

import "net"
import "fmt"
import "bufio"

func info(cmd, resp string, e error) {
	fmt.Printf(
		"#####\t%v\nRESPO\t%v\nERROR\t%v\n",
		cmd,
		resp,
		e,
	)
}
func main() {

	var conn net.Conn
	var reader *bufio.Reader
	var resp []byte
	var err error
	var length int
	var rerr error

	resp = make([]byte, 1024)

	conn, _ = net.Dial("tcp", "localhost:6379")
	reader = bufio.NewReaderSize(conn, 1024)

	fmt.Fprintf(conn, "*3\r\n$3\r\nSET\r\n$7\r\ntainaka\r\n$5\r\nritsu\r\n")

	// "SET mysample002 true\r\n" の結果の表示
	length, rerr = reader.Read(resp)
	info("SET", string(resp), err)
	fmt.Printf("READ STRING %v %v\n", length, rerr)
}
