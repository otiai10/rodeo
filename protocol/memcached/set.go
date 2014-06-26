package memcached

import "strings"
import "regexp"
import "fmt"

// CommandSet provides TCP communication of `set`.
type CommandSet struct {
	key   string
	value string
	CommandDefault
}

// Build builds TCP message by initialized parameters.
func (cmd CommandSet) Build() []byte {
	words := []string{
		cmdSET,
		cmd.key,
		FLAG,
		"0", // Expire
		cmd.strlen(cmd.value),
	}
	joined := strings.Join(words, sep) + suffix
	return []byte(joined + cmd.value + suffix)
}

// Parse parses TCP response.
func (cmd CommandSet) Parse(res []byte) (result string, e error) {
	// TODO: DO NOT CODE IT HARD
	if ok, _ := regexp.Match(setOK, res); ok {
		// TODO: validate
		return "OK", nil
	}
	e = fmt.Errorf("Response to `SET` is not OK")
	return
}
