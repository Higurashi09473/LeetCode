package main

func containsDuplicate(nums []int) bool {
	array := map[int]bool{}
	for i := 0; i < len(nums); i++ {
		array[nums[i]] = true
	}
	return len(nums) != len(array)
}
