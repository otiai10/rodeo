package redis

import "strings"

// CommandPublish provides TCP communication of `PUBLISH`.
type CommandPublish struct {
	chanName string
	message  string
	CommandDefault
}

// Build builds TCP message by initialized parameters.
func (cmd CommandPublish) Build() []byte {
	words := []string{
		"*3",
		cmd.getLenStr(cmdPUBLISH),
		cmdPUBLISH,
		cmd.getLenStr(cmd.chanName),
		cmd.chanName,
		cmd.getLenStr(cmd.message),
		cmd.message,
	}
	joined := strings.Join(words, sep) + sep
	return []byte(joined)
}

// Parse parses TCP response.
func (cmd CommandPublish) Parse(res []byte) (result string, e error) {
	return
}
