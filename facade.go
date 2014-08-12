package rodeo

import "net"
import "github.com/otiai10/rodeo/protocol"
import "encoding/json"
import "strconv"
import "strings"
import "reflect"

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
func (fcd *pFacade) ZAdd(key string, score int64, obj interface{}) (e error) {
	var bs []byte
	bs, e = json.Marshal(obj)
	if e != nil {
		return
	}
	result := fcd.Protcol.Request("ZADD", key, strconv.FormatInt(score, 10), string(bs)).Execute(fcd.Conn).ToResult()
	return result.Error
}
func (fcd *pFacade) ZCount(key string, args ...int64) (count int, e error) {
	min, max := "-inf", "+inf"
	if 0 < len(args) {
		min = strconv.FormatInt(args[0], 10)
	}
	if 1 < len(args) {
		max = strconv.FormatInt(args[1], 10)
	}
	result := fcd.Protcol.Request(
		"ZCOUNT",
		key,
		min,
		max,
	).Execute(fcd.Conn).ToResult()
	if count, e = strconv.Atoi(result.Response); e == nil {
		return
	}
	e = result.Error
	return
}

type scoredValue struct {
	Score int64
	Value interface{}
}

func (fcd *pFacade) ZRange(key string, args []int, dest interface{}) (vals []scoredValue) {
	start, stop := "0", "-1"
	if 0 < len(args) {
		start = strconv.Itoa(args[0])
	}
	if 1 < len(args) {
		stop = strconv.Itoa(args[1])
	}
	result := fcd.Protcol.Request(
		"ZRANGE",
		key,
		start,
		stop,
		"WITHSCORES",
	).Execute(fcd.Conn).ToResult()
	rows := strings.Split(result.Response, "\n")
	if len(rows) < 2 {
		return
	}
	for i := range rows {
		val := scoredValue{}
		if i%2 != 0 {
			continue
		}
		score, e := strconv.ParseInt(rows[i+1], 10, 64)
		if e != nil {
			// TODO: log?
			continue
		}
		val.Score = score
		obj := reflect.New(reflect.TypeOf(dest)).Interface()
		json.Unmarshal([]byte(rows[i]), obj)
		val.Value = obj
		vals = append(vals, val)
	}
	return
}
func (fcd *pFacade) ZRangeByScore(key string, min int64, max int64, dest interface{}) (vals []scoredValue) {
	result := fcd.Protcol.Request(
		"ZRANGEBYSCORE",
		key,
		strconv.FormatInt(min, 10),
		strconv.FormatInt(max, 10),
		"WITHSCORES",
	).Execute(fcd.Conn).ToResult()
	rows := strings.Split(result.Response, "\n")
	if len(rows) < 2 {
		return
	}
	for i := range rows {
		val := scoredValue{}
		if i%2 != 0 {
			continue
		}
		score, e := strconv.ParseInt(rows[i+1], 10, 64)
		if e != nil {
			// TODO: log?
			continue
		}
		val.Score = score
		obj := reflect.New(reflect.TypeOf(dest)).Interface()
		json.Unmarshal([]byte(rows[i]), obj)
		val.Value = obj
		vals = append(vals, val)
	}
	return
}
func (fcd *pFacade) ZRemRangeByScore(key string, min, max int64) (e error) {
	result := fcd.Protcol.Request(
		"ZREMRANGEBYSCORE",
		key,
		strconv.FormatInt(min, 10),
		strconv.FormatInt(max, 10),
	).Execute(fcd.Conn).ToResult()
	return result.Error
}
func (fcd *pFacade) ZRem(key string, val interface{}) (e error) {
	var bs []byte
	bs, e = json.Marshal(val)
	if e != nil {
		return e
	}
	result := fcd.Protcol.Request(
		"ZREM",
		key,
		string(bs),
	).Execute(fcd.Conn).ToResult()
	return result.Error
}

func (fcd *pFacade) Listen(chanName string, ch *chan string) {
	fcd.Protcol.Request("SUBSCRIBE", chanName).WaitFor(fcd.Conn, ch)
}
func (fcd *pFacade) Message(chanName, mess string) {
	fcd.Protcol.Request("PUBLISH", chanName, mess).Execute(fcd.Conn)
}
