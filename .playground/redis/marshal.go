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

	conn, err = net.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("Error on net.Dial:\t", err)
		return
	}
	reader = bufio.NewReaderSize(conn, 1024)

	fmt.Fprintf(conn, "*3\r\n$3\r\nSET\r\n$7\r\ntainaka\r\n$13\r\n{\"Hoge\":true}\r\n")

	// "SET mysample002 true\r\n" の結果の表示
	length, rerr = reader.Read(resp)

	info("SET", string(resp), err)
	fmt.Printf("READ STRING %v %v\n", length, rerr)

	fmt.Fprintf(conn, "*2\r\n$3\r\nGET\r\n$7\r\ntainaka\r\n")
	length, rerr = reader.Read(resp)

	info("GET", string(resp), err)
	fmt.Printf("READ STRING %v %v\n", length, rerr)
}
