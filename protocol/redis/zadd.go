package redis

import "regexp"
import "fmt"
import "strings"

// CommandZadd provides TCP communication of `ZADD`.
type CommandZadd struct {
	key   string
	score string
	value string
	CommandDefault
}

// Build builds TCP message by initialized parameters.
func (cmd CommandZadd) Build() []byte {
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

// Parse parses TCP response.
func (cmd CommandZadd) Parse(res []byte) (result string, e error) {
	// TODO: DO NOT CODE IT HARD
	if ok, _ := regexp.Match(":1", res); ok {
		// TODO: validate
		return "OK", nil
	}
	e = fmt.Errorf("Response to `ZADD` is not :1, but %s", string(res))
	return
}
