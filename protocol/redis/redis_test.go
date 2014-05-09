package redis_test

import "github.com/otiai10/rodeo/protocol/redis"

import "testing"
import "reflect"
import "fmt"

import "net"

func TestRedisProtocol(t *testing.T) {

	redisProtocol := redis.RedisProtocol{}

	if reflect.TypeOf(redisProtocol).String() != "redis.RedisProtocol" {
		fmt.Println("TypeError: ", reflect.TypeOf(redisProtocol).String())
		t.Fail()
		return
	}
}

func TestRedisProtocol_Execute(t *testing.T) {

	conn, _ := net.Dial("tcp", "localhost:6379")

	var redisProtocol redis.RedisProtocol
	redisProtocol = redis.RedisProtocol{}

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

	var redisProtocol redis.RedisProtocol
	redisProtocol = redis.RedisProtocol{}

	result := redisProtocol.Request("SET", "mykey", "Hello!!").Execute(conn).ToResult()
	if result.Response != "OK" {
		fmt.Printf("%+v", result)
		t.Fail()
		return
	}
	result = redisProtocol.Request("GET", "mykey").Execute(conn).ToResult()
	if result.Response != "Hello!!" {
		fmt.Printf("%+v", result)
		t.Fail()
		return
	}
}
