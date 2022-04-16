package main

import (
	"fmt"
	"github.com/matiaslindgren/AoC2020/util"
	"sort"
)

func countArrangements(jolts []int) int {
	exists := util.Set(jolts)
	// Accessing a Go-map with a non-existing key returns 0
	numPerm := map[int]int{}
	numPerm[0] = 1
	for _, jolt := range jolts[1:] {
		if exists[jolt] {
			for i := 1; i <= 3; i++ {
				numPerm[jolt] += numPerm[jolt-i]
			}
		}
	}
	maxJolt := jolts[len(jolts)-1]
	return numPerm[maxJolt]
}

func search(jolts []int) (int, int) {
	sort.Ints(jolts)
	jolts = append([]int{0}, jolts...)
	jolts = append(jolts, jolts[len(jolts)-1]+3)
	delta1, delta3 := 0, 0
	for i := 1; i < len(jolts); i++ {
		switch jolts[i] - jolts[i-1] {
		case 1:
			delta1++
		case 3:
			delta3++
		}
	}
	a, b := delta1*delta3, countArrangements(jolts)
	return a, b
}

func main() {
	jolts := util.ParseIntArray(util.SlurpStdinLines())
	a, b := search(jolts)
	fmt.Println(a, b)
}
