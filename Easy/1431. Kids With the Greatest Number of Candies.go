package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	mas := make([]bool, len(candies))
	max := 0
	for _, v := range candies {
		if max < v {
			max = v
		}
	}
	for i, v := range candies {
		if v+extraCandies >= max {
			mas[i] = true
		} else {
			mas[i] = false
		}
	}
	return mas
}
