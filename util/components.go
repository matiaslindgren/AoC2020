package util

import (
	"regexp"
	"strings"
)

// day 02
func CountChars(line string, c byte) (int) {
	n := 0
	for i := range line {
		if line[i] == c {
			n++
		}
	}
	return n
}

// day 03
func Intersect(lines []string, dx, dy int, o byte) (int) {
	n := 0
	xMax, yMax := len(lines[0]), len(lines)
	for x, y := 0, 0; y < yMax; x, y = (x+dx)%xMax, y+dy {
		if lines[y][x] == o {
			n++
		}
	}
	return n
}

// day 04
func Match(pattern, text string) (bool) {
	if matched, err := regexp.Match(pattern, []byte(text)); err != nil {
		panic(err)
	} else {
		return matched
	}
}

// day 05
func DecodeBitChars(from, on string) (int) {
	x, n := 0, len(from)
	for i, ch := range from {
		if strings.ContainsAny(string(ch), on) {
			x += 1 << (n-i-1)
		}
	}
	return x
}
