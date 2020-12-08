package main

import (
	"fmt"
	"strings"
	"github.com/matiaslindgren/AoC2020/util"
)

func parseInstruction(line string) (string, int) {
	parts := strings.Split(line, " ")
	return parts[0], util.ParseInt(parts[1])
}

func run(lines []string, mutateLine int) (int, bool) {
	acc := 0
	visited := map[int]bool{}
	i := 0
	for ; !visited[i] && i < len(lines); i++ {
		visited[i] = true
		ins, num := parseInstruction(lines[i])
		if i == mutateLine {
			switch ins {
			case "nop":
				ins = "jmp"
			case "jmp":
				ins = "nop"
			}
		}
		switch ins {
		case "acc":
			acc += num
		case "jmp":
			i += num - 1
		}
	}
	return acc, i == len(lines)
}

func search(lines []string) (int, int) {
	a, _ := run(lines, -1)
	b := 0
	for mutateLine, line := range lines {
		ins, _ := parseInstruction(line)
		switch ins {
		case "nop":
		case "jmp":
		default:
			continue
		}
		if acc, visitedAll := run(lines, mutateLine); visitedAll {
			b = acc
			break
		}
	}
	return a, b
}

func main() {
	input := util.SlurpStdinLines()
	a, b := search(input)
	fmt.Println(a, b)
}
