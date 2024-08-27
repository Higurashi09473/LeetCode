package main

func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums)-1
	mid := 0
	for low <= high {
		mid = (low + high) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	if nums[mid] < target {
		return mid + 1
	}
	return mid
}
