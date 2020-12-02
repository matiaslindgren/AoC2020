package aocutil

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func slurpStdinLines() ([]string) {
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
	lines := slurpStdinLines()
	v := make([]int, len(lines))
	for i := range v {
		v[i] = parseInt(lines[i])
	}
	return v
}

func ParseStringsTable() ([][]string) {
	lines := slurpStdinLines()
	v := make([][]string, len(lines))
	for i := range v {
		v[i] = strings.Split(lines[i], " ")
	}
	return v
}
