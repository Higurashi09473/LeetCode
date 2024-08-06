package main

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool, len(nums))
	for _, num := range nums {
		if _, ok := m[num]; ok {
			return true
		}
		m[num] = true
	}
	return false
}
