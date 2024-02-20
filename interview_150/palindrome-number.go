package interview_150

import "strconv"

func isPalindromeNumber(x int) bool {
	if x < 0 {
		return false
	}

	s := strconv.Itoa(x)
	for i := 0; i < len(s); i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}

	return true
}
