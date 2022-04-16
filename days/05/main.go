package main

import (
	"fmt"
	"github.com/matiaslindgren/AoC2020/util"
	"strings"
)

func decodeBitChars(from, on string) int {
	x, n := 0, len(from)
	for i, ch := range from {
		if strings.ContainsAny(string(ch), on) {
			x += 1 << (n - i - 1)
		}
	}
	return x
}

func search(lines []string) (int, int) {
	idExists := map[int]bool{}
	for _, line := range lines {
		id := decodeBitChars(line, "BR")
		idExists[id] = true
	}

	minId, maxId := 1<<63-1, 0
	for id := range idExists {
		minId = util.Min(minId, id)
		maxId = util.Max(maxId, id)
	}

	a, b := maxId, 0
	for id := minId; id <= maxId; id++ {
		if !idExists[id] {
			b = id
		}
	}

	return a, b
}

func main() {
	input := util.SlurpStdinLines()
	a, b := search(input)
	fmt.Println(a, b)
}
