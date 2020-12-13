package main

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/matiaslindgren/AoC2020/util"
)

func searchB(ids []*big.Int) int64 {
	N := big.NewInt(1)
	for _, n := range ids {
		if n != nil {
			N.Mul(N, n)
		}
	}

	one := big.NewInt(1)
	k := big.NewInt(int64(len(ids)-1))
	var i, a, m1, m2, n2, gcd big.Int
	sum := big.NewInt(0)

	for j, n1 := range ids {
		if n1 == nil {
			continue
		}
		n2.Set(N)
		n2.Div(&n2, n1)
		gcd.GCD(&m1, &m2, n1, &n2)
		if gcd.Cmp(one) != 0 {
			panic("n1 and n2 must both be primes")
		}
		a.Set(k)
		a.Sub(&a, i.SetInt64(int64(j)))
		sum.Add(sum, a.Mul(a.Mul(&a, &m2), &n2))
	}

	return sum.Mod(sum, N).Sub(sum, k).Int64()
}

func searchA(ids []*big.Int, earliest int64) int64 {
	a := int64(0)
	min := int64(1<<63-1)
	for _, nBig := range ids {
		if nBig == nil {
			continue
		}
		n := nBig.Int64()
		if x := n - earliest%n; x < min {
			min = x
			a = n * x
		}
	}
	return int64(a)
}

func parseBigInts(line string) []*big.Int {
	strIds := strings.Split(line, ",")
	ids := make([]*big.Int, len(strIds))
	for i, strId := range strIds {
		if strId != "x" {
			ids[i] = big.NewInt(int64(util.ParseInt(strId)))
		}
	}
	return ids
}

func search(lines []string) (int64, int64) {
	earliest, ids := int64(util.ParseInt(lines[0])), parseBigInts(lines[1])
	a := searchA(ids, earliest)
	b := searchB(ids)
	return a, b
}

func main() {
	input := util.SlurpStdinLines()
	a, b := search(input)
	fmt.Println(a, b)
}
