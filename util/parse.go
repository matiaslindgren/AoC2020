package util

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func SlurpStdin() string {
	if f, err := ioutil.ReadFile("/dev/stdin"); err != nil {
		panic(err)
	} else {
		return string(f)
	}
}

func SlurpStdinLines() []string {
	return strings.Split(strings.TrimSpace(SlurpStdin()), "\n")
}

func SlurpStdinSections() []string {
	return strings.Split(strings.TrimSpace(SlurpStdin()), "\n\n")
}

func ParseInt(s string) int {
	if x, err := strconv.ParseInt(s, 10, 64); err != nil {
		panic(err)
	} else {
		return int(x)
	}
}

func ParseIntPair(s string) (int, int) {
	pair := strings.Split(s, "-")
	return ParseInt(pair[0]), ParseInt(pair[1])
}

func ParseIntArray(lines []string) []int {
	v := make([]int, len(lines))
	for i, line := range lines {
		v[i] = ParseInt(line)
	}
	return v
}

func ParseStringsTable() [][]string {
	lines := SlurpStdinLines()
	v := make([][]string, len(lines))
	for i, line := range lines {
		v[i] = strings.Split(line, " ")
	}
	return v
}

func Match(pattern, text string) bool {
	if matched, err := regexp.Match(pattern, []byte(text)); err != nil {
		panic(err)
	} else {
		return matched
	}
}
