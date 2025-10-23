package main

import "strings"

func hasSameDigits(s string) bool {
	if len(s) == 2 {
		return s[0] == s[1]
	}

	var sb strings.Builder

	for i := 1; i < len(s); i++ {
		prev := int(s[i-1] - '0')
		curr := int(s[i] - '0')
		res := (prev + curr) % 10
		sb.WriteByte(byte(res + '0'))
	}
	return hasSameDigits(sb.String())

}
