package main

import (
	"fmt"
	"github.com/matiaslindgren/AoC2020/util"
	"strings"
)

func parseLine(line string) (int, int, byte, string) {
	parts := strings.Split(line, " ")
	ij, req, rule := strings.Split(parts[0], "-"), parts[1][0], parts[2]
	i, j := util.ParseInt(ij[0]), util.ParseInt(ij[1])
	return i, j, req, rule
}

func search(input []string) (int, int) {
	a, b := 0, 0
	for _, line := range input {
		i, j, reqChar, rule := parseLine(line)
		numReq := strings.Count(rule, string(reqChar))
		if i <= numReq && numReq <= j {
			a++
		}
		if (rule[i-1] == reqChar) != (rule[j-1] == reqChar) {
			b++
		}
	}
	return a, b
}

func main() {
	input := util.SlurpStdinLines()
	a, b := search(input)
	fmt.Println(a, b)
}
