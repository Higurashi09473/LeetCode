package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	ss, tt := countChar(s), countChar(t)
	for k, v := range ss {
		v2, ok := tt[k]
		if !ok || v2 != v {
			return false
		}
	}
	return true
}

func countChar(x string) map[rune]int {
	Xmap := make(map[rune]int)
	for _, v := range x {
		if n, ok := Xmap[v]; ok {
			Xmap[v] = n + 1
		} else {
			Xmap[v] = 1
		}
	}
	return Xmap
}
