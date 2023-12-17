package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	quick, slow := 0, 0
	target := 0
	for ; quick < len(nums); quick++ {
		if slow == 0 || nums[quick] != target {
			nums[slow] = nums[quick]
			target = nums[slow]
			slow++
		}
	}
	return slow
}

func removeDuplicatesV2(nums []int) int {
  if len(nums) <= 1 {
    return len(nums)
  }

  count := 1
  for i := 1; i < len(nums); i++ {
    if nums[i] != nums[i-1] {
      nums[count] = nums[i]
      count++
    }
  }
  return count
}


func main() {
	cases := []struct {
		input  []int
		expect []int
		size   int
	}{
		{
			[]int{1, 1, 2},
			[]int{1, 2},
			2,
		},
    {
      []int{1, 1, 2, 2, 2, 3, 4, 4, 5},
      []int{1,2,3,4,5},
      5,
    },
	}
	for _, c := range cases {
		// s := removeDuplicates(c.input)
		s := removeDuplicatesV2(c.input)
		f := func(in, out []int, s int) bool {
			for i := 0; i < s; i++ {
				if out[i] != in[i] {
					return false
				}
			}
			return true
		}

		if !f(c.input, c.expect, s) {
			fmt.Printf("input:%v expect:%v size:%d\n", c.input, c.expect, s)
		}
	}

  fmt.Println("success")
}
