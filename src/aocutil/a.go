package aocutil

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func ParseStdinLines() ([]int) {
	f, err := ioutil.ReadFile("/dev/stdin")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(f)), "\n")
	v := make([]int, len(lines))
	for i := range v {
		x, err := strconv.ParseInt(lines[i], 10, 64)
		if err != nil {
			panic(err)
		}
		v[i] = int(x)
	}
	return v
}
