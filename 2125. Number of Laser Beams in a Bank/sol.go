package main

func numberOfBeams(bank []string) int {
	total := 0
	curr := 0
	prev := 0

	for _, row := range bank {
		for _, c := range row {
			if c == '1' {
				curr++
			}
		}
		if curr > 0 {
			total += prev * curr
			prev = curr
		}
		curr = 0
	}

	return total
}
