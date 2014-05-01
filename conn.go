package rodeo

import "net"

func connect(host, port string) (conn net.Conn, e error) {
	conn, e = net.Dial("tcp", host+":"+port)
	return
}
