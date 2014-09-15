package redis

import "strings"
import "fmt"
import "regexp"
import "net"

// CommandDel provides TCP communication of `DEL`.
type CommandDel struct {
	key string
	commandDefault
}

func (cmd CommandDel) build() []byte {
	words := []string{
		"*2",
		cmd.strlen(cmdDEL),
		cmdDEL, // TODO: DRY
		cmd.strlen(cmd.key),
		cmd.key,
	}
	joined := strings.Join(words, sep) + sep
	return []byte(joined)
}

func (cmd CommandDel) parse(res []byte) (result string, e error) {

	// TODO: DO NOT CODE IT HARD
	if ok, _ := regexp.Match("\\$.+\\r\\n", res); ok {
		lines := strings.Split(string(res), "\r\n")
		// TODO: validate
		result = lines[1]
		return
	}
	e = fmt.Errorf("Response to `DEL` is `%v`", string(res))
	return
}

func (cmd CommandDel) hoge(conn net.Conn) (res []byte) {
	return
}
