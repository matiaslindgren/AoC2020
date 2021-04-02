package main

import (
	"fmt"
	"github.com/matiaslindgren/AoC2020/util"
)

func countTrees(lines []string, dx, dy int) int {
	n := 0
	xMax, yMax := len(lines[0]), len(lines)
	for x, y := 0, 0; y < yMax; x, y = (x+dx)%xMax, y+dy {
		if lines[y][x] == '#' {
			n++
		}
	}
	return n
}

func search(lines []string) (int, int) {
	a := countTrees(lines, 3, 1)
	b := a
	delta := [][]int{{1, 1}, {5, 1}, {7, 1}, {1, 2}}
	for _, pair := range delta {
		b *= countTrees(lines, pair[0], pair[1])
	}
	return a, b
}

func main() {
	a, b := search(util.SlurpStdinLines())
	fmt.Println(a, b)
}
