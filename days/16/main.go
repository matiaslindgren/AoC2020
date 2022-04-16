package main

import (
	"fmt"
	"github.com/matiaslindgren/AoC2020/util"
	"strings"
)

var Split = strings.Split

type Rule struct {
	Name                   string
	aMin, aMax, bMin, bMax int
}

func parseRule(r string) Rule {
	parts := Split(r, ": ")
	name, parts := parts[0], Split(parts[1], " ")
	a, b := Split(parts[0], "-"), Split(parts[2], "-")
	f := util.ParseInt
	return Rule{name, f(a[0]), f(a[1]), f(b[0]), f(b[1])}
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

func parseTicket(t string) []int {
	return util.ParseIntArray(Split(t, ","))
}

func (r Rule) valid(x int) bool {
	return (r.aMin <= x && x <= r.aMax) || (r.bMin <= x && x <= r.bMax)
}

func hasValidRule(rules []Rule, x int) bool {
	for _, rule := range rules {
		if rule.valid(x) {
			return true
		}
	}
	return false
}

func makeFilterTable(a, b int) [][]bool {
	T := make([][]bool, a)
	for i := 0; i < a; i++ {
		T[i] = make([]bool, b)
		for j := range T[i] {
			T[i][j] = true
		}
	}
	return T
}

func search(input string) (int, int) {
	rules, myTicket, tickets := parseSections(input)

	a := 0
	validTickets := [][]int{}
	for _, strTicket := range tickets {
		ticket := parseTicket(strTicket)
		ok := true
		for _, x := range ticket {
			if !hasValidRule(rules, x) {
				a += x
				ok = false
				break
			}
		}
		if ok {
			validTickets = append(validTickets, ticket)
		}
	}

	ticketLen := len(validTickets[0])
	pos2validRules := makeFilterTable(ticketLen, len(rules))
	for pos, valid := range pos2validRules {
		for r, rule := range rules {
			for _, ticket := range validTickets {
				valid[r] = valid[r] && rule.valid(ticket[pos])
			}
		}
	}

	pos2rule := make([]Rule, ticketLen)
	for ruleTaken := map[int]bool{}; len(ruleTaken) < len(rules); {
		for pos, validRules := range pos2validRules {
			numFree, prev := 0, 0
			for r, valid := range validRules {
				if valid && !ruleTaken[r] {
					numFree++
					prev = r
				}
			}
			if numFree == 1 {
				pos2rule[pos] = rules[prev]
				ruleTaken[prev] = true
			}
		}
	}

	b := 1
	for pos, x := range util.ParseIntArray(Split(myTicket, ",")) {
		if strings.Contains(pos2rule[pos].Name, "departure") {
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
