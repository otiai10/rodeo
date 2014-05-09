package rodeo

import "encoding/json"

import "fmt"

// type `Vaquero`
// manages configuration and connection,
// and gives interface to access storage.
type Vaquero struct {
	Conf   Conf // should not be exported?
	client TcpClient
}

type Conf struct {
	Host string
	Port string
}

func TheVaquero(conf Conf, args ...string) (v *Vaquero, e error) {
	var client TcpClient
	client, e = connect(conf.Host, conf.Port)
	if e != nil {
		return
	}
	v = &Vaquero{
		conf,
		client,
	}
	return
}

func (v *Vaquero) Set(key string, val interface{}) (e error) {
	return
}
func (v *Vaquero) Get(key string) (val string) {
	return "12345"
}
func (v *Vaquero) Store(key string, obj interface{}) (e error) {
	var bs []byte
	bs, e = json.Marshal(obj)
	// debug
	fmt.Printf("%T > %v\n", string(bs), string(bs))
	return
}
func (v *Vaquero) Cast(key string, dest interface{}) (e error) {
	return json.Unmarshal([]byte("{\"Foo\":\"Hello, rodeo\"}"), dest)
}
