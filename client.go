package rodeo

import "net"
import "github.com/otiai10/rodeo/protocol"

// type Client
// convert types of key and value
// to use (string only) KVS
type TcpClient struct {
	Conn    net.Conn
	Protcol protocol.Protocol
}

func (client *TcpClient) GetStringAnyway(key string) (value string) {
	value = "12345"
	return
}
func (client *TcpClient) Set(key string, value string) (e error) {
	return
}
