package main

import (
	"fmt"
	"math/bits"
	"strings"
	"github.com/matiaslindgren/AoC2020/util"
)

func search(sections []string) (int, int) {
	a, b := 0, 0
	for _, section := range sections {
		any, all := uint(0), ^uint(0)
		for _, ans := range strings.Split(section, "\n") {
			total := uint(0)
			for _, c := range ans {
				total |= 1 << ('z'-c)
			}
			any |= total
			all &= total
		}
		a += bits.OnesCount(any)
		b += bits.OnesCount(all)
	}
	return a, b
}

func main() {
	input := util.SlurpStdinSections()
	a, b := search(input)
	fmt.Println(a, b)
}
