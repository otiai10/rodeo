package redis

import "strings"
import "net"
import "bufio"

// CommandSet provides TCP communication of `SET`.
type CommandSet struct {
	key   string
	value string
	commandDefault
}

func (cmd CommandSet) build() []byte {
	words := []string{
		"*3",
		cmd.strlen(cmdSET),
		cmdSET,
		cmd.strlen(cmd.key),
		cmd.key,
		cmd.strlen(cmd.value),
		cmd.value,
	}
	joined := strings.Join(words, sep) + sep
	return []byte(joined)
}

func (cmd CommandSet) parse(res []byte) (result string, e error) {
	return string(res), nil
}

func (cmd CommandSet) hoge(conn net.Conn) (res []byte) {
	scanner := bufio.NewScanner(conn)
	if ok := scanner.Scan(); !ok {
		return
	}
	if m := RESP["string"].FindSubmatch(scanner.Bytes()); len(m) > 1 {
		return m[1]
	}
	return
}
