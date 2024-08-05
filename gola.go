package main

import "fmt"

func main() {
	s := "anagram"
	t := "nagaram"
	if len(s) != len(t) {
		fmt.Println(false)
	}
	ss, tt := countChar(s), countChar(t)
	for k, v := range ss {
		v2, ok := tt[k]
		if !ok || v2 != v {
			fmt.Println(false)
		}
	}
	fmt.Println(true)
}

func countChar(s string) map[rune]int {
	ss := make(map[rune]int)
	for _, v := range s {
		if n, ok := ss[v]; ok {
			ss[v] = n + 1
		} else {
			ss[v] = 1
		}
	}
	return ss
}
