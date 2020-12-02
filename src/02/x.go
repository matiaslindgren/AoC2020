package main

import (
	"aocutil"
	"fmt"
)

func search(input [][]string) (int, int) {
	a, b := 0, 0
	for _, s := range input {
		i, j := aocutil.ParseIntPair(s[0])
		reqChar := s[1][0]
		numReq := 0
		for i := range s[2] {
			if s[2][i] == reqChar {
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
	input := aocutil.ParseStringsTable()
	a, b := search(input)
	fmt.Println(a, b)
}
