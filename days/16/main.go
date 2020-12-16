package main

import (
	"fmt"
	"strings"

	"github.com/matiaslindgren/AoC2020/util"
)

var Split = strings.Split

type Rule struct {
	Name string
	R [4]int
}

func parseRule(r string) Rule {
	parts := Split(r, ": ")
	name, parts := parts[0], Split(parts[1], " ")
	a, b := Split(parts[0], "-"), Split(parts[2], "-")
	f := util.ParseInt
	return Rule{name, [...]int{f(a[0]), f(a[1]), f(b[0]), f(b[1])}}
}

func parseSections(input string) ([]Rule, string, []string) {
	sections := Split(input, "\n\n")
	rules := []Rule{}
	for _, rule := range Split(sections[0], "\n") {
		rules = append(rules, parseRule(rule))
	}
	ticket := Split(sections[1], "\n")[1]
	tickets := Split(strings.TrimSpace(sections[2]), "\n")[1:]
	return rules, ticket, tickets
}

func (r Rule) valid(x int) bool {
	return (r.R[0] <= x && x <= r.R[1]) || (r.R[2] <= x && x <= r.R[3])
}

func valid(rules []Rule, x int) (bool) {
	for _, rule := range rules {
		if (rule.R[0] <= x && x <= rule.R[1]) || (rule.R[2] <= x && x <= rule.R[3]) {
			return true
		}
	}
	return false
}

func numtrue(m map[int]bool) (int, int) {
	n := 0
	i := 0
	for k, v := range m {
		if v {
			n++
			i = k
		}
	}
	return i, n
}

func search(input string) (int, int) {
	rules, myTicket, tickets := parseSections(input)

	a := 0
	validTickets := [][]int{}
	for _, t := range tickets {
		v := util.ParseIntArray(Split(t, ","))
		ok := true
		for _, x := range v {
			if !valid(rules, x) {
				a += x
				ok = false
				break
			}
		}
		if ok {
			validTickets = append(validTickets, v)
		}
	}

	candidates := map[int]map[int]bool{}
	for i := 0; i < len(validTickets[0]); i++ {
		candidates[i] = map[int]bool{}
	}
	for i, rule := range rules {
		for col := 0; col < len(validTickets[0]); col++ {
			for _, t := range validTickets {
				x := t[col]
				if rule.valid(x) {
					candidates[col][i] = true
				} else {
					candidates[col][i] = false
					break
				}
			}
		}
	}

	col2rule := map[int]Rule{}
	for i := 0; len(col2rule) < len(validTickets[0]); {
		for col, cand := range candidates {
			j, numt := numtrue(cand)
			if numt == 1 {
				col2rule[col] = rules[j]
				i = j
				break
			}
			if numt == 0 {
				delete(candidates, col)
			}
		}
		for _, cand := range candidates {
			delete(cand, i)
		}
	}

	b := 1
	for col, x := range util.ParseIntArray(Split(myTicket, ",")) {
		if strings.Contains(col2rule[col].Name, "departure") {
			b *= x
		}
	}
	return a, b
}

func main() {
	input := util.SlurpStdin()
	a, b := search(input)
	fmt.Println(a, b)
}
