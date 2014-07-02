package rodeo

// Vaquero handles simple I/O and provide active model.
type Vaquero struct {
	Conf   Conf
	facade pFacade
}

// Conf is definitions of configuration.
// TODO: delete
type Conf struct {
	Host string
	Port string
}

// NewVaquero provide new Vaquero instance.
// TODO: change name to NewVaquero
func NewVaquero(conf Conf, args ...string) (v *Vaquero, e error) {
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

// Set can store string by key.
func (v *Vaquero) Set(key string, val string) (e error) {
	return v.facade.SetString(key, val)
}

// Get can find string by key.
func (v *Vaquero) Get(key string) (val string) {
	return v.facade.GetStringAnyway(key)
}

// Delete can delete string and object by key.
func (v *Vaquero) Delete(key string) (e error) {
	return v.facade.DeleteKey(key)
}

// Store can store object by key.
func (v *Vaquero) Store(key string, obj interface{}) (e error) {
	return v.facade.SetStruct(key, obj)
}

// Cast can find and map object by key.
func (v *Vaquero) Cast(key string, dest interface{}) (e error) {
	return v.facade.GetStruct(key, dest)
}

// Sub provides channel listening to given channel name.
func (v *Vaquero) Sub(chanName string) (ch chan string) {
	ch = make(chan string)
	v.facade.Listen(chanName, &ch)
	return
}

// Pub publishes message to given channel name.
func (v *Vaquero) Pub(chanName string, message string) (e error) {
	v.facade.Message(chanName, message)
	return
}
