package main

import (
	"fmt"
	"container/ring"
	"github.com/matiaslindgren/AoC2020/util"
)

type Ring struct {
	n int
	r *ring.Ring
	val2node map[int]*ring.Ring
}

func newRing(v []int, n int) Ring {
	ring := Ring{n, ring.New(n), map[int]*ring.Ring{}}
	r := ring.r
	for i := 0; i < n; i++ {
		if i < len(v) {
			r.Value = v[i]
		} else {
			r.Value = i+1
		}
		ring.val2node[r.Value.(int)] = r
		r = r.Next()
	}
	return ring
}

func contains(r *ring.Ring, x int) bool {
	yes := false
	r.Do(func(val interface{}) { yes = yes || val.(int)==x })
	return yes
}

func (r *Ring) step() {
	cup1 := r.r.Value.(int)
	three := r.r.Unlink(3)
	cup2 := cup1
	for cup1 == cup2 || contains(three, cup2) {
		cup2 = ((cup2-2) % r.n + r.n) % r.n + 1
	}
	r.val2node[cup2].Link(three)
	r.r = r.val2node[cup1].Next()
}

func searchA(cups []int) string {
	r := newRing(cups, len(cups))
	for i := 0; i < 100; i++ {
		r.step()
	}
	res := ""
	r.val2node[1].Do(func(p interface{}) {
		res += fmt.Sprint(p.(int))
	})
	return res[1:]
}

func searchB(cups []int) int {
	r := newRing(cups, 1_000_000)
	for i := 0; i < 10_000_000; i++ {
		r.step()
	}
	after1 := r.val2node[1].Next()
	return after1.Value.(int) * after1.Next().Value.(int)
}

func main() {
	input := util.SlurpStdinArray("")
	a := searchA(input)
	b := searchB(input)
	fmt.Println(a, b)
}
