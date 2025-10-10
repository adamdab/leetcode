package main

import (
	"math"
)

func maximumEnergy(energy []int, k int) int {
	maxValue := math.MinInt
	for end := len(energy) - k; end < len(energy); end++ {
		curr := 0
		for index := end; index >= 0; index -= k {
			curr = curr + energy[index]
			maxValue = max(maxValue, curr)
		}
	}
	return maxValue
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
