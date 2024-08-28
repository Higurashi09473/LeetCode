package main

func getConcatenation(nums []int) []int {
	dl := len(nums)
	for i := 0; i < dl; i++ {
		nums = append(nums, nums[i])
	}
	return nums
}

// func getConcatenation(nums []int) []int {
//     return append(nums, nums...)
// }
