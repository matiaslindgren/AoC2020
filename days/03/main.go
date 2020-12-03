package main

import (
	"github.com/matiaslindgren/AoC2020/util"
	"fmt"
)

func numTrees(lines []string, dy, dx int) (int) {
	a := 0
	for x, y := 0, 0; y < len(lines); x, y = (x+dx)%len(lines[0]), y+dy {
		if lines[y][x] == '#' {
			a++
		}
	}
	return a
}

func search(lines []string) (int, int) {
	a := numTrees(lines[:], 1, 3)
	b := a
	delta := [][]int{{1, 1}, {1, 5}, {1, 7}, {2, 1}}
	for _, pair := range delta {
		b *= numTrees(lines[:], pair[0], pair[1])
	}
	return a, b
}


func main() {
	input := util.SlurpStdinLines()
	a, b := search(input)
	fmt.Println(a, b)
}
