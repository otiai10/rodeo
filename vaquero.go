package rodeo

import "github.com/robfig/config"

type Vaquero struct {
	Conf C // should not be exported?
}

type C struct {
	Host string
	Port string
}

func TheVaquero(conf *config.Config, args ...string) (v *Vaquero, e error) {
	var c C
	c, e = ensureConf(conf, args)
	v = &Vaquero{
		c,
	}
	return
}

func ensureConf(conf *config.Config, args []string) (c C, e error) {
	var host, port string
	host, e = conf.String("test", "host")
	port, e = conf.String("test", "port")
	c = C{
		host,
		port,
	}
	return
}
