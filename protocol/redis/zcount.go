package redis

import "strings"
import "net"
import "bufio"

// CommandZcount provides TCP communication of `ZCOUNT`.
type CommandZcount struct {
	key string
	min string
	max string
	commandDefault
}

func (cmd CommandZcount) build() []byte {
	words := []string{
		"*4",
		cmd.strlen(cmdZCOUNT),
		cmdZCOUNT,
		cmd.strlen(cmd.key),
		cmd.key,
		cmd.strlen(cmd.min),
		cmd.min,
		cmd.strlen(cmd.max),
		cmd.max,
	}
	joined := strings.Join(words, sep) + sep
	return []byte(joined)
}

func (cmd CommandZcount) parse(res []byte) (result string, e error) {
	return string(res), e
}

func (cmd CommandZcount) hoge(conn net.Conn) (res []byte) {
	scanner := bufio.NewScanner(conn)
	if ok := scanner.Scan(); !ok {
		return
	}
	if m := RESP["int"].FindSubmatch(scanner.Bytes()); len(m) > 1 {
		return m[1]
	}
	return
}
