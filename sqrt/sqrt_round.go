package main

func main() {
	//x := mySqrt(10)
	//fmt.Println(x)
}

func mySqrt(x int) int {
	if x == 1 || x == 0 {
		return x * 1
	}

	left := 0
	right := x / 2

	for left <= right {
		mid := (left + right) / 2
		switch {
		case mid*mid == x:
			return mid
		case mid*mid < x:
			left = mid + 1
		case mid*mid > x:
			right = mid - 1

		}

	}
	return right

}
