package redis

import "github.com/otiai10/rodeo/protocol"

import "errors"
import "strings"
import "fmt"
import "regexp"

func (p *RedisProtocol) generateGetResponse(res []byte) protocol.Result {
	result := protocol.Result{}
	if ok, _ := regexp.Match("\\$.\\r\\n", res); ok {
		lines := strings.Split(string(res), "\r\n")
		result.Response = lines[1]
		return result
	}
	result.Response = string(res)
	result.Error = errors.New(
		fmt.Sprintf("Response to `Get` is `%v`", string(res)),
	)
	return result
}
func (p *RedisProtocol) generateGetMessage(key string) protocol.Protocol {
	words := []string{
		"*2",
		p.getLenStr(CMD_GET),
		CMD_GET,
		p.getLenStr(key),
		key,
	}
	joined := strings.Join(words, sep) + sep
	p.message = []byte(joined)
	return p
}
