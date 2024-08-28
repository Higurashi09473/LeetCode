package main

func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
	var count int
	for _, v := range hours {
		if v >= target {
			count++
		}
	}
	return count
}
