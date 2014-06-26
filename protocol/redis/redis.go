package redis

import "github.com/otiai10/rodeo/protocol"

import "net"
import "strconv"
import "bufio"
import "fmt"

// RedisProtocol knows the way to message TCP.
type RedisProtocol struct {
	message  []byte
	response []byte
	Command  Command
	Error    error
}

const (
	markerLength    = "$"
	markerNonExists = "$-1"
	sep             = "\r\n"
	bufSize         = 1024
	cmdGET          = "GET"
	cmdSET          = "SET"
	cmdDEL          = "DEL"
	cmdZADD         = "ZADD"
	cmdZCOUNT       = "ZCOUNT"
	cmdZRANGE       = "ZRANGE"
	cmdSUBSCRIBE    = "SUBSCRIBE"
	cmdPUBLISH      = "PUBLISH"
	// ErrorHeader is header of error messages.
	ErrorHeader = "RedisProtocol: "
)

// Command interface.
type Command interface {
	build() []byte
	parse(res []byte) (string, error)
}

// CommandDefault defines default functionalities.
type CommandDefault struct{}

// TODO: change method name
func (d CommandDefault) strlen(str string) string {
	return markerLength + strconv.Itoa(len(str))
}

// Request is interface to call commands.
// TODO: だから全てのメソッドがexportedじゃなくて良い気がする
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
	case cmdGET:
		return CommandGet{key: cmds[1]}, nil
	case cmdSET:
		return CommandSet{key: cmds[1], value: cmds[2]}, nil
	case cmdDEL:
		return CommandDel{key: cmds[1]}, nil
	case cmdSUBSCRIBE:
		return CommandSubscribe{chanName: cmds[1]}, nil
	case cmdPUBLISH:
		return CommandPublish{chanName: cmds[1], message: cmds[2]}, nil
	case cmdZADD:
		return CommandZadd{key: cmds[1], score: cmds[2], value: cmds[3]}, nil
	case cmdZCOUNT:
		return CommandZcount{key: cmds[1], min: cmds[2], max: cmds[3]}, nil
	case cmdZRANGE:
		return CommandZrange{key: cmds[1], start: cmds[2], stop: cmds[3], opt: cmds[4]}, nil
	}
	e = fmt.Errorf("Command not found for `%s`", cmds[0])
	return
}

// Execute command.
func (p *RedisProtocol) Execute(conn net.Conn) protocol.Protocol {

	if p.Error != nil {
		return p
	}

	message := p.Command.build()

	if p.Error != nil {
		return p
	}

	tcpConnReader := bufio.NewReaderSize(conn, bufSize)

	fmt.Fprintf(conn, string(message))

	response := make([]byte, bufSize)
	_, rerr := tcpConnReader.Read(response)

	if rerr != nil {
		return p.isError(rerr.Error())
	}

	p.response = response
	return p
}

// WaitFor is io waiter for pub/sub model.
func (p *RedisProtocol) WaitFor(conn net.Conn, reciever *chan string) {

	message := p.Command.build()

	tcpConnReader := bufio.NewReaderSize(conn, bufSize)

	go func() {
		fmt.Fprintf(conn, string(message))
		response := make([]byte, bufSize)
		for {
			_, _ = tcpConnReader.Read(response)
			res, e := p.Command.parse(response)

			if e != nil {
				continue
			}
			*reciever <- res
		}
	}()
}

// ToResult parses TCP response.
func (p *RedisProtocol) ToResult() (result protocol.Result) {
	if p.Error != nil {
		return protocol.Result{Error: p.Error}
	}
	res, _ := p.Command.parse(p.response)
	return protocol.Result{Response: res}
}
func (p *RedisProtocol) isError(errMessage string) protocol.Protocol {
	p.Error = fmt.Errorf(ErrorHeader + errMessage)
	return p
}
