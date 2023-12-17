package main

func sortedSquares(nums []int) []int {
	l, r := 0, len(nums)-1
	ans := make([]int, len(nums))
	pos := len(nums) - 1

	for l <= r {
		vl := nums[l] * nums[l]
		vr := nums[r] * nums[r]
		if vr > vl {
			ans[pos] = vr
			r--
		} else {
			ans[pos] = vl
			l++
		}
		pos--
	}
	return ans
}
