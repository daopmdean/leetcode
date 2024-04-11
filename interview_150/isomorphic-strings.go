package interview_150

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	ms := map[byte]byte{}
	mt := map[byte]byte{}

	for i := 0; i < len(s); i++ {
		if _, ok := ms[s[i]]; !ok {
			ms[s[i]] = t[i]
		}

		if _, ok := mt[t[i]]; !ok {
			mt[t[i]] = s[i]
		}

		if s[i] != mt[t[i]] || t[i] != ms[s[i]] {
			return false
		}
	}

	return true
}
