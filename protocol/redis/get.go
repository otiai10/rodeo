package redis

import "strings"
import "regexp" // TODO: DRY001
import "net"
import "bufio"
import "strconv"

// CommandGet provides TCP communication of `GET`.
type CommandGet struct {
	key string
	commandDefault
}

func (cmd CommandGet) build() []byte {
	words := []string{
		"*2",
		cmd.strlen(cmdGET),
		cmdGET,
		cmd.strlen(cmd.key),
		cmd.key,
	}
	joined := strings.Join(words, sep) + sep
	return []byte(joined)
}

func (cmd CommandGet) parse(res []byte) (result string, e error) {
	return string(res), nil
}

func (cmd CommandGet) hoge(conn net.Conn) (res []byte) {
	// WARN: Do not lock connection
	// TODO: Timeout
	scanner := bufio.NewScanner(conn)
	lenExp := regexp.MustCompile("\\$(-?[0-9]+)")
	for {
		if ok := scanner.Scan(); !ok {
			return
		}
		if m := lenExp.FindSubmatch(scanner.Bytes()); len(m) > 1 {
			num, _ := strconv.Atoi(string(m[1]))
			if num < 0 {
				// not found
				return
			}
			continue
		}
		return scanner.Bytes()
	}
	return
}
