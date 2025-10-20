package main

func finalValueAfterOperations(operations []string) int {
	x := 0
	for _, str := range operations {
		if contains(str, '+') {
			x++
		} else {
			x--
		}
	}
	return x
}

func contains(str string, char rune) bool {
	for _, c := range str {
		if c == char {
			return true
		}
	}
	return false
}
