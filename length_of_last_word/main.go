package main

import "fmt"

func main() {
	num := lengthOfLastWord("a ")
	fmt.Println(num)
}

func lengthOfLastWord(s string) int {
	lenNum := 0
	runeS := []rune(s)
	flag := false
	for i := len(s) - 1; i >= 0; i-- {
		if runeS[i] == ' ' {
			if flag {
				continue
			}

		}
		if runeS[i] != ' ' {
			flag = true
			lenNum++
			if i == 0 {
				break

			}
			if runeS[i-1] == ' ' {
				break
			}
		}

	}
	return lenNum
}
