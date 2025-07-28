package main

func OR(nums []int) int {
	res := 0
	for _, n := range nums {
		res = res | n
	}
	return res
}

func generateSubstest(lookedAt int, arr []int, rollingOR, maxOR int, include bool) int {
	if include {
		rollingOR = rollingOR | arr[lookedAt]
	}
	if lookedAt == len(arr)-1 {
		if rollingOR == maxOR {
			return 1
		}
		return 0
	}
	return generateSubstest(lookedAt+1, arr, rollingOR, maxOR, false) + generateSubstest(lookedAt+1, arr, rollingOR, maxOR, true)
}

func countMaxOrSubsets(nums []int) int {
	maxOR := OR(nums)
	return generateSubstest(-1, nums, 0, maxOR, false)
}
