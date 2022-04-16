package main

import (
	"fmt"
	"github.com/matiaslindgren/AoC2020/util"
	"strings"
)

func parseCards(section string) []int {
	return util.ParseIntArray(strings.Split(section, "\n")[1:])
}

func score(cards []int) int {
	s, n := 0, len(cards)
	for i, c := range cards {
		s += c * (n - i)
	}
	return s
}

func gameA(cards1, cards2 []int) (int, int) {
	for len(cards1) > 0 && len(cards2) > 0 {
		c1, c2 := cards1[0], cards2[0]
		cards1, cards2 = cards1[1:], cards2[1:]
		if c1 > c2 {
			cards1 = append(cards1, c1, c2)
		} else {
			cards2 = append(cards2, c2, c1)
		}
	}
	return score(cards1), score(cards2)
}

func gameB(cards1, cards2 []int) (int, int) {
	seen := map[string]bool{}
	for len(cards1) > 0 && len(cards2) > 0 {
		if key := fmt.Sprint(cards1, cards2); seen[key] {
			return 0, -1
		} else {
			seen[key] = true
		}
		c1, c2 := cards1[0], cards2[0]
		cards1, cards2 = cards1[1:], cards2[1:]
		var oneWins bool
		if len(cards1) >= c1 && len(cards2) >= c2 {
			s1, s2 := gameB(
				append([]int{}, cards1[:c1]...),
				append([]int{}, cards2[:c2]...))
			oneWins = s1 > s2
		} else {
			oneWins = c1 > c2
		}
		if oneWins {
			cards1 = append(cards1, c1, c2)
		} else {
			cards2 = append(cards2, c2, c1)
		}
	}
	return score(cards1), score(cards2)
}

func search(sections []string) (int, int) {
	cards1, cards2 := parseCards(sections[0]), parseCards(sections[1])
	s1, s2 := gameA(cards1, cards2)
	a := util.Max(s1, s2)
	s1, s2 = gameB(cards1, cards2)
	b := util.Max(s1, s2)
	return a, b
}

func main() {
	input := util.SlurpStdinSections()
	a, b := search(input)
	fmt.Println(a, b)
}
