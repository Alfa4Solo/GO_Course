package Pr1

import (
	"unicode"
)

func IsPalindrome(s string) bool {
	var normalized []rune

	for _, ch := range s {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			normalized = append(normalized, unicode.ToLower(ch))
		}
	}

	n := len(normalized)
	for i := 0; i < n/2; i++ {
		if normalized[i] != normalized[n-1-i] {
			return false
		}
	}
	return true
}
