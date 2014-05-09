package rodeo

import "net"
import "github.com/otiai10/rodeo/protocol"

// type pFacade
// convert types of key and value
// to use (string only) KVS.
// というか、pFacadeってexportしなくてよくない？？
// インターフェースじゃないじゃん。
type pFacade struct {
	Conn    net.Conn
	Protcol protocol.Protocol
}

func (fcd *pFacade) GetStringAnyway(key string) (value string) {
	result := fcd.Protcol.Request("GET", key).Execute(fcd.Conn).ToResult()
	return result.Response
}
func (fcd *pFacade) SetString(key string, value string) (e error) {
	result := fcd.Protcol.Request("SET", key, value).Execute(fcd.Conn).ToResult()
	return result.Error
}
