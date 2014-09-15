package redis

// import "regexp"
// import "fmt"
import "strings"
import "net"
import "bufio"

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
	return string(res), e
}

func (cmd CommandZadd) hoge(conn net.Conn) (res []byte) {
	scanner := bufio.NewScanner(conn)
	if ok := scanner.Scan(); !ok {
		return
	}
	if m := RESP["int"].FindSubmatch(scanner.Bytes()); len(m) > 1 {
		return m[1]
	}
	return
}
