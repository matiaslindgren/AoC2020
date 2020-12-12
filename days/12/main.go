package main

import (
	"fmt"
	"math"
	"github.com/matiaslindgren/AoC2020/util"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func rad(deg int) float64 {
	return float64(deg) * math.Pi / 180.0
}

func rot(x0, y0, deg int) (int, int) {
	cos := int(math.Round(math.Cos(rad(deg))))
	sin := int(math.Round(math.Sin(rad(deg))))
	x1 := x0 * cos - y0 * sin
	y1 := x0 * sin + y0 * cos
	return x1, y1
}

func move(direction byte, x, y, dist int) (int, int) {
	switch direction {
	case 'N':
		y += dist
	case 'E':
		x += dist
	case 'S':
		y -= dist
	case 'W':
		x -= dist
	}
	return x, y
}

func searchB(lines []string) int {
	x, y := 0, 0
	xW, yW := 10, 1
	for _, line := range lines {
		num := util.ParseInt(line[1:])
		ch := byte(line[0])
		switch ch {
		case 'F':
			x += num * xW
			y += num * yW
		case 'L':
			xW, yW = rot(xW, yW, num)
		case 'R':
			xW, yW = rot(xW, yW, -num)
		default:
			xW, yW = move(ch, xW, yW, num)
		}
	}
	return abs(x) + abs(y)
}

func searchA(lines []string) int {
	x, y := 0, 0
	dirs := []byte("NESW")
	dir := 1
	for _, line := range lines {
		num := util.ParseInt(line[1:])
		ch := byte(line[0])
		if ch == 'F' {
			ch = dirs[dir]
		}
		switch ch {
		case 'L':
			dir = (dir + 4 - (num/90))%4
		case 'R':
			dir = (dir + (num/90))%4
		default:
			x, y = move(ch, x, y, num)
		}
	}
	return abs(x) + abs(y)
}

func main() {
	input := util.SlurpStdinLines()
	a, b := searchA(input), searchB(input)
	fmt.Println(a, b)
}
