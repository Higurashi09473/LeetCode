package main

func countBits(n int) []int {
	mas := make([]int, n+1)
	var count, j int
	for i := 0; i <= n; i++ {
		count = 0
		j = i
		for j != 0 {
			if j%2 != 0 {
				count++
			}
			j /= 2
		}
		mas[i] = count
	}
	return mas
}
