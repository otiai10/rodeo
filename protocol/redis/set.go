package redis

import "fmt"
import "strings"
import "regexp"

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
	// TODO: DO NOT CODE IT HARD
	if ok, _ := regexp.Match("\\+OK", res); ok {
		// TODO: validate
		return "OK", nil
	}
	e = fmt.Errorf("Response to `SET` is not OK")
	return
}
