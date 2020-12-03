package util

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func SlurpStdinLines() ([]string) {
	if f, err := ioutil.ReadFile("/dev/stdin"); err != nil {
		panic(err)
	} else {
		return strings.Split(strings.TrimSpace(string(f)), "\n")
	}
}

func parseInt(s string) (int) {
	if x, err := strconv.ParseInt(s, 10, 64); err != nil {
		panic(err)
	} else {
		return int(x)
	}
}

func ParseIntPair(s string) (int, int) {
	pair := strings.Split(s, "-")
	return parseInt(pair[0]), parseInt(pair[1])
}

func ParseIntArray() ([]int) {
	lines := SlurpStdinLines()
	v := make([]int, len(lines))
	for i, line := range lines {
		v[i] = parseInt(line)
	}
	return v
}

func ParseStringsTable() ([][]string) {
	lines := SlurpStdinLines()
	v := make([][]string, len(lines))
	for i, line := range lines {
		v[i] = strings.Split(line, " ")
	}
	return v
}
