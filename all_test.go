package rodeo_test

import "github.com/otiai10/rodeo"

import . "github.com/otiai10/mint"
import "testing"
import "time"
import "fmt"

var conf = rodeo.Conf{
	Host: "localhost",
	Port: "6379",
}

type tStruct0 struct {
	Foo string
}
type tStruct1 struct {
	Foo Foo
}
type Foo struct {
	Bar string
	Buz int
}

func TestTheVaquero(t *testing.T) {

	vaquero, e := rodeo.TheVaquero(conf, "test")
	Expect(t, e).ToBe(nil)
	Expect(t, vaquero.Conf.Port).ToBe("6379")
}

func TestVaquero_Set(t *testing.T) {

	vaquero, e := rodeo.TheVaquero(conf, "test")

	e = vaquero.Set("mykey", "12345")
	Expect(t, e).ToBe(nil)
	val := vaquero.Get("mykey")
	Expect(t, val).ToBe("12345")

	e = vaquero.Set("mykey", "67890")
	Expect(t, e).ToBe(nil)

	val = vaquero.Get("mykey")
	Expect(t, val).ToBe("67890")

	e = vaquero.Delete("mykey")
	Expect(t, e).ToBe(nil)

	val = vaquero.Get("mykey")
	Expect(t, val).ToBe("")
}

func TestVaquero_Store(t *testing.T) {

	vaquero, e := rodeo.TheVaquero(conf, "test")

	key0 := "mykey0"
	obj0 := tStruct0{"Hello, rodeo"}
	e = vaquero.Store(key0, obj0)
	Expect(t, e).ToBe(nil)

	var dest0 tStruct0
	e = vaquero.Cast(key0, &dest0)
	Expect(t, e).ToBe(nil)
	Expect(t, "Hello, rodeo").ToBe(dest0.Foo)

	key1 := "mykey1"
	obj1 := tStruct1{Foo: Foo{Bar: "This is Bar", Buz: 1001}}
	e = vaquero.Store(key1, obj1)
	Expect(t, e).ToBe(nil)

	var dest1 tStruct1
	e = vaquero.Cast(key1, &dest1)
	Expect(t, nil).ToBe(e)
	Expect(t, dest1.Foo.Bar).ToBe("This is Bar")
	Expect(t, dest1.Foo.Buz).ToBe(1001)
}

func TestVaquero_PubSub(t *testing.T) {

	fin := make(chan string)

	vaqueroA, _ := rodeo.TheVaquero(conf, "test")
	vaqueroB, _ := rodeo.TheVaquero(conf, "test")

	subscriber := vaqueroA.Sub("mychan")

	go func() {
		for {
			message := <-subscriber
			fin <- message
			continue
		}
	}()

	time.Sleep(1 * time.Second)
	_ = vaqueroB.Pub("mychan", "Hi, this is VaqueroB 000")
	time.Sleep(1 * time.Second)
	_ = vaqueroB.Pub("mychan", "Hi, this is VaqueroB 001")
	time.Sleep(1 * time.Second)
	_ = vaqueroB.Pub("mychan", "Hi, this is VaqueroB 002")

	var count int
	for {
		result := <-fin
		switch count {
		case 0:
			Expect(t, result).ToBe("Hi, this is VaqueroB 000")
			break
		case 1:
			Expect(t, result).ToBe("Hi, this is VaqueroB 001")
			break
		case 2:
			Expect(t, result).ToBe("Hi, this is VaqueroB 002")
			return
		}
		count++
	}
}

type User struct {
	Name string
	Age  int
}

func (u *User) Greet() string {
	return fmt.Sprintf("Hi, I'm %s. %d years old.", u.Name, u.Age)
}
func TestVaquero_Tame(t *testing.T) {
	vaquero, _ := rodeo.TheVaquero(conf, "test")

	// truncate
	vaquero.Delete("test.users")

	users, e := vaquero.Tame("test.users", &User{})
	Expect(t, e).ToBe(nil)
	Expect(t, users).TypeOf("*rodeo.Group")

	Expect(t, users.Count()).ToBe(0)
}
