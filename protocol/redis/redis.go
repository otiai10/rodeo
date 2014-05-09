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
	Command  Command
	Error    error
}

type Command interface {
	Build() []byte
	Parse(res []byte) (string, error)
}
type CommandDefault struct{}

func (d CommandDefault) getLenStr(str string) string {
	return marker_len + strconv.Itoa(len(str))
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
	command, e := getCommand(args)
	if e != nil {
		return p.isError(e.Error())
	}
	p.Command = command
	return p
}
func getCommand(cmds []string) (command Command, e error) {
	switch cmds[0] {
	case CMD_GET:
		return CommandGet{key: cmds[1]}, nil
	case CMD_SET:
		return CommandSet{key: cmds[1], value: cmds[2]}, nil
	}
	e = errors.New(fmt.Sprintf("Command not found for `%s`", cmds[0]))
	return
}
func (p *RedisProtocol) Execute(conn net.Conn) protocol.Protocol {

	message := p.Command.Build()

	if p.Error != nil {
		return p
	}

	tcpConnReader := bufio.NewReaderSize(conn, buf_size)

	fmt.Fprintf(conn, string(message))

	response := make([]byte, buf_size)
	_, rerr := tcpConnReader.Read(response)

	if rerr != nil {
		return p.isError(rerr.Error())
	}

	p.response = response
	return p
}
func (p *RedisProtocol) ToResult() (result protocol.Result) {
	res, _ := p.Command.Parse(p.response)
	return protocol.Result{Response: res}
}
func (p *RedisProtocol) isError(errMessage string) protocol.Protocol {
	p.Error = errors.New(E_Header + errMessage)
	return p
}