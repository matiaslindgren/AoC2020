package main

import (
	"github.com/matiaslindgren/AoC2020/util"
	"fmt"
)

func search(lines []string) (int, int) {
	ids := map[int]bool{}
	for _, line := range lines {
		id := util.DecodeBitChars(line, "BR")
		ids[id] = true
	}

	minId, maxId := 1<<63-1, 0
	for id, exists := range ids {
		if exists {
			if id < minId {
				minId = id
			}
			if id > maxId {
				maxId = id
			}
		}
	}

	a, b := maxId, 0
	for id := minId; id <= maxId; id++ {
		if _, exists := ids[id]; !exists {
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
