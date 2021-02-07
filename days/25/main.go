package main

import (
	"fmt"
	"github.com/matiaslindgren/AoC2020/util"
)

const M = 20201227

func transform(subject, loopSize int) int {
	x := 1
	for l := 1; l <= loopSize; l++ {
		x *= subject
		x %= M
	}
	return x
}

func searchLoopSize(subject, key int) int {
	x := 1
	for l := 0; ; l++ {
		if x == key {
			return l
		}
		x *= subject
		x %= M
	}
}

func search(cardKey, doorKey int) int {
	cardL := searchLoopSize(7, cardKey)
	doorL := searchLoopSize(7, doorKey)
	encKey := transform(cardKey, doorL)
	if k := transform(doorKey, cardL); k != encKey {
		panic(fmt.Sprintf("mismatching encryption keys: %d != %d", encKey, k))
	}
	return encKey
}

func main() {
	input := util.SlurpStdinLines()
	cardKey, doorKey := util.ParseInt(input[0]), util.ParseInt(input[1])
	fmt.Println(search(cardKey, doorKey))
}
