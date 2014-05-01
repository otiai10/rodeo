package rodeo

import "github.com/robfig/config"

type Vaquero struct {
}

func TheVaquero(conf *config.Config) (v *Vaquero, e error) {
	v = &Vaquero{}
	return
}
