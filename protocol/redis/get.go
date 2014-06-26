package redis

import "strings"
import "fmt"
import "regexp"

// CommandGet provides TCP communication of `GET`.
type CommandGet struct {
	key string
	commandDefault
}

func (cmd CommandGet) build() []byte {
	words := []string{
		"*2",
		cmd.strlen(cmdGET),
		cmdGET,
		cmd.strlen(cmd.key),
		cmd.key,
	}
	joined := strings.Join(words, sep) + sep
	return []byte(joined)
}

func (cmd CommandGet) parse(res []byte) (result string, e error) {
	// TODO: DO NOT CODE IT HARD

	if ok, _ := regexp.Match("\\$.+\\r\\n", res); ok {
		lines := strings.Split(string(res), "\r\n")
		if lines[0] == markerNonExists {
			return
		}
		// TODO: validate
		result = lines[1]
		return
	}
	e = fmt.Errorf("Response to `Get` is `%v`", string(res))
	return
}
