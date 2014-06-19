package redis

import "strings"
import "fmt"
import "regexp"

type CommandZcount struct {
	key string
	min string
	max string
	CommandDefault
}

func (this CommandZcount) Build() []byte {
	words := []string{
		"*4",
		this.getLenStr(CMD_ZCOUNT),
		CMD_ZCOUNT,
		this.getLenStr(this.key),
		this.key,
		this.getLenStr(this.min),
		this.min,
		this.getLenStr(this.max),
		this.max,
	}
	joined := strings.Join(words, sep) + sep
	return []byte(joined)
}
func (this CommandZcount) Parse(res []byte) (result string, e error) {
	re := regexp.MustCompile(":([0-9]+)")
	if matches := re.FindStringSubmatch(string(res)); len(matches) > 1 {
		result = matches[1]
		return
	}
	e = fmt.Errorf("Response to `ZCOUNT` is `%v`", string(res))
	return
}
