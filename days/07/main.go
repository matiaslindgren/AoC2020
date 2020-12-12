package main

import (
	"fmt"
	"github.com/matiaslindgren/AoC2020/util"
	"strings"
)

func bagId(e []string) (string) {
	return strings.Join(e, "-")
}

type Deps map[string]map[string]int

func (deps *Deps) add(bag, sub string, count int) {
	if dep, exists := (*deps)[bag]; exists {
		dep[sub] = count
	} else {
		(*deps)[bag] = map[string]int{sub: count}
	}
}

func (deps Deps) countUniqueDeps(bag string, visited map[string]bool) int {
	if visited[bag] {
		return 0
	} else {
		visited[bag] = true
	}
	n := 1
	for subBag, count := range deps[bag] {
		n += count * deps.countUniqueDeps(subBag, visited)
	}
	return n
}

func (deps Deps) countAllDeps(bag string) int {
	n := 1
	for subBag, count := range deps[bag] {
		n += count * deps.countAllDeps(subBag)
	}
	return n
}

func search(lines []string) (int, int) {
	deps := Deps{}
	rdeps := Deps{}

	for _, line := range lines {
		parts := strings.Split(strings.TrimRight(line, "."), " ")
		if parts[4] == "no" {
			continue
		}
		bag := bagId(parts[:2])
		for i := 4; i < len(parts); i += 4 {
			subBag := bagId(parts[i+1:i+3])
			bagCount := util.ParseInt(parts[i])
			deps.add(bag, subBag, bagCount)
			rdeps.add(subBag, bag, 1)
		}
	}

	a := rdeps.countUniqueDeps("shiny-gold", map[string]bool{})
	b := deps.countAllDeps("shiny-gold")
	return a-1, b-1
}

func main() {
	input := util.SlurpStdinLines()
	a, b := search(input)
	fmt.Println(a, b)
}
