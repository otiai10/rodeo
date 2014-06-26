package redis

import "strings"

type CommandPublish struct {
	chanName string
	message  string
	CommandDefault
}

func (this CommandPublish) Build() []byte {
	words := []string{
		"*3",
		this.getLenStr(cmdPUBLISH),
		cmdPUBLISH,
		this.getLenStr(this.chanName),
		this.chanName,
		this.getLenStr(this.message),
		this.message,
	}
	joined := strings.Join(words, sep) + sep
	return []byte(joined)
}
func (this CommandPublish) Parse(res []byte) (result string, e error) {
	return
}
