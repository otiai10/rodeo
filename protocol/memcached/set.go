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

func (cmd CommandSet) build() []byte {
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

func (cmd CommandSet) parse(res []byte) (result string, e error) {
	// TODO: DO NOT CODE IT HARD
	if ok, _ := regexp.Match(setOK, res); ok {
		// TODO: validate
		return "OK", nil
	}
	e = fmt.Errorf("Response to `SET` is not OK")
	return
}
