package memcached_test

import "github.com/otiai10/rodeo/protocol/memcached"

import "testing"
import "reflect"
import "fmt"

import "net"

func TestMemcachedProtocol(t *testing.T) {

	memcachedProtocol := memcached.MemcachedProtocol{}

	if reflect.TypeOf(memcachedProtocol).String() != "memcached.MemcachedProtocol" {
		fmt.Println("TypeError: ", reflect.TypeOf(memcachedProtocol).String())
		t.Fail()
		return
	}
}

func TestMemcachedProtocol_Execute(t *testing.T) {

	conn, _ := net.Dial("tcp", "localhost:11211")

	var memcachedProtocol memcached.MemcachedProtocol
	memcachedProtocol = memcached.MemcachedProtocol{}

	_ = memcachedProtocol.Request("set", "mykey", "Hello!!").Execute(conn)
	if memcachedProtocol.Error != nil {
		fmt.Println(memcachedProtocol.Error.Error())
		t.Fail()
		return
	}

	_ = memcachedProtocol.Request("get", "mykey").Execute(conn)
	if memcachedProtocol.Error != nil {
		fmt.Println(memcachedProtocol.Error.Error())
		t.Fail()
		return
	}
}
func TestMemcachedProtocol_ToResult(t *testing.T) {

	conn, _ := net.Dial("tcp", "localhost:11211")

	var memcachedProtocol memcached.MemcachedProtocol
	memcachedProtocol = memcached.MemcachedProtocol{}

	result := memcachedProtocol.Request("set", "mykey", "Hello!!").Execute(conn).ToResult()
	if result.Response != "OK" {
		fmt.Printf("Expected `OK`, Actual `%+v`", result)
		t.Fail()
		return
	}
	result = memcachedProtocol.Request("get", "mykey").Execute(conn).ToResult()
	if result.Response != "Hello!!" {
		fmt.Printf("Expected `Hello!!`, Actual `%+v`", result.Response)
		t.Fail()
		return
	}
}
