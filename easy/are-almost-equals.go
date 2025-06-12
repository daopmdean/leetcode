package easy

func areAlmostEqual(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return fromBeginning(s1[i:], s2[i:])
		}
	}

	return true
}

func fromBeginning(s1 string, s2 string) bool {
	for i := 0; i < len(s1)-1; i++ {
		for j := i + 1; j < len(s1); j++ {
			bts := []byte(s1)
			tmp := bts[i]
			bts[i] = bts[j]
			bts[j] = tmp
			if string(bts) == s2 {
				return true
			}
		}
	}
	return false
}
