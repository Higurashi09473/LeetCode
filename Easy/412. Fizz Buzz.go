package main

import "fmt"

func fizzBuzz(n int) []string {
	mas := make([]string, n)
	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			mas[i-1] = "FizzBuzz"
		} else if i%3 == 0 {
			mas[i-1] = "Fizz"
		} else if i%5 == 0 {
			mas[i-1] = "Buzz"
		} else {
			mas[i-1] = fmt.Sprint(i)
		}
	}
	return mas
}
