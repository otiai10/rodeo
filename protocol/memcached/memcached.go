package memcached

import "github.com/otiai10/rodeo/protocol"

import "net"
import "strconv"
import "bufio"
import "fmt"

// MemcachedProtocol knows the way to message TCP.
type MemcachedProtocol struct {
	message  []byte
	response []byte
	Command  Command
	Error    error
}

// Command interface.
type Command interface {
	Build() []byte
	Parse(res []byte) (string, error)
}

// CommandDefault defines default functionalities.
type CommandDefault struct{}

// TODO: change method name
func (d CommandDefault) strlen(str string) string {
	return strconv.Itoa(len(str))
}

var (
	sep     = " "
	suffix  = "\r\n"
	bufSize = 1024
	// FLAG is `flag` of memcached.
	FLAG   = "1001"
	cmdGET = "get"
	cmdSET = "set"
	setOK  = "STORED"
	// ErrorHeader is header of error messages.
	ErrorHeader = "MemcachedProtocol: "
)

// Request is interface to call commands.
// TODO: だから全てのメソッドがexportedじゃなくて良い気がする
func (p *MemcachedProtocol) Request(args ...string) protocol.Protocol {
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
	}
	e = fmt.Errorf("Command not found for `%s`", cmds[0])
	return
}

// Execute command.
func (p *MemcachedProtocol) Execute(conn net.Conn) protocol.Protocol {

	message := p.Command.Build()

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
func (p *MemcachedProtocol) WaitFor(conn net.Conn, reciever *chan string) {
}

// ToResult parses TCP response.
func (p *MemcachedProtocol) ToResult() (result protocol.Result) {
	res, _ := p.Command.Parse(p.response)
	return protocol.Result{Response: res}
}
func (p *MemcachedProtocol) isError(errMessage string) protocol.Protocol {
	p.Error = fmt.Errorf(ErrorHeader + errMessage)
	return p
}
