package main

import "fmt"

func main() {

	fmt.Println(twoSumV3([]int{3, 2, 4}, 6))
	fmt.Println(twoSumV3([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSumV3([]int{3, 2, 3}, 6))
	fmt.Println(twoSumV3([]int{1, 2, 3}, 5))
	fmt.Println(twoSumV3([]int{1, 1, 1, 1, 1, 4, 1, 1, 1, 1, 1, 7, 1, 1, 1, 1, 1}, 11))
}

func twoSumV1(nums []int, target int) []int {
	// 30ms
	if len(nums) == 2 {
		return []int{0, 1}
	}
	k := 1
	for i := 0; i < len(nums); i++ {
		arrcopy := nums[i+1:]
		for j := 0; j < len(arrcopy); j++ {
			if nums[i]+arrcopy[j] == target {
				return []int{i, k + j}
			}
		}
		k++
	}
	return nil

}

func twoSumV2(nums []int, target int) []int {
	//20ms
	if len(nums) == 2 {
		return []int{0, 1}
	}
	hashMap := make(map[int]bool)
	k := 1
	for i := 0; i < len(nums); i++ {
		if _, ok := hashMap[nums[i]]; ok {
			k++
			continue
		}
		hashMap[nums[i]] = true
		arrcopy := nums[i+1:]
		for j := 0; j < len(arrcopy); j++ {
			if nums[i]+arrcopy[j] == target {
				return []int{i, k + j}
			}
		}
		k++
	}
	return nil

}

func twoSumV3(nums []int, target int) []int {
	// 0ms
	hashMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		tr := target - nums[i]
		trInd, isPres := hashMap[tr]

		if isPres {
			return []int{trInd, i}
		}

		hashMap[nums[i]] = i
	}
	return nil
}
