package interview_150

func canConstruct(ransomNote string, magazine string) bool {
	m := map[rune]int{}

	for _, c := range magazine {
		m[c]++
	}

	for _, c := range ransomNote {
		if m[c] > 0 {
			m[c]--
		} else {
			return false
		}
	}

	return true
}
