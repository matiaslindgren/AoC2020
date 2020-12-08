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

func tryRun(lines []string, mutateLine int) (int, bool) {
	acc, i := 0, 0
	for seen := map[int]bool{}; !seen[i] && i < len(lines); i++ {
		seen[i] = true
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
	return acc, i < len(lines)
}

func search(lines []string) (int, int) {
	a, _ := tryRun(lines, -1)
	b := 0
	for mutateLine, line := range lines {
		ins, _ := parseInstruction(line)
		if ins == "nop" || ins == "jmp" {
			if acc, earlyExit := tryRun(lines, mutateLine); !earlyExit {
				b = acc
				break
			}
		}
	}
	return a, b
}

func main() {
	input := util.SlurpStdinLines()
	a, b := search(input)
	fmt.Println(a, b)
}
