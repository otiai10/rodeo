package rodeo_test

import . "github.com/otiai10/rodeo"

import "testing"
import "fmt"

import "net"
import "reflect"
import "strconv"

func TestTcpClient(t *testing.T) {

	client := getTestClient()

	if reflect.TypeOf(client).String() != "rodeo.TcpClient" {
		fmt.Println(reflect.TypeOf(client).String())
		t.Fail()
		return
	}

	var key string = "mykey"
	var val int = 12345

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
	if gotVal != strconv.Itoa(val) {
		fmt.Printf(
			"`%s` got for key `%s` is not `%s`",
			gotVal, key, val,
		)
		t.Fail()
		return
	}
}

func getTestClient() TcpClient {
	var client TcpClient
	conn, _ := net.Dial(
		"tcp",
		"localhost:6379",
	)
	client = TcpClient{
		conn,
		&RedisProtocol{},
	}
	return client
}
