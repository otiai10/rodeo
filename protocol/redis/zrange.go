package redis

import "strings"
import "regexp"
import "strconv"

// CommandZrange provides TCP communication of `ZRANGE`.
type CommandZrange struct {
	key   string
	start string
	stop  string
	opt   string
	CommandDefault
}

// Build builds TCP message by initialized parameters.
func (cmd CommandZrange) Build() []byte {
	words := []string{
		"*5",
		cmd.getLenStr(cmdZRANGE),
		cmdZRANGE,
		cmd.getLenStr(cmd.key),
		cmd.key,
		cmd.getLenStr(cmd.start),
		cmd.start,
		cmd.getLenStr(cmd.stop),
		cmd.stop,
		cmd.getLenStr(cmd.opt),
		cmd.opt,
	}
	joined := strings.Join(words, sep) + sep
	return []byte(joined)
}

// Parse parses TCP response.
func (cmd CommandZrange) Parse(res []byte) (result string, e error) {
	re := regexp.MustCompile("\\*([0-9]+)")
	var recordsCount int
	if matches := re.FindStringSubmatch(string(res)); len(matches) > 1 {
		responseCount, _ := strconv.Atoi(matches[1])
		recordsCount = responseCount / 2
	}
	if recordsCount < 1 {
		return
	}
	pool := make([]string, recordsCount*2)
	lines := strings.Split(string(res), "\r\n")[1:]
	for i := 0; i < recordsCount; i++ {
		indexVal := i*4 + 1
		indexScore := i*4 + 3
		pool[i*2] = lines[indexVal]
		pool[i*2+1] = lines[indexScore]
	}
	result = strings.Join(pool, "\n")
	return
}
