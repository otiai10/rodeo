package memcached

import "errors"
import "strings"
import "fmt"
import "regexp"

type CommandGet struct {
	key string
	CommandDefault
}

func (this CommandGet) Build() []byte {
	words := []string{
		CMD_GET,
		this.key,
	}
	joined := strings.Join(words, sep) + suffix
	return []byte(joined)
}
func (this CommandGet) Parse(res []byte) (result string, e error) {
	// TODO: DO NOT CODE IT HARD
	if ok, _ := regexp.Match("\\r\\n", res); ok {
		lines := strings.Split(string(res), "\r\n")
		// TODO: validate
		result = lines[1]
		return
	}
	e = errors.New(
		fmt.Sprintf("Response to `Get` is `%v`", string(res)),
	)
	return
}
