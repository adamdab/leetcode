package main

func unique(nums []int) []int {
	size := 0
	for _, n := range nums {
		if n != nums[size] {
			size++
			nums[size] = n
		}
	}
	return nums[:size+1]
}

func isValley(left, middle, right int) bool {
	return left > middle && middle < right
}

func isHill(left, middle, right int) bool {
	return left < middle && middle > right
}

func countHillValley(nums []int) int {
	res := 0
	_unique := unique(nums)
	for i := 1; i < len(_unique)-1; i++ {
		if isValley(_unique[i-1], _unique[i], _unique[i+1]) || isHill(_unique[i-1], _unique[i], _unique[i+1]) {
			res++
		}
	}
	return res
}
