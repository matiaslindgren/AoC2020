package main

import (
	"fmt"
	"github.com/matiaslindgren/AoC2020/util"
)

func search(v []int) (int, int) {
	age1, age2 := map[int]int{}, map[int]int{}
	t, prev := 1, 0
	for _, x := range v {
		age1[x], age2[x] = t, t
		t, prev = t+1, x
	}
	a := 0
	for ; t <= 30_000_000; t++ {
		prev = age2[prev]-age1[prev]
		age1[prev], age2[prev] = age2[prev], t
		if age1[prev] == 0 {
			age1[prev] = t
		}
		if t == 2020 {
			a = prev
		}
	}
	return a, prev
}

func main() {
	input := util.SlurpStdinArray(",")
	a, b := search(input)
	fmt.Println(a, b)
}
