package rodeo_test

import "github.com/otiai10/rodeo"

import . "github.com/otiai10/mint"
import "testing"
import "time"
import "fmt"

var host = "localhost"
var port = "6379"

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

func TestNewVaquero(t *testing.T) {

	vaquero, e := rodeo.NewVaquero(host, port, "test")
	Expect(t, e).ToBe(nil)
	Expect(t, vaquero.Conf.Port).ToBe("6379")
}

func TestVaquero_Set(t *testing.T) {

	vaquero, e := rodeo.NewVaquero(host, port, "test")

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

	vaquero, e := rodeo.NewVaquero(host, port, "test")

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

	vaqueroA, _ := rodeo.NewVaquero(host, port, "test")
	vaqueroB, _ := rodeo.NewVaquero(host, port, "test")

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
	vaquero, _ := rodeo.NewVaquero(host, port, "test")

	// truncate
	vaquero.Delete("test.users")

	users, e := vaquero.Tame("test.users", User{})

	Expect(t, e).ToBe(nil)
	Expect(t, users).TypeOf("*rodeo.Group")
	count, e := users.Count()
	Expect(t, e).ToBe(nil)
	Expect(t, count).ToBe(0)

	u0 := &User{"Mary", 28}
	u1 := &User{"John", 24}
	u2 := &User{"Steve", 32}
	u3 := &User{"Anne", 10}

	users.Add(int64(u0.Age), u0)
	users.Add(int64(u1.Age), u1)
	users.Add(int64(u2.Age), u2)
	users.Add(int64(u3.Age), u3)

	count, e = users.Count()
	Expect(t, e).ToBe(nil)
	Expect(t, count).ToBe(4)

	elms := users.Range() // find all
	Expect(t, elms).TypeOf("[]*rodeo.Element")
	Expect(t, len(elms)).ToBe(4)

	Expect(t, elms[0].Retrieve()).TypeOf("*rodeo_test.User")
	anne := elms[0].Retrieve().(*User)
	Expect(t, anne).Deeply().ToBe(u3)
	Expect(t, anne.Greet()).ToBe("Hi, I'm Anne. 10 years old.")
}
