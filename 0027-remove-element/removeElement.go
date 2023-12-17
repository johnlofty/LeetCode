package main

import "fmt"


func removeElement(nums []int, val int) int  {
  quick, slow := 0, 0
  for ; quick < len(nums); quick ++ {
    if nums[quick] != val {
      nums[slow] = nums[quick]
      slow++
    }
  }
  return slow
}

func main() {
  arrays := []int{0,1,2,2,3,0,4,2}
  fmt.Printf("removeElement len:%d content:%v\n", removeElement(arrays, 2), arrays)
}
