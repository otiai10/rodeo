package rodeo

import "net"
import "github.com/otiai10/rodeo/protocol"
import "encoding/json"

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
func (fcd *pFacade) DeleteKey(key string) (e error) {
	result := fcd.Protcol.Request("DEL", key).Execute(fcd.Conn).ToResult()
	return result.Error
}
func (fcd *pFacade) GetStruct(key string, dest interface{}) (e error) {
	result := fcd.Protcol.Request("GET", key).Execute(fcd.Conn).ToResult()
	e = json.Unmarshal([]byte(result.Response), dest)
	return e
}
func (fcd *pFacade) SetStruct(key string, obj interface{}) (e error) {
	var bs []byte
	bs, e = json.Marshal(obj)
	if e != nil {
		return
	}
	result := fcd.Protcol.Request("SET", key, string(bs)).Execute(fcd.Conn).ToResult()
	return result.Error
}
func (fcd *pFacade) Listen(chanName string, ch *chan string) {
	fcd.Protcol.Request("SUBSCRIBE", chanName).WaitFor(fcd.Conn, ch)
}
func (fcd *pFacade) Message(chanName, mess string) {
	fcd.Protcol.Request("PUBLISH", chanName, mess).Execute(fcd.Conn)
}
