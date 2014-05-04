package rodeo

import "net"
import "errors"
import "strings"
import "strconv"

// 相手がredisであることを知っている
// redis特有の文字列整形を知っている
// redisのロジックはすべてここに閉じ込める
type RedisProtocol struct {
	message []byte
	Error   error
}

var (
	bulklen = "$"
	sep     = "\r\n"
)
var (
	CMD_GET = "GET"
	CMD_SET = "SET"
)

func (p *RedisProtocol) Request(args ...string) Protocol {
	if len(args) < 2 {
		p.Error = errors.New("Too short params for redis protocol")
		return p
	}
	switch args[0] {
	case CMD_GET:
		return p.generateGetMessage(args[1])
	}
	return p
}
func (p *RedisProtocol) Execute(conn net.Conn) Protocol {
	return p
}
func (p *RedisProtocol) ToResult() Result {
	return Result{}
}
func (p *RedisProtocol) generateGetMessage(key string) Protocol {
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
func (p *RedisProtocol) generateSetMessage() []byte {
	return []byte("bbb")
}
func (p *RedisProtocol) getLenStr(str string) string {
	return bulklen + strconv.Itoa(len(str))
}
func (p *RedisProtocol) GetMessage() string {
	return string(p.message)
}
