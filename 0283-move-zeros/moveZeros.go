package main


func moveZeros(nums []int) {
  slow, quick := 0, 0
  for ; quick < len(nums); quick++ {
    if nums[quick] != 0 {
      nums[slow] = nums[quick]
      slow++
    }
  }
  for ; slow < len(nums); slow ++ {
    nums[slow] = 0
  }
}
