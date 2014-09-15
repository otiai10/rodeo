package redis

import "strings"

// import "regexp"
import "strconv"
import "net"
import "bufio"

// CommandZrange provides TCP communication of `ZRANGE`.
type CommandZrange struct {
	key   string
	start string
	stop  string
	opt   string
	commandDefault
}

func (cmd CommandZrange) build() []byte {
	words := []string{
		"*5",
		cmd.strlen(cmdZRANGE),
		cmdZRANGE,
		cmd.strlen(cmd.key),
		cmd.key,
		cmd.strlen(cmd.start),
		cmd.start,
		cmd.strlen(cmd.stop),
		cmd.stop,
		cmd.strlen(cmd.opt),
		cmd.opt,
	}
	joined := strings.Join(words, sep) + sep
	return []byte(joined)
}

func (cmd CommandZrange) parse(res []byte) (result string, e error) {
	return string(res), e
	/*
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
	*/
}

func (cmd CommandZrange) hoge(conn net.Conn) (res []byte) {
	scanner := bufio.NewScanner(conn)
	if ok := scanner.Scan(); !ok {
		return
	}
	var count int = 0
	var pool = make([]string, count*2)
	if m := RESP["array"].FindSubmatch(scanner.Bytes()); len(m) > 1 {
		lineCount, _ := strconv.Atoi(string(m[1]))
		count = lineCount / 2
	}
	for i := 0; i < count; i++ {
		scanner.Scan() // ${bulk}
		scanner.Scan() // {val}
		pool = append(pool, scanner.Text())
		scanner.Scan() // ${bulk}
		scanner.Scan() // {val}
		pool = append(pool, scanner.Text())
	}
	return []byte(strings.Join(pool, "\n"))
}
