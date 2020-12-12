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
	if x, err := strconv.Atoi(s); err != nil {
		panic(err)
	} else {
		return int(x)
	}
}

func ParseIntArray(lines []string) []int {
	v := make([]int, len(lines))
	for i, line := range lines {
		v[i] = ParseInt(line)
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

func ParseGrid(lines []string) Grid {
	out := make(Grid, len(lines))
	for y, line := range lines {
		out[y] = []byte(line)
	}
	return out
}
