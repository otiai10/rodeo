package rodeo

import "net"

// 相手がredisであることを知っている
// redis特有の文字列整形を知っている
// redisのロジックはすべてここに閉じ込める
type RedisProtocol struct{}

func (p *RedisProtocol) Request(args ...interface{}) Protocol {
	return p
}
func (p *RedisProtocol) Execute(conn net.Conn) Protocol {
	return p
}
func (p *RedisProtocol) ToResult() Result {
	return Result{}
}
func (p *RedisProtocol) generateGetMessage() string {
	return ""
}
func (p *RedisProtocol) generateSetMessage() string {
	return ""
}
