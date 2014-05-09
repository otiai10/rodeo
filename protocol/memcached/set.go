package memcached

import "errors"
import "strings"
import "regexp"

type CommandSet struct {
	key   string
	value string
	CommandDefault
}

func (this CommandSet) Build() []byte {
	words := []string{
		CMD_SET,
		this.key,
		FLAG,
		"0", // Expire
		this.getLenStr(this.value),
	}
	joined := strings.Join(words, sep) + suffix
	return []byte(joined + this.value + suffix)
}
func (this CommandSet) Parse(res []byte) (result string, e error) {
	// TODO: DO NOT CODE IT HARD
	if ok, _ := regexp.Match(SET_OK, res); ok {
		// TODO: validate
		return "OK", nil
	}
	e = errors.New("Response to `SET` is not OK")
	return
}
