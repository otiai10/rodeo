package redis

import "strings"
import "fmt"
import "regexp"

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

	// TODO: DO NOT CODE IT HARD
	if ok, _ := regexp.Match("\\$.+\\r\\n", res); ok {
		lines := strings.Split(string(res), "\r\n")
		// TODO: validate
		result = lines[1]
		return
	}
	e = fmt.Errorf("Response to `ZREMRANGEBYSCORE` is `%v`", string(res))
	return
}
