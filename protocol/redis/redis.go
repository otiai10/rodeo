package redis

import "github.com/otiai10/rodeo/protocol"

import "net"
import "errors"
import "strconv"
import "bufio"
import "fmt"

// 相手がredisであることを知っている
// redis特有の文字列整形を知っている
// redisのロジックはすべてここに閉じ込める
type RedisProtocol struct {
	message  []byte
	response []byte
	Command  string
	Error    error
}

var (
	marker_len = "$"
	marker_ss  = "+"
	sep        = "\r\n"
	buf_size   = 1024
)
var (
	CMD_GET = "GET"
	CMD_SET = "SET"
)
var (
	E_Header = "RedisProtocol: "
)

func (p *RedisProtocol) Request(args ...string) protocol.Protocol {
	lenArgs := len(args)
	if lenArgs < 2 {
		return p.isError("Too short params")
	}
	p.Command = args[0]
	switch p.Command {
	case CMD_GET:
		return p.generateGetMessage(args[1])
	case CMD_SET:
		if lenArgs < 3 {
			return p.isError("Too short params for `SET` command")
		}
		return p.generateSetMessage(args[1], args[2])
	}
	return p.isError(fmt.Sprintf("Command not found for `%s`", p.Command))
}
func (p *RedisProtocol) Execute(conn net.Conn) protocol.Protocol {

	if p.Error != nil {
		return p
	}

	tcpConnReader := bufio.NewReaderSize(conn, buf_size)

	fmt.Fprintf(conn, string(p.message))

	response := make([]byte, buf_size)
	_, rerr := tcpConnReader.Read(response)

	if rerr != nil {
		return p.isError(rerr.Error())
	}

	p.response = response
	return p
}
func (p *RedisProtocol) ToResult() (result protocol.Result) {
	switch p.Command {
	case CMD_GET:
		return p.generateGetResponse(p.response)
	case CMD_SET:
		return p.generateSetResponse(p.response)
	}
	return
}
func (p *RedisProtocol) getLenStr(str string) string {
	return marker_len + strconv.Itoa(len(str))
}
func (p *RedisProtocol) isError(errMessage string) protocol.Protocol {
	p.Error = errors.New(E_Header + errMessage)
	return p
}
func (p *RedisProtocol) GetMessage() string {
	return string(p.message)
}
