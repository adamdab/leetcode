package main

import "sort"

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions[:])
	res := make([]int, len(spells))
	for index, spell := range spells {
		minPotion := (success + int64(spell) - 1) / int64(spell)
		res[index] = len(potions) - binarySearchAtLease(potions, minPotion)
	}
	return res
}

func binarySearchAtLease(slice []int, element int64) int {
	low := 0
	high := len(slice) - 1

	if int64(slice[low]) >= element {
		return 0
	}
	if int64(slice[high]) < element {
		return len(slice)
	}

	for low < high {
		mid := low + (high-low)/2

		if int64(slice[mid]) >= element {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}
