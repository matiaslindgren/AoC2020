package main

import (
	"aocutil"
	"fmt"
)

func search(input []int) (int, int) {
	var a = -1
	var b = -1
	for i, x1 := range input {
		for j, x2 := range input[i+1:] {
			if (x1 + x2 == 2020) {
				a = x1 * x2
			}
			for _, x3 := range input[j+i+2:] {
				if (x1 + x2 + x3 == 2020) {
					b = x1 * x2 * x3
				}
				if (a != -1 && b != -1) {
					return a, b
				}
			}
		}
	}
	return a, b
}

func main() {
	input := aocutil.ParseStdinLines()
	a, b := search(input[:])
	fmt.Println(a, b)
}
