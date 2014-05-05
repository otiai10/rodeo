package rodeo_test

import "github.com/otiai10/rodeo"
import "github.com/otiai10/rodeo/protocol/redis"

import "testing"
import "fmt"

import "net"
import "reflect"

func TestTcpClient(t *testing.T) {

	client := getTestClient()

	if reflect.TypeOf(client).String() != "rodeo.TcpClient" {
		fmt.Println(reflect.TypeOf(client).String())
		t.Fail()
		return
	}

	var key string = "mykey"
	var val string = "12345"

	var e error
	e = client.Set(key, val)
	if e != nil {
		fmt.Println("Set error is not nil", e)
		t.Fail()
		return
	}

	gotVal := client.GetStringAnyway(key)
	if reflect.TypeOf(gotVal).String() != "string" {
		fmt.Println("Got type is not string ", reflect.TypeOf(gotVal).String())
		t.Fail()
		return
	}
	if gotVal != val {
		fmt.Printf(
			"`%s` got for key `%s` is not `%s`",
			gotVal, key, val,
		)
		t.Fail()
		return
	}
}

func getTestClient() rodeo.TcpClient {
	var client rodeo.TcpClient
	conn, _ := net.Dial(
		"tcp",
		"localhost:6379",
	)
	client = rodeo.TcpClient{
		conn,
		&redis.RedisProtocol{},
	}
	return client
}
