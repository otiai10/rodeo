package rodeo

import "net"
import "fmt"

var f_location = "%s:%s"

func connect(host, port string) (client TcpClient, e error) {
	conn, e := net.Dial(
		"tcp",
		fmt.Sprintf(f_location, host, port),
	)
	client = TcpClient{
		conn,
	}
	return
}
