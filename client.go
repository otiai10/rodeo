package rodeo

import "net"

type TcpClient struct {
	conn net.Conn
}
