package redis

import "strings"
import "fmt"
import "regexp"

// CommandZcount provides TCP communication of `ZCOUNT`.
type CommandZcount struct {
	key string
	min string
	max string
	CommandDefault
}

// Build builds TCP message by initialized parameters.
func (cmd CommandZcount) Build() []byte {
	words := []string{
		"*4",
		cmd.getLenStr(cmdZCOUNT),
		cmdZCOUNT,
		cmd.getLenStr(cmd.key),
		cmd.key,
		cmd.getLenStr(cmd.min),
		cmd.min,
		cmd.getLenStr(cmd.max),
		cmd.max,
	}
	joined := strings.Join(words, sep) + sep
	return []byte(joined)
}

// Parse parses TCP response.
func (cmd CommandZcount) Parse(res []byte) (result string, e error) {
	re := regexp.MustCompile(":([0-9]+)")
	if matches := re.FindStringSubmatch(string(res)); len(matches) > 1 {
		result = matches[1]
		return
	}
	e = fmt.Errorf("Response to `ZCOUNT` is `%v`", string(res))
	return
}
