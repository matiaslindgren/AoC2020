package main

import (
	"github.com/matiaslindgren/AoC2020/util"
	"fmt"
)

func search(input [][]string) (int, int) {
	a, b := 0, 0
	for _, s := range input {
		i, j := util.ParseIntPair(s[0])
		reqChar := s[1][0]
		numReq := 0
		for k := range s[2] {
			if s[2][k] == reqChar {
				numReq++
			}
		}
		if i <= numReq && numReq <= j {
			a++
		}
		if (s[2][i-1] == reqChar) != (s[2][j-1] == reqChar) {
			b++
		}
	}
	return a, b
}

func main() {
	input := util.ParseStringsTable()
	a, b := search(input)
	fmt.Println(a, b)
}
