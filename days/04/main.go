package main

import (
	"fmt"
	"strings"
	"github.com/matiaslindgren/AoC2020/util"
)

type Passport map[string]string

func emptyPassport() Passport {
	return Passport{
		"byr": "",
		"cid": "",
		"ecl": "",
		"eyr": "",
		"hcl": "",
		"hgt": "",
		"iyr": "",
		"pid": "",
	}
}

func (p Passport) isFilled() bool {
	for k, v := range p {
		if k != "cid" && len(v) == 0 {
			return false
		}
	}
	return true
}

func (p Passport) isValid() bool {
	if !p.isFilled() {
		return false
	}
	for k, v := range p {
		if !isValidField(k, v) {
			return false
		}
	}
	return true
}

func isValidField(key, val string) bool {
	switch key {
	case "cid":
		return true
	case "byr":
		if util.Match(`^\d{4}$`, val) {
			byr := util.ParseInt(val)
			return 1920 <= byr && byr <= 2002
		}
	case "iyr":
		if util.Match(`^\d{4}$`, val) {
			iyr := util.ParseInt(val)
			return 2010 <= iyr && iyr <= 2020
		}
	case "eyr":
		if util.Match(`^\d{4}$`, val) {
			eyr := util.ParseInt(val)
			return 2020 <= eyr && eyr <= 2030
		}
	case "hgt":
		if util.Match(`^\d{2,3}(cm|in)$`, val) {
			hgt := util.ParseInt(val[:len(val)-2])
			switch val[len(val)-2:] {
			case "cm":
				return 150 <= hgt && hgt <= 193
			case "in":
				return 59 <= hgt && hgt <= 76
			}
		}
	case "hcl":
		if util.Match(`^#[0-9a-f]{6}$`, val) {
			return true
		}
	case "ecl":
		if util.Match(`^(amb|blu|brn|gry|grn|hzl|oth)$`, val) {
			return true
		}
	case "pid":
		if util.Match(`^\d{9}$`, val) {
			return true
		}
	}
	return false
}

func search(sections []string) (int, int) {
	a, b := 0, 0
	for _, section := range sections {
		p := emptyPassport()
		for _, line := range strings.Split(section, "\n") {
			line = strings.TrimSpace(line)
			for _, token := range strings.Split(line, " ") {
				keyval := strings.Split(strings.TrimSpace(token), ":")
				p[keyval[0]] = keyval[1]
			}
		}
		if p.isFilled() {
			a++
		}
		if p.isValid() {
			b++
		}
	}
	return a, b
}

func main() {
	input := util.SlurpStdinSections()
	a, b := search(input)
	fmt.Println(a, b)
}
