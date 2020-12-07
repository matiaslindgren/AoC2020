package main

import (
	"fmt"
	"github.com/matiaslindgren/AoC2020/util"
	"strings"
)

func getParents(child2parents map[string][]string, key string) ([]string) {
	parents := make([]string, 0)
	for _, parent := range child2parents[key] {
		parents = append(parents, parent)
		for _, p := range getParents(child2parents, parent) {
			parents = append(parents, p)
		}
	}
	return parents
}

func getChildren(bag2counts map[string][]Count, key string) (int) {
	n := 1
	for _, count := range bag2counts[key] {
		n += count.n * getChildren(bag2counts, count.bag)
	}
	return n
}

func join(e []string) (string) {
	return strings.Join(e, "-")
}

func contains(elems []string, key string) (bool) {
	for _, e := range elems {
		if e == key {
			return true
		}
	}
	return false
}

type Count struct {
	bag string;
	n int;
}

func search(lines []string) (int, int) {
	child2parents := map[string][]string{}
	bag2counts := map[string][]Count{}
	for _, line := range lines {
		parts := strings.Split(strings.TrimRight(line, "."), " ")
		if parts[4] == "no" {
			continue
		}
		parent := join(parts[:2])
		for i := 4; i < len(parts); i += 4 {
			child := join(parts[i+1:i+3])
			if !contains(child2parents[child], parent) {
				child2parents[child] = append(child2parents[child], parent)
			}
			n := util.ParseInt(parts[i])
			count := Count{child, n}
			bag2counts[parent] = append(bag2counts[parent], count)
		}
	}

	parents := make([]string, 0)
	for _, p := range getParents(child2parents, "shiny-gold") {
		if !contains(parents, p) {
			parents = append(parents, p)
		}
	}
	a := len(parents)
	b := getChildren(bag2counts, "shiny-gold") - 1

	return a, b
}

func main() {
	input := util.SlurpStdinLines()
	a, b := search(input)
	fmt.Println(a, b)
}
