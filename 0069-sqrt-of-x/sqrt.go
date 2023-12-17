package main

import (
  "fmt"
)

func mySqrt(x int) int {
  left, right := 1, x
  for left <= right {
    mid := left + (right - left) / 2
    val := mid * mid
    if val > x {
      right = mid - 1
    } else {
      left = mid + 1
    }
  }
  return left - 1
}

func main() {
  fmt.Printf("sqrt root of 2:%d\n", mySqrt(2))
  fmt.Printf("sqrt root of 4:%d\n", mySqrt(4))
  fmt.Printf("sqrt root of 8:%d\n", mySqrt(8))
}
