package protocol

import "net"

// Protocol supports any transactions rodeo.pFacade requires.
type Protocol interface {
	Request(args ...string) Protocol
	Execute(conn net.Conn) Protocol
	WaitFor(conn net.Conn, reciever *chan string)
	ToResult() Result
}

// Result pass TCP response.
type Result struct {
	Response string
	Error    error
}
