package redis

import "strings"
import "fmt"
import "regexp"
import "errors"

type CommandSubscribe struct {
	chanName string
	CommandDefault
}

func (this CommandSubscribe) Build() []byte {
	words := []string{
		"*2",
		this.getLenStr(CMD_SUBSCRIBE),
		CMD_SUBSCRIBE,
		this.getLenStr(this.chanName),
		this.chanName,
	}
	joined := strings.Join(words, sep) + sep
	return []byte(joined)
}
func (this CommandSubscribe) Parse(res []byte) (result string, e error) {
	// TODO: DO NOT CODE IT HARD
	if ok, _ := regexp.Match("\\\\*.+\\r\\n", res); ok {
		lines := strings.Split(string(res), "\r\n")
		if lines[2] != "message" {
			e = errors.New("not message event")
			return
		}
		result = lines[6]
		return
	}
	e = errors.New(
		fmt.Sprintf("Response to `Get` is `%v`", string(res)),
	)
	return
}
