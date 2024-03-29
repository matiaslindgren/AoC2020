package main

import (
	"fmt"
	"github.com/matiaslindgren/AoC2020/util"
	"regexp"
	"strings"
)

func parseMemoryAccess(line string) (int, int) {
	re := regexp.MustCompile(`mem\[(\d+)\] = (\d+)`)
	m := re.FindSubmatch([]byte(line))
	return util.ParseInt(string(m[1])), util.ParseInt(string(m[2]))
}

func parseSection(s string) (string, []string) {
	lines := strings.Split(strings.TrimSpace(s), "\n")
	return lines[0], lines[1:]
}

type Memory map[int]int

func (m Memory) sum() int {
	s := 0
	for _, x := range m {
		s += x
	}
	return s
}

func applyValMask(mask string, val int) int {
	for i, n := 0, len(mask); i < n; i++ {
		bit := 1 << (n - i - 1)
		switch mask[i] {
		case '1':
			val |= bit
		case '0':
			val &= ^bit
		}
	}
	return val
}

func applyIdxMask(mask string, idx int) []int {
	idxCombinations := []int{0}
	idxMasked := 0
	for i, n := 0, len(mask); i < n; i++ {
		bit := 1 << (n - i - 1)
		switch mask[i] {
		case 'X':
			tmp := []int{}
			for _, idx := range idxCombinations {
				tmp = append(tmp, idx, idx|bit)
			}
			idxCombinations = tmp
		case '1':
			idxMasked |= bit
		case '0':
			idxMasked |= bit & idx
		}
	}
	res := make([]int, len(idxCombinations))
	for i, idx1 := range idxCombinations {
		res[i] = idx1 | idxMasked
	}
	return res
}

func search(sections []string) (int, int) {
	memA, memB := Memory{}, Memory{}
	for _, section := range sections {
		mask, accesses := parseSection(section)
		for _, line := range accesses {
			idx, val := parseMemoryAccess(line)
			memA[idx] = applyValMask(mask, val)
			for _, idx1 := range applyIdxMask(mask, idx) {
				memB[idx1] = val
			}
		}
	}
	return memA.sum(), memB.sum()
}

func main() {
	input := util.SlurpStdin()
	sections := strings.Split(input, "mask = ")
	a, b := search(sections)
	fmt.Println(a, b)
}
