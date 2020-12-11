package main

import (
	"fmt"
	"github.com/matiaslindgren/AoC2020/util"
)

func stepA(in util.Grid) (util.Grid, bool) {
	out, changed := in.Copy(), false
	for y, row := range in {
		for x, pos := range row {
			switch pos {
			case 'L':
				if in.NumAdjacent(y, x, '#') == 0 {
					out[y][x], changed = byte('#'), true
				}
			case '#':
				if in.NumAdjacent(y, x, '#') >= 4 {
					out[y][x], changed = byte('L'), true
				}
			}
		}
	}
	return out, changed
}

func stepB(in util.Grid) (util.Grid, bool) {
	out, changed := in.Copy(), false
	for y, row := range in {
		for x, pos := range row {
			switch pos {
			case 'L':
				if in.NumOnQueensPath(y, x, '#') == 0 {
					out[y][x], changed = byte('#'), true
				}
			case '#':
				if in.NumOnQueensPath(y, x, '#') >= 5 {
					out[y][x], changed = byte('L'), true
				}
			}
		}
	}
	return out, changed
}

func search(grid util.Grid) (int, int) {
	var g util.Grid
	for tmp,c := stepA(grid); c; tmp,c = stepA(tmp) {
		g = tmp
	}
	a := g.Count('#')
	for tmp,c := stepB(grid); c; tmp,c = stepB(tmp) {
		g = tmp
	}
	b := g.Count('#')
	return a, b
}

func main() {
	input := util.ParseGrid(util.SlurpStdinLines())
	a, b := search(input)
	fmt.Println(a, b)
}
