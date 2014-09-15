package redis

import "strings"

// import "fmt"
// import "regexp"
import "net"
import "bufio"

// CommandDel provides TCP communication of `DEL`.
type CommandZRemRangeByScore struct {
	key string
	min string
	max string
	commandDefault
}

func (cmd CommandZRemRangeByScore) build() []byte {
	words := []string{
		"*4",
		cmd.strlen(cmdZREMRANGEBYSCORE),
		cmdZREMRANGEBYSCORE, // TODO: DRY
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

func (cmd CommandZRemRangeByScore) parse(res []byte) (result string, e error) {
	return string(res), e
}

func (cmd CommandZRemRangeByScore) hoge(conn net.Conn) (res []byte) {
	scanner := bufio.NewScanner(conn)
	if ok := scanner.Scan(); !ok {
		return
	}
	if m := RESP["int"].FindSubmatch(scanner.Bytes()); len(m) > 1 {
		return m[1]
	}
	return
}
