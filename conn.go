package rodeo

import "net"
import "fmt"
import "github.com/otiai10/rodeo/protocol/redis"

const serverFormat = "%s:%s"

func connect(host, port string) (facade pFacade, e error) {
	conn, e := net.Dial(
		"tcp",
		fmt.Sprintf(serverFormat, host, port),
	)
	facade = pFacade{
		conn,
		&redis.RedisProtocol{},
	}
	return
}
