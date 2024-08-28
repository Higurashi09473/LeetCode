package main

func finalValueAfterOperations(operations []string) int {
	x := 0
	for _, value := range operations {
		if value[1] == '+' {
			x++
		} else {
			x--
		}
	}
	return x
}
