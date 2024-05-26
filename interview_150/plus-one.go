package interview_150

func plusOne(digits []int) []int {
	hold := false
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+1 < 10 {
			digits[i]++
			hold = false
			break
		} else {
			digits[i] = 0
			hold = true
		}
	}

	if hold {
		digits = append([]int{1}, digits...)
	}

	return digits
}
