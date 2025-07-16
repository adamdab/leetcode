package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maximumLength(nums []int) int {
	zeros := 0
	ones := 0
	alternatingOne := 0
	alternatingZero := 0
	alternateZero := false
	alternateOne := false
	for _, n := range nums {
		if n%2 == 0 {
			zeros++
			if !alternateZero {
				alternateZero = true
				alternatingZero++
			}
			if alternateOne {
				alternateOne = false
				alternatingOne++
			}
		} else {
			ones++
			if alternateZero {
				alternateZero = false
				alternatingZero++
			}
			if !alternateOne {
				alternateOne = true
				alternatingOne++
			}
		}
	}
	return max(max(ones, zeros), max(alternatingZero, alternatingOne))
}
