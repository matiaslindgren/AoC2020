package main

import (
	"fmt"
	"github.com/matiaslindgren/AoC2020/util"
)

// https://en.wikipedia.org/wiki/Shunting-yard_algorithm#The_algorithm_in_detail
func infix2postfix(input string, equalPrecedence bool) string {
	ops := []byte{}
	out := []byte{}
	for i := 0; i < len(input); i++ {
		token := input[i]
		switch token {
		case ' ':
		case ')':
			for i := len(ops)-1; ops[i] != '('; i-- {
				out = append(out, ops[i])
				ops = ops[:i]
			}
			ops = ops[:len(ops)-1]
		case '*':
			fallthrough
		case '+':
			for i := len(ops)-1; i >= 0 && (ops[i] == '+' || (equalPrecedence && ops[i] == '*')); i-- {
				out = append(out, ops[i])
				ops = ops[:i]
			}
			fallthrough
		case '(':
			ops = append(ops, token)
		default:
			out = append(out, token)
		}
	}
	for i := len(ops)-1; i >= 0; i-- {
		out = append(out, ops[i])
	}
	return string(out)
}

func sumLines(lines []string, equalPrecedence bool) int {
	s := 0
	for _, infix := range lines {
		postfix := infix2postfix(infix, equalPrecedence)
		res := []int{}
		for i := 0; i < len(postfix); i++ {
			n := len(res)
			switch postfix[i] {
			case '*':
				res[n-2] *= res[n-1]
				res = res[:n-1]
			case '+':
				res[n-2] += res[n-1]
				res = res[:n-1]
			default:
				res = append(res, int(postfix[i]-48))
			}
		}
		s += res[0]
	}
	return s
}

func main() {
	input := util.SlurpStdinLines()
	a := sumLines(input, true)
	b := sumLines(input, false)
	fmt.Println(a, b)
}
