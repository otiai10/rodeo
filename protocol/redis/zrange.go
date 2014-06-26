package redis

import "strings"
import "regexp"
import "strconv"

type CommandZrange struct {
	key   string
	start string
	stop  string
	opt   string
	CommandDefault
}

func (this CommandZrange) Build() []byte {
	words := []string{
		"*5",
		this.getLenStr(CMD_ZRANGE),
		CMD_ZRANGE,
		this.getLenStr(this.key),
		this.key,
		this.getLenStr(this.start),
		this.start,
		this.getLenStr(this.stop),
		this.stop,
		this.getLenStr(this.opt),
		this.opt,
	}
	joined := strings.Join(words, sep) + sep
	return []byte(joined)
}

func (this CommandZrange) Parse(res []byte) (result string, e error) {
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
