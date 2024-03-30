package interview_150

func lengthOfLastWord(s string) int {

	length := 0

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if length == 0 {
				continue
			}
			break
		}
		length++
	}

	return length
}
