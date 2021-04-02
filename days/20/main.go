package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/matiaslindgren/AoC2020/util"
)

const (
	UP = 0
	RIGHT = 1
	DOWN = 2
	LEFT = 3
)
var monsterPattern = [3]string{
	"..................#..",
	"#....##....##....###.",
	".#..#..#..#..#..#....",
}
var monsterRegexp = [3]*regexp.Regexp{
	regexp.MustCompile(monsterPattern[0]),
	regexp.MustCompile(monsterPattern[1]),
	regexp.MustCompile(monsterPattern[2]),
}

func parseTile(section string) Tile {
	t := Tile{}

	lines := strings.Split(section, "\n")
	t.id = lines[0][5:9]

	t.borders[UP] = lines[1]
	left, right := []byte{}, []byte{}
	for _, line := range lines[1:] {
		left = append(left, line[0])
		right = append(right, line[len(line)-1])
	}
	t.borders[RIGHT] = string(right)
	t.borders[DOWN] = reverse(lines[len(lines)-1])
	t.borders[LEFT] = reverse(string(left))

	for _, line := range lines[2:len(lines)-1] {
		t.img = append(t.img, line[1:len(line)-1])
	}

	return t
}

func reverse(s string) string {
	n := len(s)
	r := make([]byte, n)
	for i := 0; i < n; i++ {
		r[n-i-1] = s[i]
	}
	return string(r)
}

type Tile struct {
	id string
	img []string
	borders [4]string
}

type TileMap map[string]Tile

func rotate(t1 Tile) Tile {
	n := len(t1.img)
	t2 := Tile{
		id: t1.id,
		img: make([]string, n),
	}

	for y := 0; y < n; y++ {
		row := make([]byte, n)
		for x := 0; x < n; x++ {
			row[n-x-1] = t1.img[x][y]
		}
		t2.img[y] = string(row)
	}

	for side, b := range t1.borders {
		t2.borders[(side+1)%4] = b
	}

	return t2
}

func flip(t1 Tile) Tile {
	n := len(t1.img)
	t2 := Tile{
		id: t1.id,
		img: make([]string, n),
	}

	for y, row := range t1.img {
		t2.img[n-y-1] = row
	}

	for i2, i1 := range []int{DOWN, RIGHT, UP, LEFT} {
		t2.borders[i2] = reverse(t1.borders[i1])
	}

	return t2
}

func (in *TileMap) copy() *TileMap {
	out := TileMap{}
	for id, t := range *in {
		out[id] = t
	}
	return &out
}

func (t1 Tile) findMatchOrientation(t2 Tile, side int) (Tile, bool) {
	for flips := 0; flips < 2; flips++ {
		for rotations := 0; rotations < 4; rotations++ {
			if t1.borders[side] == reverse(t2.borders[(side+2)%4]) {
				return t2, true
			}
			t2 = rotate(t2)
		}
		t2 = flip(t2)
	}
	return t2, false
}

type Grid map[int]map[int]Tile

func (g *Grid) set(x, y int, t Tile) {
	if col, exists := (*g)[x]; exists {
		col[y] = t
	} else {
		(*g)[x] = map[int]Tile{y: t}
	}
}

func (g *Grid) cornerIndexes() (int, int, int, int) {
	minX := 1<<63-1
	maxX := -minX
	maxY := maxX
	minY := minX
	for x, col := range *g {
		maxX = util.Max(maxX, x)
		minX = util.Min(minX, x)
		for y := range col {
			maxY = util.Max(maxY, y)
			minY = util.Min(minY, y)
		}
	}
	return minX, maxX, minY, maxY
}

func (g *Grid) count() int {
	minX, maxX, minY, maxY := g.cornerIndexes()
	n := 0
	for y := maxY; y >= minY; y-- {
		for x := minX; x <= maxX; x++ {
			if len((*g)[x][y].id) > 0 {
				n++
			}
		}
	}
	return n
}

func (g *Grid) corners() [4]Tile {
	minX, maxX, minY, maxY := g.cornerIndexes()
	return [4]Tile{
		(*g)[minX][minY],
		(*g)[minX][maxY],
		(*g)[maxX][minY],
		(*g)[maxX][maxY],
	}
}

func posDelta(side int) (int, int) {
	switch side {
	case UP:
		return 0, 1
	case RIGHT:
		return 1, 0
	case DOWN:
		return 0, -1
	case LEFT:
		return -1, 0
	default:
		panic(fmt.Sprint("unknown side ", side))
	}
}

func buildGrid(t1 Tile, x, y int, tiles *TileMap, frozen *map[string]bool, grid *Grid) *Grid {
	(*frozen)[t1.id] = true
	grid.set(x, y, t1)
	for side := 0; side < 4; side++ {
		for _, t := range *tiles {
			if (*frozen)[t.id] {
				continue
			}
			if t2, ok := t1.findMatchOrientation(t, side); ok {
				dx, dy := posDelta(side)
				grid = buildGrid(t2, x+dx, y+dy, tiles, frozen, grid)
			}
		}
	}
	return grid
}

func (g *Grid) buildImage(tileWidth int) []string {
	img := []string{}
	minX, maxX, minY, maxY := g.cornerIndexes()
	for y := maxY; y >= minY; y-- {
		for yTile := 0; yTile < tileWidth; yTile++ {
			row := []string{}
			for x := minX; x <= maxX; x++ {
				if t, exists := (*g)[x][y]; exists {
					row = append(row, t.img[yTile])
				}
			}
			if len(row) > 0 {
				img = append(img, strings.Join(row, ""))
			}
		}
	}
	return img
}

func hasMonster(img []string, x, y, w int) bool {
	for i, re := range monsterRegexp {
		if !re.MatchString(img[y+i][x:x+w]) {
			return false
		}
	}
	return true
}

func countMonsters(img []string) int {
	n := 0
	h, w := len(monsterPattern), len(monsterPattern[0])
	for y := 0; y < len(img)-h; y++ {
		for x := 0; x < len(img[0])-w; x++ {
			if hasMonster(img, x, y, w) {
				n++
				x += w-1
			}
		}
	}
	return n
}

func countRoughness(img []string) int {
	if numMonsters := countMonsters(img); numMonsters > 0 {
		numAll := strings.Count(strings.Join(img, ""), "#")
		numMonster := strings.Count(strings.Join(monsterPattern[:], ""), "#")
		return numAll - numMonsters * numMonster
	}
	return 0
}

func findMonsterGrid(tiles TileMap) (*Grid, []string) {
	for _, t := range tiles {
		for flips := 0; flips < 2; flips++ {
			for rotations := 0; rotations < 4; rotations++ {
				grid := buildGrid(t, 0, 0, tiles.copy(), &map[string]bool{}, &Grid{})
				if grid.count() == len(tiles) {
					img := grid.buildImage(len(t.img))
					if n := countMonsters(img); n > 0 {
						return grid, img
					}
				}
				t = rotate(t)
			}
			t = flip(t)
		}
	}
	return nil, nil
}

func search(input []string) (int, int) {
	tiles := TileMap{}
	for _, section := range input {
		t := parseTile(section)
		tiles[t.id] = t
	}

	grid, img := findMonsterGrid(tiles)

	a := 1
	for _, c := range grid.corners() {
		a *= util.ParseInt(c.id)
	}
	b := countRoughness(img)
	return a, b
}

func main() {
	input := util.SlurpStdinSections()
	a, b := search(input)
	fmt.Println(a, b)
}
