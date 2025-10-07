package main

func avoidFlood(rains []int) []int {
	ans := make([]int, 0, len(rains))
	full_lakes := make(map[int]int) // lake to last not dried day
	dry_days := []int{}

	for len(ans) < len(rains) {
		ans = append(ans, -1)
	}

	for index, val := range rains {
		if val == 0 { // dry days before any rain are useless
			if len(full_lakes) > 0 {
				dry_days = append(dry_days, index)
			}
			ans[index] = 1 // by default we put 1 as it can not be -1
		} else {
			if day, ok := full_lakes[val]; ok {
				if len(dry_days) == 0 {
					return []int{}
				}
				// is full and now it rains
				// check if we can dry it before last rain
				first_avalaible_dry_day := findFirstIndexOfElementNotSmallerThan(dry_days, day)
				if first_avalaible_dry_day == -1 {
					return []int{}
				}
				ans[dry_days[first_avalaible_dry_day]] = val
				dry_days = deleteElement(dry_days, first_avalaible_dry_day)
				full_lakes[val] = index
			} else {
				full_lakes[val] = index
			}
		}
	}

	return ans
}

func findFirstIndexOfElementNotSmallerThan(slice []int, element int) int {
	for i, e := range slice {
		if e >= element {
			return i
		}
	}
	return -1
}

func deleteElement(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}
