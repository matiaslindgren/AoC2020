package main

import (
	"fmt"
	"github.com/matiaslindgren/AoC2020/util"
	"math/bits"
	"strings"
)

func search(sections []string) (int, int) {
	a, b := 0, 0
	for _, section := range sections {
		any, all := uint32(0), ^uint32(0)
		for _, ans := range strings.Split(section, "\n") {
			person := uint32(0)
			for _, c := range ans {
				person |= 1 << ('z'-c)
			}
			any |= person
			all &= person
		}
		a += bits.OnesCount32(any)
		b += bits.OnesCount32(all)
	}
	return a, b
}

func main() {
	input := util.SlurpStdinSections()
	a, b := search(input)
	fmt.Println(a, b)
}
