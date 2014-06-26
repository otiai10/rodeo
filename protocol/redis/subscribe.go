package redis

import "strings"
import "fmt"
import "regexp"

// CommandSubscribe provides TCP communication of `SUBSCRIBE`.
type CommandSubscribe struct {
	chanName string
	CommandDefault
}

// Build builds TCP message by initialized parameters.
func (cmd CommandSubscribe) Build() []byte {
	words := []string{
		"*2",
		cmd.getLenStr(cmdSUBSCRIBE),
		cmdSUBSCRIBE,
		cmd.getLenStr(cmd.chanName),
		cmd.chanName,
	}
	joined := strings.Join(words, sep) + sep
	return []byte(joined)
}

// Parse parses TCP response.
func (cmd CommandSubscribe) Parse(res []byte) (result string, e error) {
	// TODO: DO NOT CODE IT HARD
	if ok, _ := regexp.Match("\\\\*.+\\r\\n", res); ok {
		lines := strings.Split(string(res), "\r\n")
		if lines[2] != "message" {
			e = fmt.Errorf("not message event")
			return
		}
		result = lines[6]
		return
	}
	e = fmt.Errorf("Response to `Get` is `%v`", string(res))
	return
}
