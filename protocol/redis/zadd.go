package redis

import "regexp"
import "fmt"
import "strings"

type CommandZadd struct {
	key   string
	score string
	value string
	CommandDefault
}

func (this CommandZadd) Build() []byte {
	words := []string{
		"*4",
		this.getLenStr(CMD_ZADD),
		CMD_ZADD,
		this.getLenStr(this.key),
		this.key,
		this.getLenStr(this.score),
		this.score,
		this.getLenStr(this.value),
		this.value,
	}
	joined := strings.Join(words, sep) + sep
	return []byte(joined)
}

func (this CommandZadd) Parse(res []byte) (result string, e error) {
	// TODO: DO NOT CODE IT HARD
	if ok, _ := regexp.Match(":1", res); ok {
		// TODO: validate
		return "OK", nil
	}
	e = fmt.Errorf("Response to `ZADD` is not :1, but %s", string(res))
	return
}
