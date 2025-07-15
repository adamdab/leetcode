package main

func isBetween(char, leftBound, rightBound rune) bool {
	return char >= leftBound && char <= rightBound
}

func isValidRune(char rune) bool {
	var valid = false
	valid = valid || isBetween(char, '0', '9')
	valid = valid || isBetween(char, 'a', 'z')
	valid = valid || isBetween(char, 'A', 'Z')

	return valid
}

func vowels() []rune {
	return []rune{'a', 'A', 'e', 'E', 'i', 'I', 'o', 'O', 'u', 'U'}
}

func isVowel(char rune) bool {
	return slices.Contains(vowels(), char)
}

func isConsonant(char rune) bool {
	// does not check if rune is valid rune
	return !unicode.IsDigit(char) && !isVowel(char)
}

func isValid(word string) bool {
	var hasVowel = false
	var hasConsonant = false
	if len(word) < 3 {
		return false
	}
	for _, char := range word {
		if !isValidRune(char) {
			return false
		} else {
			hasVowel = hasVowel || isVowel(char)
			hasConsonant = hasConsonant || isConsonant(char)
		}
	}
	return hasVowel && hasConsonant
}