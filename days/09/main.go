package main

import (
	"fmt"
	"github.com/matiaslindgren/AoC2020/util"
)

func hasSumPair(v []int, val int) (bool) {
	for i, x := range v {
		for _, y := range v[i+1:] {
			if x + y == val {
				return true
			}
		}
	}
	return false
}

func subseqSumRange(v []int, a int) (int, int) {
	s, l, r := v[0], 0, 1
	for ; l < r && r < len(v); {
		switch {
		case s > a:
			s -= v[l]
			l++
		case s == a:
			return l, r
		case s < a:
			s += v[r]
			r++
		}
	}
	return -1, -1
}

func search(v []int, preambleSize int) (int, int) {
	a := 0
	for i, x := range v[preambleSize+1:] {
		preamble := v[i:i+preambleSize+1]
		if !hasSumPair(preamble, x) {
			a = x
			break
		}
	}

	l, r := subseqSumRange(v, a)
	min, max := 1<<63-1, 0
	for _, x := range v[l:r+1] {
		if x > max {
			max = x
		}
		if x < min {
			min = x
		}
	}
	b := min + max

	return a, b
}

func main() {
	input := util.ParseIntArray()
	a, b := search(input, 25)
	fmt.Println(a, b)
}
