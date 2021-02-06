package main

import (
	"fmt"
	"regexp"
	"strings"
	"github.com/matiaslindgren/AoC2020/util"
)

var directions = []string{"e","se","sw","w","nw","ne"}

// Axial coordinates, see
// https://www.redblobgames.com/grids/hexagons/#map-storage
type Pos struct {
	q, r int
}

// The Hex grid is a mapping from axial coordinates to 'isBlack' booleans
type HexGrid map[Pos]bool

func (p Pos) step(direction string) Pos {
	switch direction {
	case "e":
		p.q++
	case "se":
		p.r++
	case "sw":
		p.q--
		p.r++
	case "w":
		p.q--
	case "nw":
		p.r--
	case "ne":
		p.q++
		p.r--
	default:
		panic(fmt.Sprintf("unknown direction '%s'", direction))
	}
	return p
}

func walk(steps []string) Pos {
	p := Pos{}
	for _, d := range steps {
		p = p.step(d)
	}
	return p
}

func (g HexGrid) countBlack() int {
	n := 0
	for _, isBlack := range g {
		if isBlack {
			n++
		}
	}
	return n
}

func (g0 HexGrid) expandCopy() HexGrid {
	g1 := HexGrid{}
	for p := range g0 {
		g1[p] = g0[p]
		for _, d := range directions {
			p1 := p.step(d)
			g1[p1] = g0[p1]
		}
	}
	return g1
}

func (g0 HexGrid) step() HexGrid {
	g1 := g0.expandCopy()
	g2 := g0.expandCopy()
	for p, isBlack := range g1 {
		nBlack := 0
		for _, d := range directions {
			if g1[p.step(d)] {
				nBlack++
			}
		}
		if isBlack {
			if nBlack == 0 || nBlack > 2 {
				g2[p] = false
			}
		} else {
			if nBlack == 2 {
				g2[p] = true
			}
		}
	}
	return g2
}

func parseGrid(lines []string) HexGrid {
	grid := HexGrid{{}: false}
	re := regexp.MustCompile(strings.Join(directions, "|"))
	for _, line := range lines {
		steps := re.FindAllString(line, -1)
		pos := walk(steps)
		grid[pos] = !grid[pos]
	}
	return grid
}

func search(g HexGrid) (int, int) {
	a := g.countBlack()
	for i := 0; i < 100; i++ {
		g = g.step()
	}
	b := g.countBlack()
	return a, b
}

func main() {
	input := util.SlurpStdinLines()
	a, b := search(parseGrid(input))
	fmt.Println(a, b)
}
