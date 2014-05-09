package rodeo

import "net"
import "fmt"
import "github.com/otiai10/rodeo/protocol/redis"

var f_location = "%s:%s"

func connect(host, port string) (facade pFacade, e error) {
	conn, e := net.Dial(
		"tcp",
		fmt.Sprintf(f_location, host, port),
	)
	facade = pFacade{
		conn,
		&redis.RedisProtocol{},
	}
	return
}
