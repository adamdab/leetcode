package main

func removeAnagrams(words []string) []string {

	if len(words) == 1 {
		return words
	}

	result := []string{}

	first := 0
	second := 1

	for second < len(words) {
		if isAnagram(words[first], words[second]) {
			second++
		} else {
			result = append(result, words[first])
			first = second
			second++
		}

		if second == len(words) {
			result = append(result, words[first])
		}
	}

	return result
}

func isAnagram(first, second string) bool {
	counter := [26]int{}

	for _, char := range first {
		counter[char-'a']++
	}
	for _, char := range second {
		counter[char-'a']--
	}
	for _, count := range counter {
		if count != 0 {
			return false
		}
	}
	return true
}
