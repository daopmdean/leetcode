package interview_150

func rangeBitwiseAnd(left int, right int) int {
	for right > left {
		right = right & (right - 1)
	}
	return left & right
}
