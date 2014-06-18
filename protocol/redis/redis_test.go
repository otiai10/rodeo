package redis_test

import "github.com/otiai10/rodeo/protocol/redis"

import . "github.com/otiai10/mint"
import "testing"
import "net"

func TestRedisProtocol(t *testing.T) {

	redisProtocol := redis.RedisProtocol{}

	Expect(t, redisProtocol).TypeOf("redis.RedisProtocol")
}

func TestRedisProtocol_Execute(t *testing.T) {

	conn, e := net.Dial("tcp", "localhost:6379")
	Expect(t, e).ToBe(nil)

	var redisProtocol redis.RedisProtocol
	redisProtocol = redis.RedisProtocol{}

	_ = redisProtocol.Request("SET", "mykey", "Hello!!").Execute(conn)
	Expect(t, redisProtocol.Error).ToBe(nil)

	_ = redisProtocol.Request("GET", "mykey").Execute(conn)
	Expect(t, redisProtocol.Error).ToBe(nil)

	_ = redisProtocol.Request("DEL", "mykey").Execute(conn)
	Expect(t, redisProtocol.Error).ToBe(nil)

	result := redisProtocol.Request("ZADD", "hoge", "10000", "This is 10000").Execute(conn).ToResult()
	Expect(t, result.Error).ToBe(nil)
}
func TestRedisProtocol_ToResult(t *testing.T) {

	conn, _ := net.Dial("tcp", "localhost:6379")

	var redisProtocol redis.RedisProtocol
	redisProtocol = redis.RedisProtocol{}

	result := redisProtocol.Request("SET", "mykey", "Hello!!").Execute(conn).ToResult()
	Expect(t, result.Response).ToBe("OK")

	result = redisProtocol.Request("GET", "mykey").Execute(conn).ToResult()
	Expect(t, result.Response).ToBe("Hello!!")
}
