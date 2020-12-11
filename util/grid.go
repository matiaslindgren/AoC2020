package util


type Grid [][]byte

func (g Grid) Copy() Grid {
	out := make(Grid, len(g))
	for y, row := range g {
		out[y] = make([]byte, len(row))
		copy(out[y], row)
	}
	return out
}

func (g Grid) IsOff(y, x int) bool {
	h, w := len(g), len(g[0])
	return y < 0 || y >= h || x < 0 || x >= w
}

func (g Grid) Count(ch byte) int {
	n := 0
	for _, row := range g {
		for _, pos := range row {
			if pos == ch {
				n++
			}
		}
	}
	return n
}

func (g Grid) NumOnQueensPath(y0, x0 int, ch byte) int {
	n := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			for y,x := y0+dy,x0+dx; !g.IsOff(y,x); y,x = y+dy,x+dx {
				switch g[y][x] {
				case '.':
					continue
				case ch:
					n++
				}
				break
			}
		}
	}
	return n
}

func (g Grid) NumAdjacent(y0, x0 int, ch byte) int {
	n := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			y, x := y0+dy, x0+dx
			if !g.IsOff(y,x) && g[y][x] == ch {
				n++
			}
		}
	}
	return n
}
