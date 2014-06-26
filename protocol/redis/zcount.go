package redis

import "strings"
import "fmt"
import "regexp"

// CommandZcount provides TCP communication of `ZCOUNT`.
type CommandZcount struct {
	key string
	min string
	max string
	CommandDefault
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
	re := regexp.MustCompile(":([0-9]+)")
	if matches := re.FindStringSubmatch(string(res)); len(matches) > 1 {
		result = matches[1]
		return
	}
	e = fmt.Errorf("Response to `ZCOUNT` is `%v`", string(res))
	return
}
