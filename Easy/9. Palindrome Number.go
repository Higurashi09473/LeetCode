package main

import "fmt"

func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    array, arrReversed := fmt.Sprint(x) , ""
    for x != 0{
        arrReversed += fmt.Sprint(x % 10)
        x /= 10
    } 
    if array == arrReversed || array == "0"{
        return true
    } else{
        return false
    }
}