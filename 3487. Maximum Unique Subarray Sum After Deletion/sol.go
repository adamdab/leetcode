package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSum(nums []int) int {
	seen := [101]bool{}
	_max := -100
	for _, n := range nums {
		_max = max(_max, n)
		if n > 0 {
			seen[n] = true
		}
	}
	res := _max
	for i := _max - 1; i > 0; i-- {
		if seen[i] {
			res += i
		}
	}
	return res
}
