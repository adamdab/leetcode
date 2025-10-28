package main

func countValidSelections(nums []int) int {
	left := make([]int, len(nums))
	right := make([]int, len(nums))
	counter := 0
	for i, val := range nums {
		counter += val
		left[i] = counter
	}

	counter = 0
	for i := len(nums) - 1; i > -1; i-- {
		counter += nums[i]
		right[i] = counter
	}
	count := 0

	for i := range nums {
		if nums[i] == 0 {
			prev := getWithDefault(left, i-1, 0)
			next := getWithDefault(right, i+1, 0)
			if prev == next {
				count += 2
			}
			if abs(prev-next) == 1 {
				count++
			}
		}
	}

	return count
}

func getWithDefault(arr []int, index, def int) int {
	if index < 0 || index >= len(arr) {
		return def
	}
	return arr[index]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
