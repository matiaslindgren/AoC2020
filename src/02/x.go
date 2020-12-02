package main

import (
	"aocutil"
	"fmt"
)

func a(input [][]string) (int) {
	numValid := 0
	for _, s := range input {
		min, max := aocutil.ParseIntPair(s[0])
		reqChar := s[1][0]
		numReq := 0
		for i := range s[2] {
			if s[2][i] == reqChar {
				numReq++
			}
		}
		if (min <= numReq && numReq <= max) {
			numValid++
		}
	}
	return numValid
}

func b(input [][]string) (int) {
	numValid := 0
	for _, s := range input {
		i, j := aocutil.ParseIntPair(s[0])
		reqChar := s[1][0]
		if ((s[2][i-1] == reqChar) != (s[2][j-1] == reqChar)) {
			numValid++
		}
	}
	return numValid
}

func main() {
	input := aocutil.ParseStringsTable()

	fmt.Println(a(input))
	fmt.Println(b(input))
}
