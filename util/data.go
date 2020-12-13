package util

func Set(v []int) map[int]bool {
	s := map[int]bool{}
	for _, x := range v {
		s[x] = true
	}
	return s
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinMax(v []int) (int, int) {
	min, max := v[0], v[0]
	for _, x := range v[1:] {
		min = Min(min, x)
		max = Max(max, x)
	}
	return min, max
}
