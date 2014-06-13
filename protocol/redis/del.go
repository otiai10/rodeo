package redis

import "errors"
import "strings"
import "fmt"
import "regexp"

type CommandDel struct {
	key string
	CommandDefault
}

// TODO: accept multi keys
func (this CommandDel) Build() []byte {
	words := []string{
		"*2",
		this.getLenStr(CMD_DEL),
		CMD_DEL, // TODO: DRY
		this.getLenStr(this.key),
		this.key,
	}
	joined := strings.Join(words, sep) + sep
	return []byte(joined)
}
func (this CommandDel) Parse(res []byte) (result string, e error) {

	// TODO: DO NOT CODE IT HARD
	if ok, _ := regexp.Match("\\$.+\\r\\n", res); ok {
		lines := strings.Split(string(res), "\r\n")
		// TODO: validate
		result = lines[1]
		return
	}
	e = errors.New(
		fmt.Sprintf("Response to `DEL` is `%v`", string(res)),
	)
	return
}
