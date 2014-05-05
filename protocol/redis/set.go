package redis

import "github.com/otiai10/rodeo/protocol"

import "errors"
import "strings"
import "regexp"

func (p *RedisProtocol) generateSetResponse(res []byte) protocol.Result {
	result := protocol.Result{}
	if ok, _ := regexp.Match("\\+OK", res); ok {
		result.Response = "OK"
		return result
	}
	result.Response = string(res)
	result.Error = errors.New("Response to `SET` is not OK")
	return result
}
func (p *RedisProtocol) generateSetMessage(key, value string) protocol.Protocol {
	words := []string{
		"*3",
		p.getLenStr(CMD_SET),
		CMD_SET,
		p.getLenStr(key),
		key,
		p.getLenStr(value),
		value,
	}
	joined := strings.Join(words, sep) + sep
	p.message = []byte(joined)
	return p
}
