package main

import "fmt"

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > target {
      right = mid - 1
		} else if nums[mid] < target {
      left = mid + 1
    } else {
      return mid
    }
	}
  return right + 1
  
}

func main() {
	fmt.Println("ans:", searchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println("ans:", searchInsert([]int{1, 3, 5, 6}, 2))
	fmt.Println("ans:", searchInsert([]int{1, 3, 5, 6}, 7))
}
