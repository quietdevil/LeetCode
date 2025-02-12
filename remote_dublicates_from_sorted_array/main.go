package main

import "fmt"

func main() {
	k := removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
	fmt.Println(k)
}

func removeDuplicates(nums []int) int {
	i := 0

	for j := 0; j < len(nums); {
		if nums[i] != nums[j] {
			nums[i+1] = nums[j]
			i++
		}
		j++
	}

	return i + 1

}
