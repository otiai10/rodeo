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
func (fcd *pFacade) Listen(ch *chan string) {
	// TODO: Protcol経由で、tcpをReadするgoroutineをつくる
	// Readすべきものが発生したら、それをparseして
	// chに流し込む
	/*
	   go func(){
	       for {
	           time.Sleep(5 * time.Second)
	           println("001")
	           *ch<- "これはredisから来たメッセージ想定"
	           println("002")
	       }
	   }()
	*/
}
func (fcd *pFacade) Message(ch *chan string, mess string) {
	*ch <- mess
}
