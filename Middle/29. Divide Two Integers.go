package main

func divide(dividend int, divisor int) int {
	if dividend/divisor < -(1 << 31) {
		return -(1 << 31)
	} else if dividend/divisor > 1<<31-1 {
		return 1<<31 - 1
	}
	return dividend / divisor
}
