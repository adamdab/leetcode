package main

func maxIncreasingSubarrays(nums []int) int {
	maxLen := 1
	prev := 0
	curr := 0
	for i := range nums {
		curr++
		if (i == len(nums)-1) || (nums[i] >= nums[i+1]) {
			maxLen = max(max(maxLen, curr/2), min(prev, curr))
			prev = curr
			curr = 0
		}
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
