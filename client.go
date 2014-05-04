package rodeo

import "net"
import "github.com/otiai10/rodeo/protocol"

// TcpClientって何すんだろう
// key string, val string のkvsすべてに対応できる
// Memcachedでも動くように設計する
type TcpClient struct {
	Conn    net.Conn
	Protcol protocol.Protocol
}

func (client *TcpClient) GetStringAnyway(key string) (value string) {
	value = "12345"
	return
}
func (client *TcpClient) Set(key string, value interface{}) (e error) {
	return
}
