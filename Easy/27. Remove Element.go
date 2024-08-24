package main

func removeElement(nums []int, val int) int {
	var n int
	for _, num := range nums {
		if num != val {
			nums[n] = num
			n++
		}
	}
	return n
}
