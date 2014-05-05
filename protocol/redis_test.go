package protocol_test

import "github.com/otiai10/rodeo/protocol"

import "testing"
import "reflect"
import "fmt"

import "net"

func TestRedisProtocol(t *testing.T) {

	redisProtocol := protocol.RedisProtocol{}

	if reflect.TypeOf(redisProtocol).String() != "protocol.RedisProtocol" {
		t.Fail()
		return
	}
}

func TestRedisProtocol_Request(t *testing.T) {

	var redisProtocol protocol.RedisProtocol
	var message string

	redisProtocol = protocol.RedisProtocol{}

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

	var redisProtocol protocol.RedisProtocol
	redisProtocol = protocol.RedisProtocol{}

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
func TestRedisProtocol_ToResult(t *testing.T) {

	conn, _ := net.Dial("tcp", "localhost:6379")

	var redisProtocol protocol.RedisProtocol
	redisProtocol = protocol.RedisProtocol{}

	result := redisProtocol.Request("SET", "mykey", "Hello!!").Execute(conn).ToResult()
	fmt.Printf("%+v", result)
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
