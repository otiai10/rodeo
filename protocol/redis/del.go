package redis

import "strings"
import "net"
import "bufio"

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
	return string(res), nil
}

func (cmd CommandDel) hoge(conn net.Conn) (res []byte) {
	scanner := bufio.NewScanner(conn)
	if ok := scanner.Scan(); !ok {
		return
	}
	if m := RESP["int"].FindSubmatch(scanner.Bytes()); len(m) > 1 {
		return m[1]
	}
	return
}
