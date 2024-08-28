package main

func topKFrequent(nums []int, k int) []int {
	cart := make(map[int]int)
	for _, value := range nums {
		cart[value] += 1
	}
	otv := []int{}
	max := 0
	for i := 0; i < k; i++ {
		for _, j := range nums {
			if cart[max] < cart[j] {
				max = j
			}
		}
		otv = append(otv, max)
		cart[max] = 0
		max = 0
	}
	return otv
}
