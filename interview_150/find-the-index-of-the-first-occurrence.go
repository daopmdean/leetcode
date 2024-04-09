package interview_150

func strStr(haystack string, needle string) int {
	for i := 0; i < len(haystack); i++ {
		if match(haystack[i:], needle) {
			return i
		}
	}

	return -1
}

func match(haystack, needle string) bool {
	if len(needle) > len(haystack) {
		return false
	}

	for i := 0; i < len(needle); i++ {
		if needle[i] != haystack[i] {
			return false
		}
	}

	return true
}
