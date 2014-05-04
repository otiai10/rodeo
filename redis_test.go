package rodeo_test

import "github.com/otiai10/rodeo"

import "testing"
import "fmt"
import "reflect"

func TestRedisProtocol(t *testing.T) {

	redisProtocol := rodeo.RedisProtocol{}

	if reflect.TypeOf(redisProtocol).String() != "rodeo.RedisProtocol" {
		t.Fail()
		return
	}
}

func TestRedisProtocol_Request(t *testing.T) {

	redisProtocol := rodeo.RedisProtocol{}

	_ = redisProtocol.Request("GET", "mykey")

	message := redisProtocol.GetMessage()

	fmt.Println(message)
	if message != "*2\r\n$3\r\nGET\r\n$5\r\nmykey\r\n" {
		t.Fail()
		return
	}
}
