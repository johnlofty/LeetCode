package main

import "fmt"

func isPerfectSquare(num int) bool {
	left, right := 0, num
	for left <= right {
		mid := left + (right-left)>>1
		tmp := mid * mid
		if tmp > num {
			right = mid - 1
		} else if tmp < num {
			left = mid + 1
		} else {
			return true
		}
	}
	return false
}

func main() {
  fmt.Printf("%d is perfect square %t\n", 8, isPerfectSquare(8))
  fmt.Printf("%d is perfect square %t\n", 4, isPerfectSquare(4))
}
