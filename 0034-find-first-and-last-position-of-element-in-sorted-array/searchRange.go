package main

import "fmt"

/* the crux of the problem is to find the left border and right border
of the target in arrays.

to find the left/right border of the target we can use binary_search
*/

// we need to find the target and the biggest index.
// with binary search we can find any index of the target, if there are
// multiple targets in arrays.

func searchRightBorder(nums []int, target int) int {
  ans := -2  // default to -2 to differ from -1
  left, right := 0, len(nums) - 1
  
  // the [left, right] search range.
  for left <= right {
    mid := (left + right) / 2
    fmt.Printf("left:%d right:%d mid:%d\n", left, right, mid)
    if nums[mid] > target {
      right = mid - 1
    } else {
      left = mid + 1
      ans = left
    }
  }
  return ans
}

func searchLeftBorder(nums []int, target int) int {
  ans := -2 
  left, right := 0, len(nums)-1 
  for left <= right {
    mid := (left + right) / 2
    if nums[mid] >= target {
      right = mid - 1   
      ans = right
    } else {
      left = mid + 1
    }
  }
  return ans
}


func searchRange(nums []int, target int) []int {
  leftBorder, rightBorder := searchLeftBorder(nums, target), searchRightBorder(nums, target)
  fmt.Printf("==> left:%d right:%d\n", leftBorder,rightBorder)
  if leftBorder == -2 || rightBorder == -2 {
    return []int{-1, -1}
  }
  if (rightBorder - leftBorder) > 1 {
    return []int{leftBorder + 1, rightBorder -1}
  }
  return []int{-1,-1}
}

func main() {
  fmt.Printf("ans:%v\n", searchRange([]int{1,2,3,3,5,6}, 3))

}
