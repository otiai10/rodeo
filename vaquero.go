package rodeo

import "github.com/robfig/config"
import "net"

type Vaquero struct {
	Conf C // should not be exported?
	Conn net.Conn
}

type C struct {
	Host string
	Port string
}

func TheVaquero(conf *config.Config, args ...string) (v *Vaquero, e error) {
	var c C
	c, e = ensureConf(conf, args)
	if e != nil {
		return
	}
	var conn net.Conn
	conn, e = connect(c.Host, c.Port)
	if e != nil {
		return
	}
	v = &Vaquero{
		c,
		conn,
	}
	return
}

func ensureConf(conf *config.Config, args []string) (c C, e error) {

	if len(args) == 0 {
		return confDefault(conf)
	}
	var host, port string
	host, e = conf.String(args[0], "host")
	port, e = conf.String(args[0], "port")
	c = C{
		host,
		port,
	}
	return
}

func confDefault(conf *config.Config) (c C, e error) {
	var host, port string
	host, e = conf.String("default", "host")
	port, e = conf.String("default", "port")
	if e != nil {
		return
	}
	c = C{
		host,
		port,
	}
	return
}
