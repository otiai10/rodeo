package redis

import "regexp"
import "fmt"
import "strings"
import "net"

// CommandZadd provides TCP communication of `ZADD`.
type CommandZadd struct {
	key   string
	score string
	value string
	commandDefault
}

func (cmd CommandZadd) build() []byte {
	words := []string{
		"*4",
		cmd.strlen(cmdZADD),
		cmdZADD,
		cmd.strlen(cmd.key),
		cmd.key,
		cmd.strlen(cmd.score),
		cmd.score,
		cmd.strlen(cmd.value),
		cmd.value,
	}
	joined := strings.Join(words, sep) + sep
	return []byte(joined)
}

func (cmd CommandZadd) parse(res []byte) (result string, e error) {
	// TODO: DO NOT CODE IT HARD
	if ok, _ := regexp.Match(":1", res); ok {
		// TODO: validate
		return "OK", nil
	}
	e = fmt.Errorf("Response to `ZADD` is not :1, but %s", string(res))
	return
}

func (cmd CommandZadd) hoge(conn net.Conn) (res []byte) {
	return
}
