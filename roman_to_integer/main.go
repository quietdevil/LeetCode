package main

import "fmt"

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000
func romanToInt(s string) int {
	mapRoman := make(map[string]int)
	mapRoman["I"] = 1
	mapRoman["V"] = 5
	mapRoman["X"] = 10
	mapRoman["L"] = 50
	mapRoman["C"] = 100
	mapRoman["D"] = 500
	mapRoman["M"] = 1000
	lenS := len(s)
	c, _ := mapRoman[string(s[lenS-1])]
	var integer int = c
	prev_value := c
	for i := lenS - 2; i > -1; i-- {
		v, _ := mapRoman[string(s[i])]
		if v < prev_value {
			integer -= v
		} else {
			integer += v
		}

		prev_value = v
	}
	return integer
}
