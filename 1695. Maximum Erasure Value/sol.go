package main

func first(nums []int, left, right int) int {
	for i := left; i < right; i++ {
		if nums[i] == nums[right] {
			return i
		}
	}
	return right
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maximumUniqueSubarray(nums []int) int {
	curr := 0
	_max := 0
	left, right := 0, 0
	for right < len(nums) {
		if f := first(nums, left, right); f == right {
			curr += nums[right]
		} else {
			_max = max(_max, curr)
			for i := left; i < f; i++ {
				curr -= nums[i]
			}
			left = f + 1
		}
		right++
	}
	_max = max(_max, curr)
	return _max
}
