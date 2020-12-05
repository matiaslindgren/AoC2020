package main

import (
	"github.com/matiaslindgren/AoC2020/util"
	"fmt"
)

func search(input [][]string) (int, int) {
	a, b := 0, 0
	for _, s := range input {
		pair, reqChar, line := s[0], s[1][0], s[2]
		i, j := util.ParseIntPair(pair)
		numReq := util.CountChars(line, reqChar)
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
