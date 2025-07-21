package main

import (
	"strings"
)

func makeFancyString(s string) string {
	counter := 0
	const max = 2
	var current rune
	var sb strings.Builder
	for _, c := range s {
		if c == current {
			counter++
			if counter <= max {
				sb.WriteRune(c)
			}
		} else {
			current = c
			counter = 1
			sb.WriteRune(c)
		}
	}
	return sb.String()
}
