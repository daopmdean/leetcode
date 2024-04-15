package interview_150

import "strings"

func wordPattern(pattern string, s string) bool {
	slc := strings.Split(s, " ")
	if len(pattern) != len(slc) {
		return false
	}

	m1 := map[rune]string{}
	m2 := map[string]rune{}

	for i, b := range pattern {
		if _, ok := m1[b]; !ok {
			m1[b] = slc[i]
		}

		if _, ok := m2[slc[i]]; !ok {
			m2[slc[i]] = b
		}

		if m1[b] != slc[i] || m2[slc[i]] != b {
			return false
		}
	}

	return true
}
