package memcached

import "strings"
import "fmt"
import "regexp"

// CommandGet provides TCP communication of `get`.
type CommandGet struct {
	key string
	CommandDefault
}

// Build builds TCP message by initialized parameters.
func (cmd CommandGet) Build() []byte {
	words := []string{
		cmdGET,
		cmd.key,
	}
	joined := strings.Join(words, sep) + suffix
	return []byte(joined)
}

// Parse parses TCP response.
func (cmd CommandGet) Parse(res []byte) (result string, e error) {
	// TODO: DO NOT CODE IT HARD
	if ok, _ := regexp.Match("\\r\\n", res); ok {
		lines := strings.Split(string(res), "\r\n")
		// TODO: validate
		result = lines[1]
		return
	}
	e = fmt.Errorf("Response to `Get` is `%v`", string(res))
	return
}
