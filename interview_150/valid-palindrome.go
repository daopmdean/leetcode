package interview_150

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	newS := ""
	for _, c := range s {
		if unicode.IsLetter(c) || unicode.IsNumber(c) {
			newS += string(c)
		}
	}

	start, last := 0, len(newS)-1
	for start <= last {
		if newS[start] != newS[last] {
			return false
		}
		start++
		last--
	}

	return true
}
