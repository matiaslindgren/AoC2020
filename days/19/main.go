package main

import (
	"fmt"
	"sort"
	"strings"
	"github.com/matiaslindgren/AoC2020/util"
)

type Rule [][]string
type CNF map[string]Rule

func parseRule(r string) (string, Rule) {
	parts := strings.Split(r, ":")
	key := parts[0]
	parts = strings.Split(strings.TrimSpace(parts[1]), "|")
	rule := make(Rule, len(parts))
	for i, symbols := range parts {
		rule[i] = strings.Split(strings.TrimSpace(symbols), " ")
	}
	return key, rule
}

func isTerminal(symbols []string) bool {
	for _, symbol := range symbols {
		if symbol[0] == '"' {
			return true
		}
	}
	return false
}

// https://en.wikipedia.org/wiki/Chomsky_normal_form
func parseCNF(ruleLines []string) CNF {
	grammar := CNF{}
	for _, line := range ruleLines {
		key, rule := parseRule(line)
		grammar[key] = rule
	}
	grammar = grammar.removeUnitRules()
	return grammar
}

func (grammar CNF) removeUnitRules() CNF {
	for key, rule := range grammar {
		r := Rule{}
		for _, symbols := range rule {
			if len(symbols) == 1 && !isTerminal(symbols) {
				for _, syms := range grammar[symbols[0]] {
					r = append(r, syms)
				}
			} else {
				r = append(r, symbols)
			}
		}
		grammar[key] = r
	}
	return grammar
}

func (grammar CNF) maxKey() int {
	max := 0
	for k := range grammar {
		max = util.Max(max, util.ParseInt(k))
	}
	return max
}

// https://en.wikipedia.org/wiki/CYK_algorithm
func (grammar CNF) isInLanguage(input string) bool {
	n, r := len(input), len(grammar)
	P := make([][][]bool, n+1)
	for i := 0; i <= n; i++ {
		P[i] = make([][]bool, n+1)
		for j := 0; j <= n; j++ {
			P[i][j] = make([]bool, r+1)
		}
	}

	key2idx := map[string]int{}
	keys := []string{}
	for k := range grammar {
		keys = append(keys, k)
	}
	sort.Slice(keys, func (i, j int) bool {
		return util.ParseInt(keys[i]) < util.ParseInt(keys[j]) })
	for _, k := range keys {
		key2idx[k] = len(key2idx)+1
	}
	tableKey := func(ruleKey string) int { return key2idx[ruleKey] }

	for s := 1; s <= n; s++ {
		for k, rule := range grammar {
			v := tableKey(k)
			for _, symbols := range rule {
				if isTerminal(symbols) && symbols[0][1] == input[s-1] {
					P[1][s][v] = true
				}
			}
		}
	}

	for l := 2; l <= n; l++ {
		for s := 1; s <= n-l+1; s++ {
			for p := 1; p <= l-1; p++ {
				for k, rule := range grammar {
					a := tableKey(k)
					for _, symbols := range rule {
						if !isTerminal(symbols) {
							b, c := tableKey(symbols[0]), tableKey(symbols[1])
							if P[p][s][b] && P[l-p][s+p][c] {
								P[l][s][a] = true
							}
						}
					}
				}
			}
		}
	}

	return P[n][1][1]
}

func search(ruleSection, inputSection string) (int, int) {
	rules := strings.Split(ruleSection, "\n")
	input := strings.Split(inputSection, "\n")
	grammar := parseCNF(rules)

	a := 0
	for _, line := range input {
		if grammar.isInLanguage(strings.TrimSpace(line)) {
			a++
		}
	}

	grammar["8"] = Rule{[]string{"42"}, []string{"42", "8"}}
	newKey := fmt.Sprint(grammar.maxKey()+1)
	grammar["11"] = Rule{[]string{"42", "31"}, []string{"42", newKey}}
	grammar[newKey] = Rule{[]string{"11", "31"}}
	grammar = grammar.removeUnitRules()
	b := 0
	for _, line := range input {
		if grammar.isInLanguage(strings.TrimSpace(line)) {
			b++
		}
	}

	return a, b
}

func main() {
	input := util.SlurpStdinSections()
	a, b := search(input[0], input[1])
	fmt.Println(a, b)
}
