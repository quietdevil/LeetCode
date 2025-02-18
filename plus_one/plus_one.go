package main

import "fmt"

func main() {
	slice := plusOne([]int{1, 9, 9})
	fmt.Println(slice)
}

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+1 == 10 {
			digits[i] = 0
		} else {
			digits[i]++
			return digits
		}

	}
	digits[0] = 1
	digits = append(digits, 0)
	return digits
}
