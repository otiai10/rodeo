package redis

import "strings"
import "fmt"
import "regexp"

// CommandDel provides TCP communication of `DEL`.
type CommandDel struct {
	key string
	CommandDefault
}

// Build builds TCP message by initialized parameters.
// TODO: accept multi keys
func (cmd CommandDel) Build() []byte {
	words := []string{
		"*2",
		cmd.getLenStr(cmdDEL),
		cmdDEL, // TODO: DRY
		cmd.getLenStr(cmd.key),
		cmd.key,
	}
	joined := strings.Join(words, sep) + sep
	return []byte(joined)
}

// Parse parses TCP response.
func (cmd CommandDel) Parse(res []byte) (result string, e error) {

	// TODO: DO NOT CODE IT HARD
	if ok, _ := regexp.Match("\\$.+\\r\\n", res); ok {
		lines := strings.Split(string(res), "\r\n")
		// TODO: validate
		result = lines[1]
		return
	}
	e = fmt.Errorf("Response to `DEL` is `%v`", string(res))
	return
}
