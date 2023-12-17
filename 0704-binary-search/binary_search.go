package main

import (
  "fmt"
)


func search(nums []int, target int) int {
	left, right := 0, len(nums) -1

	for left <= right {
		mid := (right + left) / 2
    fmt.Printf("start:%d end:%d mid:%d\n", left, right, mid)
		if nums[mid] > target {
			right = mid -1
		} else if nums[mid] <  target {
      left = mid + 1
		} else {
			return mid
		}

	}
	return -1
}

func main() {
  fmt.Println("ans", search([]int{-1, 0, 3,5,9, 12}, 9))
  fmt.Println("ans", search([]int{-1, 0, 3,5,9, 12}, 2))
  fmt.Println("ans", search([]int{5}, 5))
}
