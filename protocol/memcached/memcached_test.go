package memcached_test

import "github.com/otiai10/rodeo/protocol/memcached"

import . "github.com/otiai10/mint"
import "testing"
import "net"

func TestMemcachedProtocol(t *testing.T) {

	memcachedProtocol := memcached.MemcachedProtocol{}

	Expect(t, memcachedProtocol).TypeOf("memcached.MemcachedProtocol")
}

func TestMemcachedProtocol_Execute(t *testing.T) {

	conn, e := net.Dial("tcp", "localhost:11211")
	Expect(t, e).ToBe(nil)

	var memcachedProtocol memcached.MemcachedProtocol
	memcachedProtocol = memcached.MemcachedProtocol{}

	_ = memcachedProtocol.Request("set", "mykey", "Hello!!").Execute(conn)
	Expect(t, memcachedProtocol.Error).ToBe(nil)

	_ = memcachedProtocol.Request("get", "mykey").Execute(conn)
	Expect(t, memcachedProtocol.Error).ToBe(nil)
}
func TestMemcachedProtocol_ToResult(t *testing.T) {

	conn, _ := net.Dial("tcp", "localhost:11211")

	var memcachedProtocol memcached.MemcachedProtocol
	memcachedProtocol = memcached.MemcachedProtocol{}

	result := memcachedProtocol.Request("set", "mykey", "Hello!!").Execute(conn).ToResult()
	Expect(t, result.Response).ToBe("OK")

	result = memcachedProtocol.Request("get", "mykey").Execute(conn).ToResult()
	Expect(t, result.Response).ToBe("Hello!!")
}
