package redis

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
		"*3",
		this.getLenStr(CMD_SET),
		CMD_SET,
		this.getLenStr(this.key),
		this.key,
		this.getLenStr(this.value),
		this.value,
	}
	joined := strings.Join(words, sep) + sep
	return []byte(joined)
}
func (this CommandSet) Parse(res []byte) (result string, e error) {
	// TODO: DO NOT CODE IT HARD
	if ok, _ := regexp.Match("\\+OK", res); ok {
		// TODO: validate
		return "OK", nil
	}
	e = errors.New("Response to `SET` is not OK")
	return
}
