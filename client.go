package rodeo

import "net"

type TcpClient struct {
	Conn net.Conn
}

func (client *TcpClient) GetStringAnyway(key string) (value string) {
	value = "12345"
	return
}
func (client *TcpClient) Set(key string, value interface{}) (e error) {
	return
}
