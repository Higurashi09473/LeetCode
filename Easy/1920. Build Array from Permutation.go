package main

func buildArray(nums []int) []int {
	mas := make([]int, len(nums))
	for i, v := range nums {
		mas[i] = nums[v]
	}
	return mas
}
