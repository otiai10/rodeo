package rodeo

import "net"
import "errors"
import "strings"
import "strconv"
import "fmt"
import "bufio"

// 相手がredisであることを知っている
// redis特有の文字列整形を知っている
// redisのロジックはすべてここに閉じ込める
type RedisProtocol struct {
	message []byte
	Error   error
}

var (
	bulklen  = "$"
	sep      = "\r\n"
	buf_size = 1024
)
var (
	CMD_GET = "GET"
	CMD_SET = "SET"
)
var (
	E_Header = "RedisProtocol: "
)

func (p *RedisProtocol) Request(args ...string) Protocol {
	lenArgs := len(args)
	if lenArgs < 2 {
		return p.isError("Too short params")
	}
	switch args[0] {
	case CMD_GET:
		return p.generateGetMessage(args[1])
	case CMD_SET:
		if lenArgs < 3 {
			return p.isError("Too short params for `SET` command")
		}
		return p.generateSetMessage(args[1], args[2])
	}
	return p.isError(fmt.Sprintf("Command not found for `%s`", args[0]))
}
func (p *RedisProtocol) Execute(conn net.Conn) Protocol {

	if p.Error != nil {
		return p
	}

	tcpConnReader := bufio.NewReaderSize(conn, buf_size)

	fmt.Fprintf(conn, string(p.message))

	responseBuffer := make([]byte, buf_size)
	_, rerr := tcpConnReader.Read(responseBuffer)

	if rerr != nil {
		return p.isError(rerr.Error())
	}
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
func (p *RedisProtocol) generateSetMessage(key, value string) Protocol {
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
func (p *RedisProtocol) getLenStr(str string) string {
	return bulklen + strconv.Itoa(len(str))
}
func (p *RedisProtocol) isError(errMessage string) Protocol {
	p.Error = errors.New(E_Header + errMessage)
	return p
}
func (p *RedisProtocol) GetMessage() string {
	return string(p.message)
}
