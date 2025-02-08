package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(0))
}

func isPalindrome(x int) bool {
	if x == 0 {
		return true
	}
	z := 0
	xcopy := x
	for x != 0 && x > 0 {
		a := x % 10
		z += a
		x = x / 10
		if z == xcopy {
			return true
		}
		z *= 10
	}
	return false
}
