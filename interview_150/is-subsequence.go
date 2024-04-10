package interview_150

func isSubsequence(s string, t string) bool {
	if len(t) < len(s) {
		return false
	}

	i, j := 0, 0
	for {
		if i >= len(s) {
			return true
		}
		if j >= len(t) {
			return false
		}
		if s[i] == t[j] {
			i++
			j++
			continue
		}
		j++
	}

	return false
}
