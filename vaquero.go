package rodeo

// type `Vaquero`
// manages configuration and connection,
// and gives interface to access storage.
type Vaquero struct {
	Conf   Conf
	facade pFacade
}

type Conf struct {
	Host string
	Port string
}

func TheVaquero(conf Conf, args ...string) (v *Vaquero, e error) {
	var facade pFacade
	facade, e = connect(conf.Host, conf.Port)
	if e != nil {
		return
	}
	v = &Vaquero{
		conf,
		facade,
	}
	return
}

var chanRegistry = make(map[string]chan string)

func (v *Vaquero) Set(key string, val string) (e error) {
	return v.facade.SetString(key, val)
}
func (v *Vaquero) Get(key string) (val string) {
	return v.facade.GetStringAnyway(key)
}
func (v *Vaquero) Store(key string, obj interface{}) (e error) {
	return v.facade.SetStruct(key, obj)
}
func (v *Vaquero) Cast(key string, dest interface{}) (e error) {
	return v.facade.GetStruct(key, dest)
}

func (v *Vaquero) Sub(chanName string) (ch chan string) {
	ch = make(chan string)
	chanRegistry[chanName] = ch
	v.facade.Listen(&ch)
	return
}
func (v *Vaquero) Pub(chanName string, message string) (e error) {
	ch := chanRegistry[chanName]
	v.facade.Message(&ch, message)
	return
}
