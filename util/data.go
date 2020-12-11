package util

func Set(v []int) map[int]bool {
	s := map[int]bool{}
	for _, x := range v {
		s[x] = true
	}
	return s
}
