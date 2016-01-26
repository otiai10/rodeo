package redis

import (
	"net"
	"testing"

	. "github.com/otiai10/mint"
)

var conn net.Conn

func init() {
	c, err := net.Dial("tcp", "localhost:6379")
	if err != nil {
		panic(err)
	}
	conn = c
}

func TestPing_Send(t *testing.T) {
	ping := Ping{}
	Expect(t, ping.Send(conn)).ToBe([]byte("+PONG"))
}
