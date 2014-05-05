package protocol

import "net"
import "errors"
import "strings"
import "strconv"
import "fmt"
import "bufio"
import "regexp"

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

func (p *RedisProtocol) Request(args ...string) Protocol {
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
func (p *RedisProtocol) Execute(conn net.Conn) Protocol {

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
func (p *RedisProtocol) ToResult() (result Result) {
	switch p.Command {
	case CMD_GET:
		return p.generateGetResponse(p.response)
	case CMD_SET:
		return p.generateSetResponse(p.response)
	}
	return
}
func (p *RedisProtocol) generateGetResponse(res []byte) Result {
	result := Result{}
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
func (p *RedisProtocol) generateSetResponse(res []byte) Result {
	result := Result{}
	if ok, _ := regexp.Match("\\+OK", res); ok {
		result.Response = "OK"
		return result
	}
	result.Response = string(res)
	result.Error = errors.New("Response to `SET` is not OK")
	return result
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
	return marker_len + strconv.Itoa(len(str))
}
func (p *RedisProtocol) isError(errMessage string) Protocol {
	p.Error = errors.New(E_Header + errMessage)
	return p
}
func (p *RedisProtocol) GetMessage() string {
	return string(p.message)
}
