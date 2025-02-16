package main

func main() {
	searchInsert([]int{1, 2, 3}, 2)
}

func searchInsert(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := (low + high) / 2
		guess := nums[mid]
		switch {
		case guess == target:
			return mid
		case guess > target:
			high = mid - 1
		case guess < target:
			low = mid + 1
		}
	}
	return low
}
