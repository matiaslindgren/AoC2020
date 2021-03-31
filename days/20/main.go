package main

import (
	"fmt"
	"strings"

	"github.com/matiaslindgren/AoC2020/util"
)

func parseTile(section string) Tile {
	t := Tile{}

	lines := strings.Split(section, "\n")
	t.id = lines[0][5:9]

	t.borders[0] = lines[1]
	left, right := []byte{}, []byte{}
	for _, line := range lines[1:] {
		left = append(left, line[0])
		right = append(right, line[len(line)-1])
	}
	t.borders[1] = string(right)
	t.borders[2] = reverse(lines[len(lines)-1])
	t.borders[3] = reverse(string(left))

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

type Match struct {
	tile1, tile2 Tile
	border1, border2 int
}

type Tiles map[string]Tile

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

	for i, b := range t1.borders {
		t2.borders[(i+1)%4] = b
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

	t2.borders[0] = reverse(t1.borders[2])
	t2.borders[1] = reverse(t1.borders[1])
	t2.borders[2] = reverse(t1.borders[0])
	t2.borders[3] = reverse(t1.borders[3])

	return t2
}

// func (t Tile) print() {
// 	fmt.Println(t.id)
// 	for i, b := range t.borders {
// 		fmt.Println(i, b)
// 	}
// 	fmt.Println()
// 	for i, line := range t.img {
// 		fmt.Println(i, line)
// 	}
// 	fmt.Println()
// }

func matchBorders(t1, t2 Tile) (int, int, bool) {
	for i1, b1 := range t1.borders {
		for i2, b2 := range t2.borders {
			if b1 == b2 || b1 == reverse(b2) {
				return i1, i2, true
			}
		}
	}
	return -1, -1, false
}

func findCorners(tiles Tiles) [4]string {
	tile2matchCount := map[string]int{}
	for _, t1 := range tiles {
		for _, t2 := range tiles {
			if t1.id == t2.id {
				continue
			}
			if _, _, ok := matchBorders(t1, t2); ok {
				tile2matchCount[t1.id]++
			}
		}
	}
	i, corners := 0, [4]string{}
	for id, c := range tile2matchCount {
		if c == 2 {
			corners[i] = id
			i++
		}
	}
	if i != 4 {
		panic(fmt.Sprintf("wrong amount of corners %d != 4", i))
	}
	return corners
}

func buildGrid(tiles Tiles, start Tile) ([]string, bool) {
}

func search(input []string) (int, int) {
	originalTiles := Tiles{}
	for _, section := range input {
		t := parseTile(section)
		originalTiles[t.id] = t
	}

	corners := findCorners(originalTiles)
	a := 1
	for _, id := range corners {
		a *= util.ParseInt(id)
	}

	for _, id := range corners {
		grid, ok := buildGrid(originalTiles, originalTiles[id])
		if ok {
			fmt.Println(grid)
		}
	}

	b := 0
	return a, b
}

func main() {
	input := util.SlurpStdinSections()
	a, b := search(input)
	fmt.Println(a, b)
	fmt.Println(
`1951 2311 3079
2729 1427 2473
2971 1489 1171`)
}
