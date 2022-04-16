package main

import (
	"fmt"
	"github.com/matiaslindgren/AoC2020/util"
)

type Row map[int]bool
type Grid2D map[int]Row
type Grid3D map[int]Grid2D
type Grid4D map[int]Grid3D

func initEmptyGrid3D(n int) Grid3D {
	g := Grid3D{}
	for z := -n; z <= n; z++ {
		g[z] = Grid2D{}
		for y := -n; y <= n; y++ {
			g[z][y] = Row{}
		}
	}
	return g
}

func initGrid3D(lines []string) Grid3D {
	n := len(lines) / 2
	g := initEmptyGrid3D(n)
	for y, row := range lines {
		for x, ch := range row {
			g[0][y-n][x-n] = ch == '#'
		}
	}
	return g
}

func (in Grid3D) countActiveAdj(x0, y0, z0 int) int {
	n := 0
	for dz := -1; dz <= 1; dz++ {
		for dy := -1; dy <= 1; dy++ {
			for dx := -1; dx <= 1; dx++ {
				if in[z0+dz][y0+dy][x0+dx] {
					n++
				}
			}
		}
	}
	if in[z0][y0][x0] {
		n--
	}
	return n
}

func (in Grid3D) countActive() int {
	n := 0
	for _, zdim := range in {
		for _, ydim := range zdim {
			for _, active := range ydim {
				if active {
					n++
				}
			}
		}
	}
	return n
}

func (in Grid3D) step() Grid3D {
	n := len(in)/2 + 1
	out := initEmptyGrid3D(n)
	for z := -n; z <= n; z++ {
		for y := -n; y <= n; y++ {
			for x := -n; x <= n; x++ {
				if a := in.countActiveAdj(x, y, z); in[z][y][x] {
					out[z][y][x] = a == 2 || a == 3
				} else {
					out[z][y][x] = a == 3
				}
			}
		}
	}
	return out
}

func initEmptyGrid4D(n int) Grid4D {
	g := Grid4D{}
	for w := -n; w <= n; w++ {
		g[w] = initEmptyGrid3D(n)
	}
	return g
}

func initGrid4D(lines []string) Grid4D {
	n := len(lines) / 2
	g := initEmptyGrid4D(n)
	for y, row := range lines {
		for x, ch := range row {
			g[0][0][y-n][x-n] = ch == '#'
		}
	}
	return g
}

func (in Grid4D) countActiveAdj(x0, y0, z0, w0 int) int {
	n := 0
	for dw := -1; dw <= 1; dw++ {
		for dz := -1; dz <= 1; dz++ {
			for dy := -1; dy <= 1; dy++ {
				for dx := -1; dx <= 1; dx++ {
					if in[w0+dw][z0+dz][y0+dy][x0+dx] {
						n++
					}
				}
			}
		}
	}
	if in[w0][z0][y0][x0] {
		n--
	}
	return n
}

func (in Grid4D) countActive() int {
	n := 0
	for _, wdim := range in {
		n += wdim.countActive()
	}
	return n
}

func (in Grid4D) step() Grid4D {
	n := len(in)/2 + 1
	out := initEmptyGrid4D(n)
	for w := -n; w <= n; w++ {
		for z := -n; z <= n; z++ {
			for y := -n; y <= n; y++ {
				for x := -n; x <= n; x++ {
					if a := in.countActiveAdj(x, y, z, w); in[w][z][y][x] {
						out[w][z][y][x] = a == 2 || a == 3
					} else {
						out[w][z][y][x] = a == 3
					}
				}
			}
		}
	}
	return out
}

func search(lines []string) (int, int) {
	g3 := initGrid3D(lines)
	g4 := initGrid4D(lines)
	for i := 0; i < 6; i++ {
		g3 = g3.step()
		g4 = g4.step()
	}
	return g3.countActive(), g4.countActive()
}

func main() {
	input := util.SlurpStdinLines()
	a, b := search(input)
	fmt.Println(a, b)
}
