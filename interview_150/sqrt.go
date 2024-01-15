package interview_150

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}

	for i := 1; i <= x; i++ {
		double := i * i
		if double == x {
			return i
		} else if double > x {
			return i - 1
		}
	}

	return 1
}
