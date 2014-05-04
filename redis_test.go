package rodeo_test

import "github.com/otiai10/rodeo"

import "testing"
import "reflect"
import "fmt"

import "net"

func TestRedisProtocol(t *testing.T) {

	redisProtocol := rodeo.RedisProtocol{}

	if reflect.TypeOf(redisProtocol).String() != "rodeo.RedisProtocol" {
		t.Fail()
		return
	}
}

func TestRedisProtocol_Request(t *testing.T) {

	var redisProtocol rodeo.RedisProtocol
	var message string

	redisProtocol = rodeo.RedisProtocol{}

	_ = redisProtocol.Request("GET", "mykey")

	message = redisProtocol.GetMessage()
	if message != "*2\r\n$3\r\nGET\r\n$5\r\nmykey\r\n" {
		t.Fail()
		return
	}

	_ = redisProtocol.Request("SET", "mykey", "12345")

	message = redisProtocol.GetMessage()
	if message != "*3\r\n$3\r\nSET\r\n$5\r\nmykey\r\n$5\r\n12345\r\n" {
		t.Fail()
		return
	}
}

func TestRedisProtocol_Execute(t *testing.T) {

	conn, _ := net.Dial("tcp", "localhost:6379")

	var redisProtocol rodeo.RedisProtocol
	redisProtocol = rodeo.RedisProtocol{}

	_ = redisProtocol.Request("SET", "mykey", "Hello!!").Execute(conn)
	if redisProtocol.Error != nil {
		fmt.Println(redisProtocol.Error.Error())
		t.Fail()
		return
	}
	_ = redisProtocol.Request("GET", "mykey").Execute(conn)
	if redisProtocol.Error != nil {
		fmt.Println(redisProtocol.Error.Error())
		t.Fail()
		return
	}
}
