package interview_150

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := map[rune]int{}
	for _, b := range s {
		m[b]++
	}

	for _, b := range t {
		if v, ok := m[b]; ok {
			if v == 0 {
				return false
			}
			m[b]--
		} else {
			return false
		}
	}

	return true
}
