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
	value = "12345"
	return
}
func (fcd *pFacade) Set(key string, value string) (e error) {
	return
}
