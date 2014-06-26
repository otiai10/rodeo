package redis

import "fmt"
import "strings"
import "regexp"

// CommandSet provides TCP communication of `SET`.
type CommandSet struct {
	key   string
	value string
	CommandDefault
}

// Build builds TCP message by initialized parameters.
func (cmd CommandSet) Build() []byte {
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

// Parse parses TCP response.
func (cmd CommandSet) Parse(res []byte) (result string, e error) {
	// TODO: DO NOT CODE IT HARD
	if ok, _ := regexp.Match("\\+OK", res); ok {
		// TODO: validate
		return "OK", nil
	}
	e = fmt.Errorf("Response to `SET` is not OK")
	return
}
