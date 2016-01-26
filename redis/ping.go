package redis

import (
	"bufio"
	"net"
)

// Ping ...
type Ping struct{}

// Send ...
func (ping Ping) Send(conn net.Conn) []byte {
	scanner := bufio.NewScanner(conn)
	conn.Write([]byte("PING\r\n"))
	scanner.Scan()
	return scanner.Bytes()
}
