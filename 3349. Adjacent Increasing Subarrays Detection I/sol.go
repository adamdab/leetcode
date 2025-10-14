package main

func hasIncreasingSubarrays(nums []int, k int) bool {

	if len(nums) < 2*k {
		return false
	}

	for start := 0; start < len(nums)-k; {

		sizeFirst := getMaxSizeOfIncreasing(start, nums)
		if sizeFirst >= 2*k {
			return true
		}

		if sizeFirst >= k {
			sizeSecond := getMaxSizeOfIncreasing(start+sizeFirst, nums)
			if sizeSecond >= k {
				return true
			} else {
				start += sizeFirst + sizeSecond
			}
		} else {
			start += sizeFirst
		}
	}

	return false
}

func getMaxSizeOfIncreasing(start int, nums []int) int {
	for i := 1; i < len(nums)-start; i++ {
		if nums[start+i-1] >= nums[start+i] {
			return i
		}
	}

	return len(nums) - start
}
