package redis

import "strings"

// CommandPublish provides TCP communication of `PUBLISH`.
type CommandPublish struct {
	chanName string
	message  string
	commandDefault
}

func (cmd CommandPublish) build() []byte {
	words := []string{
		"*3",
		cmd.strlen(cmdPUBLISH),
		cmdPUBLISH,
		cmd.strlen(cmd.chanName),
		cmd.chanName,
		cmd.strlen(cmd.message),
		cmd.message,
	}
	joined := strings.Join(words, sep) + sep
	return []byte(joined)
}

func (cmd CommandPublish) parse(res []byte) (result string, e error) {
	return
}
